package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Implement holds the schema definition for the Implement entity.
type Implement struct {
	ent.Schema
}

func (Implement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Implement.
func (Implement) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Annotations(entproto.Field(4)),
	}
}

// Edges of the Implement.
func (Implement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("location", Location.Type).
			Ref("implement").
			Unique().
			Required().
			Annotations(entproto.Field(5)),
		edge.From("category", Category.Type).
			Ref("implement").
			Unique().
			Annotations(entproto.Field(7)),
	}
}

func (Implement) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
