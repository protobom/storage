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
	"github.com/protobom/storage/internal/backends/ent/externalreference"
	"github.com/protobom/storage/internal/backends/ent/hashesentry"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// HashesEntryQuery is the builder for querying HashesEntry entities.
type HashesEntryQuery struct {
	config
	ctx                    *QueryContext
	order                  []hashesentry.OrderOption
	inters                 []Interceptor
	predicates             []predicate.HashesEntry
	withDocument           *DocumentQuery
	withExternalReferences *ExternalReferenceQuery
	withNodes              *NodeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the HashesEntryQuery builder.
func (heq *HashesEntryQuery) Where(ps ...predicate.HashesEntry) *HashesEntryQuery {
	heq.predicates = append(heq.predicates, ps...)
	return heq
}

// Limit the number of records to be returned by this query.
func (heq *HashesEntryQuery) Limit(limit int) *HashesEntryQuery {
	heq.ctx.Limit = &limit
	return heq
}

// Offset to start from.
func (heq *HashesEntryQuery) Offset(offset int) *HashesEntryQuery {
	heq.ctx.Offset = &offset
	return heq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (heq *HashesEntryQuery) Unique(unique bool) *HashesEntryQuery {
	heq.ctx.Unique = &unique
	return heq
}

// Order specifies how the records should be ordered.
func (heq *HashesEntryQuery) Order(o ...hashesentry.OrderOption) *HashesEntryQuery {
	heq.order = append(heq.order, o...)
	return heq
}

// QueryDocument chains the current query on the "document" edge.
func (heq *HashesEntryQuery) QueryDocument() *DocumentQuery {
	query := (&DocumentClient{config: heq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := heq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := heq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(hashesentry.Table, hashesentry.FieldID, selector),
			sqlgraph.To(document.Table, document.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, hashesentry.DocumentTable, hashesentry.DocumentColumn),
		)
		fromU = sqlgraph.SetNeighbors(heq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryExternalReferences chains the current query on the "external_references" edge.
func (heq *HashesEntryQuery) QueryExternalReferences() *ExternalReferenceQuery {
	query := (&ExternalReferenceClient{config: heq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := heq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := heq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(hashesentry.Table, hashesentry.FieldID, selector),
			sqlgraph.To(externalreference.Table, externalreference.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, hashesentry.ExternalReferencesTable, hashesentry.ExternalReferencesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(heq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNodes chains the current query on the "nodes" edge.
func (heq *HashesEntryQuery) QueryNodes() *NodeQuery {
	query := (&NodeClient{config: heq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := heq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := heq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(hashesentry.Table, hashesentry.FieldID, selector),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, hashesentry.NodesTable, hashesentry.NodesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(heq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first HashesEntry entity from the query.
// Returns a *NotFoundError when no HashesEntry was found.
func (heq *HashesEntryQuery) First(ctx context.Context) (*HashesEntry, error) {
	nodes, err := heq.Limit(1).All(setContextOp(ctx, heq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{hashesentry.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (heq *HashesEntryQuery) FirstX(ctx context.Context) *HashesEntry {
	node, err := heq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first HashesEntry ID from the query.
// Returns a *NotFoundError when no HashesEntry ID was found.
func (heq *HashesEntryQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = heq.Limit(1).IDs(setContextOp(ctx, heq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{hashesentry.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (heq *HashesEntryQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := heq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single HashesEntry entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one HashesEntry entity is found.
// Returns a *NotFoundError when no HashesEntry entities are found.
func (heq *HashesEntryQuery) Only(ctx context.Context) (*HashesEntry, error) {
	nodes, err := heq.Limit(2).All(setContextOp(ctx, heq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{hashesentry.Label}
	default:
		return nil, &NotSingularError{hashesentry.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (heq *HashesEntryQuery) OnlyX(ctx context.Context) *HashesEntry {
	node, err := heq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only HashesEntry ID in the query.
// Returns a *NotSingularError when more than one HashesEntry ID is found.
// Returns a *NotFoundError when no entities are found.
func (heq *HashesEntryQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = heq.Limit(2).IDs(setContextOp(ctx, heq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{hashesentry.Label}
	default:
		err = &NotSingularError{hashesentry.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (heq *HashesEntryQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := heq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of HashesEntries.
func (heq *HashesEntryQuery) All(ctx context.Context) ([]*HashesEntry, error) {
	ctx = setContextOp(ctx, heq.ctx, ent.OpQueryAll)
	if err := heq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*HashesEntry, *HashesEntryQuery]()
	return withInterceptors[[]*HashesEntry](ctx, heq, qr, heq.inters)
}

// AllX is like All, but panics if an error occurs.
func (heq *HashesEntryQuery) AllX(ctx context.Context) []*HashesEntry {
	nodes, err := heq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of HashesEntry IDs.
func (heq *HashesEntryQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if heq.ctx.Unique == nil && heq.path != nil {
		heq.Unique(true)
	}
	ctx = setContextOp(ctx, heq.ctx, ent.OpQueryIDs)
	if err = heq.Select(hashesentry.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (heq *HashesEntryQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := heq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (heq *HashesEntryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, heq.ctx, ent.OpQueryCount)
	if err := heq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, heq, querierCount[*HashesEntryQuery](), heq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (heq *HashesEntryQuery) CountX(ctx context.Context) int {
	count, err := heq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (heq *HashesEntryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, heq.ctx, ent.OpQueryExist)
	switch _, err := heq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (heq *HashesEntryQuery) ExistX(ctx context.Context) bool {
	exist, err := heq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the HashesEntryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (heq *HashesEntryQuery) Clone() *HashesEntryQuery {
	if heq == nil {
		return nil
	}
	return &HashesEntryQuery{
		config:                 heq.config,
		ctx:                    heq.ctx.Clone(),
		order:                  append([]hashesentry.OrderOption{}, heq.order...),
		inters:                 append([]Interceptor{}, heq.inters...),
		predicates:             append([]predicate.HashesEntry{}, heq.predicates...),
		withDocument:           heq.withDocument.Clone(),
		withExternalReferences: heq.withExternalReferences.Clone(),
		withNodes:              heq.withNodes.Clone(),
		// clone intermediate query.
		sql:  heq.sql.Clone(),
		path: heq.path,
	}
}

// WithDocument tells the query-builder to eager-load the nodes that are connected to
// the "document" edge. The optional arguments are used to configure the query builder of the edge.
func (heq *HashesEntryQuery) WithDocument(opts ...func(*DocumentQuery)) *HashesEntryQuery {
	query := (&DocumentClient{config: heq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	heq.withDocument = query
	return heq
}

// WithExternalReferences tells the query-builder to eager-load the nodes that are connected to
// the "external_references" edge. The optional arguments are used to configure the query builder of the edge.
func (heq *HashesEntryQuery) WithExternalReferences(opts ...func(*ExternalReferenceQuery)) *HashesEntryQuery {
	query := (&ExternalReferenceClient{config: heq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	heq.withExternalReferences = query
	return heq
}

// WithNodes tells the query-builder to eager-load the nodes that are connected to
// the "nodes" edge. The optional arguments are used to configure the query builder of the edge.
func (heq *HashesEntryQuery) WithNodes(opts ...func(*NodeQuery)) *HashesEntryQuery {
	query := (&NodeClient{config: heq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	heq.withNodes = query
	return heq
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
//	client.HashesEntry.Query().
//		GroupBy(hashesentry.FieldDocumentID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (heq *HashesEntryQuery) GroupBy(field string, fields ...string) *HashesEntryGroupBy {
	heq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &HashesEntryGroupBy{build: heq}
	grbuild.flds = &heq.ctx.Fields
	grbuild.label = hashesentry.Label
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
//	client.HashesEntry.Query().
//		Select(hashesentry.FieldDocumentID).
//		Scan(ctx, &v)
func (heq *HashesEntryQuery) Select(fields ...string) *HashesEntrySelect {
	heq.ctx.Fields = append(heq.ctx.Fields, fields...)
	sbuild := &HashesEntrySelect{HashesEntryQuery: heq}
	sbuild.label = hashesentry.Label
	sbuild.flds, sbuild.scan = &heq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a HashesEntrySelect configured with the given aggregations.
func (heq *HashesEntryQuery) Aggregate(fns ...AggregateFunc) *HashesEntrySelect {
	return heq.Select().Aggregate(fns...)
}

func (heq *HashesEntryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range heq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, heq); err != nil {
				return err
			}
		}
	}
	for _, f := range heq.ctx.Fields {
		if !hashesentry.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if heq.path != nil {
		prev, err := heq.path(ctx)
		if err != nil {
			return err
		}
		heq.sql = prev
	}
	return nil
}

func (heq *HashesEntryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*HashesEntry, error) {
	var (
		nodes       = []*HashesEntry{}
		_spec       = heq.querySpec()
		loadedTypes = [3]bool{
			heq.withDocument != nil,
			heq.withExternalReferences != nil,
			heq.withNodes != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*HashesEntry).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &HashesEntry{config: heq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, heq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := heq.withDocument; query != nil {
		if err := heq.loadDocument(ctx, query, nodes, nil,
			func(n *HashesEntry, e *Document) { n.Edges.Document = e }); err != nil {
			return nil, err
		}
	}
	if query := heq.withExternalReferences; query != nil {
		if err := heq.loadExternalReferences(ctx, query, nodes,
			func(n *HashesEntry) { n.Edges.ExternalReferences = []*ExternalReference{} },
			func(n *HashesEntry, e *ExternalReference) {
				n.Edges.ExternalReferences = append(n.Edges.ExternalReferences, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := heq.withNodes; query != nil {
		if err := heq.loadNodes(ctx, query, nodes,
			func(n *HashesEntry) { n.Edges.Nodes = []*Node{} },
			func(n *HashesEntry, e *Node) { n.Edges.Nodes = append(n.Edges.Nodes, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (heq *HashesEntryQuery) loadDocument(ctx context.Context, query *DocumentQuery, nodes []*HashesEntry, init func(*HashesEntry), assign func(*HashesEntry, *Document)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*HashesEntry)
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
func (heq *HashesEntryQuery) loadExternalReferences(ctx context.Context, query *ExternalReferenceQuery, nodes []*HashesEntry, init func(*HashesEntry), assign func(*HashesEntry, *ExternalReference)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*HashesEntry)
	nids := make(map[uuid.UUID]map[*HashesEntry]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(hashesentry.ExternalReferencesTable)
		s.Join(joinT).On(s.C(externalreference.FieldID), joinT.C(hashesentry.ExternalReferencesPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(hashesentry.ExternalReferencesPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(hashesentry.ExternalReferencesPrimaryKey[1]))
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
					nids[inValue] = map[*HashesEntry]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*ExternalReference](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "external_references" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (heq *HashesEntryQuery) loadNodes(ctx context.Context, query *NodeQuery, nodes []*HashesEntry, init func(*HashesEntry), assign func(*HashesEntry, *Node)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*HashesEntry)
	nids := make(map[uuid.UUID]map[*HashesEntry]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(hashesentry.NodesTable)
		s.Join(joinT).On(s.C(node.FieldID), joinT.C(hashesentry.NodesPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(hashesentry.NodesPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(hashesentry.NodesPrimaryKey[1]))
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
					nids[inValue] = map[*HashesEntry]struct{}{byID[outValue]: {}}
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

func (heq *HashesEntryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := heq.querySpec()
	_spec.Node.Columns = heq.ctx.Fields
	if len(heq.ctx.Fields) > 0 {
		_spec.Unique = heq.ctx.Unique != nil && *heq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, heq.driver, _spec)
}

func (heq *HashesEntryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(hashesentry.Table, hashesentry.Columns, sqlgraph.NewFieldSpec(hashesentry.FieldID, field.TypeUUID))
	_spec.From = heq.sql
	if unique := heq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if heq.path != nil {
		_spec.Unique = true
	}
	if fields := heq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hashesentry.FieldID)
		for i := range fields {
			if fields[i] != hashesentry.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if heq.withDocument != nil {
			_spec.Node.AddColumnOnce(hashesentry.FieldDocumentID)
		}
	}
	if ps := heq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := heq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := heq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := heq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (heq *HashesEntryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(heq.driver.Dialect())
	t1 := builder.Table(hashesentry.Table)
	columns := heq.ctx.Fields
	if len(columns) == 0 {
		columns = hashesentry.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if heq.sql != nil {
		selector = heq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if heq.ctx.Unique != nil && *heq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range heq.predicates {
		p(selector)
	}
	for _, p := range heq.order {
		p(selector)
	}
	if offset := heq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := heq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// HashesEntryGroupBy is the group-by builder for HashesEntry entities.
type HashesEntryGroupBy struct {
	selector
	build *HashesEntryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hegb *HashesEntryGroupBy) Aggregate(fns ...AggregateFunc) *HashesEntryGroupBy {
	hegb.fns = append(hegb.fns, fns...)
	return hegb
}

// Scan applies the selector query and scans the result into the given value.
func (hegb *HashesEntryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hegb.build.ctx, ent.OpQueryGroupBy)
	if err := hegb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HashesEntryQuery, *HashesEntryGroupBy](ctx, hegb.build, hegb, hegb.build.inters, v)
}

func (hegb *HashesEntryGroupBy) sqlScan(ctx context.Context, root *HashesEntryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(hegb.fns))
	for _, fn := range hegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*hegb.flds)+len(hegb.fns))
		for _, f := range *hegb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*hegb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hegb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// HashesEntrySelect is the builder for selecting fields of HashesEntry entities.
type HashesEntrySelect struct {
	*HashesEntryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (hes *HashesEntrySelect) Aggregate(fns ...AggregateFunc) *HashesEntrySelect {
	hes.fns = append(hes.fns, fns...)
	return hes
}

// Scan applies the selector query and scans the result into the given value.
func (hes *HashesEntrySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hes.ctx, ent.OpQuerySelect)
	if err := hes.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HashesEntryQuery, *HashesEntrySelect](ctx, hes.HashesEntryQuery, hes, hes.inters, v)
}

func (hes *HashesEntrySelect) sqlScan(ctx context.Context, root *HashesEntryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(hes.fns))
	for _, fn := range hes.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*hes.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
