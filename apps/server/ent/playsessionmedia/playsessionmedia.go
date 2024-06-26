// Code generated by ent, DO NOT EDIT.

package playsessionmedia

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the playsessionmedia type in the database.
	Label = "play_session_media"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldResolutions holds the string denoting the resolutions field in the database.
	FieldResolutions = "resolutions"
	// EdgeAudioTracks holds the string denoting the audio_tracks edge name in mutations.
	EdgeAudioTracks = "audio_tracks"
	// EdgeVideo holds the string denoting the video edge name in mutations.
	EdgeVideo = "video"
	// EdgeSession holds the string denoting the session edge name in mutations.
	EdgeSession = "session"
	// EdgeVideoCodecs holds the string denoting the video_codecs edge name in mutations.
	EdgeVideoCodecs = "video_codecs"
	// EdgeAudioCodecs holds the string denoting the audio_codecs edge name in mutations.
	EdgeAudioCodecs = "audio_codecs"
	// Table holds the table name of the playsessionmedia in the database.
	Table = "play_session_media"
	// AudioTracksTable is the table that holds the audio_tracks relation/edge.
	AudioTracksTable = "audio_tracks"
	// AudioTracksInverseTable is the table name for the AudioTrack entity.
	// It exists in this package in order to avoid circular dependency with the "audiotrack" package.
	AudioTracksInverseTable = "audio_tracks"
	// AudioTracksColumn is the table column denoting the audio_tracks relation/edge.
	AudioTracksColumn = "play_session_media_audio_tracks"
	// VideoTable is the table that holds the video relation/edge.
	VideoTable = "play_session_media"
	// VideoInverseTable is the table name for the Video entity.
	// It exists in this package in order to avoid circular dependency with the "video" package.
	VideoInverseTable = "videos"
	// VideoColumn is the table column denoting the video relation/edge.
	VideoColumn = "video_play_session_medias"
	// SessionTable is the table that holds the session relation/edge.
	SessionTable = "play_session_media"
	// SessionInverseTable is the table name for the PlaySession entity.
	// It exists in this package in order to avoid circular dependency with the "playsession" package.
	SessionInverseTable = "play_sessions"
	// SessionColumn is the table column denoting the session relation/edge.
	SessionColumn = "play_session_media"
	// VideoCodecsTable is the table that holds the video_codecs relation/edge.
	VideoCodecsTable = "video_codecs"
	// VideoCodecsInverseTable is the table name for the VideoCodec entity.
	// It exists in this package in order to avoid circular dependency with the "videocodec" package.
	VideoCodecsInverseTable = "video_codecs"
	// VideoCodecsColumn is the table column denoting the video_codecs relation/edge.
	VideoCodecsColumn = "play_session_media_video_codecs"
	// AudioCodecsTable is the table that holds the audio_codecs relation/edge.
	AudioCodecsTable = "audio_codecs"
	// AudioCodecsInverseTable is the table name for the AudioCodec entity.
	// It exists in this package in order to avoid circular dependency with the "audiocodec" package.
	AudioCodecsInverseTable = "audio_codecs"
	// AudioCodecsColumn is the table column denoting the audio_codecs relation/edge.
	AudioCodecsColumn = "play_session_media_audio_codecs"
)

// Columns holds all SQL columns for playsessionmedia fields.
var Columns = []string{
	FieldID,
	FieldResolutions,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "play_session_media"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"play_session_media",
	"video_play_session_medias",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultResolutions holds the default value on creation for the "resolutions" field.
	DefaultResolutions []string
)

// OrderOption defines the ordering options for the PlaySessionMedia queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAudioTracksCount orders the results by audio_tracks count.
func ByAudioTracksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAudioTracksStep(), opts...)
	}
}

// ByAudioTracks orders the results by audio_tracks terms.
func ByAudioTracks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAudioTracksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByVideoField orders the results by video field.
func ByVideoField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVideoStep(), sql.OrderByField(field, opts...))
	}
}

// BySessionField orders the results by session field.
func BySessionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSessionStep(), sql.OrderByField(field, opts...))
	}
}

// ByVideoCodecsCount orders the results by video_codecs count.
func ByVideoCodecsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVideoCodecsStep(), opts...)
	}
}

// ByVideoCodecs orders the results by video_codecs terms.
func ByVideoCodecs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVideoCodecsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAudioCodecsCount orders the results by audio_codecs count.
func ByAudioCodecsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAudioCodecsStep(), opts...)
	}
}

// ByAudioCodecs orders the results by audio_codecs terms.
func ByAudioCodecs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAudioCodecsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAudioTracksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AudioTracksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AudioTracksTable, AudioTracksColumn),
	)
}
func newVideoStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VideoInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, VideoTable, VideoColumn),
	)
}
func newSessionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SessionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, SessionTable, SessionColumn),
	)
}
func newVideoCodecsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VideoCodecsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, VideoCodecsTable, VideoCodecsColumn),
	)
}
func newAudioCodecsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AudioCodecsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AudioCodecsTable, AudioCodecsColumn),
	)
}
