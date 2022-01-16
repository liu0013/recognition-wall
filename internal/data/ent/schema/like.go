package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Like holds the schema definition for the Like entity.
type Like struct {
	ent.Schema
}

// Fields of the Like.
func (Like) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int32("count"),
	}
}

// Edges of the Like.
func (Like) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("feedback", Feedback.Type).
			Ref("like").
			Unique().
			Required(),
	}
}
