// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/altierawr/vstreamer/ent/audiocodec"
	"github.com/altierawr/vstreamer/ent/playsessionmedia"
)

// AudioCodec is the model entity for the AudioCodec schema.
type AudioCodec struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Mime holds the value of the "mime" field.
	Mime string `json:"mime,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AudioCodecQuery when eager-loading is set.
	Edges                           AudioCodecEdges `json:"edges"`
	play_session_media_audio_codecs *int
	selectValues                    sql.SelectValues
}

// AudioCodecEdges holds the relations/edges for other nodes in the graph.
type AudioCodecEdges struct {
	// Streams holds the value of the streams edge.
	Streams []*Stream `json:"streams,omitempty"`
	// AudioTracks holds the value of the audio_tracks edge.
	AudioTracks []*AudioTrack `json:"audio_tracks,omitempty"`
	// Media holds the value of the media edge.
	Media *PlaySessionMedia `json:"media,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedStreams     map[string][]*Stream
	namedAudioTracks map[string][]*AudioTrack
}

// StreamsOrErr returns the Streams value or an error if the edge
// was not loaded in eager-loading.
func (e AudioCodecEdges) StreamsOrErr() ([]*Stream, error) {
	if e.loadedTypes[0] {
		return e.Streams, nil
	}
	return nil, &NotLoadedError{edge: "streams"}
}

// AudioTracksOrErr returns the AudioTracks value or an error if the edge
// was not loaded in eager-loading.
func (e AudioCodecEdges) AudioTracksOrErr() ([]*AudioTrack, error) {
	if e.loadedTypes[1] {
		return e.AudioTracks, nil
	}
	return nil, &NotLoadedError{edge: "audio_tracks"}
}

// MediaOrErr returns the Media value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AudioCodecEdges) MediaOrErr() (*PlaySessionMedia, error) {
	if e.Media != nil {
		return e.Media, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: playsessionmedia.Label}
	}
	return nil, &NotLoadedError{edge: "media"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AudioCodec) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case audiocodec.FieldID:
			values[i] = new(sql.NullInt64)
		case audiocodec.FieldName, audiocodec.FieldMime:
			values[i] = new(sql.NullString)
		case audiocodec.ForeignKeys[0]: // play_session_media_audio_codecs
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AudioCodec fields.
func (ac *AudioCodec) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case audiocodec.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ac.ID = int(value.Int64)
		case audiocodec.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ac.Name = value.String
			}
		case audiocodec.FieldMime:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mime", values[i])
			} else if value.Valid {
				ac.Mime = value.String
			}
		case audiocodec.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field play_session_media_audio_codecs", value)
			} else if value.Valid {
				ac.play_session_media_audio_codecs = new(int)
				*ac.play_session_media_audio_codecs = int(value.Int64)
			}
		default:
			ac.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AudioCodec.
// This includes values selected through modifiers, order, etc.
func (ac *AudioCodec) Value(name string) (ent.Value, error) {
	return ac.selectValues.Get(name)
}

// QueryStreams queries the "streams" edge of the AudioCodec entity.
func (ac *AudioCodec) QueryStreams() *StreamQuery {
	return NewAudioCodecClient(ac.config).QueryStreams(ac)
}

// QueryAudioTracks queries the "audio_tracks" edge of the AudioCodec entity.
func (ac *AudioCodec) QueryAudioTracks() *AudioTrackQuery {
	return NewAudioCodecClient(ac.config).QueryAudioTracks(ac)
}

// QueryMedia queries the "media" edge of the AudioCodec entity.
func (ac *AudioCodec) QueryMedia() *PlaySessionMediaQuery {
	return NewAudioCodecClient(ac.config).QueryMedia(ac)
}

// Update returns a builder for updating this AudioCodec.
// Note that you need to call AudioCodec.Unwrap() before calling this method if this AudioCodec
// was returned from a transaction, and the transaction was committed or rolled back.
func (ac *AudioCodec) Update() *AudioCodecUpdateOne {
	return NewAudioCodecClient(ac.config).UpdateOne(ac)
}

// Unwrap unwraps the AudioCodec entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ac *AudioCodec) Unwrap() *AudioCodec {
	_tx, ok := ac.config.driver.(*txDriver)
	if !ok {
		panic("ent: AudioCodec is not a transactional entity")
	}
	ac.config.driver = _tx.drv
	return ac
}

// String implements the fmt.Stringer.
func (ac *AudioCodec) String() string {
	var builder strings.Builder
	builder.WriteString("AudioCodec(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ac.ID))
	builder.WriteString("name=")
	builder.WriteString(ac.Name)
	builder.WriteString(", ")
	builder.WriteString("mime=")
	builder.WriteString(ac.Mime)
	builder.WriteByte(')')
	return builder.String()
}

// NamedStreams returns the Streams named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ac *AudioCodec) NamedStreams(name string) ([]*Stream, error) {
	if ac.Edges.namedStreams == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ac.Edges.namedStreams[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ac *AudioCodec) appendNamedStreams(name string, edges ...*Stream) {
	if ac.Edges.namedStreams == nil {
		ac.Edges.namedStreams = make(map[string][]*Stream)
	}
	if len(edges) == 0 {
		ac.Edges.namedStreams[name] = []*Stream{}
	} else {
		ac.Edges.namedStreams[name] = append(ac.Edges.namedStreams[name], edges...)
	}
}

// NamedAudioTracks returns the AudioTracks named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ac *AudioCodec) NamedAudioTracks(name string) ([]*AudioTrack, error) {
	if ac.Edges.namedAudioTracks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ac.Edges.namedAudioTracks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ac *AudioCodec) appendNamedAudioTracks(name string, edges ...*AudioTrack) {
	if ac.Edges.namedAudioTracks == nil {
		ac.Edges.namedAudioTracks = make(map[string][]*AudioTrack)
	}
	if len(edges) == 0 {
		ac.Edges.namedAudioTracks[name] = []*AudioTrack{}
	} else {
		ac.Edges.namedAudioTracks[name] = append(ac.Edges.namedAudioTracks[name], edges...)
	}
}

// AudioCodecs is a parsable slice of AudioCodec.
type AudioCodecs []*AudioCodec
