// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/altierawr/vstreamer/ent/playbackclient"
	"github.com/altierawr/vstreamer/ent/playsession"
	"github.com/altierawr/vstreamer/ent/playsessionmedia"
)

// PlaySessionCreate is the builder for creating a PlaySession entity.
type PlaySessionCreate struct {
	config
	mutation *PlaySessionMutation
	hooks    []Hook
}

// SetCurrentTime sets the "current_time" field.
func (psc *PlaySessionCreate) SetCurrentTime(i int) *PlaySessionCreate {
	psc.mutation.SetCurrentTime(i)
	return psc
}

// SetNillableCurrentTime sets the "current_time" field if the given value is not nil.
func (psc *PlaySessionCreate) SetNillableCurrentTime(i *int) *PlaySessionCreate {
	if i != nil {
		psc.SetCurrentTime(*i)
	}
	return psc
}

// SetState sets the "state" field.
func (psc *PlaySessionCreate) SetState(pl playsession.State) *PlaySessionCreate {
	psc.mutation.SetState(pl)
	return psc
}

// SetNillableState sets the "state" field if the given value is not nil.
func (psc *PlaySessionCreate) SetNillableState(pl *playsession.State) *PlaySessionCreate {
	if pl != nil {
		psc.SetState(*pl)
	}
	return psc
}

// AddClientIDs adds the "clients" edge to the PlaybackClient entity by IDs.
func (psc *PlaySessionCreate) AddClientIDs(ids ...int) *PlaySessionCreate {
	psc.mutation.AddClientIDs(ids...)
	return psc
}

// AddClients adds the "clients" edges to the PlaybackClient entity.
func (psc *PlaySessionCreate) AddClients(p ...*PlaybackClient) *PlaySessionCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psc.AddClientIDs(ids...)
}

// SetMediaID sets the "media" edge to the PlaySessionMedia entity by ID.
func (psc *PlaySessionCreate) SetMediaID(id int) *PlaySessionCreate {
	psc.mutation.SetMediaID(id)
	return psc
}

// SetNillableMediaID sets the "media" edge to the PlaySessionMedia entity by ID if the given value is not nil.
func (psc *PlaySessionCreate) SetNillableMediaID(id *int) *PlaySessionCreate {
	if id != nil {
		psc = psc.SetMediaID(*id)
	}
	return psc
}

// SetMedia sets the "media" edge to the PlaySessionMedia entity.
func (psc *PlaySessionCreate) SetMedia(p *PlaySessionMedia) *PlaySessionCreate {
	return psc.SetMediaID(p.ID)
}

// Mutation returns the PlaySessionMutation object of the builder.
func (psc *PlaySessionCreate) Mutation() *PlaySessionMutation {
	return psc.mutation
}

// Save creates the PlaySession in the database.
func (psc *PlaySessionCreate) Save(ctx context.Context) (*PlaySession, error) {
	psc.defaults()
	return withHooks(ctx, psc.sqlSave, psc.mutation, psc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (psc *PlaySessionCreate) SaveX(ctx context.Context) *PlaySession {
	v, err := psc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (psc *PlaySessionCreate) Exec(ctx context.Context) error {
	_, err := psc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psc *PlaySessionCreate) ExecX(ctx context.Context) {
	if err := psc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psc *PlaySessionCreate) defaults() {
	if _, ok := psc.mutation.State(); !ok {
		v := playsession.DefaultState
		psc.mutation.SetState(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psc *PlaySessionCreate) check() error {
	if _, ok := psc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "PlaySession.state"`)}
	}
	if v, ok := psc.mutation.State(); ok {
		if err := playsession.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "PlaySession.state": %w`, err)}
		}
	}
	return nil
}

func (psc *PlaySessionCreate) sqlSave(ctx context.Context) (*PlaySession, error) {
	if err := psc.check(); err != nil {
		return nil, err
	}
	_node, _spec := psc.createSpec()
	if err := sqlgraph.CreateNode(ctx, psc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	psc.mutation.id = &_node.ID
	psc.mutation.done = true
	return _node, nil
}

func (psc *PlaySessionCreate) createSpec() (*PlaySession, *sqlgraph.CreateSpec) {
	var (
		_node = &PlaySession{config: psc.config}
		_spec = sqlgraph.NewCreateSpec(playsession.Table, sqlgraph.NewFieldSpec(playsession.FieldID, field.TypeInt))
	)
	if value, ok := psc.mutation.CurrentTime(); ok {
		_spec.SetField(playsession.FieldCurrentTime, field.TypeInt, value)
		_node.CurrentTime = value
	}
	if value, ok := psc.mutation.State(); ok {
		_spec.SetField(playsession.FieldState, field.TypeEnum, value)
		_node.State = value
	}
	if nodes := psc.mutation.ClientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   playsession.ClientsTable,
			Columns: []string{playsession.ClientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playbackclient.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psc.mutation.MediaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   playsession.MediaTable,
			Columns: []string{playsession.MediaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playsessionmedia.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PlaySessionCreateBulk is the builder for creating many PlaySession entities in bulk.
type PlaySessionCreateBulk struct {
	config
	err      error
	builders []*PlaySessionCreate
}

// Save creates the PlaySession entities in the database.
func (pscb *PlaySessionCreateBulk) Save(ctx context.Context) ([]*PlaySession, error) {
	if pscb.err != nil {
		return nil, pscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pscb.builders))
	nodes := make([]*PlaySession, len(pscb.builders))
	mutators := make([]Mutator, len(pscb.builders))
	for i := range pscb.builders {
		func(i int, root context.Context) {
			builder := pscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlaySessionMutation)
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
					_, err = mutators[i+1].Mutate(root, pscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pscb *PlaySessionCreateBulk) SaveX(ctx context.Context) []*PlaySession {
	v, err := pscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pscb *PlaySessionCreateBulk) Exec(ctx context.Context) error {
	_, err := pscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pscb *PlaySessionCreateBulk) ExecX(ctx context.Context) {
	if err := pscb.Exec(ctx); err != nil {
		panic(err)
	}
}
