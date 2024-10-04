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
	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/externalreference"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// ExternalReferenceQuery is the builder for querying ExternalReference entities.
type ExternalReferenceQuery struct {
	config
	ctx          *QueryContext
	order        []externalreference.OrderOption
	inters       []Interceptor
	predicates   []predicate.ExternalReference
	withDocument *DocumentQuery
	withNode     *NodeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ExternalReferenceQuery builder.
func (erq *ExternalReferenceQuery) Where(ps ...predicate.ExternalReference) *ExternalReferenceQuery {
	erq.predicates = append(erq.predicates, ps...)
	return erq
}

// Limit the number of records to be returned by this query.
func (erq *ExternalReferenceQuery) Limit(limit int) *ExternalReferenceQuery {
	erq.ctx.Limit = &limit
	return erq
}

// Offset to start from.
func (erq *ExternalReferenceQuery) Offset(offset int) *ExternalReferenceQuery {
	erq.ctx.Offset = &offset
	return erq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (erq *ExternalReferenceQuery) Unique(unique bool) *ExternalReferenceQuery {
	erq.ctx.Unique = &unique
	return erq
}

// Order specifies how the records should be ordered.
func (erq *ExternalReferenceQuery) Order(o ...externalreference.OrderOption) *ExternalReferenceQuery {
	erq.order = append(erq.order, o...)
	return erq
}

// QueryDocument chains the current query on the "document" edge.
func (erq *ExternalReferenceQuery) QueryDocument() *DocumentQuery {
	query := (&DocumentClient{config: erq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := erq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := erq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(externalreference.Table, externalreference.FieldID, selector),
			sqlgraph.To(document.Table, document.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, externalreference.DocumentTable, externalreference.DocumentColumn),
		)
		fromU = sqlgraph.SetNeighbors(erq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNode chains the current query on the "node" edge.
func (erq *ExternalReferenceQuery) QueryNode() *NodeQuery {
	query := (&NodeClient{config: erq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := erq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := erq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(externalreference.Table, externalreference.FieldID, selector),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, externalreference.NodeTable, externalreference.NodeColumn),
		)
		fromU = sqlgraph.SetNeighbors(erq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ExternalReference entity from the query.
// Returns a *NotFoundError when no ExternalReference was found.
func (erq *ExternalReferenceQuery) First(ctx context.Context) (*ExternalReference, error) {
	nodes, err := erq.Limit(1).All(setContextOp(ctx, erq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{externalreference.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (erq *ExternalReferenceQuery) FirstX(ctx context.Context) *ExternalReference {
	node, err := erq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ExternalReference ID from the query.
// Returns a *NotFoundError when no ExternalReference ID was found.
func (erq *ExternalReferenceQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = erq.Limit(1).IDs(setContextOp(ctx, erq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{externalreference.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (erq *ExternalReferenceQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := erq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ExternalReference entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ExternalReference entity is found.
// Returns a *NotFoundError when no ExternalReference entities are found.
func (erq *ExternalReferenceQuery) Only(ctx context.Context) (*ExternalReference, error) {
	nodes, err := erq.Limit(2).All(setContextOp(ctx, erq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{externalreference.Label}
	default:
		return nil, &NotSingularError{externalreference.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (erq *ExternalReferenceQuery) OnlyX(ctx context.Context) *ExternalReference {
	node, err := erq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ExternalReference ID in the query.
// Returns a *NotSingularError when more than one ExternalReference ID is found.
// Returns a *NotFoundError when no entities are found.
func (erq *ExternalReferenceQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = erq.Limit(2).IDs(setContextOp(ctx, erq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{externalreference.Label}
	default:
		err = &NotSingularError{externalreference.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (erq *ExternalReferenceQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := erq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ExternalReferences.
func (erq *ExternalReferenceQuery) All(ctx context.Context) ([]*ExternalReference, error) {
	ctx = setContextOp(ctx, erq.ctx, ent.OpQueryAll)
	if err := erq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ExternalReference, *ExternalReferenceQuery]()
	return withInterceptors[[]*ExternalReference](ctx, erq, qr, erq.inters)
}

// AllX is like All, but panics if an error occurs.
func (erq *ExternalReferenceQuery) AllX(ctx context.Context) []*ExternalReference {
	nodes, err := erq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ExternalReference IDs.
func (erq *ExternalReferenceQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if erq.ctx.Unique == nil && erq.path != nil {
		erq.Unique(true)
	}
	ctx = setContextOp(ctx, erq.ctx, ent.OpQueryIDs)
	if err = erq.Select(externalreference.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (erq *ExternalReferenceQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := erq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (erq *ExternalReferenceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, erq.ctx, ent.OpQueryCount)
	if err := erq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, erq, querierCount[*ExternalReferenceQuery](), erq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (erq *ExternalReferenceQuery) CountX(ctx context.Context) int {
	count, err := erq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (erq *ExternalReferenceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, erq.ctx, ent.OpQueryExist)
	switch _, err := erq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (erq *ExternalReferenceQuery) ExistX(ctx context.Context) bool {
	exist, err := erq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ExternalReferenceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (erq *ExternalReferenceQuery) Clone() *ExternalReferenceQuery {
	if erq == nil {
		return nil
	}
	return &ExternalReferenceQuery{
		config:       erq.config,
		ctx:          erq.ctx.Clone(),
		order:        append([]externalreference.OrderOption{}, erq.order...),
		inters:       append([]Interceptor{}, erq.inters...),
		predicates:   append([]predicate.ExternalReference{}, erq.predicates...),
		withDocument: erq.withDocument.Clone(),
		withNode:     erq.withNode.Clone(),
		// clone intermediate query.
		sql:  erq.sql.Clone(),
		path: erq.path,
	}
}

// WithDocument tells the query-builder to eager-load the nodes that are connected to
// the "document" edge. The optional arguments are used to configure the query builder of the edge.
func (erq *ExternalReferenceQuery) WithDocument(opts ...func(*DocumentQuery)) *ExternalReferenceQuery {
	query := (&DocumentClient{config: erq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	erq.withDocument = query
	return erq
}

// WithNode tells the query-builder to eager-load the nodes that are connected to
// the "node" edge. The optional arguments are used to configure the query builder of the edge.
func (erq *ExternalReferenceQuery) WithNode(opts ...func(*NodeQuery)) *ExternalReferenceQuery {
	query := (&NodeClient{config: erq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	erq.withNode = query
	return erq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DocumentID uuid.UUID `json:"document_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ExternalReference.Query().
//		GroupBy(externalreference.FieldDocumentID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (erq *ExternalReferenceQuery) GroupBy(field string, fields ...string) *ExternalReferenceGroupBy {
	erq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ExternalReferenceGroupBy{build: erq}
	grbuild.flds = &erq.ctx.Fields
	grbuild.label = externalreference.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DocumentID uuid.UUID `json:"document_id,omitempty"`
//	}
//
//	client.ExternalReference.Query().
//		Select(externalreference.FieldDocumentID).
//		Scan(ctx, &v)
func (erq *ExternalReferenceQuery) Select(fields ...string) *ExternalReferenceSelect {
	erq.ctx.Fields = append(erq.ctx.Fields, fields...)
	sbuild := &ExternalReferenceSelect{ExternalReferenceQuery: erq}
	sbuild.label = externalreference.Label
	sbuild.flds, sbuild.scan = &erq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ExternalReferenceSelect configured with the given aggregations.
func (erq *ExternalReferenceQuery) Aggregate(fns ...AggregateFunc) *ExternalReferenceSelect {
	return erq.Select().Aggregate(fns...)
}

func (erq *ExternalReferenceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range erq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, erq); err != nil {
				return err
			}
		}
	}
	for _, f := range erq.ctx.Fields {
		if !externalreference.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if erq.path != nil {
		prev, err := erq.path(ctx)
		if err != nil {
			return err
		}
		erq.sql = prev
	}
	return nil
}

func (erq *ExternalReferenceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ExternalReference, error) {
	var (
		nodes       = []*ExternalReference{}
		_spec       = erq.querySpec()
		loadedTypes = [2]bool{
			erq.withDocument != nil,
			erq.withNode != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ExternalReference).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ExternalReference{config: erq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, erq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := erq.withDocument; query != nil {
		if err := erq.loadDocument(ctx, query, nodes, nil,
			func(n *ExternalReference, e *Document) { n.Edges.Document = e }); err != nil {
			return nil, err
		}
	}
	if query := erq.withNode; query != nil {
		if err := erq.loadNode(ctx, query, nodes, nil,
			func(n *ExternalReference, e *Node) { n.Edges.Node = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (erq *ExternalReferenceQuery) loadDocument(ctx context.Context, query *DocumentQuery, nodes []*ExternalReference, init func(*ExternalReference), assign func(*ExternalReference, *Document)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ExternalReference)
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
func (erq *ExternalReferenceQuery) loadNode(ctx context.Context, query *NodeQuery, nodes []*ExternalReference, init func(*ExternalReference), assign func(*ExternalReference, *Node)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*ExternalReference)
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

func (erq *ExternalReferenceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := erq.querySpec()
	_spec.Node.Columns = erq.ctx.Fields
	if len(erq.ctx.Fields) > 0 {
		_spec.Unique = erq.ctx.Unique != nil && *erq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, erq.driver, _spec)
}

func (erq *ExternalReferenceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(externalreference.Table, externalreference.Columns, sqlgraph.NewFieldSpec(externalreference.FieldID, field.TypeUUID))
	_spec.From = erq.sql
	if unique := erq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if erq.path != nil {
		_spec.Unique = true
	}
	if fields := erq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, externalreference.FieldID)
		for i := range fields {
			if fields[i] != externalreference.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if erq.withDocument != nil {
			_spec.Node.AddColumnOnce(externalreference.FieldDocumentID)
		}
		if erq.withNode != nil {
			_spec.Node.AddColumnOnce(externalreference.FieldNodeID)
		}
	}
	if ps := erq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := erq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := erq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := erq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (erq *ExternalReferenceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(erq.driver.Dialect())
	t1 := builder.Table(externalreference.Table)
	columns := erq.ctx.Fields
	if len(columns) == 0 {
		columns = externalreference.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if erq.sql != nil {
		selector = erq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if erq.ctx.Unique != nil && *erq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range erq.predicates {
		p(selector)
	}
	for _, p := range erq.order {
		p(selector)
	}
	if offset := erq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := erq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ExternalReferenceGroupBy is the group-by builder for ExternalReference entities.
type ExternalReferenceGroupBy struct {
	selector
	build *ExternalReferenceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ergb *ExternalReferenceGroupBy) Aggregate(fns ...AggregateFunc) *ExternalReferenceGroupBy {
	ergb.fns = append(ergb.fns, fns...)
	return ergb
}

// Scan applies the selector query and scans the result into the given value.
func (ergb *ExternalReferenceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ergb.build.ctx, ent.OpQueryGroupBy)
	if err := ergb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExternalReferenceQuery, *ExternalReferenceGroupBy](ctx, ergb.build, ergb, ergb.build.inters, v)
}

func (ergb *ExternalReferenceGroupBy) sqlScan(ctx context.Context, root *ExternalReferenceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ergb.fns))
	for _, fn := range ergb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ergb.flds)+len(ergb.fns))
		for _, f := range *ergb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ergb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ergb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ExternalReferenceSelect is the builder for selecting fields of ExternalReference entities.
type ExternalReferenceSelect struct {
	*ExternalReferenceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ers *ExternalReferenceSelect) Aggregate(fns ...AggregateFunc) *ExternalReferenceSelect {
	ers.fns = append(ers.fns, fns...)
	return ers
}

// Scan applies the selector query and scans the result into the given value.
func (ers *ExternalReferenceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ers.ctx, ent.OpQuerySelect)
	if err := ers.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExternalReferenceQuery, *ExternalReferenceSelect](ctx, ers.ExternalReferenceQuery, ers, ers.inters, v)
}

func (ers *ExternalReferenceSelect) sqlScan(ctx context.Context, root *ExternalReferenceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ers.fns))
	for _, fn := range ers.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ers.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ers.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
