// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/altierawr/vstreamer/ent/stream"
)

// Stream is the model entity for the Stream schema.
type Stream struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Width holds the value of the "width" field.
	Width int `json:"width,omitempty"`
	// Height holds the value of the "height" field.
	Height int `json:"height,omitempty"`
	// Container holds the value of the "container" field.
	Container string `json:"container,omitempty"`
	// VideoCodec holds the value of the "video_codec" field.
	VideoCodec string `json:"video_codec,omitempty"`
	// AudioCodec holds the value of the "audio_codec" field.
	AudioCodec string `json:"audio_codec,omitempty"`
	// SegmentDuration holds the value of the "segment_duration" field.
	SegmentDuration int `json:"segment_duration,omitempty"`
	// Quality holds the value of the "quality" field.
	Quality stream.Quality `json:"quality,omitempty"`
	// Type holds the value of the "type" field.
	Type         stream.Type `json:"type,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Stream) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case stream.FieldID, stream.FieldWidth, stream.FieldHeight, stream.FieldSegmentDuration:
			values[i] = new(sql.NullInt64)
		case stream.FieldContainer, stream.FieldVideoCodec, stream.FieldAudioCodec, stream.FieldQuality, stream.FieldType:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Stream fields.
func (s *Stream) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case stream.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case stream.FieldWidth:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field width", values[i])
			} else if value.Valid {
				s.Width = int(value.Int64)
			}
		case stream.FieldHeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field height", values[i])
			} else if value.Valid {
				s.Height = int(value.Int64)
			}
		case stream.FieldContainer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field container", values[i])
			} else if value.Valid {
				s.Container = value.String
			}
		case stream.FieldVideoCodec:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field video_codec", values[i])
			} else if value.Valid {
				s.VideoCodec = value.String
			}
		case stream.FieldAudioCodec:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field audio_codec", values[i])
			} else if value.Valid {
				s.AudioCodec = value.String
			}
		case stream.FieldSegmentDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field segment_duration", values[i])
			} else if value.Valid {
				s.SegmentDuration = int(value.Int64)
			}
		case stream.FieldQuality:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field quality", values[i])
			} else if value.Valid {
				s.Quality = stream.Quality(value.String)
			}
		case stream.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				s.Type = stream.Type(value.String)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Stream.
// This includes values selected through modifiers, order, etc.
func (s *Stream) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this Stream.
// Note that you need to call Stream.Unwrap() before calling this method if this Stream
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Stream) Update() *StreamUpdateOne {
	return NewStreamClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Stream entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Stream) Unwrap() *Stream {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Stream is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Stream) String() string {
	var builder strings.Builder
	builder.WriteString("Stream(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("width=")
	builder.WriteString(fmt.Sprintf("%v", s.Width))
	builder.WriteString(", ")
	builder.WriteString("height=")
	builder.WriteString(fmt.Sprintf("%v", s.Height))
	builder.WriteString(", ")
	builder.WriteString("container=")
	builder.WriteString(s.Container)
	builder.WriteString(", ")
	builder.WriteString("video_codec=")
	builder.WriteString(s.VideoCodec)
	builder.WriteString(", ")
	builder.WriteString("audio_codec=")
	builder.WriteString(s.AudioCodec)
	builder.WriteString(", ")
	builder.WriteString("segment_duration=")
	builder.WriteString(fmt.Sprintf("%v", s.SegmentDuration))
	builder.WriteString(", ")
	builder.WriteString("quality=")
	builder.WriteString(fmt.Sprintf("%v", s.Quality))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", s.Type))
	builder.WriteByte(')')
	return builder.String()
}

// Streams is a parsable slice of Stream.
type Streams []*Stream
