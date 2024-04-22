// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/altierawr/vstreamer/ent/playsessionmedia"
	"github.com/altierawr/vstreamer/ent/videocodec"
)

// VideoCodec is the model entity for the VideoCodec schema.
type VideoCodec struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Mime holds the value of the "mime" field.
	Mime string `json:"mime,omitempty"`
	// DynamicRange holds the value of the "dynamic_range" field.
	DynamicRange videocodec.DynamicRange `json:"dynamic_range,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VideoCodecQuery when eager-loading is set.
	Edges                           VideoCodecEdges `json:"edges"`
	play_session_media_video_codecs *int
	selectValues                    sql.SelectValues
}

// VideoCodecEdges holds the relations/edges for other nodes in the graph.
type VideoCodecEdges struct {
	// Media holds the value of the media edge.
	Media *PlaySessionMedia `json:"media,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// MediaOrErr returns the Media value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VideoCodecEdges) MediaOrErr() (*PlaySessionMedia, error) {
	if e.Media != nil {
		return e.Media, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: playsessionmedia.Label}
	}
	return nil, &NotLoadedError{edge: "media"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VideoCodec) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case videocodec.FieldID:
			values[i] = new(sql.NullInt64)
		case videocodec.FieldName, videocodec.FieldMime, videocodec.FieldDynamicRange:
			values[i] = new(sql.NullString)
		case videocodec.ForeignKeys[0]: // play_session_media_video_codecs
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VideoCodec fields.
func (vc *VideoCodec) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case videocodec.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			vc.ID = int(value.Int64)
		case videocodec.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				vc.Name = value.String
			}
		case videocodec.FieldMime:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mime", values[i])
			} else if value.Valid {
				vc.Mime = value.String
			}
		case videocodec.FieldDynamicRange:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dynamic_range", values[i])
			} else if value.Valid {
				vc.DynamicRange = videocodec.DynamicRange(value.String)
			}
		case videocodec.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field play_session_media_video_codecs", value)
			} else if value.Valid {
				vc.play_session_media_video_codecs = new(int)
				*vc.play_session_media_video_codecs = int(value.Int64)
			}
		default:
			vc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the VideoCodec.
// This includes values selected through modifiers, order, etc.
func (vc *VideoCodec) Value(name string) (ent.Value, error) {
	return vc.selectValues.Get(name)
}

// QueryMedia queries the "media" edge of the VideoCodec entity.
func (vc *VideoCodec) QueryMedia() *PlaySessionMediaQuery {
	return NewVideoCodecClient(vc.config).QueryMedia(vc)
}

// Update returns a builder for updating this VideoCodec.
// Note that you need to call VideoCodec.Unwrap() before calling this method if this VideoCodec
// was returned from a transaction, and the transaction was committed or rolled back.
func (vc *VideoCodec) Update() *VideoCodecUpdateOne {
	return NewVideoCodecClient(vc.config).UpdateOne(vc)
}

// Unwrap unwraps the VideoCodec entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vc *VideoCodec) Unwrap() *VideoCodec {
	_tx, ok := vc.config.driver.(*txDriver)
	if !ok {
		panic("ent: VideoCodec is not a transactional entity")
	}
	vc.config.driver = _tx.drv
	return vc
}

// String implements the fmt.Stringer.
func (vc *VideoCodec) String() string {
	var builder strings.Builder
	builder.WriteString("VideoCodec(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vc.ID))
	builder.WriteString("name=")
	builder.WriteString(vc.Name)
	builder.WriteString(", ")
	builder.WriteString("mime=")
	builder.WriteString(vc.Mime)
	builder.WriteString(", ")
	builder.WriteString("dynamic_range=")
	builder.WriteString(fmt.Sprintf("%v", vc.DynamicRange))
	builder.WriteByte(')')
	return builder.String()
}

// VideoCodecs is a parsable slice of VideoCodec.
type VideoCodecs []*VideoCodec
