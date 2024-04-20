// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/altierawr/vstreamer/ent/playbackclient"
	"github.com/altierawr/vstreamer/ent/playsession"
	"github.com/altierawr/vstreamer/ent/predicate"
)

// PlaybackClientQuery is the builder for querying PlaybackClient entities.
type PlaybackClientQuery struct {
	config
	ctx         *QueryContext
	order       []playbackclient.OrderOption
	inters      []Interceptor
	predicates  []predicate.PlaybackClient
	withSession *PlaySessionQuery
	withFKs     bool
	modifiers   []func(*sql.Selector)
	loadTotal   []func(context.Context, []*PlaybackClient) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PlaybackClientQuery builder.
func (pcq *PlaybackClientQuery) Where(ps ...predicate.PlaybackClient) *PlaybackClientQuery {
	pcq.predicates = append(pcq.predicates, ps...)
	return pcq
}

// Limit the number of records to be returned by this query.
func (pcq *PlaybackClientQuery) Limit(limit int) *PlaybackClientQuery {
	pcq.ctx.Limit = &limit
	return pcq
}

// Offset to start from.
func (pcq *PlaybackClientQuery) Offset(offset int) *PlaybackClientQuery {
	pcq.ctx.Offset = &offset
	return pcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pcq *PlaybackClientQuery) Unique(unique bool) *PlaybackClientQuery {
	pcq.ctx.Unique = &unique
	return pcq
}

// Order specifies how the records should be ordered.
func (pcq *PlaybackClientQuery) Order(o ...playbackclient.OrderOption) *PlaybackClientQuery {
	pcq.order = append(pcq.order, o...)
	return pcq
}

// QuerySession chains the current query on the "session" edge.
func (pcq *PlaybackClientQuery) QuerySession() *PlaySessionQuery {
	query := (&PlaySessionClient{config: pcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(playbackclient.Table, playbackclient.FieldID, selector),
			sqlgraph.To(playsession.Table, playsession.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, playbackclient.SessionTable, playbackclient.SessionColumn),
		)
		fromU = sqlgraph.SetNeighbors(pcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PlaybackClient entity from the query.
// Returns a *NotFoundError when no PlaybackClient was found.
func (pcq *PlaybackClientQuery) First(ctx context.Context) (*PlaybackClient, error) {
	nodes, err := pcq.Limit(1).All(setContextOp(ctx, pcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{playbackclient.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pcq *PlaybackClientQuery) FirstX(ctx context.Context) *PlaybackClient {
	node, err := pcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PlaybackClient ID from the query.
// Returns a *NotFoundError when no PlaybackClient ID was found.
func (pcq *PlaybackClientQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pcq.Limit(1).IDs(setContextOp(ctx, pcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{playbackclient.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pcq *PlaybackClientQuery) FirstIDX(ctx context.Context) int {
	id, err := pcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PlaybackClient entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PlaybackClient entity is found.
// Returns a *NotFoundError when no PlaybackClient entities are found.
func (pcq *PlaybackClientQuery) Only(ctx context.Context) (*PlaybackClient, error) {
	nodes, err := pcq.Limit(2).All(setContextOp(ctx, pcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{playbackclient.Label}
	default:
		return nil, &NotSingularError{playbackclient.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pcq *PlaybackClientQuery) OnlyX(ctx context.Context) *PlaybackClient {
	node, err := pcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PlaybackClient ID in the query.
// Returns a *NotSingularError when more than one PlaybackClient ID is found.
// Returns a *NotFoundError when no entities are found.
func (pcq *PlaybackClientQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pcq.Limit(2).IDs(setContextOp(ctx, pcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{playbackclient.Label}
	default:
		err = &NotSingularError{playbackclient.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pcq *PlaybackClientQuery) OnlyIDX(ctx context.Context) int {
	id, err := pcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PlaybackClients.
func (pcq *PlaybackClientQuery) All(ctx context.Context) ([]*PlaybackClient, error) {
	ctx = setContextOp(ctx, pcq.ctx, "All")
	if err := pcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PlaybackClient, *PlaybackClientQuery]()
	return withInterceptors[[]*PlaybackClient](ctx, pcq, qr, pcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pcq *PlaybackClientQuery) AllX(ctx context.Context) []*PlaybackClient {
	nodes, err := pcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PlaybackClient IDs.
func (pcq *PlaybackClientQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pcq.ctx.Unique == nil && pcq.path != nil {
		pcq.Unique(true)
	}
	ctx = setContextOp(ctx, pcq.ctx, "IDs")
	if err = pcq.Select(playbackclient.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pcq *PlaybackClientQuery) IDsX(ctx context.Context) []int {
	ids, err := pcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pcq *PlaybackClientQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pcq.ctx, "Count")
	if err := pcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pcq, querierCount[*PlaybackClientQuery](), pcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pcq *PlaybackClientQuery) CountX(ctx context.Context) int {
	count, err := pcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pcq *PlaybackClientQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pcq.ctx, "Exist")
	switch _, err := pcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pcq *PlaybackClientQuery) ExistX(ctx context.Context) bool {
	exist, err := pcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PlaybackClientQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pcq *PlaybackClientQuery) Clone() *PlaybackClientQuery {
	if pcq == nil {
		return nil
	}
	return &PlaybackClientQuery{
		config:      pcq.config,
		ctx:         pcq.ctx.Clone(),
		order:       append([]playbackclient.OrderOption{}, pcq.order...),
		inters:      append([]Interceptor{}, pcq.inters...),
		predicates:  append([]predicate.PlaybackClient{}, pcq.predicates...),
		withSession: pcq.withSession.Clone(),
		// clone intermediate query.
		sql:  pcq.sql.Clone(),
		path: pcq.path,
	}
}

// WithSession tells the query-builder to eager-load the nodes that are connected to
// the "session" edge. The optional arguments are used to configure the query builder of the edge.
func (pcq *PlaybackClientQuery) WithSession(opts ...func(*PlaySessionQuery)) *PlaybackClientQuery {
	query := (&PlaySessionClient{config: pcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pcq.withSession = query
	return pcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		IsBuffered bool `json:"is_buffered,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PlaybackClient.Query().
//		GroupBy(playbackclient.FieldIsBuffered).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pcq *PlaybackClientQuery) GroupBy(field string, fields ...string) *PlaybackClientGroupBy {
	pcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PlaybackClientGroupBy{build: pcq}
	grbuild.flds = &pcq.ctx.Fields
	grbuild.label = playbackclient.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		IsBuffered bool `json:"is_buffered,omitempty"`
//	}
//
//	client.PlaybackClient.Query().
//		Select(playbackclient.FieldIsBuffered).
//		Scan(ctx, &v)
func (pcq *PlaybackClientQuery) Select(fields ...string) *PlaybackClientSelect {
	pcq.ctx.Fields = append(pcq.ctx.Fields, fields...)
	sbuild := &PlaybackClientSelect{PlaybackClientQuery: pcq}
	sbuild.label = playbackclient.Label
	sbuild.flds, sbuild.scan = &pcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PlaybackClientSelect configured with the given aggregations.
func (pcq *PlaybackClientQuery) Aggregate(fns ...AggregateFunc) *PlaybackClientSelect {
	return pcq.Select().Aggregate(fns...)
}

func (pcq *PlaybackClientQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pcq); err != nil {
				return err
			}
		}
	}
	for _, f := range pcq.ctx.Fields {
		if !playbackclient.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pcq.path != nil {
		prev, err := pcq.path(ctx)
		if err != nil {
			return err
		}
		pcq.sql = prev
	}
	return nil
}

func (pcq *PlaybackClientQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PlaybackClient, error) {
	var (
		nodes       = []*PlaybackClient{}
		withFKs     = pcq.withFKs
		_spec       = pcq.querySpec()
		loadedTypes = [1]bool{
			pcq.withSession != nil,
		}
	)
	if pcq.withSession != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, playbackclient.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PlaybackClient).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PlaybackClient{config: pcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(pcq.modifiers) > 0 {
		_spec.Modifiers = pcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pcq.withSession; query != nil {
		if err := pcq.loadSession(ctx, query, nodes, nil,
			func(n *PlaybackClient, e *PlaySession) { n.Edges.Session = e }); err != nil {
			return nil, err
		}
	}
	for i := range pcq.loadTotal {
		if err := pcq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pcq *PlaybackClientQuery) loadSession(ctx context.Context, query *PlaySessionQuery, nodes []*PlaybackClient, init func(*PlaybackClient), assign func(*PlaybackClient, *PlaySession)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PlaybackClient)
	for i := range nodes {
		if nodes[i].play_session_clients == nil {
			continue
		}
		fk := *nodes[i].play_session_clients
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(playsession.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "play_session_clients" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pcq *PlaybackClientQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pcq.querySpec()
	if len(pcq.modifiers) > 0 {
		_spec.Modifiers = pcq.modifiers
	}
	_spec.Node.Columns = pcq.ctx.Fields
	if len(pcq.ctx.Fields) > 0 {
		_spec.Unique = pcq.ctx.Unique != nil && *pcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pcq.driver, _spec)
}

func (pcq *PlaybackClientQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(playbackclient.Table, playbackclient.Columns, sqlgraph.NewFieldSpec(playbackclient.FieldID, field.TypeInt))
	_spec.From = pcq.sql
	if unique := pcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pcq.path != nil {
		_spec.Unique = true
	}
	if fields := pcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, playbackclient.FieldID)
		for i := range fields {
			if fields[i] != playbackclient.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pcq *PlaybackClientQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pcq.driver.Dialect())
	t1 := builder.Table(playbackclient.Table)
	columns := pcq.ctx.Fields
	if len(columns) == 0 {
		columns = playbackclient.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pcq.sql != nil {
		selector = pcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pcq.ctx.Unique != nil && *pcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pcq.predicates {
		p(selector)
	}
	for _, p := range pcq.order {
		p(selector)
	}
	if offset := pcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PlaybackClientGroupBy is the group-by builder for PlaybackClient entities.
type PlaybackClientGroupBy struct {
	selector
	build *PlaybackClientQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pcgb *PlaybackClientGroupBy) Aggregate(fns ...AggregateFunc) *PlaybackClientGroupBy {
	pcgb.fns = append(pcgb.fns, fns...)
	return pcgb
}

// Scan applies the selector query and scans the result into the given value.
func (pcgb *PlaybackClientGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pcgb.build.ctx, "GroupBy")
	if err := pcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PlaybackClientQuery, *PlaybackClientGroupBy](ctx, pcgb.build, pcgb, pcgb.build.inters, v)
}

func (pcgb *PlaybackClientGroupBy) sqlScan(ctx context.Context, root *PlaybackClientQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pcgb.fns))
	for _, fn := range pcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pcgb.flds)+len(pcgb.fns))
		for _, f := range *pcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PlaybackClientSelect is the builder for selecting fields of PlaybackClient entities.
type PlaybackClientSelect struct {
	*PlaybackClientQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pcs *PlaybackClientSelect) Aggregate(fns ...AggregateFunc) *PlaybackClientSelect {
	pcs.fns = append(pcs.fns, fns...)
	return pcs
}

// Scan applies the selector query and scans the result into the given value.
func (pcs *PlaybackClientSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pcs.ctx, "Select")
	if err := pcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PlaybackClientQuery, *PlaybackClientSelect](ctx, pcs.PlaybackClientQuery, pcs, pcs.inters, v)
}

func (pcs *PlaybackClientSelect) sqlScan(ctx context.Context, root *PlaybackClientQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pcs.fns))
	for _, fn := range pcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}