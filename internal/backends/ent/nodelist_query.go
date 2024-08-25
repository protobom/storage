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
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/nodelist"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// NodeListQuery is the builder for querying NodeList entities.
type NodeListQuery struct {
	config
	ctx          *QueryContext
	order        []nodelist.OrderOption
	inters       []Interceptor
	predicates   []predicate.NodeList
	withNodes    *NodeQuery
	withDocument *DocumentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NodeListQuery builder.
func (nlq *NodeListQuery) Where(ps ...predicate.NodeList) *NodeListQuery {
	nlq.predicates = append(nlq.predicates, ps...)
	return nlq
}

// Limit the number of records to be returned by this query.
func (nlq *NodeListQuery) Limit(limit int) *NodeListQuery {
	nlq.ctx.Limit = &limit
	return nlq
}

// Offset to start from.
func (nlq *NodeListQuery) Offset(offset int) *NodeListQuery {
	nlq.ctx.Offset = &offset
	return nlq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nlq *NodeListQuery) Unique(unique bool) *NodeListQuery {
	nlq.ctx.Unique = &unique
	return nlq
}

// Order specifies how the records should be ordered.
func (nlq *NodeListQuery) Order(o ...nodelist.OrderOption) *NodeListQuery {
	nlq.order = append(nlq.order, o...)
	return nlq
}

// QueryNodes chains the current query on the "nodes" edge.
func (nlq *NodeListQuery) QueryNodes() *NodeQuery {
	query := (&NodeClient{config: nlq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nlq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(nodelist.Table, nodelist.FieldID, selector),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, nodelist.NodesTable, nodelist.NodesColumn),
		)
		fromU = sqlgraph.SetNeighbors(nlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDocument chains the current query on the "document" edge.
func (nlq *NodeListQuery) QueryDocument() *DocumentQuery {
	query := (&DocumentClient{config: nlq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nlq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(nodelist.Table, nodelist.FieldID, selector),
			sqlgraph.To(document.Table, document.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, nodelist.DocumentTable, nodelist.DocumentColumn),
		)
		fromU = sqlgraph.SetNeighbors(nlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first NodeList entity from the query.
// Returns a *NotFoundError when no NodeList was found.
func (nlq *NodeListQuery) First(ctx context.Context) (*NodeList, error) {
	nodes, err := nlq.Limit(1).All(setContextOp(ctx, nlq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{nodelist.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nlq *NodeListQuery) FirstX(ctx context.Context) *NodeList {
	node, err := nlq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NodeList ID from the query.
// Returns a *NotFoundError when no NodeList ID was found.
func (nlq *NodeListQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = nlq.Limit(1).IDs(setContextOp(ctx, nlq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{nodelist.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nlq *NodeListQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := nlq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NodeList entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one NodeList entity is found.
// Returns a *NotFoundError when no NodeList entities are found.
func (nlq *NodeListQuery) Only(ctx context.Context) (*NodeList, error) {
	nodes, err := nlq.Limit(2).All(setContextOp(ctx, nlq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{nodelist.Label}
	default:
		return nil, &NotSingularError{nodelist.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nlq *NodeListQuery) OnlyX(ctx context.Context) *NodeList {
	node, err := nlq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NodeList ID in the query.
// Returns a *NotSingularError when more than one NodeList ID is found.
// Returns a *NotFoundError when no entities are found.
func (nlq *NodeListQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = nlq.Limit(2).IDs(setContextOp(ctx, nlq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{nodelist.Label}
	default:
		err = &NotSingularError{nodelist.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nlq *NodeListQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := nlq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NodeLists.
func (nlq *NodeListQuery) All(ctx context.Context) ([]*NodeList, error) {
	ctx = setContextOp(ctx, nlq.ctx, ent.OpQueryAll)
	if err := nlq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*NodeList, *NodeListQuery]()
	return withInterceptors[[]*NodeList](ctx, nlq, qr, nlq.inters)
}

// AllX is like All, but panics if an error occurs.
func (nlq *NodeListQuery) AllX(ctx context.Context) []*NodeList {
	nodes, err := nlq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NodeList IDs.
func (nlq *NodeListQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if nlq.ctx.Unique == nil && nlq.path != nil {
		nlq.Unique(true)
	}
	ctx = setContextOp(ctx, nlq.ctx, ent.OpQueryIDs)
	if err = nlq.Select(nodelist.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nlq *NodeListQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := nlq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nlq *NodeListQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, nlq.ctx, ent.OpQueryCount)
	if err := nlq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, nlq, querierCount[*NodeListQuery](), nlq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (nlq *NodeListQuery) CountX(ctx context.Context) int {
	count, err := nlq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nlq *NodeListQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, nlq.ctx, ent.OpQueryExist)
	switch _, err := nlq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (nlq *NodeListQuery) ExistX(ctx context.Context) bool {
	exist, err := nlq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NodeListQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nlq *NodeListQuery) Clone() *NodeListQuery {
	if nlq == nil {
		return nil
	}
	return &NodeListQuery{
		config:       nlq.config,
		ctx:          nlq.ctx.Clone(),
		order:        append([]nodelist.OrderOption{}, nlq.order...),
		inters:       append([]Interceptor{}, nlq.inters...),
		predicates:   append([]predicate.NodeList{}, nlq.predicates...),
		withNodes:    nlq.withNodes.Clone(),
		withDocument: nlq.withDocument.Clone(),
		// clone intermediate query.
		sql:  nlq.sql.Clone(),
		path: nlq.path,
	}
}

// WithNodes tells the query-builder to eager-load the nodes that are connected to
// the "nodes" edge. The optional arguments are used to configure the query builder of the edge.
func (nlq *NodeListQuery) WithNodes(opts ...func(*NodeQuery)) *NodeListQuery {
	query := (&NodeClient{config: nlq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nlq.withNodes = query
	return nlq
}

// WithDocument tells the query-builder to eager-load the nodes that are connected to
// the "document" edge. The optional arguments are used to configure the query builder of the edge.
func (nlq *NodeListQuery) WithDocument(opts ...func(*DocumentQuery)) *NodeListQuery {
	query := (&DocumentClient{config: nlq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nlq.withDocument = query
	return nlq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ProtoMessage *sbom.NodeList `json:"proto_message,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.NodeList.Query().
//		GroupBy(nodelist.FieldProtoMessage).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (nlq *NodeListQuery) GroupBy(field string, fields ...string) *NodeListGroupBy {
	nlq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NodeListGroupBy{build: nlq}
	grbuild.flds = &nlq.ctx.Fields
	grbuild.label = nodelist.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ProtoMessage *sbom.NodeList `json:"proto_message,omitempty"`
//	}
//
//	client.NodeList.Query().
//		Select(nodelist.FieldProtoMessage).
//		Scan(ctx, &v)
func (nlq *NodeListQuery) Select(fields ...string) *NodeListSelect {
	nlq.ctx.Fields = append(nlq.ctx.Fields, fields...)
	sbuild := &NodeListSelect{NodeListQuery: nlq}
	sbuild.label = nodelist.Label
	sbuild.flds, sbuild.scan = &nlq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NodeListSelect configured with the given aggregations.
func (nlq *NodeListQuery) Aggregate(fns ...AggregateFunc) *NodeListSelect {
	return nlq.Select().Aggregate(fns...)
}

func (nlq *NodeListQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range nlq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, nlq); err != nil {
				return err
			}
		}
	}
	for _, f := range nlq.ctx.Fields {
		if !nodelist.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nlq.path != nil {
		prev, err := nlq.path(ctx)
		if err != nil {
			return err
		}
		nlq.sql = prev
	}
	return nil
}

func (nlq *NodeListQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*NodeList, error) {
	var (
		nodes       = []*NodeList{}
		_spec       = nlq.querySpec()
		loadedTypes = [2]bool{
			nlq.withNodes != nil,
			nlq.withDocument != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*NodeList).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &NodeList{config: nlq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nlq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := nlq.withNodes; query != nil {
		if err := nlq.loadNodes(ctx, query, nodes,
			func(n *NodeList) { n.Edges.Nodes = []*Node{} },
			func(n *NodeList, e *Node) { n.Edges.Nodes = append(n.Edges.Nodes, e) }); err != nil {
			return nil, err
		}
	}
	if query := nlq.withDocument; query != nil {
		if err := nlq.loadDocument(ctx, query, nodes, nil,
			func(n *NodeList, e *Document) { n.Edges.Document = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (nlq *NodeListQuery) loadNodes(ctx context.Context, query *NodeQuery, nodes []*NodeList, init func(*NodeList), assign func(*NodeList, *Node)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*NodeList)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(node.FieldNodeListID)
	}
	query.Where(predicate.Node(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(nodelist.NodesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.NodeListID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "node_list_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (nlq *NodeListQuery) loadDocument(ctx context.Context, query *DocumentQuery, nodes []*NodeList, init func(*NodeList), assign func(*NodeList, *Document)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*NodeList)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(document.FieldNodeListID)
	}
	query.Where(predicate.Document(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(nodelist.DocumentColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.NodeListID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "node_list_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (nlq *NodeListQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nlq.querySpec()
	_spec.Node.Columns = nlq.ctx.Fields
	if len(nlq.ctx.Fields) > 0 {
		_spec.Unique = nlq.ctx.Unique != nil && *nlq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, nlq.driver, _spec)
}

func (nlq *NodeListQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(nodelist.Table, nodelist.Columns, sqlgraph.NewFieldSpec(nodelist.FieldID, field.TypeUUID))
	_spec.From = nlq.sql
	if unique := nlq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if nlq.path != nil {
		_spec.Unique = true
	}
	if fields := nlq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nodelist.FieldID)
		for i := range fields {
			if fields[i] != nodelist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nlq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nlq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nlq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nlq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nlq *NodeListQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nlq.driver.Dialect())
	t1 := builder.Table(nodelist.Table)
	columns := nlq.ctx.Fields
	if len(columns) == 0 {
		columns = nodelist.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nlq.sql != nil {
		selector = nlq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nlq.ctx.Unique != nil && *nlq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range nlq.predicates {
		p(selector)
	}
	for _, p := range nlq.order {
		p(selector)
	}
	if offset := nlq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nlq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NodeListGroupBy is the group-by builder for NodeList entities.
type NodeListGroupBy struct {
	selector
	build *NodeListQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (nlgb *NodeListGroupBy) Aggregate(fns ...AggregateFunc) *NodeListGroupBy {
	nlgb.fns = append(nlgb.fns, fns...)
	return nlgb
}

// Scan applies the selector query and scans the result into the given value.
func (nlgb *NodeListGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, nlgb.build.ctx, ent.OpQueryGroupBy)
	if err := nlgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NodeListQuery, *NodeListGroupBy](ctx, nlgb.build, nlgb, nlgb.build.inters, v)
}

func (nlgb *NodeListGroupBy) sqlScan(ctx context.Context, root *NodeListQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(nlgb.fns))
	for _, fn := range nlgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*nlgb.flds)+len(nlgb.fns))
		for _, f := range *nlgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*nlgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nlgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NodeListSelect is the builder for selecting fields of NodeList entities.
type NodeListSelect struct {
	*NodeListQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (nls *NodeListSelect) Aggregate(fns ...AggregateFunc) *NodeListSelect {
	nls.fns = append(nls.fns, fns...)
	return nls
}

// Scan applies the selector query and scans the result into the given value.
func (nls *NodeListSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, nls.ctx, ent.OpQuerySelect)
	if err := nls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NodeListQuery, *NodeListSelect](ctx, nls.NodeListQuery, nls, nls.inters, v)
}

func (nls *NodeListSelect) sqlScan(ctx context.Context, root *NodeListQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(nls.fns))
	for _, fn := range nls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*nls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
