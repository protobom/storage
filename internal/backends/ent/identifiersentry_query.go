// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/identifiersentry"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// IdentifiersEntryQuery is the builder for querying IdentifiersEntry entities.
type IdentifiersEntryQuery struct {
	config
	ctx          *QueryContext
	order        []identifiersentry.OrderOption
	inters       []Interceptor
	predicates   []predicate.IdentifiersEntry
	withDocument *DocumentQuery
	withNodes    *NodeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the IdentifiersEntryQuery builder.
func (ieq *IdentifiersEntryQuery) Where(ps ...predicate.IdentifiersEntry) *IdentifiersEntryQuery {
	ieq.predicates = append(ieq.predicates, ps...)
	return ieq
}

// Limit the number of records to be returned by this query.
func (ieq *IdentifiersEntryQuery) Limit(limit int) *IdentifiersEntryQuery {
	ieq.ctx.Limit = &limit
	return ieq
}

// Offset to start from.
func (ieq *IdentifiersEntryQuery) Offset(offset int) *IdentifiersEntryQuery {
	ieq.ctx.Offset = &offset
	return ieq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ieq *IdentifiersEntryQuery) Unique(unique bool) *IdentifiersEntryQuery {
	ieq.ctx.Unique = &unique
	return ieq
}

// Order specifies how the records should be ordered.
func (ieq *IdentifiersEntryQuery) Order(o ...identifiersentry.OrderOption) *IdentifiersEntryQuery {
	ieq.order = append(ieq.order, o...)
	return ieq
}

// QueryDocument chains the current query on the "document" edge.
func (ieq *IdentifiersEntryQuery) QueryDocument() *DocumentQuery {
	query := (&DocumentClient{config: ieq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ieq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ieq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(identifiersentry.Table, identifiersentry.FieldID, selector),
			sqlgraph.To(document.Table, document.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, identifiersentry.DocumentTable, identifiersentry.DocumentColumn),
		)
		fromU = sqlgraph.SetNeighbors(ieq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNodes chains the current query on the "nodes" edge.
func (ieq *IdentifiersEntryQuery) QueryNodes() *NodeQuery {
	query := (&NodeClient{config: ieq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ieq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ieq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(identifiersentry.Table, identifiersentry.FieldID, selector),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, identifiersentry.NodesTable, identifiersentry.NodesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(ieq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first IdentifiersEntry entity from the query.
// Returns a *NotFoundError when no IdentifiersEntry was found.
func (ieq *IdentifiersEntryQuery) First(ctx context.Context) (*IdentifiersEntry, error) {
	nodes, err := ieq.Limit(1).All(setContextOp(ctx, ieq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{identifiersentry.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ieq *IdentifiersEntryQuery) FirstX(ctx context.Context) *IdentifiersEntry {
	node, err := ieq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first IdentifiersEntry ID from the query.
// Returns a *NotFoundError when no IdentifiersEntry ID was found.
func (ieq *IdentifiersEntryQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ieq.Limit(1).IDs(setContextOp(ctx, ieq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{identifiersentry.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ieq *IdentifiersEntryQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ieq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single IdentifiersEntry entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one IdentifiersEntry entity is found.
// Returns a *NotFoundError when no IdentifiersEntry entities are found.
func (ieq *IdentifiersEntryQuery) Only(ctx context.Context) (*IdentifiersEntry, error) {
	nodes, err := ieq.Limit(2).All(setContextOp(ctx, ieq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{identifiersentry.Label}
	default:
		return nil, &NotSingularError{identifiersentry.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ieq *IdentifiersEntryQuery) OnlyX(ctx context.Context) *IdentifiersEntry {
	node, err := ieq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only IdentifiersEntry ID in the query.
// Returns a *NotSingularError when more than one IdentifiersEntry ID is found.
// Returns a *NotFoundError when no entities are found.
func (ieq *IdentifiersEntryQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ieq.Limit(2).IDs(setContextOp(ctx, ieq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{identifiersentry.Label}
	default:
		err = &NotSingularError{identifiersentry.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ieq *IdentifiersEntryQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ieq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of IdentifiersEntries.
func (ieq *IdentifiersEntryQuery) All(ctx context.Context) ([]*IdentifiersEntry, error) {
	ctx = setContextOp(ctx, ieq.ctx, ent.OpQueryAll)
	if err := ieq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*IdentifiersEntry, *IdentifiersEntryQuery]()
	return withInterceptors[[]*IdentifiersEntry](ctx, ieq, qr, ieq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ieq *IdentifiersEntryQuery) AllX(ctx context.Context) []*IdentifiersEntry {
	nodes, err := ieq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of IdentifiersEntry IDs.
func (ieq *IdentifiersEntryQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if ieq.ctx.Unique == nil && ieq.path != nil {
		ieq.Unique(true)
	}
	ctx = setContextOp(ctx, ieq.ctx, ent.OpQueryIDs)
	if err = ieq.Select(identifiersentry.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ieq *IdentifiersEntryQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ieq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ieq *IdentifiersEntryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ieq.ctx, ent.OpQueryCount)
	if err := ieq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ieq, querierCount[*IdentifiersEntryQuery](), ieq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ieq *IdentifiersEntryQuery) CountX(ctx context.Context) int {
	count, err := ieq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ieq *IdentifiersEntryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ieq.ctx, ent.OpQueryExist)
	switch _, err := ieq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ieq *IdentifiersEntryQuery) ExistX(ctx context.Context) bool {
	exist, err := ieq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the IdentifiersEntryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ieq *IdentifiersEntryQuery) Clone() *IdentifiersEntryQuery {
	if ieq == nil {
		return nil
	}
	return &IdentifiersEntryQuery{
		config:       ieq.config,
		ctx:          ieq.ctx.Clone(),
		order:        append([]identifiersentry.OrderOption{}, ieq.order...),
		inters:       append([]Interceptor{}, ieq.inters...),
		predicates:   append([]predicate.IdentifiersEntry{}, ieq.predicates...),
		withDocument: ieq.withDocument.Clone(),
		withNodes:    ieq.withNodes.Clone(),
		// clone intermediate query.
		sql:  ieq.sql.Clone(),
		path: ieq.path,
	}
}

// WithDocument tells the query-builder to eager-load the nodes that are connected to
// the "document" edge. The optional arguments are used to configure the query builder of the edge.
func (ieq *IdentifiersEntryQuery) WithDocument(opts ...func(*DocumentQuery)) *IdentifiersEntryQuery {
	query := (&DocumentClient{config: ieq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ieq.withDocument = query
	return ieq
}

// WithNodes tells the query-builder to eager-load the nodes that are connected to
// the "nodes" edge. The optional arguments are used to configure the query builder of the edge.
func (ieq *IdentifiersEntryQuery) WithNodes(opts ...func(*NodeQuery)) *IdentifiersEntryQuery {
	query := (&NodeClient{config: ieq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ieq.withNodes = query
	return ieq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DocumentID uuid.UUID `json:"-"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.IdentifiersEntry.Query().
//		GroupBy(identifiersentry.FieldDocumentID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ieq *IdentifiersEntryQuery) GroupBy(field string, fields ...string) *IdentifiersEntryGroupBy {
	ieq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &IdentifiersEntryGroupBy{build: ieq}
	grbuild.flds = &ieq.ctx.Fields
	grbuild.label = identifiersentry.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DocumentID uuid.UUID `json:"-"`
//	}
//
//	client.IdentifiersEntry.Query().
//		Select(identifiersentry.FieldDocumentID).
//		Scan(ctx, &v)
func (ieq *IdentifiersEntryQuery) Select(fields ...string) *IdentifiersEntrySelect {
	ieq.ctx.Fields = append(ieq.ctx.Fields, fields...)
	sbuild := &IdentifiersEntrySelect{IdentifiersEntryQuery: ieq}
	sbuild.label = identifiersentry.Label
	sbuild.flds, sbuild.scan = &ieq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a IdentifiersEntrySelect configured with the given aggregations.
func (ieq *IdentifiersEntryQuery) Aggregate(fns ...AggregateFunc) *IdentifiersEntrySelect {
	return ieq.Select().Aggregate(fns...)
}

func (ieq *IdentifiersEntryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ieq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ieq); err != nil {
				return err
			}
		}
	}
	for _, f := range ieq.ctx.Fields {
		if !identifiersentry.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ieq.path != nil {
		prev, err := ieq.path(ctx)
		if err != nil {
			return err
		}
		ieq.sql = prev
	}
	return nil
}

func (ieq *IdentifiersEntryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*IdentifiersEntry, error) {
	var (
		nodes       = []*IdentifiersEntry{}
		_spec       = ieq.querySpec()
		loadedTypes = [2]bool{
			ieq.withDocument != nil,
			ieq.withNodes != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*IdentifiersEntry).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &IdentifiersEntry{config: ieq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ieq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ieq.withDocument; query != nil {
		if err := ieq.loadDocument(ctx, query, nodes, nil,
			func(n *IdentifiersEntry, e *Document) { n.Edges.Document = e }); err != nil {
			return nil, err
		}
	}
	if query := ieq.withNodes; query != nil {
		if err := ieq.loadNodes(ctx, query, nodes,
			func(n *IdentifiersEntry) { n.Edges.Nodes = []*Node{} },
			func(n *IdentifiersEntry, e *Node) { n.Edges.Nodes = append(n.Edges.Nodes, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ieq *IdentifiersEntryQuery) loadDocument(ctx context.Context, query *DocumentQuery, nodes []*IdentifiersEntry, init func(*IdentifiersEntry), assign func(*IdentifiersEntry, *Document)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*IdentifiersEntry)
	for i := range nodes {
		fk := nodes[i].DocumentID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(document.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "document_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ieq *IdentifiersEntryQuery) loadNodes(ctx context.Context, query *NodeQuery, nodes []*IdentifiersEntry, init func(*IdentifiersEntry), assign func(*IdentifiersEntry, *Node)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*IdentifiersEntry)
	nids := make(map[uuid.UUID]map[*IdentifiersEntry]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(identifiersentry.NodesTable)
		s.Join(joinT).On(s.C(node.FieldID), joinT.C(identifiersentry.NodesPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(identifiersentry.NodesPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(identifiersentry.NodesPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*IdentifiersEntry]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Node](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "nodes" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (ieq *IdentifiersEntryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ieq.querySpec()
	_spec.Node.Columns = ieq.ctx.Fields
	if len(ieq.ctx.Fields) > 0 {
		_spec.Unique = ieq.ctx.Unique != nil && *ieq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ieq.driver, _spec)
}

func (ieq *IdentifiersEntryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(identifiersentry.Table, identifiersentry.Columns, sqlgraph.NewFieldSpec(identifiersentry.FieldID, field.TypeUUID))
	_spec.From = ieq.sql
	if unique := ieq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ieq.path != nil {
		_spec.Unique = true
	}
	if fields := ieq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, identifiersentry.FieldID)
		for i := range fields {
			if fields[i] != identifiersentry.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ieq.withDocument != nil {
			_spec.Node.AddColumnOnce(identifiersentry.FieldDocumentID)
		}
	}
	if ps := ieq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ieq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ieq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ieq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ieq *IdentifiersEntryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ieq.driver.Dialect())
	t1 := builder.Table(identifiersentry.Table)
	columns := ieq.ctx.Fields
	if len(columns) == 0 {
		columns = identifiersentry.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ieq.sql != nil {
		selector = ieq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ieq.ctx.Unique != nil && *ieq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ieq.predicates {
		p(selector)
	}
	for _, p := range ieq.order {
		p(selector)
	}
	if offset := ieq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ieq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// IdentifiersEntryGroupBy is the group-by builder for IdentifiersEntry entities.
type IdentifiersEntryGroupBy struct {
	selector
	build *IdentifiersEntryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (iegb *IdentifiersEntryGroupBy) Aggregate(fns ...AggregateFunc) *IdentifiersEntryGroupBy {
	iegb.fns = append(iegb.fns, fns...)
	return iegb
}

// Scan applies the selector query and scans the result into the given value.
func (iegb *IdentifiersEntryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, iegb.build.ctx, ent.OpQueryGroupBy)
	if err := iegb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IdentifiersEntryQuery, *IdentifiersEntryGroupBy](ctx, iegb.build, iegb, iegb.build.inters, v)
}

func (iegb *IdentifiersEntryGroupBy) sqlScan(ctx context.Context, root *IdentifiersEntryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(iegb.fns))
	for _, fn := range iegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*iegb.flds)+len(iegb.fns))
		for _, f := range *iegb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*iegb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := iegb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// IdentifiersEntrySelect is the builder for selecting fields of IdentifiersEntry entities.
type IdentifiersEntrySelect struct {
	*IdentifiersEntryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ies *IdentifiersEntrySelect) Aggregate(fns ...AggregateFunc) *IdentifiersEntrySelect {
	ies.fns = append(ies.fns, fns...)
	return ies
}

// Scan applies the selector query and scans the result into the given value.
func (ies *IdentifiersEntrySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ies.ctx, ent.OpQuerySelect)
	if err := ies.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IdentifiersEntryQuery, *IdentifiersEntrySelect](ctx, ies.IdentifiersEntryQuery, ies, ies.inters, v)
}

func (ies *IdentifiersEntrySelect) sqlScan(ctx context.Context, root *IdentifiersEntryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ies.fns))
	for _, fn := range ies.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ies.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ies.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
