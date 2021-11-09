package client

import (
	"github.com/open-farms/inventory/ent/proto/entpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *Client) GetEquipment(request *entpb.GetEquipmentRequest) (*entpb.Equipment, error) {
	client := entpb.NewEquipmentServiceClient(c.connection)
	resp, err := client.Get(c.ctx, request)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (c *Client) CreateEquipment(request *entpb.CreateEquipmentRequest) (*entpb.Equipment, error) {
	client := entpb.NewEquipmentServiceClient(c.connection)
	resp, err := client.Create(c.ctx, request)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (c *Client) UpdateEquipment(request *entpb.UpdateEquipmentRequest) (*entpb.Equipment, error) {
	client := entpb.NewEquipmentServiceClient(c.connection)
	resp, err := client.Update(c.ctx, request)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (c *Client) DeleteEquipment(request *entpb.DeleteEquipmentRequest) (*emptypb.Empty, error) {
	client := entpb.NewEquipmentServiceClient(c.connection)
	resp, err := client.Delete(c.ctx, request)
	if err != nil {
		return nil, err
	}

	return resp, err
}
