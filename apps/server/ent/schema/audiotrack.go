package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AudioTrack holds the schema definition for the AudioTrack entity.
type AudioTrack struct {
	ent.Schema
}

// Fields of the AudioTrack.
func (AudioTrack) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("nr_channels"),
		field.String("channel_layout"),
		field.String("language").
			Optional(),
	}
}

// Edges of the AudioTrack.
func (AudioTrack) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("codecs", AudioCodec.Type),
		edge.From("media", PlaySessionMedia.Type).
			Ref("audio_tracks").
			Unique().
			Required(),
	}
}
