package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Feedback holds the schema definition for the Feedback entity.
type Feedback struct {
	ent.Schema
}

// Fields of the Feedback.
func (Feedback) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name"),
		field.String("content"),
		field.Time("created_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("updated_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

// Edges of the Feedback.
func (Feedback) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("like", Like.Type).
			Unique(),
	}
}
