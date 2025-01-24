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
	"github.com/protobom/storage/internal/backends/ent/metadata"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/person"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// PersonQuery is the builder for querying Person entities.
type PersonQuery struct {
	config
	ctx                 *QueryContext
	order               []person.OrderOption
	inters              []Interceptor
	predicates          []predicate.Person
	withContactOwner    *PersonQuery
	withContacts        *PersonQuery
	withDocuments       *DocumentQuery
	withMetadata        *MetadataQuery
	withOriginatorNodes *NodeQuery
	withSupplierNodes   *NodeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PersonQuery builder.
func (pq *PersonQuery) Where(ps ...predicate.Person) *PersonQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PersonQuery) Limit(limit int) *PersonQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PersonQuery) Offset(offset int) *PersonQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PersonQuery) Unique(unique bool) *PersonQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PersonQuery) Order(o ...person.OrderOption) *PersonQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryContactOwner chains the current query on the "contact_owner" edge.
func (pq *PersonQuery) QueryContactOwner() *PersonQuery {
	query := (&PersonClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(person.Table, person.FieldID, selector),
			sqlgraph.To(person.Table, person.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, person.ContactOwnerTable, person.ContactOwnerPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryContacts chains the current query on the "contacts" edge.
func (pq *PersonQuery) QueryContacts() *PersonQuery {
	query := (&PersonClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(person.Table, person.FieldID, selector),
			sqlgraph.To(person.Table, person.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, person.ContactsTable, person.ContactsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDocuments chains the current query on the "documents" edge.
func (pq *PersonQuery) QueryDocuments() *DocumentQuery {
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
			sqlgraph.From(person.Table, person.FieldID, selector),
			sqlgraph.To(document.Table, document.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, person.DocumentsTable, person.DocumentsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMetadata chains the current query on the "metadata" edge.
func (pq *PersonQuery) QueryMetadata() *MetadataQuery {
	query := (&MetadataClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(person.Table, person.FieldID, selector),
			sqlgraph.To(metadata.Table, metadata.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, person.MetadataTable, person.MetadataPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOriginatorNodes chains the current query on the "originator_nodes" edge.
func (pq *PersonQuery) QueryOriginatorNodes() *NodeQuery {
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
			sqlgraph.From(person.Table, person.FieldID, selector),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, person.OriginatorNodesTable, person.OriginatorNodesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySupplierNodes chains the current query on the "supplier_nodes" edge.
func (pq *PersonQuery) QuerySupplierNodes() *NodeQuery {
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
			sqlgraph.From(person.Table, person.FieldID, selector),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, person.SupplierNodesTable, person.SupplierNodesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Person entity from the query.
// Returns a *NotFoundError when no Person was found.
func (pq *PersonQuery) First(ctx context.Context) (*Person, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{person.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PersonQuery) FirstX(ctx context.Context) *Person {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Person ID from the query.
// Returns a *NotFoundError when no Person ID was found.
func (pq *PersonQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{person.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PersonQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Person entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Person entity is found.
// Returns a *NotFoundError when no Person entities are found.
func (pq *PersonQuery) Only(ctx context.Context) (*Person, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{person.Label}
	default:
		return nil, &NotSingularError{person.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PersonQuery) OnlyX(ctx context.Context) *Person {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Person ID in the query.
// Returns a *NotSingularError when more than one Person ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PersonQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{person.Label}
	default:
		err = &NotSingularError{person.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PersonQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Persons.
func (pq *PersonQuery) All(ctx context.Context) ([]*Person, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryAll)
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Person, *PersonQuery]()
	return withInterceptors[[]*Person](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PersonQuery) AllX(ctx context.Context) []*Person {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Person IDs.
func (pq *PersonQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryIDs)
	if err = pq.Select(person.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PersonQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PersonQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryCount)
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PersonQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PersonQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PersonQuery) Exist(ctx context.Context) (bool, error) {
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
func (pq *PersonQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PersonQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PersonQuery) Clone() *PersonQuery {
	if pq == nil {
		return nil
	}
	return &PersonQuery{
		config:              pq.config,
		ctx:                 pq.ctx.Clone(),
		order:               append([]person.OrderOption{}, pq.order...),
		inters:              append([]Interceptor{}, pq.inters...),
		predicates:          append([]predicate.Person{}, pq.predicates...),
		withContactOwner:    pq.withContactOwner.Clone(),
		withContacts:        pq.withContacts.Clone(),
		withDocuments:       pq.withDocuments.Clone(),
		withMetadata:        pq.withMetadata.Clone(),
		withOriginatorNodes: pq.withOriginatorNodes.Clone(),
		withSupplierNodes:   pq.withSupplierNodes.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithContactOwner tells the query-builder to eager-load the nodes that are connected to
// the "contact_owner" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PersonQuery) WithContactOwner(opts ...func(*PersonQuery)) *PersonQuery {
	query := (&PersonClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withContactOwner = query
	return pq
}

// WithContacts tells the query-builder to eager-load the nodes that are connected to
// the "contacts" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PersonQuery) WithContacts(opts ...func(*PersonQuery)) *PersonQuery {
	query := (&PersonClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withContacts = query
	return pq
}

// WithDocuments tells the query-builder to eager-load the nodes that are connected to
// the "documents" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PersonQuery) WithDocuments(opts ...func(*DocumentQuery)) *PersonQuery {
	query := (&DocumentClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withDocuments = query
	return pq
}

// WithMetadata tells the query-builder to eager-load the nodes that are connected to
// the "metadata" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PersonQuery) WithMetadata(opts ...func(*MetadataQuery)) *PersonQuery {
	query := (&MetadataClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withMetadata = query
	return pq
}

// WithOriginatorNodes tells the query-builder to eager-load the nodes that are connected to
// the "originator_nodes" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PersonQuery) WithOriginatorNodes(opts ...func(*NodeQuery)) *PersonQuery {
	query := (&NodeClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withOriginatorNodes = query
	return pq
}

// WithSupplierNodes tells the query-builder to eager-load the nodes that are connected to
// the "supplier_nodes" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PersonQuery) WithSupplierNodes(opts ...func(*NodeQuery)) *PersonQuery {
	query := (&NodeClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withSupplierNodes = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ProtoMessage *sbom.Person `json:"-"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Person.Query().
//		GroupBy(person.FieldProtoMessage).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PersonQuery) GroupBy(field string, fields ...string) *PersonGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PersonGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = person.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ProtoMessage *sbom.Person `json:"-"`
//	}
//
//	client.Person.Query().
//		Select(person.FieldProtoMessage).
//		Scan(ctx, &v)
func (pq *PersonQuery) Select(fields ...string) *PersonSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PersonSelect{PersonQuery: pq}
	sbuild.label = person.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PersonSelect configured with the given aggregations.
func (pq *PersonQuery) Aggregate(fns ...AggregateFunc) *PersonSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PersonQuery) prepareQuery(ctx context.Context) error {
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
		if !person.ValidColumn(f) {
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

func (pq *PersonQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Person, error) {
	var (
		nodes       = []*Person{}
		_spec       = pq.querySpec()
		loadedTypes = [6]bool{
			pq.withContactOwner != nil,
			pq.withContacts != nil,
			pq.withDocuments != nil,
			pq.withMetadata != nil,
			pq.withOriginatorNodes != nil,
			pq.withSupplierNodes != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Person).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Person{config: pq.config}
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
	if query := pq.withContactOwner; query != nil {
		if err := pq.loadContactOwner(ctx, query, nodes,
			func(n *Person) { n.Edges.ContactOwner = []*Person{} },
			func(n *Person, e *Person) { n.Edges.ContactOwner = append(n.Edges.ContactOwner, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withContacts; query != nil {
		if err := pq.loadContacts(ctx, query, nodes,
			func(n *Person) { n.Edges.Contacts = []*Person{} },
			func(n *Person, e *Person) { n.Edges.Contacts = append(n.Edges.Contacts, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withDocuments; query != nil {
		if err := pq.loadDocuments(ctx, query, nodes,
			func(n *Person) { n.Edges.Documents = []*Document{} },
			func(n *Person, e *Document) { n.Edges.Documents = append(n.Edges.Documents, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withMetadata; query != nil {
		if err := pq.loadMetadata(ctx, query, nodes,
			func(n *Person) { n.Edges.Metadata = []*Metadata{} },
			func(n *Person, e *Metadata) { n.Edges.Metadata = append(n.Edges.Metadata, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withOriginatorNodes; query != nil {
		if err := pq.loadOriginatorNodes(ctx, query, nodes,
			func(n *Person) { n.Edges.OriginatorNodes = []*Node{} },
			func(n *Person, e *Node) { n.Edges.OriginatorNodes = append(n.Edges.OriginatorNodes, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withSupplierNodes; query != nil {
		if err := pq.loadSupplierNodes(ctx, query, nodes,
			func(n *Person) { n.Edges.SupplierNodes = []*Node{} },
			func(n *Person, e *Node) { n.Edges.SupplierNodes = append(n.Edges.SupplierNodes, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PersonQuery) loadContactOwner(ctx context.Context, query *PersonQuery, nodes []*Person, init func(*Person), assign func(*Person, *Person)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Person)
	nids := make(map[uuid.UUID]map[*Person]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(person.ContactOwnerTable)
		s.Join(joinT).On(s.C(person.FieldID), joinT.C(person.ContactOwnerPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(person.ContactOwnerPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(person.ContactOwnerPrimaryKey[1]))
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
					nids[inValue] = map[*Person]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Person](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "contact_owner" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (pq *PersonQuery) loadContacts(ctx context.Context, query *PersonQuery, nodes []*Person, init func(*Person), assign func(*Person, *Person)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Person)
	nids := make(map[uuid.UUID]map[*Person]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(person.ContactsTable)
		s.Join(joinT).On(s.C(person.FieldID), joinT.C(person.ContactsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(person.ContactsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(person.ContactsPrimaryKey[0]))
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
					nids[inValue] = map[*Person]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Person](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "contacts" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (pq *PersonQuery) loadDocuments(ctx context.Context, query *DocumentQuery, nodes []*Person, init func(*Person), assign func(*Person, *Document)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Person)
	nids := make(map[uuid.UUID]map[*Person]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(person.DocumentsTable)
		s.Join(joinT).On(s.C(document.FieldID), joinT.C(person.DocumentsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(person.DocumentsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(person.DocumentsPrimaryKey[1]))
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
					nids[inValue] = map[*Person]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Document](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "documents" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (pq *PersonQuery) loadMetadata(ctx context.Context, query *MetadataQuery, nodes []*Person, init func(*Person), assign func(*Person, *Metadata)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Person)
	nids := make(map[uuid.UUID]map[*Person]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(person.MetadataTable)
		s.Join(joinT).On(s.C(metadata.FieldID), joinT.C(person.MetadataPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(person.MetadataPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(person.MetadataPrimaryKey[1]))
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
					nids[inValue] = map[*Person]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Metadata](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "metadata" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (pq *PersonQuery) loadOriginatorNodes(ctx context.Context, query *NodeQuery, nodes []*Person, init func(*Person), assign func(*Person, *Node)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Person)
	nids := make(map[uuid.UUID]map[*Person]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(person.OriginatorNodesTable)
		s.Join(joinT).On(s.C(node.FieldID), joinT.C(person.OriginatorNodesPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(person.OriginatorNodesPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(person.OriginatorNodesPrimaryKey[1]))
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
					nids[inValue] = map[*Person]struct{}{byID[outValue]: {}}
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
			return fmt.Errorf(`unexpected "originator_nodes" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (pq *PersonQuery) loadSupplierNodes(ctx context.Context, query *NodeQuery, nodes []*Person, init func(*Person), assign func(*Person, *Node)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Person)
	nids := make(map[uuid.UUID]map[*Person]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(person.SupplierNodesTable)
		s.Join(joinT).On(s.C(node.FieldID), joinT.C(person.SupplierNodesPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(person.SupplierNodesPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(person.SupplierNodesPrimaryKey[1]))
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
					nids[inValue] = map[*Person]struct{}{byID[outValue]: {}}
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
			return fmt.Errorf(`unexpected "supplier_nodes" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (pq *PersonQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PersonQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(person.Table, person.Columns, sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, person.FieldID)
		for i := range fields {
			if fields[i] != person.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
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

func (pq *PersonQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(person.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = person.Columns
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

// PersonGroupBy is the group-by builder for Person entities.
type PersonGroupBy struct {
	selector
	build *PersonQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PersonGroupBy) Aggregate(fns ...AggregateFunc) *PersonGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PersonGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, ent.OpQueryGroupBy)
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PersonQuery, *PersonGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PersonGroupBy) sqlScan(ctx context.Context, root *PersonQuery, v any) error {
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

// PersonSelect is the builder for selecting fields of Person entities.
type PersonSelect struct {
	*PersonQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PersonSelect) Aggregate(fns ...AggregateFunc) *PersonSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PersonSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, ent.OpQuerySelect)
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PersonQuery, *PersonSelect](ctx, ps.PersonQuery, ps, ps.inters, v)
}

func (ps *PersonSelect) sqlScan(ctx context.Context, root *PersonQuery, v any) error {
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
