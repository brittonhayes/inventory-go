package main

import (
	"flag"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/open-farms/inventory/ent"
	elk "github.com/open-farms/inventory/ent/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"

	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/open-farms/inventory/ent/proto/entpb"
	"github.com/open-farms/inventory/internal/logging"
	"github.com/open-farms/inventory/internal/settings"
	"go.uber.org/zap"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "inventory"
	// Version is the version of the compiled software.
	Version string = "x.y.z"

	// flagconf is the config flag.
	flagconf string

	logger *zap.Logger
)

func init() {
	flag.StringVar(&flagconf, "config", "./config", "config path, eg: -config config.yaml")
}

func setup() (config.Config, error) {
	flag.Parse()
	logger = zap.NewExample(
		zap.Fields(zap.String("service", Name)),
	)

	cfg, err := settings.Configure(flagconf)
	if err != nil {
		return nil, err
	}

	migrate, err := cfg.Value("storage.database.migrate").Bool()
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if migrate {
		err = settings.Migrate(flagconf, false)
		if err != nil {
			return nil, err
		}
	}

	return cfg, nil
}

func grpcService(client *ent.Client, address string, timeout time.Duration) *transgrpc.Server {
	equipmentSvc := entpb.NewEquipmentService(client)
	vehicleSvc := entpb.NewVehicleService(client)
	toolsSvc := entpb.NewToolService(client)
	implementSvc := entpb.NewImplementService(client)

	grpcServer := transgrpc.NewServer(transgrpc.Address(address), transgrpc.Timeout(timeout))
	entpb.RegisterEquipmentServiceServer(grpcServer, equipmentSvc)
	entpb.RegisterVehicleServiceServer(grpcServer, vehicleSvc)
	entpb.RegisterToolServiceServer(grpcServer, toolsSvc)
	entpb.RegisterImplementServiceServer(grpcServer, implementSvc)

	return grpcServer
}

func httpService(c *ent.Client, address string, timeout time.Duration) *transhttp.Server {
	r := chi.NewRouter()
	r.Use(logging.Logger(logger))
	r.Use(middleware.StripSlashes)
	r.Route("/v1", func(r chi.Router) {
		elk.NewEquipmentHandler(c, logger).MountRoutes(r)
		elk.NewVehicleHandler(c, logger).MountRoutes(r)
		elk.NewToolHandler(c, logger).MountRoutes(r)
		elk.NewImplementHandler(c, logger).MountRoutes(r)
		elk.NewLocationHandler(c, logger).MountRoutes(r)
		elk.NewCategoryHandler(c, logger).MountRoutes(r)
	})

	httpServer := transhttp.NewServer(transhttp.Address(address))
	httpServer.HandlePrefix("/", r)
	return httpServer
}

func main() {

	// Initialize configs
	// and parse flags
	cfg, err := setup()
	if err != nil {
		logger.Fatal(err.Error())
	}

	driver, _ := cfg.Value("storage.database.driver").String()
	source, _ := cfg.Value("storage.database.source").String()
	httpAddress, _ := cfg.Value("server.http.addr").String()
	httpTimeout, _ := cfg.Value("server.http.timeout").Duration()
	grpcAddress, _ := cfg.Value("server.grpc.addr").String()
	grpcTimeout, _ := cfg.Value("server.grpc.timeout").Duration()

	// Connect to database
	c, err := ent.Open(driver, source)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer c.Close()

	// Create servers with context and graceful shutdown timer
	httpServer := httpService(c, httpAddress, httpTimeout)
	grpcServer := grpcService(c, grpcAddress, grpcTimeout)
	app := kratos.New(
		kratos.Name("inventory"),
		kratos.Server(httpServer, grpcServer),
		kratos.Version(Version),
		kratos.Logger(
			logging.NewZapLogger(
				zap.NewProductionEncoderConfig(),
				zap.NewProductionConfig().Level,
			),
		),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
