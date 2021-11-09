package schema

import (
	"regexp"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TODO
// vehicles
// w/ 1:many implement
// w/ 1:many part
// w/ 1:1 location

// Vehicle holds the schema definition for the Vehicle entity.
type Vehicle struct {
	ent.Schema
}

func (Vehicle) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Vehicle.
func (Vehicle) Fields() []ent.Field {
	return []ent.Field{
		field.String("make").
			Annotations(entproto.Field(4)),
		field.String("model").
			Annotations(entproto.Field(5)),
		field.Int64("hours").
			Default(0).
			Min(0).
			Annotations(entproto.Field(6)),
		field.Int64("year").
			Optional().
			Min(1900).
			Annotations(entproto.Field(7)),
		field.Bool("active").
			Default(true).
			Annotations(entproto.Field(8)),
		field.String("power").
			Optional().
			Default("GAS").
			Match(regexp.MustCompile("(GAS|DIESEL|ELECTRIC)")).
			Annotations(entproto.Field(9)),
	}
}

// Edges of the Vehicle.
func (Vehicle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("location", Location.Type).
			Ref("vehicle").
			Unique().
			Required().
			Annotations(entproto.Field(10)),
	}
}

func (Vehicle) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
