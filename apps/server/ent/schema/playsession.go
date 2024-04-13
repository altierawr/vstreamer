package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// PlaySession holds the schema definition for the PlaySession entity.
type PlaySession struct {
	ent.Schema
}

// Fields of the PlaySession.
func (PlaySession) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the PlaySession.
func (PlaySession) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("video", Video.Type).
			Ref("play_sessions").
			Unique(),
	}
}
