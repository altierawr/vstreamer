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
	"github.com/altierawr/vstreamer/ent/audiotrack"
	"github.com/altierawr/vstreamer/ent/playsession"
	"github.com/altierawr/vstreamer/ent/playsessionmedia"
	"github.com/altierawr/vstreamer/ent/predicate"
	"github.com/altierawr/vstreamer/ent/video"
)

// PlaySessionMediaQuery is the builder for querying PlaySessionMedia entities.
type PlaySessionMediaQuery struct {
	config
	ctx                  *QueryContext
	order                []playsessionmedia.OrderOption
	inters               []Interceptor
	predicates           []predicate.PlaySessionMedia
	withAudioTracks      *AudioTrackQuery
	withVideo            *VideoQuery
	withSession          *PlaySessionQuery
	withFKs              bool
	modifiers            []func(*sql.Selector)
	loadTotal            []func(context.Context, []*PlaySessionMedia) error
	withNamedAudioTracks map[string]*AudioTrackQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PlaySessionMediaQuery builder.
func (psmq *PlaySessionMediaQuery) Where(ps ...predicate.PlaySessionMedia) *PlaySessionMediaQuery {
	psmq.predicates = append(psmq.predicates, ps...)
	return psmq
}

// Limit the number of records to be returned by this query.
func (psmq *PlaySessionMediaQuery) Limit(limit int) *PlaySessionMediaQuery {
	psmq.ctx.Limit = &limit
	return psmq
}

// Offset to start from.
func (psmq *PlaySessionMediaQuery) Offset(offset int) *PlaySessionMediaQuery {
	psmq.ctx.Offset = &offset
	return psmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (psmq *PlaySessionMediaQuery) Unique(unique bool) *PlaySessionMediaQuery {
	psmq.ctx.Unique = &unique
	return psmq
}

// Order specifies how the records should be ordered.
func (psmq *PlaySessionMediaQuery) Order(o ...playsessionmedia.OrderOption) *PlaySessionMediaQuery {
	psmq.order = append(psmq.order, o...)
	return psmq
}

// QueryAudioTracks chains the current query on the "audio_tracks" edge.
func (psmq *PlaySessionMediaQuery) QueryAudioTracks() *AudioTrackQuery {
	query := (&AudioTrackClient{config: psmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := psmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := psmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(playsessionmedia.Table, playsessionmedia.FieldID, selector),
			sqlgraph.To(audiotrack.Table, audiotrack.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, playsessionmedia.AudioTracksTable, playsessionmedia.AudioTracksColumn),
		)
		fromU = sqlgraph.SetNeighbors(psmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVideo chains the current query on the "video" edge.
func (psmq *PlaySessionMediaQuery) QueryVideo() *VideoQuery {
	query := (&VideoClient{config: psmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := psmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := psmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(playsessionmedia.Table, playsessionmedia.FieldID, selector),
			sqlgraph.To(video.Table, video.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, playsessionmedia.VideoTable, playsessionmedia.VideoColumn),
		)
		fromU = sqlgraph.SetNeighbors(psmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySession chains the current query on the "session" edge.
func (psmq *PlaySessionMediaQuery) QuerySession() *PlaySessionQuery {
	query := (&PlaySessionClient{config: psmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := psmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := psmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(playsessionmedia.Table, playsessionmedia.FieldID, selector),
			sqlgraph.To(playsession.Table, playsession.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, playsessionmedia.SessionTable, playsessionmedia.SessionColumn),
		)
		fromU = sqlgraph.SetNeighbors(psmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PlaySessionMedia entity from the query.
// Returns a *NotFoundError when no PlaySessionMedia was found.
func (psmq *PlaySessionMediaQuery) First(ctx context.Context) (*PlaySessionMedia, error) {
	nodes, err := psmq.Limit(1).All(setContextOp(ctx, psmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{playsessionmedia.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (psmq *PlaySessionMediaQuery) FirstX(ctx context.Context) *PlaySessionMedia {
	node, err := psmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PlaySessionMedia ID from the query.
// Returns a *NotFoundError when no PlaySessionMedia ID was found.
func (psmq *PlaySessionMediaQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = psmq.Limit(1).IDs(setContextOp(ctx, psmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{playsessionmedia.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (psmq *PlaySessionMediaQuery) FirstIDX(ctx context.Context) int {
	id, err := psmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PlaySessionMedia entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PlaySessionMedia entity is found.
// Returns a *NotFoundError when no PlaySessionMedia entities are found.
func (psmq *PlaySessionMediaQuery) Only(ctx context.Context) (*PlaySessionMedia, error) {
	nodes, err := psmq.Limit(2).All(setContextOp(ctx, psmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{playsessionmedia.Label}
	default:
		return nil, &NotSingularError{playsessionmedia.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (psmq *PlaySessionMediaQuery) OnlyX(ctx context.Context) *PlaySessionMedia {
	node, err := psmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PlaySessionMedia ID in the query.
// Returns a *NotSingularError when more than one PlaySessionMedia ID is found.
// Returns a *NotFoundError when no entities are found.
func (psmq *PlaySessionMediaQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = psmq.Limit(2).IDs(setContextOp(ctx, psmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{playsessionmedia.Label}
	default:
		err = &NotSingularError{playsessionmedia.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (psmq *PlaySessionMediaQuery) OnlyIDX(ctx context.Context) int {
	id, err := psmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PlaySessionMediaSlice.
func (psmq *PlaySessionMediaQuery) All(ctx context.Context) ([]*PlaySessionMedia, error) {
	ctx = setContextOp(ctx, psmq.ctx, "All")
	if err := psmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PlaySessionMedia, *PlaySessionMediaQuery]()
	return withInterceptors[[]*PlaySessionMedia](ctx, psmq, qr, psmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (psmq *PlaySessionMediaQuery) AllX(ctx context.Context) []*PlaySessionMedia {
	nodes, err := psmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PlaySessionMedia IDs.
func (psmq *PlaySessionMediaQuery) IDs(ctx context.Context) (ids []int, err error) {
	if psmq.ctx.Unique == nil && psmq.path != nil {
		psmq.Unique(true)
	}
	ctx = setContextOp(ctx, psmq.ctx, "IDs")
	if err = psmq.Select(playsessionmedia.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (psmq *PlaySessionMediaQuery) IDsX(ctx context.Context) []int {
	ids, err := psmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (psmq *PlaySessionMediaQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, psmq.ctx, "Count")
	if err := psmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, psmq, querierCount[*PlaySessionMediaQuery](), psmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (psmq *PlaySessionMediaQuery) CountX(ctx context.Context) int {
	count, err := psmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (psmq *PlaySessionMediaQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, psmq.ctx, "Exist")
	switch _, err := psmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (psmq *PlaySessionMediaQuery) ExistX(ctx context.Context) bool {
	exist, err := psmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PlaySessionMediaQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (psmq *PlaySessionMediaQuery) Clone() *PlaySessionMediaQuery {
	if psmq == nil {
		return nil
	}
	return &PlaySessionMediaQuery{
		config:          psmq.config,
		ctx:             psmq.ctx.Clone(),
		order:           append([]playsessionmedia.OrderOption{}, psmq.order...),
		inters:          append([]Interceptor{}, psmq.inters...),
		predicates:      append([]predicate.PlaySessionMedia{}, psmq.predicates...),
		withAudioTracks: psmq.withAudioTracks.Clone(),
		withVideo:       psmq.withVideo.Clone(),
		withSession:     psmq.withSession.Clone(),
		// clone intermediate query.
		sql:  psmq.sql.Clone(),
		path: psmq.path,
	}
}

// WithAudioTracks tells the query-builder to eager-load the nodes that are connected to
// the "audio_tracks" edge. The optional arguments are used to configure the query builder of the edge.
func (psmq *PlaySessionMediaQuery) WithAudioTracks(opts ...func(*AudioTrackQuery)) *PlaySessionMediaQuery {
	query := (&AudioTrackClient{config: psmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	psmq.withAudioTracks = query
	return psmq
}

// WithVideo tells the query-builder to eager-load the nodes that are connected to
// the "video" edge. The optional arguments are used to configure the query builder of the edge.
func (psmq *PlaySessionMediaQuery) WithVideo(opts ...func(*VideoQuery)) *PlaySessionMediaQuery {
	query := (&VideoClient{config: psmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	psmq.withVideo = query
	return psmq
}

// WithSession tells the query-builder to eager-load the nodes that are connected to
// the "session" edge. The optional arguments are used to configure the query builder of the edge.
func (psmq *PlaySessionMediaQuery) WithSession(opts ...func(*PlaySessionQuery)) *PlaySessionMediaQuery {
	query := (&PlaySessionClient{config: psmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	psmq.withSession = query
	return psmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		VideoCodecs []string `json:"video_codecs,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PlaySessionMedia.Query().
//		GroupBy(playsessionmedia.FieldVideoCodecs).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (psmq *PlaySessionMediaQuery) GroupBy(field string, fields ...string) *PlaySessionMediaGroupBy {
	psmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PlaySessionMediaGroupBy{build: psmq}
	grbuild.flds = &psmq.ctx.Fields
	grbuild.label = playsessionmedia.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		VideoCodecs []string `json:"video_codecs,omitempty"`
//	}
//
//	client.PlaySessionMedia.Query().
//		Select(playsessionmedia.FieldVideoCodecs).
//		Scan(ctx, &v)
func (psmq *PlaySessionMediaQuery) Select(fields ...string) *PlaySessionMediaSelect {
	psmq.ctx.Fields = append(psmq.ctx.Fields, fields...)
	sbuild := &PlaySessionMediaSelect{PlaySessionMediaQuery: psmq}
	sbuild.label = playsessionmedia.Label
	sbuild.flds, sbuild.scan = &psmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PlaySessionMediaSelect configured with the given aggregations.
func (psmq *PlaySessionMediaQuery) Aggregate(fns ...AggregateFunc) *PlaySessionMediaSelect {
	return psmq.Select().Aggregate(fns...)
}

func (psmq *PlaySessionMediaQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range psmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, psmq); err != nil {
				return err
			}
		}
	}
	for _, f := range psmq.ctx.Fields {
		if !playsessionmedia.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if psmq.path != nil {
		prev, err := psmq.path(ctx)
		if err != nil {
			return err
		}
		psmq.sql = prev
	}
	return nil
}

func (psmq *PlaySessionMediaQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PlaySessionMedia, error) {
	var (
		nodes       = []*PlaySessionMedia{}
		withFKs     = psmq.withFKs
		_spec       = psmq.querySpec()
		loadedTypes = [3]bool{
			psmq.withAudioTracks != nil,
			psmq.withVideo != nil,
			psmq.withSession != nil,
		}
	)
	if psmq.withVideo != nil || psmq.withSession != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, playsessionmedia.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PlaySessionMedia).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PlaySessionMedia{config: psmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(psmq.modifiers) > 0 {
		_spec.Modifiers = psmq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, psmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := psmq.withAudioTracks; query != nil {
		if err := psmq.loadAudioTracks(ctx, query, nodes,
			func(n *PlaySessionMedia) { n.Edges.AudioTracks = []*AudioTrack{} },
			func(n *PlaySessionMedia, e *AudioTrack) { n.Edges.AudioTracks = append(n.Edges.AudioTracks, e) }); err != nil {
			return nil, err
		}
	}
	if query := psmq.withVideo; query != nil {
		if err := psmq.loadVideo(ctx, query, nodes, nil,
			func(n *PlaySessionMedia, e *Video) { n.Edges.Video = e }); err != nil {
			return nil, err
		}
	}
	if query := psmq.withSession; query != nil {
		if err := psmq.loadSession(ctx, query, nodes, nil,
			func(n *PlaySessionMedia, e *PlaySession) { n.Edges.Session = e }); err != nil {
			return nil, err
		}
	}
	for name, query := range psmq.withNamedAudioTracks {
		if err := psmq.loadAudioTracks(ctx, query, nodes,
			func(n *PlaySessionMedia) { n.appendNamedAudioTracks(name) },
			func(n *PlaySessionMedia, e *AudioTrack) { n.appendNamedAudioTracks(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range psmq.loadTotal {
		if err := psmq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (psmq *PlaySessionMediaQuery) loadAudioTracks(ctx context.Context, query *AudioTrackQuery, nodes []*PlaySessionMedia, init func(*PlaySessionMedia), assign func(*PlaySessionMedia, *AudioTrack)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*PlaySessionMedia)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.AudioTrack(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(playsessionmedia.AudioTracksColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.play_session_media_audio_tracks
		if fk == nil {
			return fmt.Errorf(`foreign-key "play_session_media_audio_tracks" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "play_session_media_audio_tracks" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (psmq *PlaySessionMediaQuery) loadVideo(ctx context.Context, query *VideoQuery, nodes []*PlaySessionMedia, init func(*PlaySessionMedia), assign func(*PlaySessionMedia, *Video)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PlaySessionMedia)
	for i := range nodes {
		if nodes[i].video_play_session_medias == nil {
			continue
		}
		fk := *nodes[i].video_play_session_medias
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
			return fmt.Errorf(`unexpected foreign-key "video_play_session_medias" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (psmq *PlaySessionMediaQuery) loadSession(ctx context.Context, query *PlaySessionQuery, nodes []*PlaySessionMedia, init func(*PlaySessionMedia), assign func(*PlaySessionMedia, *PlaySession)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PlaySessionMedia)
	for i := range nodes {
		if nodes[i].play_session_media == nil {
			continue
		}
		fk := *nodes[i].play_session_media
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
			return fmt.Errorf(`unexpected foreign-key "play_session_media" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (psmq *PlaySessionMediaQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := psmq.querySpec()
	if len(psmq.modifiers) > 0 {
		_spec.Modifiers = psmq.modifiers
	}
	_spec.Node.Columns = psmq.ctx.Fields
	if len(psmq.ctx.Fields) > 0 {
		_spec.Unique = psmq.ctx.Unique != nil && *psmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, psmq.driver, _spec)
}

func (psmq *PlaySessionMediaQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(playsessionmedia.Table, playsessionmedia.Columns, sqlgraph.NewFieldSpec(playsessionmedia.FieldID, field.TypeInt))
	_spec.From = psmq.sql
	if unique := psmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if psmq.path != nil {
		_spec.Unique = true
	}
	if fields := psmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, playsessionmedia.FieldID)
		for i := range fields {
			if fields[i] != playsessionmedia.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := psmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := psmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := psmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := psmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (psmq *PlaySessionMediaQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(psmq.driver.Dialect())
	t1 := builder.Table(playsessionmedia.Table)
	columns := psmq.ctx.Fields
	if len(columns) == 0 {
		columns = playsessionmedia.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if psmq.sql != nil {
		selector = psmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if psmq.ctx.Unique != nil && *psmq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range psmq.predicates {
		p(selector)
	}
	for _, p := range psmq.order {
		p(selector)
	}
	if offset := psmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := psmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedAudioTracks tells the query-builder to eager-load the nodes that are connected to the "audio_tracks"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (psmq *PlaySessionMediaQuery) WithNamedAudioTracks(name string, opts ...func(*AudioTrackQuery)) *PlaySessionMediaQuery {
	query := (&AudioTrackClient{config: psmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if psmq.withNamedAudioTracks == nil {
		psmq.withNamedAudioTracks = make(map[string]*AudioTrackQuery)
	}
	psmq.withNamedAudioTracks[name] = query
	return psmq
}

// PlaySessionMediaGroupBy is the group-by builder for PlaySessionMedia entities.
type PlaySessionMediaGroupBy struct {
	selector
	build *PlaySessionMediaQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (psmgb *PlaySessionMediaGroupBy) Aggregate(fns ...AggregateFunc) *PlaySessionMediaGroupBy {
	psmgb.fns = append(psmgb.fns, fns...)
	return psmgb
}

// Scan applies the selector query and scans the result into the given value.
func (psmgb *PlaySessionMediaGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, psmgb.build.ctx, "GroupBy")
	if err := psmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PlaySessionMediaQuery, *PlaySessionMediaGroupBy](ctx, psmgb.build, psmgb, psmgb.build.inters, v)
}

func (psmgb *PlaySessionMediaGroupBy) sqlScan(ctx context.Context, root *PlaySessionMediaQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(psmgb.fns))
	for _, fn := range psmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*psmgb.flds)+len(psmgb.fns))
		for _, f := range *psmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*psmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := psmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PlaySessionMediaSelect is the builder for selecting fields of PlaySessionMedia entities.
type PlaySessionMediaSelect struct {
	*PlaySessionMediaQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (psms *PlaySessionMediaSelect) Aggregate(fns ...AggregateFunc) *PlaySessionMediaSelect {
	psms.fns = append(psms.fns, fns...)
	return psms
}

// Scan applies the selector query and scans the result into the given value.
func (psms *PlaySessionMediaSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, psms.ctx, "Select")
	if err := psms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PlaySessionMediaQuery, *PlaySessionMediaSelect](ctx, psms.PlaySessionMediaQuery, psms, psms.inters, v)
}

func (psms *PlaySessionMediaSelect) sqlScan(ctx context.Context, root *PlaySessionMediaQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(psms.fns))
	for _, fn := range psms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*psms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := psms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
