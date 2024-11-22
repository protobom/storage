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
	"github.com/protobom/storage/internal/backends/ent/documenttype"
	"github.com/protobom/storage/internal/backends/ent/metadata"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// DocumentTypeQuery is the builder for querying DocumentType entities.
type DocumentTypeQuery struct {
	config
	ctx          *QueryContext
	order        []documenttype.OrderOption
	inters       []Interceptor
	predicates   []predicate.DocumentType
	withDocument *DocumentQuery
	withMetadata *MetadataQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DocumentTypeQuery builder.
func (dtq *DocumentTypeQuery) Where(ps ...predicate.DocumentType) *DocumentTypeQuery {
	dtq.predicates = append(dtq.predicates, ps...)
	return dtq
}

// Limit the number of records to be returned by this query.
func (dtq *DocumentTypeQuery) Limit(limit int) *DocumentTypeQuery {
	dtq.ctx.Limit = &limit
	return dtq
}

// Offset to start from.
func (dtq *DocumentTypeQuery) Offset(offset int) *DocumentTypeQuery {
	dtq.ctx.Offset = &offset
	return dtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dtq *DocumentTypeQuery) Unique(unique bool) *DocumentTypeQuery {
	dtq.ctx.Unique = &unique
	return dtq
}

// Order specifies how the records should be ordered.
func (dtq *DocumentTypeQuery) Order(o ...documenttype.OrderOption) *DocumentTypeQuery {
	dtq.order = append(dtq.order, o...)
	return dtq
}

// QueryDocument chains the current query on the "document" edge.
func (dtq *DocumentTypeQuery) QueryDocument() *DocumentQuery {
	query := (&DocumentClient{config: dtq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(documenttype.Table, documenttype.FieldID, selector),
			sqlgraph.To(document.Table, document.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, documenttype.DocumentTable, documenttype.DocumentColumn),
		)
		fromU = sqlgraph.SetNeighbors(dtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMetadata chains the current query on the "metadata" edge.
func (dtq *DocumentTypeQuery) QueryMetadata() *MetadataQuery {
	query := (&MetadataClient{config: dtq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(documenttype.Table, documenttype.FieldID, selector),
			sqlgraph.To(metadata.Table, metadata.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, documenttype.MetadataTable, documenttype.MetadataColumn),
		)
		fromU = sqlgraph.SetNeighbors(dtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DocumentType entity from the query.
// Returns a *NotFoundError when no DocumentType was found.
func (dtq *DocumentTypeQuery) First(ctx context.Context) (*DocumentType, error) {
	nodes, err := dtq.Limit(1).All(setContextOp(ctx, dtq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{documenttype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dtq *DocumentTypeQuery) FirstX(ctx context.Context) *DocumentType {
	node, err := dtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DocumentType ID from the query.
// Returns a *NotFoundError when no DocumentType ID was found.
func (dtq *DocumentTypeQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dtq.Limit(1).IDs(setContextOp(ctx, dtq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{documenttype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dtq *DocumentTypeQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DocumentType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DocumentType entity is found.
// Returns a *NotFoundError when no DocumentType entities are found.
func (dtq *DocumentTypeQuery) Only(ctx context.Context) (*DocumentType, error) {
	nodes, err := dtq.Limit(2).All(setContextOp(ctx, dtq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{documenttype.Label}
	default:
		return nil, &NotSingularError{documenttype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dtq *DocumentTypeQuery) OnlyX(ctx context.Context) *DocumentType {
	node, err := dtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DocumentType ID in the query.
// Returns a *NotSingularError when more than one DocumentType ID is found.
// Returns a *NotFoundError when no entities are found.
func (dtq *DocumentTypeQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dtq.Limit(2).IDs(setContextOp(ctx, dtq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{documenttype.Label}
	default:
		err = &NotSingularError{documenttype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dtq *DocumentTypeQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DocumentTypes.
func (dtq *DocumentTypeQuery) All(ctx context.Context) ([]*DocumentType, error) {
	ctx = setContextOp(ctx, dtq.ctx, ent.OpQueryAll)
	if err := dtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DocumentType, *DocumentTypeQuery]()
	return withInterceptors[[]*DocumentType](ctx, dtq, qr, dtq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dtq *DocumentTypeQuery) AllX(ctx context.Context) []*DocumentType {
	nodes, err := dtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DocumentType IDs.
func (dtq *DocumentTypeQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if dtq.ctx.Unique == nil && dtq.path != nil {
		dtq.Unique(true)
	}
	ctx = setContextOp(ctx, dtq.ctx, ent.OpQueryIDs)
	if err = dtq.Select(documenttype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dtq *DocumentTypeQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dtq *DocumentTypeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dtq.ctx, ent.OpQueryCount)
	if err := dtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dtq, querierCount[*DocumentTypeQuery](), dtq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dtq *DocumentTypeQuery) CountX(ctx context.Context) int {
	count, err := dtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dtq *DocumentTypeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dtq.ctx, ent.OpQueryExist)
	switch _, err := dtq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dtq *DocumentTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := dtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DocumentTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dtq *DocumentTypeQuery) Clone() *DocumentTypeQuery {
	if dtq == nil {
		return nil
	}
	return &DocumentTypeQuery{
		config:       dtq.config,
		ctx:          dtq.ctx.Clone(),
		order:        append([]documenttype.OrderOption{}, dtq.order...),
		inters:       append([]Interceptor{}, dtq.inters...),
		predicates:   append([]predicate.DocumentType{}, dtq.predicates...),
		withDocument: dtq.withDocument.Clone(),
		withMetadata: dtq.withMetadata.Clone(),
		// clone intermediate query.
		sql:  dtq.sql.Clone(),
		path: dtq.path,
	}
}

// WithDocument tells the query-builder to eager-load the nodes that are connected to
// the "document" edge. The optional arguments are used to configure the query builder of the edge.
func (dtq *DocumentTypeQuery) WithDocument(opts ...func(*DocumentQuery)) *DocumentTypeQuery {
	query := (&DocumentClient{config: dtq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dtq.withDocument = query
	return dtq
}

// WithMetadata tells the query-builder to eager-load the nodes that are connected to
// the "metadata" edge. The optional arguments are used to configure the query builder of the edge.
func (dtq *DocumentTypeQuery) WithMetadata(opts ...func(*MetadataQuery)) *DocumentTypeQuery {
	query := (&MetadataClient{config: dtq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dtq.withMetadata = query
	return dtq
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
//	client.DocumentType.Query().
//		GroupBy(documenttype.FieldDocumentID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dtq *DocumentTypeQuery) GroupBy(field string, fields ...string) *DocumentTypeGroupBy {
	dtq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DocumentTypeGroupBy{build: dtq}
	grbuild.flds = &dtq.ctx.Fields
	grbuild.label = documenttype.Label
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
//	client.DocumentType.Query().
//		Select(documenttype.FieldDocumentID).
//		Scan(ctx, &v)
func (dtq *DocumentTypeQuery) Select(fields ...string) *DocumentTypeSelect {
	dtq.ctx.Fields = append(dtq.ctx.Fields, fields...)
	sbuild := &DocumentTypeSelect{DocumentTypeQuery: dtq}
	sbuild.label = documenttype.Label
	sbuild.flds, sbuild.scan = &dtq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DocumentTypeSelect configured with the given aggregations.
func (dtq *DocumentTypeQuery) Aggregate(fns ...AggregateFunc) *DocumentTypeSelect {
	return dtq.Select().Aggregate(fns...)
}

func (dtq *DocumentTypeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dtq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dtq); err != nil {
				return err
			}
		}
	}
	for _, f := range dtq.ctx.Fields {
		if !documenttype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dtq.path != nil {
		prev, err := dtq.path(ctx)
		if err != nil {
			return err
		}
		dtq.sql = prev
	}
	return nil
}

func (dtq *DocumentTypeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DocumentType, error) {
	var (
		nodes       = []*DocumentType{}
		_spec       = dtq.querySpec()
		loadedTypes = [2]bool{
			dtq.withDocument != nil,
			dtq.withMetadata != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DocumentType).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DocumentType{config: dtq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dtq.withDocument; query != nil {
		if err := dtq.loadDocument(ctx, query, nodes, nil,
			func(n *DocumentType, e *Document) { n.Edges.Document = e }); err != nil {
			return nil, err
		}
	}
	if query := dtq.withMetadata; query != nil {
		if err := dtq.loadMetadata(ctx, query, nodes, nil,
			func(n *DocumentType, e *Metadata) { n.Edges.Metadata = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dtq *DocumentTypeQuery) loadDocument(ctx context.Context, query *DocumentQuery, nodes []*DocumentType, init func(*DocumentType), assign func(*DocumentType, *Document)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*DocumentType)
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
func (dtq *DocumentTypeQuery) loadMetadata(ctx context.Context, query *MetadataQuery, nodes []*DocumentType, init func(*DocumentType), assign func(*DocumentType, *Metadata)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*DocumentType)
	for i := range nodes {
		fk := nodes[i].MetadataID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(metadata.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "metadata_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (dtq *DocumentTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dtq.querySpec()
	_spec.Node.Columns = dtq.ctx.Fields
	if len(dtq.ctx.Fields) > 0 {
		_spec.Unique = dtq.ctx.Unique != nil && *dtq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dtq.driver, _spec)
}

func (dtq *DocumentTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(documenttype.Table, documenttype.Columns, sqlgraph.NewFieldSpec(documenttype.FieldID, field.TypeUUID))
	_spec.From = dtq.sql
	if unique := dtq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dtq.path != nil {
		_spec.Unique = true
	}
	if fields := dtq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, documenttype.FieldID)
		for i := range fields {
			if fields[i] != documenttype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if dtq.withDocument != nil {
			_spec.Node.AddColumnOnce(documenttype.FieldDocumentID)
		}
		if dtq.withMetadata != nil {
			_spec.Node.AddColumnOnce(documenttype.FieldMetadataID)
		}
	}
	if ps := dtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dtq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dtq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dtq *DocumentTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dtq.driver.Dialect())
	t1 := builder.Table(documenttype.Table)
	columns := dtq.ctx.Fields
	if len(columns) == 0 {
		columns = documenttype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dtq.sql != nil {
		selector = dtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dtq.ctx.Unique != nil && *dtq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dtq.predicates {
		p(selector)
	}
	for _, p := range dtq.order {
		p(selector)
	}
	if offset := dtq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dtq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DocumentTypeGroupBy is the group-by builder for DocumentType entities.
type DocumentTypeGroupBy struct {
	selector
	build *DocumentTypeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dtgb *DocumentTypeGroupBy) Aggregate(fns ...AggregateFunc) *DocumentTypeGroupBy {
	dtgb.fns = append(dtgb.fns, fns...)
	return dtgb
}

// Scan applies the selector query and scans the result into the given value.
func (dtgb *DocumentTypeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dtgb.build.ctx, ent.OpQueryGroupBy)
	if err := dtgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DocumentTypeQuery, *DocumentTypeGroupBy](ctx, dtgb.build, dtgb, dtgb.build.inters, v)
}

func (dtgb *DocumentTypeGroupBy) sqlScan(ctx context.Context, root *DocumentTypeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dtgb.fns))
	for _, fn := range dtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dtgb.flds)+len(dtgb.fns))
		for _, f := range *dtgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dtgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dtgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DocumentTypeSelect is the builder for selecting fields of DocumentType entities.
type DocumentTypeSelect struct {
	*DocumentTypeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dts *DocumentTypeSelect) Aggregate(fns ...AggregateFunc) *DocumentTypeSelect {
	dts.fns = append(dts.fns, fns...)
	return dts
}

// Scan applies the selector query and scans the result into the given value.
func (dts *DocumentTypeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dts.ctx, ent.OpQuerySelect)
	if err := dts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DocumentTypeQuery, *DocumentTypeSelect](ctx, dts.DocumentTypeQuery, dts, dts.inters, v)
}

func (dts *DocumentTypeSelect) sqlScan(ctx context.Context, root *DocumentTypeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(dts.fns))
	for _, fn := range dts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*dts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
