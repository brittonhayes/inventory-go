package commands

import (
	"encoding/json"
	"fmt"

	"github.com/open-farms/inventory/client"
	"github.com/open-farms/inventory/ent/proto/entpb"
)

type EquipmentCmd struct {
	Get    EquipmentGetCmd    `cmd:"" help:"Get equipment by ID"`
	Delete EquipmentDeleteCmd `cmd:"" help:"Delete equipment by ID"`
	Create EquipmentCreateCmd `cmd:"" help:"Create equipment"`
}

type EquipmentGetCmd struct {
	ID int32 `arg:"" help:"Equipment unique ID"`
}

func (e *EquipmentGetCmd) Run(globals *Globals) error {
	c, err := client.NewClient(globals.Context, globals.URL)
	if err != nil {
		return err
	}

	resp, err := c.GetEquipment(&entpb.GetEquipmentRequest{Id: e.ID})
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(&resp, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}

type EquipmentDeleteCmd struct {
	ID int32 `arg:"" help:"Equipment unique ID"`
}

func (e *EquipmentDeleteCmd) Run(globals *Globals) error {
	c, err := client.NewClient(globals.Context, globals.URL)
	if err != nil {
		return err
	}

	req := &entpb.DeleteEquipmentRequest{Id: e.ID}
	_, err = c.DeleteEquipment(req)
	if err != nil {
		return err
	}

	fmt.Println("deleted equipment with ID =", req.Id)
	return nil
}

type EquipmentCreateCmd struct {
	Equipments []string `arg:"" help:"The equipment to create"`
}

func (e *EquipmentCreateCmd) Run(globals *Globals) error {
	c, err := client.NewClient(globals.Context, globals.URL)
	if err != nil {
		return err
	}

	for _, equipment := range e.Equipments {
		eq := &entpb.Equipment{}
		err = json.Unmarshal([]byte(equipment), &eq)
		if err != nil {
			return err
		}

		_, err = c.CreateEquipment(&entpb.CreateEquipmentRequest{Equipment: eq})
		if err != nil {
			return err
		}
	}

	return nil
}
