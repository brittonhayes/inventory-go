package client

import (
	"context"

	"google.golang.org/grpc"

	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
)

type Client struct {
	ctx        context.Context
	connection *grpc.ClientConn
}

func NewClient(ctx context.Context, endpoint string) (*Client, error) {
	conn, err := transgrpc.DialInsecure(ctx, transgrpc.WithEndpoint(endpoint))
	if err != nil {
		return nil, err
	}

	return &Client{
		ctx:        ctx,
		connection: conn,
	}, nil
}
