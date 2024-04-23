package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Stream holds the schema definition for the Stream entity.
type Stream struct {
	ent.Schema
}

// Fields of the Stream.
func (Stream) Fields() []ent.Field {
	return []ent.Field{
		field.Int("width"),
		field.Int("height"),
		field.String("container"),
		field.Int("segment_duration"),
		field.Enum("quality").
			Values("maximum", "medium", "low"),
		field.Enum("type").
			Values("direct", "remux", "video_transcode", "audio_transcode", "full_transcode"),
	}
}

// Edges of the Stream.
func (Stream) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("video_codec", VideoCodec.Type).
			Ref("streams").
			Unique(),
		edge.From("audio_codec", AudioCodec.Type).
			Ref("streams").
			Unique(),
	}
}
