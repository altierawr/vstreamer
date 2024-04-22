package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PlaySessionMedia holds the schema definition for the PlaySessionMedia entity.
type PlaySessionMedia struct {
	ent.Schema
}

// Fields of the PlaySessionMedia.
func (PlaySessionMedia) Fields() []ent.Field {
	return []ent.Field{
		field.Strings("resolutions").
			Default([]string{}),
	}
}

// Edges of the PlaySessionMedia.
func (PlaySessionMedia) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("audio_tracks", AudioTrack.Type),
		edge.From("video", Video.Type).
			Ref("play_session_medias").
			Unique(),
		edge.From("session", PlaySession.Type).
			Ref("media").
			Unique().
			Required(),
		edge.To("video_codecs", VideoCodec.Type),
	}
}
