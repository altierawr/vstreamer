// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/altierawr/vstreamer/ent/playsessionmedia"
	"github.com/altierawr/vstreamer/ent/predicate"
	"github.com/altierawr/vstreamer/ent/stream"
	"github.com/altierawr/vstreamer/ent/videocodec"
)

// VideoCodecQuery is the builder for querying VideoCodec entities.
type VideoCodecQuery struct {
	config
	ctx              *QueryContext
	order            []videocodec.OrderOption
	inters           []Interceptor
	predicates       []predicate.VideoCodec
	withStreams      *StreamQuery
	withMedia        *PlaySessionMediaQuery
	withFKs          bool
	modifiers        []func(*sql.Selector)
	loadTotal        []func(context.Context, []*VideoCodec) error
	withNamedStreams map[string]*StreamQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VideoCodecQuery builder.
func (vcq *VideoCodecQuery) Where(ps ...predicate.VideoCodec) *VideoCodecQuery {
	vcq.predicates = append(vcq.predicates, ps...)
	return vcq
}

// Limit the number of records to be returned by this query.
func (vcq *VideoCodecQuery) Limit(limit int) *VideoCodecQuery {
	vcq.ctx.Limit = &limit
	return vcq
}

// Offset to start from.
func (vcq *VideoCodecQuery) Offset(offset int) *VideoCodecQuery {
	vcq.ctx.Offset = &offset
	return vcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vcq *VideoCodecQuery) Unique(unique bool) *VideoCodecQuery {
	vcq.ctx.Unique = &unique
	return vcq
}

// Order specifies how the records should be ordered.
func (vcq *VideoCodecQuery) Order(o ...videocodec.OrderOption) *VideoCodecQuery {
	vcq.order = append(vcq.order, o...)
	return vcq
}

// QueryStreams chains the current query on the "streams" edge.
func (vcq *VideoCodecQuery) QueryStreams() *StreamQuery {
	query := (&StreamClient{config: vcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(videocodec.Table, videocodec.FieldID, selector),
			sqlgraph.To(stream.Table, stream.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, videocodec.StreamsTable, videocodec.StreamsColumn),
		)
		fromU = sqlgraph.SetNeighbors(vcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMedia chains the current query on the "media" edge.
func (vcq *VideoCodecQuery) QueryMedia() *PlaySessionMediaQuery {
	query := (&PlaySessionMediaClient{config: vcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(videocodec.Table, videocodec.FieldID, selector),
			sqlgraph.To(playsessionmedia.Table, playsessionmedia.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, videocodec.MediaTable, videocodec.MediaColumn),
		)
		fromU = sqlgraph.SetNeighbors(vcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first VideoCodec entity from the query.
// Returns a *NotFoundError when no VideoCodec was found.
func (vcq *VideoCodecQuery) First(ctx context.Context) (*VideoCodec, error) {
	nodes, err := vcq.Limit(1).All(setContextOp(ctx, vcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{videocodec.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vcq *VideoCodecQuery) FirstX(ctx context.Context) *VideoCodec {
	node, err := vcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VideoCodec ID from the query.
// Returns a *NotFoundError when no VideoCodec ID was found.
func (vcq *VideoCodecQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vcq.Limit(1).IDs(setContextOp(ctx, vcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{videocodec.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vcq *VideoCodecQuery) FirstIDX(ctx context.Context) int {
	id, err := vcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VideoCodec entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one VideoCodec entity is found.
// Returns a *NotFoundError when no VideoCodec entities are found.
func (vcq *VideoCodecQuery) Only(ctx context.Context) (*VideoCodec, error) {
	nodes, err := vcq.Limit(2).All(setContextOp(ctx, vcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{videocodec.Label}
	default:
		return nil, &NotSingularError{videocodec.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vcq *VideoCodecQuery) OnlyX(ctx context.Context) *VideoCodec {
	node, err := vcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VideoCodec ID in the query.
// Returns a *NotSingularError when more than one VideoCodec ID is found.
// Returns a *NotFoundError when no entities are found.
func (vcq *VideoCodecQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vcq.Limit(2).IDs(setContextOp(ctx, vcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{videocodec.Label}
	default:
		err = &NotSingularError{videocodec.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vcq *VideoCodecQuery) OnlyIDX(ctx context.Context) int {
	id, err := vcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VideoCodecs.
func (vcq *VideoCodecQuery) All(ctx context.Context) ([]*VideoCodec, error) {
	ctx = setContextOp(ctx, vcq.ctx, "All")
	if err := vcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*VideoCodec, *VideoCodecQuery]()
	return withInterceptors[[]*VideoCodec](ctx, vcq, qr, vcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vcq *VideoCodecQuery) AllX(ctx context.Context) []*VideoCodec {
	nodes, err := vcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VideoCodec IDs.
func (vcq *VideoCodecQuery) IDs(ctx context.Context) (ids []int, err error) {
	if vcq.ctx.Unique == nil && vcq.path != nil {
		vcq.Unique(true)
	}
	ctx = setContextOp(ctx, vcq.ctx, "IDs")
	if err = vcq.Select(videocodec.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vcq *VideoCodecQuery) IDsX(ctx context.Context) []int {
	ids, err := vcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vcq *VideoCodecQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vcq.ctx, "Count")
	if err := vcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vcq, querierCount[*VideoCodecQuery](), vcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vcq *VideoCodecQuery) CountX(ctx context.Context) int {
	count, err := vcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vcq *VideoCodecQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vcq.ctx, "Exist")
	switch _, err := vcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vcq *VideoCodecQuery) ExistX(ctx context.Context) bool {
	exist, err := vcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VideoCodecQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vcq *VideoCodecQuery) Clone() *VideoCodecQuery {
	if vcq == nil {
		return nil
	}
	return &VideoCodecQuery{
		config:      vcq.config,
		ctx:         vcq.ctx.Clone(),
		order:       append([]videocodec.OrderOption{}, vcq.order...),
		inters:      append([]Interceptor{}, vcq.inters...),
		predicates:  append([]predicate.VideoCodec{}, vcq.predicates...),
		withStreams: vcq.withStreams.Clone(),
		withMedia:   vcq.withMedia.Clone(),
		// clone intermediate query.
		sql:  vcq.sql.Clone(),
		path: vcq.path,
	}
}

// WithStreams tells the query-builder to eager-load the nodes that are connected to
// the "streams" edge. The optional arguments are used to configure the query builder of the edge.
func (vcq *VideoCodecQuery) WithStreams(opts ...func(*StreamQuery)) *VideoCodecQuery {
	query := (&StreamClient{config: vcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vcq.withStreams = query
	return vcq
}

// WithMedia tells the query-builder to eager-load the nodes that are connected to
// the "media" edge. The optional arguments are used to configure the query builder of the edge.
func (vcq *VideoCodecQuery) WithMedia(opts ...func(*PlaySessionMediaQuery)) *VideoCodecQuery {
	query := (&PlaySessionMediaClient{config: vcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vcq.withMedia = query
	return vcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.VideoCodec.Query().
//		GroupBy(videocodec.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vcq *VideoCodecQuery) GroupBy(field string, fields ...string) *VideoCodecGroupBy {
	vcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VideoCodecGroupBy{build: vcq}
	grbuild.flds = &vcq.ctx.Fields
	grbuild.label = videocodec.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.VideoCodec.Query().
//		Select(videocodec.FieldName).
//		Scan(ctx, &v)
func (vcq *VideoCodecQuery) Select(fields ...string) *VideoCodecSelect {
	vcq.ctx.Fields = append(vcq.ctx.Fields, fields...)
	sbuild := &VideoCodecSelect{VideoCodecQuery: vcq}
	sbuild.label = videocodec.Label
	sbuild.flds, sbuild.scan = &vcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VideoCodecSelect configured with the given aggregations.
func (vcq *VideoCodecQuery) Aggregate(fns ...AggregateFunc) *VideoCodecSelect {
	return vcq.Select().Aggregate(fns...)
}

func (vcq *VideoCodecQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vcq); err != nil {
				return err
			}
		}
	}
	for _, f := range vcq.ctx.Fields {
		if !videocodec.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vcq.path != nil {
		prev, err := vcq.path(ctx)
		if err != nil {
			return err
		}
		vcq.sql = prev
	}
	return nil
}

func (vcq *VideoCodecQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*VideoCodec, error) {
	var (
		nodes       = []*VideoCodec{}
		withFKs     = vcq.withFKs
		_spec       = vcq.querySpec()
		loadedTypes = [2]bool{
			vcq.withStreams != nil,
			vcq.withMedia != nil,
		}
	)
	if vcq.withMedia != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, videocodec.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*VideoCodec).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &VideoCodec{config: vcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(vcq.modifiers) > 0 {
		_spec.Modifiers = vcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := vcq.withStreams; query != nil {
		if err := vcq.loadStreams(ctx, query, nodes,
			func(n *VideoCodec) { n.Edges.Streams = []*Stream{} },
			func(n *VideoCodec, e *Stream) { n.Edges.Streams = append(n.Edges.Streams, e) }); err != nil {
			return nil, err
		}
	}
	if query := vcq.withMedia; query != nil {
		if err := vcq.loadMedia(ctx, query, nodes, nil,
			func(n *VideoCodec, e *PlaySessionMedia) { n.Edges.Media = e }); err != nil {
			return nil, err
		}
	}
	for name, query := range vcq.withNamedStreams {
		if err := vcq.loadStreams(ctx, query, nodes,
			func(n *VideoCodec) { n.appendNamedStreams(name) },
			func(n *VideoCodec, e *Stream) { n.appendNamedStreams(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range vcq.loadTotal {
		if err := vcq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vcq *VideoCodecQuery) loadStreams(ctx context.Context, query *StreamQuery, nodes []*VideoCodec, init func(*VideoCodec), assign func(*VideoCodec, *Stream)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*VideoCodec)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Stream(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(videocodec.StreamsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.video_codec_streams
		if fk == nil {
			return fmt.Errorf(`foreign-key "video_codec_streams" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "video_codec_streams" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (vcq *VideoCodecQuery) loadMedia(ctx context.Context, query *PlaySessionMediaQuery, nodes []*VideoCodec, init func(*VideoCodec), assign func(*VideoCodec, *PlaySessionMedia)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*VideoCodec)
	for i := range nodes {
		if nodes[i].play_session_media_video_codecs == nil {
			continue
		}
		fk := *nodes[i].play_session_media_video_codecs
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(playsessionmedia.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "play_session_media_video_codecs" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (vcq *VideoCodecQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vcq.querySpec()
	if len(vcq.modifiers) > 0 {
		_spec.Modifiers = vcq.modifiers
	}
	_spec.Node.Columns = vcq.ctx.Fields
	if len(vcq.ctx.Fields) > 0 {
		_spec.Unique = vcq.ctx.Unique != nil && *vcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vcq.driver, _spec)
}

func (vcq *VideoCodecQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(videocodec.Table, videocodec.Columns, sqlgraph.NewFieldSpec(videocodec.FieldID, field.TypeInt))
	_spec.From = vcq.sql
	if unique := vcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vcq.path != nil {
		_spec.Unique = true
	}
	if fields := vcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, videocodec.FieldID)
		for i := range fields {
			if fields[i] != videocodec.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vcq *VideoCodecQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vcq.driver.Dialect())
	t1 := builder.Table(videocodec.Table)
	columns := vcq.ctx.Fields
	if len(columns) == 0 {
		columns = videocodec.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vcq.sql != nil {
		selector = vcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vcq.ctx.Unique != nil && *vcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vcq.predicates {
		p(selector)
	}
	for _, p := range vcq.order {
		p(selector)
	}
	if offset := vcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedStreams tells the query-builder to eager-load the nodes that are connected to the "streams"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (vcq *VideoCodecQuery) WithNamedStreams(name string, opts ...func(*StreamQuery)) *VideoCodecQuery {
	query := (&StreamClient{config: vcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if vcq.withNamedStreams == nil {
		vcq.withNamedStreams = make(map[string]*StreamQuery)
	}
	vcq.withNamedStreams[name] = query
	return vcq
}

// VideoCodecGroupBy is the group-by builder for VideoCodec entities.
type VideoCodecGroupBy struct {
	selector
	build *VideoCodecQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vcgb *VideoCodecGroupBy) Aggregate(fns ...AggregateFunc) *VideoCodecGroupBy {
	vcgb.fns = append(vcgb.fns, fns...)
	return vcgb
}

// Scan applies the selector query and scans the result into the given value.
func (vcgb *VideoCodecGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vcgb.build.ctx, "GroupBy")
	if err := vcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VideoCodecQuery, *VideoCodecGroupBy](ctx, vcgb.build, vcgb, vcgb.build.inters, v)
}

func (vcgb *VideoCodecGroupBy) sqlScan(ctx context.Context, root *VideoCodecQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vcgb.fns))
	for _, fn := range vcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vcgb.flds)+len(vcgb.fns))
		for _, f := range *vcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VideoCodecSelect is the builder for selecting fields of VideoCodec entities.
type VideoCodecSelect struct {
	*VideoCodecQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vcs *VideoCodecSelect) Aggregate(fns ...AggregateFunc) *VideoCodecSelect {
	vcs.fns = append(vcs.fns, fns...)
	return vcs
}

// Scan applies the selector query and scans the result into the given value.
func (vcs *VideoCodecSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vcs.ctx, "Select")
	if err := vcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VideoCodecQuery, *VideoCodecSelect](ctx, vcs.VideoCodecQuery, vcs, vcs.inters, v)
}

func (vcs *VideoCodecSelect) sqlScan(ctx context.Context, root *VideoCodecQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vcs.fns))
	for _, fn := range vcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
