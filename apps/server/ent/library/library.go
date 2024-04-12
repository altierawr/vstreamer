// Code generated by ent, DO NOT EDIT.

package library

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the library type in the database.
	Label = "library"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeVideos holds the string denoting the videos edge name in mutations.
	EdgeVideos = "videos"
	// Table holds the table name of the library in the database.
	Table = "libraries"
	// VideosTable is the table that holds the videos relation/edge.
	VideosTable = "videos"
	// VideosInverseTable is the table name for the Video entity.
	// It exists in this package in order to avoid circular dependency with the "video" package.
	VideosInverseTable = "videos"
	// VideosColumn is the table column denoting the videos relation/edge.
	VideosColumn = "library_videos"
)

// Columns holds all SQL columns for library fields.
var Columns = []string{
	FieldID,
	FieldPath,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/altierawr/vstreamer/ent/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the Library queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPath orders the results by the path field.
func ByPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPath, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByVideosCount orders the results by videos count.
func ByVideosCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVideosStep(), opts...)
	}
}

// ByVideos orders the results by videos terms.
func ByVideos(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVideosStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newVideosStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VideosInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, VideosTable, VideosColumn),
	)
}