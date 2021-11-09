package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Equipment holds the schema definition for the Equipment entity.
type Equipment struct {
	ent.Schema
}

func (Equipment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Equipment.
func (Equipment) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			Annotations(entproto.Field(4)),
		field.String("condition").
			Match(conditionPattern).
			Annotations(entproto.Field(5)),
	}
}

func (Equipment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}

// Edges of the Equipment.
func (Equipment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("location", Location.Type).
			Ref("equipment").
			Unique().
			Required().
			Annotations(entproto.Field(6)),
		edge.From("category", Category.Type).
			Ref("equipment").
			Unique().
			Annotations(entproto.Field(7)),
	}
}
