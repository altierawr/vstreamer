// Code generated by ent, DO NOT EDIT.

package audiocodec

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the audiocodec type in the database.
	Label = "audio_codec"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldMime holds the string denoting the mime field in the database.
	FieldMime = "mime"
	// EdgeStreams holds the string denoting the streams edge name in mutations.
	EdgeStreams = "streams"
	// EdgeAudioTracks holds the string denoting the audio_tracks edge name in mutations.
	EdgeAudioTracks = "audio_tracks"
	// EdgeMedia holds the string denoting the media edge name in mutations.
	EdgeMedia = "media"
	// Table holds the table name of the audiocodec in the database.
	Table = "audio_codecs"
	// StreamsTable is the table that holds the streams relation/edge.
	StreamsTable = "streams"
	// StreamsInverseTable is the table name for the Stream entity.
	// It exists in this package in order to avoid circular dependency with the "stream" package.
	StreamsInverseTable = "streams"
	// StreamsColumn is the table column denoting the streams relation/edge.
	StreamsColumn = "audio_codec_streams"
	// AudioTracksTable is the table that holds the audio_tracks relation/edge. The primary key declared below.
	AudioTracksTable = "audio_track_codecs"
	// AudioTracksInverseTable is the table name for the AudioTrack entity.
	// It exists in this package in order to avoid circular dependency with the "audiotrack" package.
	AudioTracksInverseTable = "audio_tracks"
	// MediaTable is the table that holds the media relation/edge.
	MediaTable = "audio_codecs"
	// MediaInverseTable is the table name for the PlaySessionMedia entity.
	// It exists in this package in order to avoid circular dependency with the "playsessionmedia" package.
	MediaInverseTable = "play_session_media"
	// MediaColumn is the table column denoting the media relation/edge.
	MediaColumn = "play_session_media_audio_codecs"
)

// Columns holds all SQL columns for audiocodec fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldMime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "audio_codecs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"play_session_media_audio_codecs",
}

var (
	// AudioTracksPrimaryKey and AudioTracksColumn2 are the table columns denoting the
	// primary key for the audio_tracks relation (M2M).
	AudioTracksPrimaryKey = []string{"audio_track_id", "audio_codec_id"}
)

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

// OrderOption defines the ordering options for the AudioCodec queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByMime orders the results by the mime field.
func ByMime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMime, opts...).ToFunc()
}

// ByStreamsCount orders the results by streams count.
func ByStreamsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStreamsStep(), opts...)
	}
}

// ByStreams orders the results by streams terms.
func ByStreams(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStreamsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
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

// ByMediaField orders the results by media field.
func ByMediaField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMediaStep(), sql.OrderByField(field, opts...))
	}
}
func newStreamsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StreamsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, StreamsTable, StreamsColumn),
	)
}
func newAudioTracksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AudioTracksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, AudioTracksTable, AudioTracksPrimaryKey...),
	)
}
func newMediaStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MediaInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MediaTable, MediaColumn),
	)
}
