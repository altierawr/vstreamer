// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/altierawr/vstreamer/ent/audiocodec"
	"github.com/altierawr/vstreamer/ent/audiotrack"
	"github.com/altierawr/vstreamer/ent/playsessionmedia"
	"github.com/altierawr/vstreamer/ent/predicate"
	"github.com/altierawr/vstreamer/ent/stream"
)

// AudioCodecUpdate is the builder for updating AudioCodec entities.
type AudioCodecUpdate struct {
	config
	hooks    []Hook
	mutation *AudioCodecMutation
}

// Where appends a list predicates to the AudioCodecUpdate builder.
func (acu *AudioCodecUpdate) Where(ps ...predicate.AudioCodec) *AudioCodecUpdate {
	acu.mutation.Where(ps...)
	return acu
}

// SetName sets the "name" field.
func (acu *AudioCodecUpdate) SetName(s string) *AudioCodecUpdate {
	acu.mutation.SetName(s)
	return acu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (acu *AudioCodecUpdate) SetNillableName(s *string) *AudioCodecUpdate {
	if s != nil {
		acu.SetName(*s)
	}
	return acu
}

// SetMime sets the "mime" field.
func (acu *AudioCodecUpdate) SetMime(s string) *AudioCodecUpdate {
	acu.mutation.SetMime(s)
	return acu
}

// SetNillableMime sets the "mime" field if the given value is not nil.
func (acu *AudioCodecUpdate) SetNillableMime(s *string) *AudioCodecUpdate {
	if s != nil {
		acu.SetMime(*s)
	}
	return acu
}

// AddStreamIDs adds the "streams" edge to the Stream entity by IDs.
func (acu *AudioCodecUpdate) AddStreamIDs(ids ...int) *AudioCodecUpdate {
	acu.mutation.AddStreamIDs(ids...)
	return acu
}

// AddStreams adds the "streams" edges to the Stream entity.
func (acu *AudioCodecUpdate) AddStreams(s ...*Stream) *AudioCodecUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return acu.AddStreamIDs(ids...)
}

// AddAudioTrackIDs adds the "audio_tracks" edge to the AudioTrack entity by IDs.
func (acu *AudioCodecUpdate) AddAudioTrackIDs(ids ...int) *AudioCodecUpdate {
	acu.mutation.AddAudioTrackIDs(ids...)
	return acu
}

// AddAudioTracks adds the "audio_tracks" edges to the AudioTrack entity.
func (acu *AudioCodecUpdate) AddAudioTracks(a ...*AudioTrack) *AudioCodecUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return acu.AddAudioTrackIDs(ids...)
}

// SetMediaID sets the "media" edge to the PlaySessionMedia entity by ID.
func (acu *AudioCodecUpdate) SetMediaID(id int) *AudioCodecUpdate {
	acu.mutation.SetMediaID(id)
	return acu
}

// SetNillableMediaID sets the "media" edge to the PlaySessionMedia entity by ID if the given value is not nil.
func (acu *AudioCodecUpdate) SetNillableMediaID(id *int) *AudioCodecUpdate {
	if id != nil {
		acu = acu.SetMediaID(*id)
	}
	return acu
}

// SetMedia sets the "media" edge to the PlaySessionMedia entity.
func (acu *AudioCodecUpdate) SetMedia(p *PlaySessionMedia) *AudioCodecUpdate {
	return acu.SetMediaID(p.ID)
}

// Mutation returns the AudioCodecMutation object of the builder.
func (acu *AudioCodecUpdate) Mutation() *AudioCodecMutation {
	return acu.mutation
}

// ClearStreams clears all "streams" edges to the Stream entity.
func (acu *AudioCodecUpdate) ClearStreams() *AudioCodecUpdate {
	acu.mutation.ClearStreams()
	return acu
}

// RemoveStreamIDs removes the "streams" edge to Stream entities by IDs.
func (acu *AudioCodecUpdate) RemoveStreamIDs(ids ...int) *AudioCodecUpdate {
	acu.mutation.RemoveStreamIDs(ids...)
	return acu
}

// RemoveStreams removes "streams" edges to Stream entities.
func (acu *AudioCodecUpdate) RemoveStreams(s ...*Stream) *AudioCodecUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return acu.RemoveStreamIDs(ids...)
}

// ClearAudioTracks clears all "audio_tracks" edges to the AudioTrack entity.
func (acu *AudioCodecUpdate) ClearAudioTracks() *AudioCodecUpdate {
	acu.mutation.ClearAudioTracks()
	return acu
}

// RemoveAudioTrackIDs removes the "audio_tracks" edge to AudioTrack entities by IDs.
func (acu *AudioCodecUpdate) RemoveAudioTrackIDs(ids ...int) *AudioCodecUpdate {
	acu.mutation.RemoveAudioTrackIDs(ids...)
	return acu
}

// RemoveAudioTracks removes "audio_tracks" edges to AudioTrack entities.
func (acu *AudioCodecUpdate) RemoveAudioTracks(a ...*AudioTrack) *AudioCodecUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return acu.RemoveAudioTrackIDs(ids...)
}

// ClearMedia clears the "media" edge to the PlaySessionMedia entity.
func (acu *AudioCodecUpdate) ClearMedia() *AudioCodecUpdate {
	acu.mutation.ClearMedia()
	return acu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (acu *AudioCodecUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, acu.sqlSave, acu.mutation, acu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (acu *AudioCodecUpdate) SaveX(ctx context.Context) int {
	affected, err := acu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (acu *AudioCodecUpdate) Exec(ctx context.Context) error {
	_, err := acu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acu *AudioCodecUpdate) ExecX(ctx context.Context) {
	if err := acu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (acu *AudioCodecUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(audiocodec.Table, audiocodec.Columns, sqlgraph.NewFieldSpec(audiocodec.FieldID, field.TypeInt))
	if ps := acu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acu.mutation.Name(); ok {
		_spec.SetField(audiocodec.FieldName, field.TypeString, value)
	}
	if value, ok := acu.mutation.Mime(); ok {
		_spec.SetField(audiocodec.FieldMime, field.TypeString, value)
	}
	if acu.mutation.StreamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   audiocodec.StreamsTable,
			Columns: []string{audiocodec.StreamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(stream.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := acu.mutation.RemovedStreamsIDs(); len(nodes) > 0 && !acu.mutation.StreamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   audiocodec.StreamsTable,
			Columns: []string{audiocodec.StreamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(stream.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := acu.mutation.StreamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   audiocodec.StreamsTable,
			Columns: []string{audiocodec.StreamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(stream.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if acu.mutation.AudioTracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   audiocodec.AudioTracksTable,
			Columns: audiocodec.AudioTracksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(audiotrack.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := acu.mutation.RemovedAudioTracksIDs(); len(nodes) > 0 && !acu.mutation.AudioTracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   audiocodec.AudioTracksTable,
			Columns: audiocodec.AudioTracksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(audiotrack.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := acu.mutation.AudioTracksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   audiocodec.AudioTracksTable,
			Columns: audiocodec.AudioTracksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(audiotrack.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if acu.mutation.MediaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   audiocodec.MediaTable,
			Columns: []string{audiocodec.MediaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playsessionmedia.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := acu.mutation.MediaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   audiocodec.MediaTable,
			Columns: []string{audiocodec.MediaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playsessionmedia.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, acu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{audiocodec.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	acu.mutation.done = true
	return n, nil
}

// AudioCodecUpdateOne is the builder for updating a single AudioCodec entity.
type AudioCodecUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AudioCodecMutation
}

// SetName sets the "name" field.
func (acuo *AudioCodecUpdateOne) SetName(s string) *AudioCodecUpdateOne {
	acuo.mutation.SetName(s)
	return acuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (acuo *AudioCodecUpdateOne) SetNillableName(s *string) *AudioCodecUpdateOne {
	if s != nil {
		acuo.SetName(*s)
	}
	return acuo
}

// SetMime sets the "mime" field.
func (acuo *AudioCodecUpdateOne) SetMime(s string) *AudioCodecUpdateOne {
	acuo.mutation.SetMime(s)
	return acuo
}

// SetNillableMime sets the "mime" field if the given value is not nil.
func (acuo *AudioCodecUpdateOne) SetNillableMime(s *string) *AudioCodecUpdateOne {
	if s != nil {
		acuo.SetMime(*s)
	}
	return acuo
}

// AddStreamIDs adds the "streams" edge to the Stream entity by IDs.
func (acuo *AudioCodecUpdateOne) AddStreamIDs(ids ...int) *AudioCodecUpdateOne {
	acuo.mutation.AddStreamIDs(ids...)
	return acuo
}

// AddStreams adds the "streams" edges to the Stream entity.
func (acuo *AudioCodecUpdateOne) AddStreams(s ...*Stream) *AudioCodecUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return acuo.AddStreamIDs(ids...)
}

// AddAudioTrackIDs adds the "audio_tracks" edge to the AudioTrack entity by IDs.
func (acuo *AudioCodecUpdateOne) AddAudioTrackIDs(ids ...int) *AudioCodecUpdateOne {
	acuo.mutation.AddAudioTrackIDs(ids...)
	return acuo
}

// AddAudioTracks adds the "audio_tracks" edges to the AudioTrack entity.
func (acuo *AudioCodecUpdateOne) AddAudioTracks(a ...*AudioTrack) *AudioCodecUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return acuo.AddAudioTrackIDs(ids...)
}

// SetMediaID sets the "media" edge to the PlaySessionMedia entity by ID.
func (acuo *AudioCodecUpdateOne) SetMediaID(id int) *AudioCodecUpdateOne {
	acuo.mutation.SetMediaID(id)
	return acuo
}

// SetNillableMediaID sets the "media" edge to the PlaySessionMedia entity by ID if the given value is not nil.
func (acuo *AudioCodecUpdateOne) SetNillableMediaID(id *int) *AudioCodecUpdateOne {
	if id != nil {
		acuo = acuo.SetMediaID(*id)
	}
	return acuo
}

// SetMedia sets the "media" edge to the PlaySessionMedia entity.
func (acuo *AudioCodecUpdateOne) SetMedia(p *PlaySessionMedia) *AudioCodecUpdateOne {
	return acuo.SetMediaID(p.ID)
}

// Mutation returns the AudioCodecMutation object of the builder.
func (acuo *AudioCodecUpdateOne) Mutation() *AudioCodecMutation {
	return acuo.mutation
}

// ClearStreams clears all "streams" edges to the Stream entity.
func (acuo *AudioCodecUpdateOne) ClearStreams() *AudioCodecUpdateOne {
	acuo.mutation.ClearStreams()
	return acuo
}

// RemoveStreamIDs removes the "streams" edge to Stream entities by IDs.
func (acuo *AudioCodecUpdateOne) RemoveStreamIDs(ids ...int) *AudioCodecUpdateOne {
	acuo.mutation.RemoveStreamIDs(ids...)
	return acuo
}

// RemoveStreams removes "streams" edges to Stream entities.
func (acuo *AudioCodecUpdateOne) RemoveStreams(s ...*Stream) *AudioCodecUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return acuo.RemoveStreamIDs(ids...)
}

// ClearAudioTracks clears all "audio_tracks" edges to the AudioTrack entity.
func (acuo *AudioCodecUpdateOne) ClearAudioTracks() *AudioCodecUpdateOne {
	acuo.mutation.ClearAudioTracks()
	return acuo
}

// RemoveAudioTrackIDs removes the "audio_tracks" edge to AudioTrack entities by IDs.
func (acuo *AudioCodecUpdateOne) RemoveAudioTrackIDs(ids ...int) *AudioCodecUpdateOne {
	acuo.mutation.RemoveAudioTrackIDs(ids...)
	return acuo
}

// RemoveAudioTracks removes "audio_tracks" edges to AudioTrack entities.
func (acuo *AudioCodecUpdateOne) RemoveAudioTracks(a ...*AudioTrack) *AudioCodecUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return acuo.RemoveAudioTrackIDs(ids...)
}

// ClearMedia clears the "media" edge to the PlaySessionMedia entity.
func (acuo *AudioCodecUpdateOne) ClearMedia() *AudioCodecUpdateOne {
	acuo.mutation.ClearMedia()
	return acuo
}

// Where appends a list predicates to the AudioCodecUpdate builder.
func (acuo *AudioCodecUpdateOne) Where(ps ...predicate.AudioCodec) *AudioCodecUpdateOne {
	acuo.mutation.Where(ps...)
	return acuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (acuo *AudioCodecUpdateOne) Select(field string, fields ...string) *AudioCodecUpdateOne {
	acuo.fields = append([]string{field}, fields...)
	return acuo
}

// Save executes the query and returns the updated AudioCodec entity.
func (acuo *AudioCodecUpdateOne) Save(ctx context.Context) (*AudioCodec, error) {
	return withHooks(ctx, acuo.sqlSave, acuo.mutation, acuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (acuo *AudioCodecUpdateOne) SaveX(ctx context.Context) *AudioCodec {
	node, err := acuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (acuo *AudioCodecUpdateOne) Exec(ctx context.Context) error {
	_, err := acuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acuo *AudioCodecUpdateOne) ExecX(ctx context.Context) {
	if err := acuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (acuo *AudioCodecUpdateOne) sqlSave(ctx context.Context) (_node *AudioCodec, err error) {
	_spec := sqlgraph.NewUpdateSpec(audiocodec.Table, audiocodec.Columns, sqlgraph.NewFieldSpec(audiocodec.FieldID, field.TypeInt))
	id, ok := acuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AudioCodec.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := acuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, audiocodec.FieldID)
		for _, f := range fields {
			if !audiocodec.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != audiocodec.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := acuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acuo.mutation.Name(); ok {
		_spec.SetField(audiocodec.FieldName, field.TypeString, value)
	}
	if value, ok := acuo.mutation.Mime(); ok {
		_spec.SetField(audiocodec.FieldMime, field.TypeString, value)
	}
	if acuo.mutation.StreamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   audiocodec.StreamsTable,
			Columns: []string{audiocodec.StreamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(stream.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := acuo.mutation.RemovedStreamsIDs(); len(nodes) > 0 && !acuo.mutation.StreamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   audiocodec.StreamsTable,
			Columns: []string{audiocodec.StreamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(stream.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := acuo.mutation.StreamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   audiocodec.StreamsTable,
			Columns: []string{audiocodec.StreamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(stream.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if acuo.mutation.AudioTracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   audiocodec.AudioTracksTable,
			Columns: audiocodec.AudioTracksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(audiotrack.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := acuo.mutation.RemovedAudioTracksIDs(); len(nodes) > 0 && !acuo.mutation.AudioTracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   audiocodec.AudioTracksTable,
			Columns: audiocodec.AudioTracksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(audiotrack.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := acuo.mutation.AudioTracksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   audiocodec.AudioTracksTable,
			Columns: audiocodec.AudioTracksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(audiotrack.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if acuo.mutation.MediaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   audiocodec.MediaTable,
			Columns: []string{audiocodec.MediaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playsessionmedia.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := acuo.mutation.MediaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   audiocodec.MediaTable,
			Columns: []string{audiocodec.MediaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playsessionmedia.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AudioCodec{config: acuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, acuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{audiocodec.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	acuo.mutation.done = true
	return _node, nil
}
