package client

import (
	"github.com/open-farms/inventory/ent/proto/entpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *Client) GetVehicle(request *entpb.GetVehicleRequest) (*entpb.Vehicle, error) {
	client := entpb.NewVehicleServiceClient(c.connection)
	resp, err := client.Get(c.ctx, request)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (c *Client) CreateVehicle(request *entpb.CreateVehicleRequest) (*entpb.Vehicle, error) {
	client := entpb.NewVehicleServiceClient(c.connection)
	resp, err := client.Create(c.ctx, request)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (c *Client) UpdateVehicle(request *entpb.UpdateVehicleRequest) (*entpb.Vehicle, error) {
	client := entpb.NewVehicleServiceClient(c.connection)
	resp, err := client.Update(c.ctx, request)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (c *Client) DeleteVehicle(request *entpb.DeleteVehicleRequest) (*emptypb.Empty, error) {
	client := entpb.NewVehicleServiceClient(c.connection)
	resp, err := client.Delete(c.ctx, request)
	if err != nil {
		return nil, err
	}

	return resp, err
}
