package schema

import (
	"regexp"
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

var conditionPattern = regexp.MustCompile("(MINT|GOOD|POOR|BROKEN)")

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("create_time").
			Default(time.Now).
			Annotations(entproto.Field(2)),
		field.Time("update_time").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(entproto.Field(3)),
	}
}
