// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/altierawr/vstreamer/ent/playsession"
	"github.com/altierawr/vstreamer/ent/predicate"
	"github.com/altierawr/vstreamer/ent/video"
)

// PlaySessionQuery is the builder for querying PlaySession entities.
type PlaySessionQuery struct {
	config
	ctx        *QueryContext
	order      []playsession.OrderOption
	inters     []Interceptor
	predicates []predicate.PlaySession
	withVideo  *VideoQuery
	withFKs    bool
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*PlaySession) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PlaySessionQuery builder.
func (psq *PlaySessionQuery) Where(ps ...predicate.PlaySession) *PlaySessionQuery {
	psq.predicates = append(psq.predicates, ps...)
	return psq
}

// Limit the number of records to be returned by this query.
func (psq *PlaySessionQuery) Limit(limit int) *PlaySessionQuery {
	psq.ctx.Limit = &limit
	return psq
}

// Offset to start from.
func (psq *PlaySessionQuery) Offset(offset int) *PlaySessionQuery {
	psq.ctx.Offset = &offset
	return psq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (psq *PlaySessionQuery) Unique(unique bool) *PlaySessionQuery {
	psq.ctx.Unique = &unique
	return psq
}

// Order specifies how the records should be ordered.
func (psq *PlaySessionQuery) Order(o ...playsession.OrderOption) *PlaySessionQuery {
	psq.order = append(psq.order, o...)
	return psq
}

// QueryVideo chains the current query on the "video" edge.
func (psq *PlaySessionQuery) QueryVideo() *VideoQuery {
	query := (&VideoClient{config: psq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := psq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := psq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(playsession.Table, playsession.FieldID, selector),
			sqlgraph.To(video.Table, video.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, playsession.VideoTable, playsession.VideoColumn),
		)
		fromU = sqlgraph.SetNeighbors(psq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PlaySession entity from the query.
// Returns a *NotFoundError when no PlaySession was found.
func (psq *PlaySessionQuery) First(ctx context.Context) (*PlaySession, error) {
	nodes, err := psq.Limit(1).All(setContextOp(ctx, psq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{playsession.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (psq *PlaySessionQuery) FirstX(ctx context.Context) *PlaySession {
	node, err := psq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PlaySession ID from the query.
// Returns a *NotFoundError when no PlaySession ID was found.
func (psq *PlaySessionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = psq.Limit(1).IDs(setContextOp(ctx, psq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{playsession.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (psq *PlaySessionQuery) FirstIDX(ctx context.Context) int {
	id, err := psq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PlaySession entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PlaySession entity is found.
// Returns a *NotFoundError when no PlaySession entities are found.
func (psq *PlaySessionQuery) Only(ctx context.Context) (*PlaySession, error) {
	nodes, err := psq.Limit(2).All(setContextOp(ctx, psq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{playsession.Label}
	default:
		return nil, &NotSingularError{playsession.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (psq *PlaySessionQuery) OnlyX(ctx context.Context) *PlaySession {
	node, err := psq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PlaySession ID in the query.
// Returns a *NotSingularError when more than one PlaySession ID is found.
// Returns a *NotFoundError when no entities are found.
func (psq *PlaySessionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = psq.Limit(2).IDs(setContextOp(ctx, psq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{playsession.Label}
	default:
		err = &NotSingularError{playsession.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (psq *PlaySessionQuery) OnlyIDX(ctx context.Context) int {
	id, err := psq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PlaySessions.
func (psq *PlaySessionQuery) All(ctx context.Context) ([]*PlaySession, error) {
	ctx = setContextOp(ctx, psq.ctx, "All")
	if err := psq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PlaySession, *PlaySessionQuery]()
	return withInterceptors[[]*PlaySession](ctx, psq, qr, psq.inters)
}

// AllX is like All, but panics if an error occurs.
func (psq *PlaySessionQuery) AllX(ctx context.Context) []*PlaySession {
	nodes, err := psq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PlaySession IDs.
func (psq *PlaySessionQuery) IDs(ctx context.Context) (ids []int, err error) {
	if psq.ctx.Unique == nil && psq.path != nil {
		psq.Unique(true)
	}
	ctx = setContextOp(ctx, psq.ctx, "IDs")
	if err = psq.Select(playsession.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (psq *PlaySessionQuery) IDsX(ctx context.Context) []int {
	ids, err := psq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (psq *PlaySessionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, psq.ctx, "Count")
	if err := psq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, psq, querierCount[*PlaySessionQuery](), psq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (psq *PlaySessionQuery) CountX(ctx context.Context) int {
	count, err := psq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (psq *PlaySessionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, psq.ctx, "Exist")
	switch _, err := psq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (psq *PlaySessionQuery) ExistX(ctx context.Context) bool {
	exist, err := psq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PlaySessionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (psq *PlaySessionQuery) Clone() *PlaySessionQuery {
	if psq == nil {
		return nil
	}
	return &PlaySessionQuery{
		config:     psq.config,
		ctx:        psq.ctx.Clone(),
		order:      append([]playsession.OrderOption{}, psq.order...),
		inters:     append([]Interceptor{}, psq.inters...),
		predicates: append([]predicate.PlaySession{}, psq.predicates...),
		withVideo:  psq.withVideo.Clone(),
		// clone intermediate query.
		sql:  psq.sql.Clone(),
		path: psq.path,
	}
}

// WithVideo tells the query-builder to eager-load the nodes that are connected to
// the "video" edge. The optional arguments are used to configure the query builder of the edge.
func (psq *PlaySessionQuery) WithVideo(opts ...func(*VideoQuery)) *PlaySessionQuery {
	query := (&VideoClient{config: psq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	psq.withVideo = query
	return psq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (psq *PlaySessionQuery) GroupBy(field string, fields ...string) *PlaySessionGroupBy {
	psq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PlaySessionGroupBy{build: psq}
	grbuild.flds = &psq.ctx.Fields
	grbuild.label = playsession.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (psq *PlaySessionQuery) Select(fields ...string) *PlaySessionSelect {
	psq.ctx.Fields = append(psq.ctx.Fields, fields...)
	sbuild := &PlaySessionSelect{PlaySessionQuery: psq}
	sbuild.label = playsession.Label
	sbuild.flds, sbuild.scan = &psq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PlaySessionSelect configured with the given aggregations.
func (psq *PlaySessionQuery) Aggregate(fns ...AggregateFunc) *PlaySessionSelect {
	return psq.Select().Aggregate(fns...)
}

func (psq *PlaySessionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range psq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, psq); err != nil {
				return err
			}
		}
	}
	for _, f := range psq.ctx.Fields {
		if !playsession.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if psq.path != nil {
		prev, err := psq.path(ctx)
		if err != nil {
			return err
		}
		psq.sql = prev
	}
	return nil
}

func (psq *PlaySessionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PlaySession, error) {
	var (
		nodes       = []*PlaySession{}
		withFKs     = psq.withFKs
		_spec       = psq.querySpec()
		loadedTypes = [1]bool{
			psq.withVideo != nil,
		}
	)
	if psq.withVideo != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, playsession.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PlaySession).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PlaySession{config: psq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(psq.modifiers) > 0 {
		_spec.Modifiers = psq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, psq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := psq.withVideo; query != nil {
		if err := psq.loadVideo(ctx, query, nodes, nil,
			func(n *PlaySession, e *Video) { n.Edges.Video = e }); err != nil {
			return nil, err
		}
	}
	for i := range psq.loadTotal {
		if err := psq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (psq *PlaySessionQuery) loadVideo(ctx context.Context, query *VideoQuery, nodes []*PlaySession, init func(*PlaySession), assign func(*PlaySession, *Video)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PlaySession)
	for i := range nodes {
		if nodes[i].video_play_sessions == nil {
			continue
		}
		fk := *nodes[i].video_play_sessions
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(video.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "video_play_sessions" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (psq *PlaySessionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := psq.querySpec()
	if len(psq.modifiers) > 0 {
		_spec.Modifiers = psq.modifiers
	}
	_spec.Node.Columns = psq.ctx.Fields
	if len(psq.ctx.Fields) > 0 {
		_spec.Unique = psq.ctx.Unique != nil && *psq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, psq.driver, _spec)
}

func (psq *PlaySessionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(playsession.Table, playsession.Columns, sqlgraph.NewFieldSpec(playsession.FieldID, field.TypeInt))
	_spec.From = psq.sql
	if unique := psq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if psq.path != nil {
		_spec.Unique = true
	}
	if fields := psq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, playsession.FieldID)
		for i := range fields {
			if fields[i] != playsession.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := psq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := psq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := psq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := psq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (psq *PlaySessionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(psq.driver.Dialect())
	t1 := builder.Table(playsession.Table)
	columns := psq.ctx.Fields
	if len(columns) == 0 {
		columns = playsession.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if psq.sql != nil {
		selector = psq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if psq.ctx.Unique != nil && *psq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range psq.predicates {
		p(selector)
	}
	for _, p := range psq.order {
		p(selector)
	}
	if offset := psq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := psq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PlaySessionGroupBy is the group-by builder for PlaySession entities.
type PlaySessionGroupBy struct {
	selector
	build *PlaySessionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (psgb *PlaySessionGroupBy) Aggregate(fns ...AggregateFunc) *PlaySessionGroupBy {
	psgb.fns = append(psgb.fns, fns...)
	return psgb
}

// Scan applies the selector query and scans the result into the given value.
func (psgb *PlaySessionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, psgb.build.ctx, "GroupBy")
	if err := psgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PlaySessionQuery, *PlaySessionGroupBy](ctx, psgb.build, psgb, psgb.build.inters, v)
}

func (psgb *PlaySessionGroupBy) sqlScan(ctx context.Context, root *PlaySessionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(psgb.fns))
	for _, fn := range psgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*psgb.flds)+len(psgb.fns))
		for _, f := range *psgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*psgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := psgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PlaySessionSelect is the builder for selecting fields of PlaySession entities.
type PlaySessionSelect struct {
	*PlaySessionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pss *PlaySessionSelect) Aggregate(fns ...AggregateFunc) *PlaySessionSelect {
	pss.fns = append(pss.fns, fns...)
	return pss
}

// Scan applies the selector query and scans the result into the given value.
func (pss *PlaySessionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pss.ctx, "Select")
	if err := pss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PlaySessionQuery, *PlaySessionSelect](ctx, pss.PlaySessionQuery, pss, pss.inters, v)
}

func (pss *PlaySessionSelect) sqlScan(ctx context.Context, root *PlaySessionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pss.fns))
	for _, fn := range pss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}