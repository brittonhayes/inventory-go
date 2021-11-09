package settings

import (
	"context"
	"os"

	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/open-farms/inventory/ent"
)

const (
	DefaultConfigPath string = "./config/config.yaml"
)

func Migrate(configPath string, dry bool) error {
	if configPath == "" {
		configPath = DefaultConfigPath
	}

	// Initialize configuration
	cfg, err := Configure(configPath)
	if err != nil {
		return err
	}

	dbDriver, err := cfg.Value("storage.database.driver").String()
	if err != nil {
		return err
	}

	dbURI, err := cfg.Value("storage.database.source").String()
	if err != nil {
		return err
	}

	dbDebug, err := cfg.Value("storage.database.debug").Bool()
	if err != nil {
		return err
	}

	// Setup client for database
	client, err := ent.Open(dbDriver, dbURI)
	if err != nil {
		return err
	}
	defer client.Close()

	if dry {
		err := client.Schema.WriteTo(context.Background(), os.Stdout)
		if err != nil {
			return err
		}
	}

	// Run the auto migration tool
	ctx := context.Background()
	if dbDebug {
		err = client.Debug().Schema.Create(ctx)
		if err != nil {
			return err
		}
		return nil
	}

	err = client.Schema.Create(ctx,
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
	)
	if err != nil {
		return err
	}

	return nil
}

func Configure(cfgFile string) (config.Config, error) {
	// Setup config from file
	source := config.WithSource(
		env.NewSource("INVENTORY_"),
		file.NewSource(cfgFile),
	)
	c := config.New(source)
	err := c.Load()
	if err != nil {
		return nil, err
	}

	return c, nil
}
