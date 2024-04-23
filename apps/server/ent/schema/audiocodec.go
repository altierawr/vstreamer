package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AudioCodec holds the schema definition for the AudioCodec entity.
type AudioCodec struct {
	ent.Schema
}

// Fields of the AudioCodec.
func (AudioCodec) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("mime"),
	}
}

// Edges of the AudioCodec.
func (AudioCodec) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("streams", Stream.Type),
		edge.From("audio_tracks", AudioTrack.Type).
			Ref("codecs"),
		edge.From("media", PlaySessionMedia.Type).
			Ref("audio_codecs").
			Unique(),
	}
}
