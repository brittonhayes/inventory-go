package commands

import (
	"context"
)

type Globals struct {
	Context context.Context `kong:"-"`
	URL     string          `short:"u" default:"127.0.0.1:9000" help:"gRPC endpoint to send requests"`

	Config string `help:"Path to config file" short:"c" default:"./config/config.yaml" type:"path"`
}
