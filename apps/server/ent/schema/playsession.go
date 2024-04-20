package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PlaySession holds the schema definition for the PlaySession entity.
type PlaySession struct {
	ent.Schema
}

// Fields of the PlaySession.
func (PlaySession) Fields() []ent.Field {
	return []ent.Field{
		field.Int("current_time").
			Optional(),
		field.Enum("state").
			Values("PLAYING", "PAUSED", "BUFFERING", "STOPPED").
			Default("STOPPED"),
	}
}

// Edges of the PlaySession.
func (PlaySession) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("clients", PlaybackClient.Type),
		edge.To("media", PlaySessionMedia.Type).
			Unique(),
	}
}
