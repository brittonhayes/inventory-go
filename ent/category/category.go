// Code generated by entc, DO NOT EDIT.

package category

import (
	"time"
)

const (
	// Label holds the string label denoting the category type in the database.
	Label = "category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeVehicle holds the string denoting the vehicle edge name in mutations.
	EdgeVehicle = "vehicle"
	// EdgeTool holds the string denoting the tool edge name in mutations.
	EdgeTool = "tool"
	// EdgeImplement holds the string denoting the implement edge name in mutations.
	EdgeImplement = "implement"
	// EdgeEquipment holds the string denoting the equipment edge name in mutations.
	EdgeEquipment = "equipment"
	// EdgeLocation holds the string denoting the location edge name in mutations.
	EdgeLocation = "location"
	// Table holds the table name of the category in the database.
	Table = "categories"
	// VehicleTable is the table that holds the vehicle relation/edge.
	VehicleTable = "vehicles"
	// VehicleInverseTable is the table name for the Vehicle entity.
	// It exists in this package in order to avoid circular dependency with the "vehicle" package.
	VehicleInverseTable = "vehicles"
	// VehicleColumn is the table column denoting the vehicle relation/edge.
	VehicleColumn = "category_vehicle"
	// ToolTable is the table that holds the tool relation/edge.
	ToolTable = "tools"
	// ToolInverseTable is the table name for the Tool entity.
	// It exists in this package in order to avoid circular dependency with the "tool" package.
	ToolInverseTable = "tools"
	// ToolColumn is the table column denoting the tool relation/edge.
	ToolColumn = "category_tool"
	// ImplementTable is the table that holds the implement relation/edge.
	ImplementTable = "implements"
	// ImplementInverseTable is the table name for the Implement entity.
	// It exists in this package in order to avoid circular dependency with the "implement" package.
	ImplementInverseTable = "implements"
	// ImplementColumn is the table column denoting the implement relation/edge.
	ImplementColumn = "category_implement"
	// EquipmentTable is the table that holds the equipment relation/edge.
	EquipmentTable = "equipment"
	// EquipmentInverseTable is the table name for the Equipment entity.
	// It exists in this package in order to avoid circular dependency with the "equipment" package.
	EquipmentInverseTable = "equipment"
	// EquipmentColumn is the table column denoting the equipment relation/edge.
	EquipmentColumn = "category_equipment"
	// LocationTable is the table that holds the location relation/edge.
	LocationTable = "locations"
	// LocationInverseTable is the table name for the Location entity.
	// It exists in this package in order to avoid circular dependency with the "location" package.
	LocationInverseTable = "locations"
	// LocationColumn is the table column denoting the location relation/edge.
	LocationColumn = "category_location"
)

// Columns holds all SQL columns for category fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)
