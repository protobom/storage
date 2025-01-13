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
	"github.com/protobom/storage/internal/backends/ent/edgetype"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/nodelist"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// EdgeTypeQuery is the builder for querying EdgeType entities.
type EdgeTypeQuery struct {
	config
	ctx           *QueryContext
	order         []edgetype.OrderOption
	inters        []Interceptor
	predicates    []predicate.EdgeType
	withDocument  *DocumentQuery
	withFrom      *NodeQuery
	withTo        *NodeQuery
	withNodeLists *NodeListQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EdgeTypeQuery builder.
func (etq *EdgeTypeQuery) Where(ps ...predicate.EdgeType) *EdgeTypeQuery {
	etq.predicates = append(etq.predicates, ps...)
	return etq
}

// Limit the number of records to be returned by this query.
func (etq *EdgeTypeQuery) Limit(limit int) *EdgeTypeQuery {
	etq.ctx.Limit = &limit
	return etq
}

// Offset to start from.
func (etq *EdgeTypeQuery) Offset(offset int) *EdgeTypeQuery {
	etq.ctx.Offset = &offset
	return etq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (etq *EdgeTypeQuery) Unique(unique bool) *EdgeTypeQuery {
	etq.ctx.Unique = &unique
	return etq
}

// Order specifies how the records should be ordered.
func (etq *EdgeTypeQuery) Order(o ...edgetype.OrderOption) *EdgeTypeQuery {
	etq.order = append(etq.order, o...)
	return etq
}

// QueryDocument chains the current query on the "document" edge.
func (etq *EdgeTypeQuery) QueryDocument() *DocumentQuery {
	query := (&DocumentClient{config: etq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := etq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := etq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(edgetype.Table, edgetype.FieldID, selector),
			sqlgraph.To(document.Table, document.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, edgetype.DocumentTable, edgetype.DocumentColumn),
		)
		fromU = sqlgraph.SetNeighbors(etq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFrom chains the current query on the "from" edge.
func (etq *EdgeTypeQuery) QueryFrom() *NodeQuery {
	query := (&NodeClient{config: etq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := etq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := etq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(edgetype.Table, edgetype.FieldID, selector),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, edgetype.FromTable, edgetype.FromColumn),
		)
		fromU = sqlgraph.SetNeighbors(etq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTo chains the current query on the "to" edge.
func (etq *EdgeTypeQuery) QueryTo() *NodeQuery {
	query := (&NodeClient{config: etq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := etq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := etq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(edgetype.Table, edgetype.FieldID, selector),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, edgetype.ToTable, edgetype.ToColumn),
		)
		fromU = sqlgraph.SetNeighbors(etq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNodeLists chains the current query on the "node_lists" edge.
func (etq *EdgeTypeQuery) QueryNodeLists() *NodeListQuery {
	query := (&NodeListClient{config: etq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := etq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := etq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(edgetype.Table, edgetype.FieldID, selector),
			sqlgraph.To(nodelist.Table, nodelist.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, edgetype.NodeListsTable, edgetype.NodeListsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(etq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EdgeType entity from the query.
// Returns a *NotFoundError when no EdgeType was found.
func (etq *EdgeTypeQuery) First(ctx context.Context) (*EdgeType, error) {
	nodes, err := etq.Limit(1).All(setContextOp(ctx, etq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{edgetype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (etq *EdgeTypeQuery) FirstX(ctx context.Context) *EdgeType {
	node, err := etq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EdgeType ID from the query.
// Returns a *NotFoundError when no EdgeType ID was found.
func (etq *EdgeTypeQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = etq.Limit(1).IDs(setContextOp(ctx, etq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{edgetype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (etq *EdgeTypeQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := etq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EdgeType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EdgeType entity is found.
// Returns a *NotFoundError when no EdgeType entities are found.
func (etq *EdgeTypeQuery) Only(ctx context.Context) (*EdgeType, error) {
	nodes, err := etq.Limit(2).All(setContextOp(ctx, etq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{edgetype.Label}
	default:
		return nil, &NotSingularError{edgetype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (etq *EdgeTypeQuery) OnlyX(ctx context.Context) *EdgeType {
	node, err := etq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EdgeType ID in the query.
// Returns a *NotSingularError when more than one EdgeType ID is found.
// Returns a *NotFoundError when no entities are found.
func (etq *EdgeTypeQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = etq.Limit(2).IDs(setContextOp(ctx, etq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{edgetype.Label}
	default:
		err = &NotSingularError{edgetype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (etq *EdgeTypeQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := etq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EdgeTypes.
func (etq *EdgeTypeQuery) All(ctx context.Context) ([]*EdgeType, error) {
	ctx = setContextOp(ctx, etq.ctx, ent.OpQueryAll)
	if err := etq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EdgeType, *EdgeTypeQuery]()
	return withInterceptors[[]*EdgeType](ctx, etq, qr, etq.inters)
}

// AllX is like All, but panics if an error occurs.
func (etq *EdgeTypeQuery) AllX(ctx context.Context) []*EdgeType {
	nodes, err := etq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EdgeType IDs.
func (etq *EdgeTypeQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if etq.ctx.Unique == nil && etq.path != nil {
		etq.Unique(true)
	}
	ctx = setContextOp(ctx, etq.ctx, ent.OpQueryIDs)
	if err = etq.Select(edgetype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (etq *EdgeTypeQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := etq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (etq *EdgeTypeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, etq.ctx, ent.OpQueryCount)
	if err := etq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, etq, querierCount[*EdgeTypeQuery](), etq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (etq *EdgeTypeQuery) CountX(ctx context.Context) int {
	count, err := etq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (etq *EdgeTypeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, etq.ctx, ent.OpQueryExist)
	switch _, err := etq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (etq *EdgeTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := etq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EdgeTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (etq *EdgeTypeQuery) Clone() *EdgeTypeQuery {
	if etq == nil {
		return nil
	}
	return &EdgeTypeQuery{
		config:        etq.config,
		ctx:           etq.ctx.Clone(),
		order:         append([]edgetype.OrderOption{}, etq.order...),
		inters:        append([]Interceptor{}, etq.inters...),
		predicates:    append([]predicate.EdgeType{}, etq.predicates...),
		withDocument:  etq.withDocument.Clone(),
		withFrom:      etq.withFrom.Clone(),
		withTo:        etq.withTo.Clone(),
		withNodeLists: etq.withNodeLists.Clone(),
		// clone intermediate query.
		sql:  etq.sql.Clone(),
		path: etq.path,
	}
}

// WithDocument tells the query-builder to eager-load the nodes that are connected to
// the "document" edge. The optional arguments are used to configure the query builder of the edge.
func (etq *EdgeTypeQuery) WithDocument(opts ...func(*DocumentQuery)) *EdgeTypeQuery {
	query := (&DocumentClient{config: etq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	etq.withDocument = query
	return etq
}

// WithFrom tells the query-builder to eager-load the nodes that are connected to
// the "from" edge. The optional arguments are used to configure the query builder of the edge.
func (etq *EdgeTypeQuery) WithFrom(opts ...func(*NodeQuery)) *EdgeTypeQuery {
	query := (&NodeClient{config: etq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	etq.withFrom = query
	return etq
}

// WithTo tells the query-builder to eager-load the nodes that are connected to
// the "to" edge. The optional arguments are used to configure the query builder of the edge.
func (etq *EdgeTypeQuery) WithTo(opts ...func(*NodeQuery)) *EdgeTypeQuery {
	query := (&NodeClient{config: etq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	etq.withTo = query
	return etq
}

// WithNodeLists tells the query-builder to eager-load the nodes that are connected to
// the "node_lists" edge. The optional arguments are used to configure the query builder of the edge.
func (etq *EdgeTypeQuery) WithNodeLists(opts ...func(*NodeListQuery)) *EdgeTypeQuery {
	query := (&NodeListClient{config: etq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	etq.withNodeLists = query
	return etq
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
//	client.EdgeType.Query().
//		GroupBy(edgetype.FieldDocumentID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (etq *EdgeTypeQuery) GroupBy(field string, fields ...string) *EdgeTypeGroupBy {
	etq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EdgeTypeGroupBy{build: etq}
	grbuild.flds = &etq.ctx.Fields
	grbuild.label = edgetype.Label
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
//	client.EdgeType.Query().
//		Select(edgetype.FieldDocumentID).
//		Scan(ctx, &v)
func (etq *EdgeTypeQuery) Select(fields ...string) *EdgeTypeSelect {
	etq.ctx.Fields = append(etq.ctx.Fields, fields...)
	sbuild := &EdgeTypeSelect{EdgeTypeQuery: etq}
	sbuild.label = edgetype.Label
	sbuild.flds, sbuild.scan = &etq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EdgeTypeSelect configured with the given aggregations.
func (etq *EdgeTypeQuery) Aggregate(fns ...AggregateFunc) *EdgeTypeSelect {
	return etq.Select().Aggregate(fns...)
}

func (etq *EdgeTypeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range etq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, etq); err != nil {
				return err
			}
		}
	}
	for _, f := range etq.ctx.Fields {
		if !edgetype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if etq.path != nil {
		prev, err := etq.path(ctx)
		if err != nil {
			return err
		}
		etq.sql = prev
	}
	return nil
}

func (etq *EdgeTypeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EdgeType, error) {
	var (
		nodes       = []*EdgeType{}
		_spec       = etq.querySpec()
		loadedTypes = [4]bool{
			etq.withDocument != nil,
			etq.withFrom != nil,
			etq.withTo != nil,
			etq.withNodeLists != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EdgeType).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EdgeType{config: etq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, etq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := etq.withDocument; query != nil {
		if err := etq.loadDocument(ctx, query, nodes, nil,
			func(n *EdgeType, e *Document) { n.Edges.Document = e }); err != nil {
			return nil, err
		}
	}
	if query := etq.withFrom; query != nil {
		if err := etq.loadFrom(ctx, query, nodes, nil,
			func(n *EdgeType, e *Node) { n.Edges.From = e }); err != nil {
			return nil, err
		}
	}
	if query := etq.withTo; query != nil {
		if err := etq.loadTo(ctx, query, nodes, nil,
			func(n *EdgeType, e *Node) { n.Edges.To = e }); err != nil {
			return nil, err
		}
	}
	if query := etq.withNodeLists; query != nil {
		if err := etq.loadNodeLists(ctx, query, nodes,
			func(n *EdgeType) { n.Edges.NodeLists = []*NodeList{} },
			func(n *EdgeType, e *NodeList) { n.Edges.NodeLists = append(n.Edges.NodeLists, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (etq *EdgeTypeQuery) loadDocument(ctx context.Context, query *DocumentQuery, nodes []*EdgeType, init func(*EdgeType), assign func(*EdgeType, *Document)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*EdgeType)
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
func (etq *EdgeTypeQuery) loadFrom(ctx context.Context, query *NodeQuery, nodes []*EdgeType, init func(*EdgeType), assign func(*EdgeType, *Node)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*EdgeType)
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
func (etq *EdgeTypeQuery) loadTo(ctx context.Context, query *NodeQuery, nodes []*EdgeType, init func(*EdgeType), assign func(*EdgeType, *Node)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*EdgeType)
	for i := range nodes {
		fk := nodes[i].ToNodeID
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
			return fmt.Errorf(`unexpected foreign-key "to_node_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (etq *EdgeTypeQuery) loadNodeLists(ctx context.Context, query *NodeListQuery, nodes []*EdgeType, init func(*EdgeType), assign func(*EdgeType, *NodeList)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*EdgeType)
	nids := make(map[uuid.UUID]map[*EdgeType]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(edgetype.NodeListsTable)
		s.Join(joinT).On(s.C(nodelist.FieldID), joinT.C(edgetype.NodeListsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(edgetype.NodeListsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(edgetype.NodeListsPrimaryKey[1]))
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
					nids[inValue] = map[*EdgeType]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*NodeList](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "node_lists" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (etq *EdgeTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := etq.querySpec()
	_spec.Node.Columns = etq.ctx.Fields
	if len(etq.ctx.Fields) > 0 {
		_spec.Unique = etq.ctx.Unique != nil && *etq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, etq.driver, _spec)
}

func (etq *EdgeTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(edgetype.Table, edgetype.Columns, sqlgraph.NewFieldSpec(edgetype.FieldID, field.TypeUUID))
	_spec.From = etq.sql
	if unique := etq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if etq.path != nil {
		_spec.Unique = true
	}
	if fields := etq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, edgetype.FieldID)
		for i := range fields {
			if fields[i] != edgetype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if etq.withDocument != nil {
			_spec.Node.AddColumnOnce(edgetype.FieldDocumentID)
		}
		if etq.withFrom != nil {
			_spec.Node.AddColumnOnce(edgetype.FieldNodeID)
		}
		if etq.withTo != nil {
			_spec.Node.AddColumnOnce(edgetype.FieldToNodeID)
		}
	}
	if ps := etq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := etq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := etq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := etq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (etq *EdgeTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(etq.driver.Dialect())
	t1 := builder.Table(edgetype.Table)
	columns := etq.ctx.Fields
	if len(columns) == 0 {
		columns = edgetype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if etq.sql != nil {
		selector = etq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if etq.ctx.Unique != nil && *etq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range etq.predicates {
		p(selector)
	}
	for _, p := range etq.order {
		p(selector)
	}
	if offset := etq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := etq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EdgeTypeGroupBy is the group-by builder for EdgeType entities.
type EdgeTypeGroupBy struct {
	selector
	build *EdgeTypeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (etgb *EdgeTypeGroupBy) Aggregate(fns ...AggregateFunc) *EdgeTypeGroupBy {
	etgb.fns = append(etgb.fns, fns...)
	return etgb
}

// Scan applies the selector query and scans the result into the given value.
func (etgb *EdgeTypeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, etgb.build.ctx, ent.OpQueryGroupBy)
	if err := etgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EdgeTypeQuery, *EdgeTypeGroupBy](ctx, etgb.build, etgb, etgb.build.inters, v)
}

func (etgb *EdgeTypeGroupBy) sqlScan(ctx context.Context, root *EdgeTypeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(etgb.fns))
	for _, fn := range etgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*etgb.flds)+len(etgb.fns))
		for _, f := range *etgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*etgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := etgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EdgeTypeSelect is the builder for selecting fields of EdgeType entities.
type EdgeTypeSelect struct {
	*EdgeTypeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ets *EdgeTypeSelect) Aggregate(fns ...AggregateFunc) *EdgeTypeSelect {
	ets.fns = append(ets.fns, fns...)
	return ets
}

// Scan applies the selector query and scans the result into the given value.
func (ets *EdgeTypeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ets.ctx, ent.OpQuerySelect)
	if err := ets.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EdgeTypeQuery, *EdgeTypeSelect](ctx, ets.EdgeTypeQuery, ets, ets.inters, v)
}

func (ets *EdgeTypeSelect) sqlScan(ctx context.Context, root *EdgeTypeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ets.fns))
	for _, fn := range ets.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ets.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ets.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
