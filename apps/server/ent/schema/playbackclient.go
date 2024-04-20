package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PlaybackClient holds the schema definition for the PlaybackClient entity.
type PlaybackClient struct {
	ent.Schema
}

// Fields of the PlaybackClient.
func (PlaybackClient) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("is_buffered").
			Default(false),
	}
}

// Edges of the PlaybackClient.
func (PlaybackClient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("session", PlaySession.Type).
			Ref("clients").
			Unique(),
	}
}
