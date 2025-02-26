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
	"github.com/google/uuid"
	"github.com/protobom/storage/internal/backends/ent/annotation"
	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// AnnotationQuery is the builder for querying Annotation entities.
type AnnotationQuery struct {
	config
	ctx          *QueryContext
	order        []annotation.OrderOption
	inters       []Interceptor
	predicates   []predicate.Annotation
	withDocument *DocumentQuery
	withNode     *NodeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AnnotationQuery builder.
func (aq *AnnotationQuery) Where(ps ...predicate.Annotation) *AnnotationQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *AnnotationQuery) Limit(limit int) *AnnotationQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *AnnotationQuery) Offset(offset int) *AnnotationQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AnnotationQuery) Unique(unique bool) *AnnotationQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *AnnotationQuery) Order(o ...annotation.OrderOption) *AnnotationQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryDocument chains the current query on the "document" edge.
func (aq *AnnotationQuery) QueryDocument() *DocumentQuery {
	query := (&DocumentClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(annotation.Table, annotation.FieldID, selector),
			sqlgraph.To(document.Table, document.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, annotation.DocumentTable, annotation.DocumentColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNode chains the current query on the "node" edge.
func (aq *AnnotationQuery) QueryNode() *NodeQuery {
	query := (&NodeClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(annotation.Table, annotation.FieldID, selector),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, annotation.NodeTable, annotation.NodeColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Annotation entity from the query.
// Returns a *NotFoundError when no Annotation was found.
func (aq *AnnotationQuery) First(ctx context.Context) (*Annotation, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{annotation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AnnotationQuery) FirstX(ctx context.Context) *Annotation {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Annotation ID from the query.
// Returns a *NotFoundError when no Annotation ID was found.
func (aq *AnnotationQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{annotation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AnnotationQuery) FirstIDX(ctx context.Context) int {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Annotation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Annotation entity is found.
// Returns a *NotFoundError when no Annotation entities are found.
func (aq *AnnotationQuery) Only(ctx context.Context) (*Annotation, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{annotation.Label}
	default:
		return nil, &NotSingularError{annotation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AnnotationQuery) OnlyX(ctx context.Context) *Annotation {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Annotation ID in the query.
// Returns a *NotSingularError when more than one Annotation ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AnnotationQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{annotation.Label}
	default:
		err = &NotSingularError{annotation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AnnotationQuery) OnlyIDX(ctx context.Context) int {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Annotations.
func (aq *AnnotationQuery) All(ctx context.Context) ([]*Annotation, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryAll)
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Annotation, *AnnotationQuery]()
	return withInterceptors[[]*Annotation](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *AnnotationQuery) AllX(ctx context.Context) []*Annotation {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Annotation IDs.
func (aq *AnnotationQuery) IDs(ctx context.Context) (ids []int, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryIDs)
	if err = aq.Select(annotation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AnnotationQuery) IDsX(ctx context.Context) []int {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AnnotationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryCount)
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*AnnotationQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AnnotationQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AnnotationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryExist)
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AnnotationQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AnnotationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AnnotationQuery) Clone() *AnnotationQuery {
	if aq == nil {
		return nil
	}
	return &AnnotationQuery{
		config:       aq.config,
		ctx:          aq.ctx.Clone(),
		order:        append([]annotation.OrderOption{}, aq.order...),
		inters:       append([]Interceptor{}, aq.inters...),
		predicates:   append([]predicate.Annotation{}, aq.predicates...),
		withDocument: aq.withDocument.Clone(),
		withNode:     aq.withNode.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithDocument tells the query-builder to eager-load the nodes that are connected to
// the "document" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AnnotationQuery) WithDocument(opts ...func(*DocumentQuery)) *AnnotationQuery {
	query := (&DocumentClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withDocument = query
	return aq
}

// WithNode tells the query-builder to eager-load the nodes that are connected to
// the "node" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AnnotationQuery) WithNode(opts ...func(*NodeQuery)) *AnnotationQuery {
	query := (&NodeClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withNode = query
	return aq
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
//	client.Annotation.Query().
//		GroupBy(annotation.FieldDocumentID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AnnotationQuery) GroupBy(field string, fields ...string) *AnnotationGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AnnotationGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = annotation.Label
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
//	client.Annotation.Query().
//		Select(annotation.FieldDocumentID).
//		Scan(ctx, &v)
func (aq *AnnotationQuery) Select(fields ...string) *AnnotationSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &AnnotationSelect{AnnotationQuery: aq}
	sbuild.label = annotation.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AnnotationSelect configured with the given aggregations.
func (aq *AnnotationQuery) Aggregate(fns ...AggregateFunc) *AnnotationSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AnnotationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !annotation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AnnotationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Annotation, error) {
	var (
		nodes       = []*Annotation{}
		_spec       = aq.querySpec()
		loadedTypes = [2]bool{
			aq.withDocument != nil,
			aq.withNode != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Annotation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Annotation{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withDocument; query != nil {
		if err := aq.loadDocument(ctx, query, nodes, nil,
			func(n *Annotation, e *Document) { n.Edges.Document = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withNode; query != nil {
		if err := aq.loadNode(ctx, query, nodes, nil,
			func(n *Annotation, e *Node) { n.Edges.Node = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AnnotationQuery) loadDocument(ctx context.Context, query *DocumentQuery, nodes []*Annotation, init func(*Annotation), assign func(*Annotation, *Document)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Annotation)
	for i := range nodes {
		if nodes[i].DocumentID == nil {
			continue
		}
		fk := *nodes[i].DocumentID
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
func (aq *AnnotationQuery) loadNode(ctx context.Context, query *NodeQuery, nodes []*Annotation, init func(*Annotation), assign func(*Annotation, *Node)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Annotation)
	for i := range nodes {
		if nodes[i].NodeID == nil {
			continue
		}
		fk := *nodes[i].NodeID
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

func (aq *AnnotationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AnnotationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(annotation.Table, annotation.Columns, sqlgraph.NewFieldSpec(annotation.FieldID, field.TypeInt))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, annotation.FieldID)
		for i := range fields {
			if fields[i] != annotation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if aq.withDocument != nil {
			_spec.Node.AddColumnOnce(annotation.FieldDocumentID)
		}
		if aq.withNode != nil {
			_spec.Node.AddColumnOnce(annotation.FieldNodeID)
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AnnotationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(annotation.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = annotation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AnnotationGroupBy is the group-by builder for Annotation entities.
type AnnotationGroupBy struct {
	selector
	build *AnnotationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AnnotationGroupBy) Aggregate(fns ...AggregateFunc) *AnnotationGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *AnnotationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, ent.OpQueryGroupBy)
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AnnotationQuery, *AnnotationGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *AnnotationGroupBy) sqlScan(ctx context.Context, root *AnnotationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AnnotationSelect is the builder for selecting fields of Annotation entities.
type AnnotationSelect struct {
	*AnnotationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AnnotationSelect) Aggregate(fns ...AggregateFunc) *AnnotationSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AnnotationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, ent.OpQuerySelect)
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AnnotationQuery, *AnnotationSelect](ctx, as.AnnotationQuery, as, as.inters, v)
}

func (as *AnnotationSelect) sqlScan(ctx context.Context, root *AnnotationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
