// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/altierawr/vstreamer/ent/audiocodec"
	"github.com/altierawr/vstreamer/ent/audiotrack"
	"github.com/altierawr/vstreamer/ent/playsessionmedia"
	"github.com/altierawr/vstreamer/ent/stream"
)

// AudioCodecCreate is the builder for creating a AudioCodec entity.
type AudioCodecCreate struct {
	config
	mutation *AudioCodecMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (acc *AudioCodecCreate) SetName(s string) *AudioCodecCreate {
	acc.mutation.SetName(s)
	return acc
}

// SetMime sets the "mime" field.
func (acc *AudioCodecCreate) SetMime(s string) *AudioCodecCreate {
	acc.mutation.SetMime(s)
	return acc
}

// AddStreamIDs adds the "streams" edge to the Stream entity by IDs.
func (acc *AudioCodecCreate) AddStreamIDs(ids ...int) *AudioCodecCreate {
	acc.mutation.AddStreamIDs(ids...)
	return acc
}

// AddStreams adds the "streams" edges to the Stream entity.
func (acc *AudioCodecCreate) AddStreams(s ...*Stream) *AudioCodecCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return acc.AddStreamIDs(ids...)
}

// AddAudioTrackIDs adds the "audio_tracks" edge to the AudioTrack entity by IDs.
func (acc *AudioCodecCreate) AddAudioTrackIDs(ids ...int) *AudioCodecCreate {
	acc.mutation.AddAudioTrackIDs(ids...)
	return acc
}

// AddAudioTracks adds the "audio_tracks" edges to the AudioTrack entity.
func (acc *AudioCodecCreate) AddAudioTracks(a ...*AudioTrack) *AudioCodecCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return acc.AddAudioTrackIDs(ids...)
}

// SetMediaID sets the "media" edge to the PlaySessionMedia entity by ID.
func (acc *AudioCodecCreate) SetMediaID(id int) *AudioCodecCreate {
	acc.mutation.SetMediaID(id)
	return acc
}

// SetNillableMediaID sets the "media" edge to the PlaySessionMedia entity by ID if the given value is not nil.
func (acc *AudioCodecCreate) SetNillableMediaID(id *int) *AudioCodecCreate {
	if id != nil {
		acc = acc.SetMediaID(*id)
	}
	return acc
}

// SetMedia sets the "media" edge to the PlaySessionMedia entity.
func (acc *AudioCodecCreate) SetMedia(p *PlaySessionMedia) *AudioCodecCreate {
	return acc.SetMediaID(p.ID)
}

// Mutation returns the AudioCodecMutation object of the builder.
func (acc *AudioCodecCreate) Mutation() *AudioCodecMutation {
	return acc.mutation
}

// Save creates the AudioCodec in the database.
func (acc *AudioCodecCreate) Save(ctx context.Context) (*AudioCodec, error) {
	return withHooks(ctx, acc.sqlSave, acc.mutation, acc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (acc *AudioCodecCreate) SaveX(ctx context.Context) *AudioCodec {
	v, err := acc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acc *AudioCodecCreate) Exec(ctx context.Context) error {
	_, err := acc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acc *AudioCodecCreate) ExecX(ctx context.Context) {
	if err := acc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acc *AudioCodecCreate) check() error {
	if _, ok := acc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "AudioCodec.name"`)}
	}
	if _, ok := acc.mutation.Mime(); !ok {
		return &ValidationError{Name: "mime", err: errors.New(`ent: missing required field "AudioCodec.mime"`)}
	}
	return nil
}

func (acc *AudioCodecCreate) sqlSave(ctx context.Context) (*AudioCodec, error) {
	if err := acc.check(); err != nil {
		return nil, err
	}
	_node, _spec := acc.createSpec()
	if err := sqlgraph.CreateNode(ctx, acc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	acc.mutation.id = &_node.ID
	acc.mutation.done = true
	return _node, nil
}

func (acc *AudioCodecCreate) createSpec() (*AudioCodec, *sqlgraph.CreateSpec) {
	var (
		_node = &AudioCodec{config: acc.config}
		_spec = sqlgraph.NewCreateSpec(audiocodec.Table, sqlgraph.NewFieldSpec(audiocodec.FieldID, field.TypeInt))
	)
	if value, ok := acc.mutation.Name(); ok {
		_spec.SetField(audiocodec.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := acc.mutation.Mime(); ok {
		_spec.SetField(audiocodec.FieldMime, field.TypeString, value)
		_node.Mime = value
	}
	if nodes := acc.mutation.StreamsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := acc.mutation.AudioTracksIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := acc.mutation.MediaIDs(); len(nodes) > 0 {
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
		_node.play_session_media_audio_codecs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AudioCodecCreateBulk is the builder for creating many AudioCodec entities in bulk.
type AudioCodecCreateBulk struct {
	config
	err      error
	builders []*AudioCodecCreate
}

// Save creates the AudioCodec entities in the database.
func (accb *AudioCodecCreateBulk) Save(ctx context.Context) ([]*AudioCodec, error) {
	if accb.err != nil {
		return nil, accb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(accb.builders))
	nodes := make([]*AudioCodec, len(accb.builders))
	mutators := make([]Mutator, len(accb.builders))
	for i := range accb.builders {
		func(i int, root context.Context) {
			builder := accb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AudioCodecMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, accb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, accb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, accb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (accb *AudioCodecCreateBulk) SaveX(ctx context.Context) []*AudioCodec {
	v, err := accb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (accb *AudioCodecCreateBulk) Exec(ctx context.Context) error {
	_, err := accb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (accb *AudioCodecCreateBulk) ExecX(ctx context.Context) {
	if err := accb.Exec(ctx); err != nil {
		panic(err)
	}
}
