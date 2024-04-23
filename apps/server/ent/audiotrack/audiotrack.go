// Code generated by ent, DO NOT EDIT.

package audiotrack

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the audiotrack type in the database.
	Label = "audio_track"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldNrChannels holds the string denoting the nr_channels field in the database.
	FieldNrChannels = "nr_channels"
	// FieldChannelLayout holds the string denoting the channel_layout field in the database.
	FieldChannelLayout = "channel_layout"
	// FieldLanguage holds the string denoting the language field in the database.
	FieldLanguage = "language"
	// EdgeCodecs holds the string denoting the codecs edge name in mutations.
	EdgeCodecs = "codecs"
	// EdgeMedia holds the string denoting the media edge name in mutations.
	EdgeMedia = "media"
	// Table holds the table name of the audiotrack in the database.
	Table = "audio_tracks"
	// CodecsTable is the table that holds the codecs relation/edge. The primary key declared below.
	CodecsTable = "audio_track_codecs"
	// CodecsInverseTable is the table name for the AudioCodec entity.
	// It exists in this package in order to avoid circular dependency with the "audiocodec" package.
	CodecsInverseTable = "audio_codecs"
	// MediaTable is the table that holds the media relation/edge.
	MediaTable = "audio_tracks"
	// MediaInverseTable is the table name for the PlaySessionMedia entity.
	// It exists in this package in order to avoid circular dependency with the "playsessionmedia" package.
	MediaInverseTable = "play_session_media"
	// MediaColumn is the table column denoting the media relation/edge.
	MediaColumn = "play_session_media_audio_tracks"
)

// Columns holds all SQL columns for audiotrack fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldNrChannels,
	FieldChannelLayout,
	FieldLanguage,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "audio_tracks"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"play_session_media_audio_tracks",
}

var (
	// CodecsPrimaryKey and CodecsColumn2 are the table columns denoting the
	// primary key for the codecs relation (M2M).
	CodecsPrimaryKey = []string{"audio_track_id", "audio_codec_id"}
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

// OrderOption defines the ordering options for the AudioTrack queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByNrChannels orders the results by the nr_channels field.
func ByNrChannels(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNrChannels, opts...).ToFunc()
}

// ByChannelLayout orders the results by the channel_layout field.
func ByChannelLayout(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChannelLayout, opts...).ToFunc()
}

// ByLanguage orders the results by the language field.
func ByLanguage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLanguage, opts...).ToFunc()
}

// ByCodecsCount orders the results by codecs count.
func ByCodecsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCodecsStep(), opts...)
	}
}

// ByCodecs orders the results by codecs terms.
func ByCodecs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCodecsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMediaField orders the results by media field.
func ByMediaField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMediaStep(), sql.OrderByField(field, opts...))
	}
}
func newCodecsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CodecsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, CodecsTable, CodecsPrimaryKey...),
	)
}
func newMediaStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MediaInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MediaTable, MediaColumn),
	)
}
