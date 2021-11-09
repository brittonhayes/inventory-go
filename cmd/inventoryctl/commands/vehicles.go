package commands

type VehiclesCmd struct {
	Get  VehiclesGetCmd  `cmd:"" help:"Get a vehicle by ID"`
	List VehiclesListCmd `cmd:"" help:"List all vehicles"`
}

type VehiclesGetCmd struct {
	ID string `arg:"" help:"Vehicle unique ID"`
}

type VehiclesListCmd struct {
	Limit int `default:"50" help:"The maximum number of vehicles to return"`
}
