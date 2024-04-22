// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/altierawr/vstreamer/ent/playsessionmedia"
	"github.com/altierawr/vstreamer/ent/videocodec"
)

// VideoCodecCreate is the builder for creating a VideoCodec entity.
type VideoCodecCreate struct {
	config
	mutation *VideoCodecMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (vcc *VideoCodecCreate) SetName(s string) *VideoCodecCreate {
	vcc.mutation.SetName(s)
	return vcc
}

// SetMime sets the "mime" field.
func (vcc *VideoCodecCreate) SetMime(s string) *VideoCodecCreate {
	vcc.mutation.SetMime(s)
	return vcc
}

// SetDynamicRange sets the "dynamic_range" field.
func (vcc *VideoCodecCreate) SetDynamicRange(vr videocodec.DynamicRange) *VideoCodecCreate {
	vcc.mutation.SetDynamicRange(vr)
	return vcc
}

// SetMediaID sets the "media" edge to the PlaySessionMedia entity by ID.
func (vcc *VideoCodecCreate) SetMediaID(id int) *VideoCodecCreate {
	vcc.mutation.SetMediaID(id)
	return vcc
}

// SetNillableMediaID sets the "media" edge to the PlaySessionMedia entity by ID if the given value is not nil.
func (vcc *VideoCodecCreate) SetNillableMediaID(id *int) *VideoCodecCreate {
	if id != nil {
		vcc = vcc.SetMediaID(*id)
	}
	return vcc
}

// SetMedia sets the "media" edge to the PlaySessionMedia entity.
func (vcc *VideoCodecCreate) SetMedia(p *PlaySessionMedia) *VideoCodecCreate {
	return vcc.SetMediaID(p.ID)
}

// Mutation returns the VideoCodecMutation object of the builder.
func (vcc *VideoCodecCreate) Mutation() *VideoCodecMutation {
	return vcc.mutation
}

// Save creates the VideoCodec in the database.
func (vcc *VideoCodecCreate) Save(ctx context.Context) (*VideoCodec, error) {
	return withHooks(ctx, vcc.sqlSave, vcc.mutation, vcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vcc *VideoCodecCreate) SaveX(ctx context.Context) *VideoCodec {
	v, err := vcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcc *VideoCodecCreate) Exec(ctx context.Context) error {
	_, err := vcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcc *VideoCodecCreate) ExecX(ctx context.Context) {
	if err := vcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vcc *VideoCodecCreate) check() error {
	if _, ok := vcc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "VideoCodec.name"`)}
	}
	if _, ok := vcc.mutation.Mime(); !ok {
		return &ValidationError{Name: "mime", err: errors.New(`ent: missing required field "VideoCodec.mime"`)}
	}
	if _, ok := vcc.mutation.DynamicRange(); !ok {
		return &ValidationError{Name: "dynamic_range", err: errors.New(`ent: missing required field "VideoCodec.dynamic_range"`)}
	}
	if v, ok := vcc.mutation.DynamicRange(); ok {
		if err := videocodec.DynamicRangeValidator(v); err != nil {
			return &ValidationError{Name: "dynamic_range", err: fmt.Errorf(`ent: validator failed for field "VideoCodec.dynamic_range": %w`, err)}
		}
	}
	return nil
}

func (vcc *VideoCodecCreate) sqlSave(ctx context.Context) (*VideoCodec, error) {
	if err := vcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	vcc.mutation.id = &_node.ID
	vcc.mutation.done = true
	return _node, nil
}

func (vcc *VideoCodecCreate) createSpec() (*VideoCodec, *sqlgraph.CreateSpec) {
	var (
		_node = &VideoCodec{config: vcc.config}
		_spec = sqlgraph.NewCreateSpec(videocodec.Table, sqlgraph.NewFieldSpec(videocodec.FieldID, field.TypeInt))
	)
	if value, ok := vcc.mutation.Name(); ok {
		_spec.SetField(videocodec.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := vcc.mutation.Mime(); ok {
		_spec.SetField(videocodec.FieldMime, field.TypeString, value)
		_node.Mime = value
	}
	if value, ok := vcc.mutation.DynamicRange(); ok {
		_spec.SetField(videocodec.FieldDynamicRange, field.TypeEnum, value)
		_node.DynamicRange = value
	}
	if nodes := vcc.mutation.MediaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   videocodec.MediaTable,
			Columns: []string{videocodec.MediaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playsessionmedia.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.play_session_media_video_codecs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VideoCodecCreateBulk is the builder for creating many VideoCodec entities in bulk.
type VideoCodecCreateBulk struct {
	config
	err      error
	builders []*VideoCodecCreate
}

// Save creates the VideoCodec entities in the database.
func (vccb *VideoCodecCreateBulk) Save(ctx context.Context) ([]*VideoCodec, error) {
	if vccb.err != nil {
		return nil, vccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(vccb.builders))
	nodes := make([]*VideoCodec, len(vccb.builders))
	mutators := make([]Mutator, len(vccb.builders))
	for i := range vccb.builders {
		func(i int, root context.Context) {
			builder := vccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VideoCodecMutation)
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
					_, err = mutators[i+1].Mutate(root, vccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vccb *VideoCodecCreateBulk) SaveX(ctx context.Context) []*VideoCodec {
	v, err := vccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vccb *VideoCodecCreateBulk) Exec(ctx context.Context) error {
	_, err := vccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vccb *VideoCodecCreateBulk) ExecX(ctx context.Context) {
	if err := vccb.Exec(ctx); err != nil {
		panic(err)
	}
}
