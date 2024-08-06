// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------
package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/predicate"
	"github.com/protobom/storage/internal/backends/ent/purpose"
)

// PurposeQuery is the builder for querying Purpose entities.
type PurposeQuery struct {
	config
	ctx          *QueryContext
	order        []purpose.OrderOption
	inters       []Interceptor
	predicates   []predicate.Purpose
	withDocument *DocumentQuery
	withNode     *NodeQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PurposeQuery builder.
func (pq *PurposeQuery) Where(ps ...predicate.Purpose) *PurposeQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PurposeQuery) Limit(limit int) *PurposeQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PurposeQuery) Offset(offset int) *PurposeQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PurposeQuery) Unique(unique bool) *PurposeQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PurposeQuery) Order(o ...purpose.OrderOption) *PurposeQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryDocument chains the current query on the "document" edge.
func (pq *PurposeQuery) QueryDocument() *DocumentQuery {
	query := (&DocumentClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(purpose.Table, purpose.FieldID, selector),
			sqlgraph.To(document.Table, document.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, purpose.DocumentTable, purpose.DocumentColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNode chains the current query on the "node" edge.
func (pq *PurposeQuery) QueryNode() *NodeQuery {
	query := (&NodeClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(purpose.Table, purpose.FieldID, selector),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, purpose.NodeTable, purpose.NodeColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Purpose entity from the query.
// Returns a *NotFoundError when no Purpose was found.
func (pq *PurposeQuery) First(ctx context.Context) (*Purpose, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{purpose.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PurposeQuery) FirstX(ctx context.Context) *Purpose {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Purpose ID from the query.
// Returns a *NotFoundError when no Purpose ID was found.
func (pq *PurposeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{purpose.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PurposeQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Purpose entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Purpose entity is found.
// Returns a *NotFoundError when no Purpose entities are found.
func (pq *PurposeQuery) Only(ctx context.Context) (*Purpose, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{purpose.Label}
	default:
		return nil, &NotSingularError{purpose.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PurposeQuery) OnlyX(ctx context.Context) *Purpose {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Purpose ID in the query.
// Returns a *NotSingularError when more than one Purpose ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PurposeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{purpose.Label}
	default:
		err = &NotSingularError{purpose.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PurposeQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Purposes.
func (pq *PurposeQuery) All(ctx context.Context) ([]*Purpose, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryAll)
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Purpose, *PurposeQuery]()
	return withInterceptors[[]*Purpose](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PurposeQuery) AllX(ctx context.Context) []*Purpose {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Purpose IDs.
func (pq *PurposeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryIDs)
	if err = pq.Select(purpose.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PurposeQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PurposeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryCount)
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PurposeQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PurposeQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PurposeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryExist)
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PurposeQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PurposeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PurposeQuery) Clone() *PurposeQuery {
	if pq == nil {
		return nil
	}
	return &PurposeQuery{
		config:       pq.config,
		ctx:          pq.ctx.Clone(),
		order:        append([]purpose.OrderOption{}, pq.order...),
		inters:       append([]Interceptor{}, pq.inters...),
		predicates:   append([]predicate.Purpose{}, pq.predicates...),
		withDocument: pq.withDocument.Clone(),
		withNode:     pq.withNode.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithDocument tells the query-builder to eager-load the nodes that are connected to
// the "document" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PurposeQuery) WithDocument(opts ...func(*DocumentQuery)) *PurposeQuery {
	query := (&DocumentClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withDocument = query
	return pq
}

// WithNode tells the query-builder to eager-load the nodes that are connected to
// the "node" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PurposeQuery) WithNode(opts ...func(*NodeQuery)) *PurposeQuery {
	query := (&NodeClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withNode = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		NodeID string `json:"node_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Purpose.Query().
//		GroupBy(purpose.FieldNodeID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PurposeQuery) GroupBy(field string, fields ...string) *PurposeGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PurposeGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = purpose.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		NodeID string `json:"node_id,omitempty"`
//	}
//
//	client.Purpose.Query().
//		Select(purpose.FieldNodeID).
//		Scan(ctx, &v)
func (pq *PurposeQuery) Select(fields ...string) *PurposeSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PurposeSelect{PurposeQuery: pq}
	sbuild.label = purpose.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PurposeSelect configured with the given aggregations.
func (pq *PurposeQuery) Aggregate(fns ...AggregateFunc) *PurposeSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PurposeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !purpose.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PurposeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Purpose, error) {
	var (
		nodes       = []*Purpose{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [2]bool{
			pq.withDocument != nil,
			pq.withNode != nil,
		}
	)
	if pq.withDocument != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, purpose.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Purpose).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Purpose{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withDocument; query != nil {
		if err := pq.loadDocument(ctx, query, nodes, nil,
			func(n *Purpose, e *Document) { n.Edges.Document = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withNode; query != nil {
		if err := pq.loadNode(ctx, query, nodes, nil,
			func(n *Purpose, e *Node) { n.Edges.Node = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PurposeQuery) loadDocument(ctx context.Context, query *DocumentQuery, nodes []*Purpose, init func(*Purpose), assign func(*Purpose, *Document)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Purpose)
	for i := range nodes {
		if nodes[i].document_id == nil {
			continue
		}
		fk := *nodes[i].document_id
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
func (pq *PurposeQuery) loadNode(ctx context.Context, query *NodeQuery, nodes []*Purpose, init func(*Purpose), assign func(*Purpose, *Node)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Purpose)
	for i := range nodes {
		fk := nodes[i].NodeID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(node.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "node_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pq *PurposeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PurposeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(purpose.Table, purpose.Columns, sqlgraph.NewFieldSpec(purpose.FieldID, field.TypeInt))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, purpose.FieldID)
		for i := range fields {
			if fields[i] != purpose.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if pq.withNode != nil {
			_spec.Node.AddColumnOnce(purpose.FieldNodeID)
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PurposeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(purpose.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = purpose.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PurposeGroupBy is the group-by builder for Purpose entities.
type PurposeGroupBy struct {
	selector
	build *PurposeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PurposeGroupBy) Aggregate(fns ...AggregateFunc) *PurposeGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PurposeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, ent.OpQueryGroupBy)
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PurposeQuery, *PurposeGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PurposeGroupBy) sqlScan(ctx context.Context, root *PurposeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PurposeSelect is the builder for selecting fields of Purpose entities.
type PurposeSelect struct {
	*PurposeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PurposeSelect) Aggregate(fns ...AggregateFunc) *PurposeSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PurposeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, ent.OpQuerySelect)
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PurposeQuery, *PurposeSelect](ctx, ps.PurposeQuery, ps, ps.inters, v)
}

func (ps *PurposeSelect) sqlScan(ctx context.Context, root *PurposeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
