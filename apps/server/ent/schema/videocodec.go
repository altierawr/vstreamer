package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// VideoCodec holds the schema definition for the VideoCodec entity.
type VideoCodec struct {
	ent.Schema
}

// Fields of the VideoCodec.
func (VideoCodec) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("mime"),
		field.Enum("dynamic_range").
			Values("sdr", "hdr"),
	}
}

// Edges of the VideoCodec.
func (VideoCodec) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("media", PlaySessionMedia.Type).
			Ref("video_codecs").
			Unique(),
	}
}
