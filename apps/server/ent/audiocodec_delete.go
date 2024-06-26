// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/altierawr/vstreamer/ent/audiocodec"
	"github.com/altierawr/vstreamer/ent/predicate"
)

// AudioCodecDelete is the builder for deleting a AudioCodec entity.
type AudioCodecDelete struct {
	config
	hooks    []Hook
	mutation *AudioCodecMutation
}

// Where appends a list predicates to the AudioCodecDelete builder.
func (acd *AudioCodecDelete) Where(ps ...predicate.AudioCodec) *AudioCodecDelete {
	acd.mutation.Where(ps...)
	return acd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (acd *AudioCodecDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, acd.sqlExec, acd.mutation, acd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (acd *AudioCodecDelete) ExecX(ctx context.Context) int {
	n, err := acd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (acd *AudioCodecDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(audiocodec.Table, sqlgraph.NewFieldSpec(audiocodec.FieldID, field.TypeInt))
	if ps := acd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, acd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	acd.mutation.done = true
	return affected, err
}

// AudioCodecDeleteOne is the builder for deleting a single AudioCodec entity.
type AudioCodecDeleteOne struct {
	acd *AudioCodecDelete
}

// Where appends a list predicates to the AudioCodecDelete builder.
func (acdo *AudioCodecDeleteOne) Where(ps ...predicate.AudioCodec) *AudioCodecDeleteOne {
	acdo.acd.mutation.Where(ps...)
	return acdo
}

// Exec executes the deletion query.
func (acdo *AudioCodecDeleteOne) Exec(ctx context.Context) error {
	n, err := acdo.acd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{audiocodec.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (acdo *AudioCodecDeleteOne) ExecX(ctx context.Context) {
	if err := acdo.Exec(ctx); err != nil {
		panic(err)
	}
}
