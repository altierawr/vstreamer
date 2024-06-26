// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/altierawr/vstreamer/ent/library"
	"github.com/altierawr/vstreamer/ent/playsessionmedia"
	"github.com/altierawr/vstreamer/ent/predicate"
	"github.com/altierawr/vstreamer/ent/video"
)

// VideoUpdate is the builder for updating Video entities.
type VideoUpdate struct {
	config
	hooks    []Hook
	mutation *VideoMutation
}

// Where appends a list predicates to the VideoUpdate builder.
func (vu *VideoUpdate) Where(ps ...predicate.Video) *VideoUpdate {
	vu.mutation.Where(ps...)
	return vu
}

// SetPath sets the "path" field.
func (vu *VideoUpdate) SetPath(s string) *VideoUpdate {
	vu.mutation.SetPath(s)
	return vu
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (vu *VideoUpdate) SetNillablePath(s *string) *VideoUpdate {
	if s != nil {
		vu.SetPath(*s)
	}
	return vu
}

// SetCreatedAt sets the "created_at" field.
func (vu *VideoUpdate) SetCreatedAt(t time.Time) *VideoUpdate {
	vu.mutation.SetCreatedAt(t)
	return vu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vu *VideoUpdate) SetNillableCreatedAt(t *time.Time) *VideoUpdate {
	if t != nil {
		vu.SetCreatedAt(*t)
	}
	return vu
}

// AddPlaySessionMediaIDs adds the "play_session_medias" edge to the PlaySessionMedia entity by IDs.
func (vu *VideoUpdate) AddPlaySessionMediaIDs(ids ...int) *VideoUpdate {
	vu.mutation.AddPlaySessionMediaIDs(ids...)
	return vu
}

// AddPlaySessionMedias adds the "play_session_medias" edges to the PlaySessionMedia entity.
func (vu *VideoUpdate) AddPlaySessionMedias(p ...*PlaySessionMedia) *VideoUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return vu.AddPlaySessionMediaIDs(ids...)
}

// SetLibraryID sets the "library" edge to the Library entity by ID.
func (vu *VideoUpdate) SetLibraryID(id int) *VideoUpdate {
	vu.mutation.SetLibraryID(id)
	return vu
}

// SetNillableLibraryID sets the "library" edge to the Library entity by ID if the given value is not nil.
func (vu *VideoUpdate) SetNillableLibraryID(id *int) *VideoUpdate {
	if id != nil {
		vu = vu.SetLibraryID(*id)
	}
	return vu
}

// SetLibrary sets the "library" edge to the Library entity.
func (vu *VideoUpdate) SetLibrary(l *Library) *VideoUpdate {
	return vu.SetLibraryID(l.ID)
}

// Mutation returns the VideoMutation object of the builder.
func (vu *VideoUpdate) Mutation() *VideoMutation {
	return vu.mutation
}

// ClearPlaySessionMedias clears all "play_session_medias" edges to the PlaySessionMedia entity.
func (vu *VideoUpdate) ClearPlaySessionMedias() *VideoUpdate {
	vu.mutation.ClearPlaySessionMedias()
	return vu
}

// RemovePlaySessionMediaIDs removes the "play_session_medias" edge to PlaySessionMedia entities by IDs.
func (vu *VideoUpdate) RemovePlaySessionMediaIDs(ids ...int) *VideoUpdate {
	vu.mutation.RemovePlaySessionMediaIDs(ids...)
	return vu
}

// RemovePlaySessionMedias removes "play_session_medias" edges to PlaySessionMedia entities.
func (vu *VideoUpdate) RemovePlaySessionMedias(p ...*PlaySessionMedia) *VideoUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return vu.RemovePlaySessionMediaIDs(ids...)
}

// ClearLibrary clears the "library" edge to the Library entity.
func (vu *VideoUpdate) ClearLibrary() *VideoUpdate {
	vu.mutation.ClearLibrary()
	return vu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *VideoUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, vu.sqlSave, vu.mutation, vu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vu *VideoUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *VideoUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *VideoUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vu *VideoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(video.Table, video.Columns, sqlgraph.NewFieldSpec(video.FieldID, field.TypeInt))
	if ps := vu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vu.mutation.Path(); ok {
		_spec.SetField(video.FieldPath, field.TypeString, value)
	}
	if value, ok := vu.mutation.CreatedAt(); ok {
		_spec.SetField(video.FieldCreatedAt, field.TypeTime, value)
	}
	if vu.mutation.PlaySessionMediasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.PlaySessionMediasTable,
			Columns: []string{video.PlaySessionMediasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playsessionmedia.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.RemovedPlaySessionMediasIDs(); len(nodes) > 0 && !vu.mutation.PlaySessionMediasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.PlaySessionMediasTable,
			Columns: []string{video.PlaySessionMediasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playsessionmedia.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.PlaySessionMediasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.PlaySessionMediasTable,
			Columns: []string{video.PlaySessionMediasColumn},
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
	if vu.mutation.LibraryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   video.LibraryTable,
			Columns: []string{video.LibraryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(library.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.LibraryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   video.LibraryTable,
			Columns: []string{video.LibraryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(library.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{video.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vu.mutation.done = true
	return n, nil
}

// VideoUpdateOne is the builder for updating a single Video entity.
type VideoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VideoMutation
}

// SetPath sets the "path" field.
func (vuo *VideoUpdateOne) SetPath(s string) *VideoUpdateOne {
	vuo.mutation.SetPath(s)
	return vuo
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (vuo *VideoUpdateOne) SetNillablePath(s *string) *VideoUpdateOne {
	if s != nil {
		vuo.SetPath(*s)
	}
	return vuo
}

// SetCreatedAt sets the "created_at" field.
func (vuo *VideoUpdateOne) SetCreatedAt(t time.Time) *VideoUpdateOne {
	vuo.mutation.SetCreatedAt(t)
	return vuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vuo *VideoUpdateOne) SetNillableCreatedAt(t *time.Time) *VideoUpdateOne {
	if t != nil {
		vuo.SetCreatedAt(*t)
	}
	return vuo
}

// AddPlaySessionMediaIDs adds the "play_session_medias" edge to the PlaySessionMedia entity by IDs.
func (vuo *VideoUpdateOne) AddPlaySessionMediaIDs(ids ...int) *VideoUpdateOne {
	vuo.mutation.AddPlaySessionMediaIDs(ids...)
	return vuo
}

// AddPlaySessionMedias adds the "play_session_medias" edges to the PlaySessionMedia entity.
func (vuo *VideoUpdateOne) AddPlaySessionMedias(p ...*PlaySessionMedia) *VideoUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return vuo.AddPlaySessionMediaIDs(ids...)
}

// SetLibraryID sets the "library" edge to the Library entity by ID.
func (vuo *VideoUpdateOne) SetLibraryID(id int) *VideoUpdateOne {
	vuo.mutation.SetLibraryID(id)
	return vuo
}

// SetNillableLibraryID sets the "library" edge to the Library entity by ID if the given value is not nil.
func (vuo *VideoUpdateOne) SetNillableLibraryID(id *int) *VideoUpdateOne {
	if id != nil {
		vuo = vuo.SetLibraryID(*id)
	}
	return vuo
}

// SetLibrary sets the "library" edge to the Library entity.
func (vuo *VideoUpdateOne) SetLibrary(l *Library) *VideoUpdateOne {
	return vuo.SetLibraryID(l.ID)
}

// Mutation returns the VideoMutation object of the builder.
func (vuo *VideoUpdateOne) Mutation() *VideoMutation {
	return vuo.mutation
}

// ClearPlaySessionMedias clears all "play_session_medias" edges to the PlaySessionMedia entity.
func (vuo *VideoUpdateOne) ClearPlaySessionMedias() *VideoUpdateOne {
	vuo.mutation.ClearPlaySessionMedias()
	return vuo
}

// RemovePlaySessionMediaIDs removes the "play_session_medias" edge to PlaySessionMedia entities by IDs.
func (vuo *VideoUpdateOne) RemovePlaySessionMediaIDs(ids ...int) *VideoUpdateOne {
	vuo.mutation.RemovePlaySessionMediaIDs(ids...)
	return vuo
}

// RemovePlaySessionMedias removes "play_session_medias" edges to PlaySessionMedia entities.
func (vuo *VideoUpdateOne) RemovePlaySessionMedias(p ...*PlaySessionMedia) *VideoUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return vuo.RemovePlaySessionMediaIDs(ids...)
}

// ClearLibrary clears the "library" edge to the Library entity.
func (vuo *VideoUpdateOne) ClearLibrary() *VideoUpdateOne {
	vuo.mutation.ClearLibrary()
	return vuo
}

// Where appends a list predicates to the VideoUpdate builder.
func (vuo *VideoUpdateOne) Where(ps ...predicate.Video) *VideoUpdateOne {
	vuo.mutation.Where(ps...)
	return vuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vuo *VideoUpdateOne) Select(field string, fields ...string) *VideoUpdateOne {
	vuo.fields = append([]string{field}, fields...)
	return vuo
}

// Save executes the query and returns the updated Video entity.
func (vuo *VideoUpdateOne) Save(ctx context.Context) (*Video, error) {
	return withHooks(ctx, vuo.sqlSave, vuo.mutation, vuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *VideoUpdateOne) SaveX(ctx context.Context) *Video {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *VideoUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VideoUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vuo *VideoUpdateOne) sqlSave(ctx context.Context) (_node *Video, err error) {
	_spec := sqlgraph.NewUpdateSpec(video.Table, video.Columns, sqlgraph.NewFieldSpec(video.FieldID, field.TypeInt))
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Video.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, video.FieldID)
		for _, f := range fields {
			if !video.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != video.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vuo.mutation.Path(); ok {
		_spec.SetField(video.FieldPath, field.TypeString, value)
	}
	if value, ok := vuo.mutation.CreatedAt(); ok {
		_spec.SetField(video.FieldCreatedAt, field.TypeTime, value)
	}
	if vuo.mutation.PlaySessionMediasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.PlaySessionMediasTable,
			Columns: []string{video.PlaySessionMediasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playsessionmedia.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.RemovedPlaySessionMediasIDs(); len(nodes) > 0 && !vuo.mutation.PlaySessionMediasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.PlaySessionMediasTable,
			Columns: []string{video.PlaySessionMediasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playsessionmedia.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.PlaySessionMediasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.PlaySessionMediasTable,
			Columns: []string{video.PlaySessionMediasColumn},
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
	if vuo.mutation.LibraryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   video.LibraryTable,
			Columns: []string{video.LibraryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(library.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.LibraryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   video.LibraryTable,
			Columns: []string{video.LibraryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(library.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Video{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{video.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vuo.mutation.done = true
	return _node, nil
}
