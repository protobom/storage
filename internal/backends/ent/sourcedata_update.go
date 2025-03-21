// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/protobom/storage/internal/backends/ent/hashesentry"
	"github.com/protobom/storage/internal/backends/ent/metadata"
	"github.com/protobom/storage/internal/backends/ent/predicate"
	"github.com/protobom/storage/internal/backends/ent/sourcedata"
)

// SourceDataUpdate is the builder for updating SourceData entities.
type SourceDataUpdate struct {
	config
	hooks    []Hook
	mutation *SourceDataMutation
}

// Where appends a list predicates to the SourceDataUpdate builder.
func (sdu *SourceDataUpdate) Where(ps ...predicate.SourceData) *SourceDataUpdate {
	sdu.mutation.Where(ps...)
	return sdu
}

// SetFormat sets the "format" field.
func (sdu *SourceDataUpdate) SetFormat(s string) *SourceDataUpdate {
	sdu.mutation.SetFormat(s)
	return sdu
}

// SetNillableFormat sets the "format" field if the given value is not nil.
func (sdu *SourceDataUpdate) SetNillableFormat(s *string) *SourceDataUpdate {
	if s != nil {
		sdu.SetFormat(*s)
	}
	return sdu
}

// SetSize sets the "size" field.
func (sdu *SourceDataUpdate) SetSize(i int64) *SourceDataUpdate {
	sdu.mutation.ResetSize()
	sdu.mutation.SetSize(i)
	return sdu
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (sdu *SourceDataUpdate) SetNillableSize(i *int64) *SourceDataUpdate {
	if i != nil {
		sdu.SetSize(*i)
	}
	return sdu
}

// AddSize adds i to the "size" field.
func (sdu *SourceDataUpdate) AddSize(i int64) *SourceDataUpdate {
	sdu.mutation.AddSize(i)
	return sdu
}

// SetURI sets the "uri" field.
func (sdu *SourceDataUpdate) SetURI(s string) *SourceDataUpdate {
	sdu.mutation.SetURI(s)
	return sdu
}

// SetNillableURI sets the "uri" field if the given value is not nil.
func (sdu *SourceDataUpdate) SetNillableURI(s *string) *SourceDataUpdate {
	if s != nil {
		sdu.SetURI(*s)
	}
	return sdu
}

// ClearURI clears the value of the "uri" field.
func (sdu *SourceDataUpdate) ClearURI() *SourceDataUpdate {
	sdu.mutation.ClearURI()
	return sdu
}

// AddHashIDs adds the "hashes" edge to the HashesEntry entity by IDs.
func (sdu *SourceDataUpdate) AddHashIDs(ids ...uuid.UUID) *SourceDataUpdate {
	sdu.mutation.AddHashIDs(ids...)
	return sdu
}

// AddHashes adds the "hashes" edges to the HashesEntry entity.
func (sdu *SourceDataUpdate) AddHashes(h ...*HashesEntry) *SourceDataUpdate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return sdu.AddHashIDs(ids...)
}

// AddMetadatumIDs adds the "metadata" edge to the Metadata entity by IDs.
func (sdu *SourceDataUpdate) AddMetadatumIDs(ids ...uuid.UUID) *SourceDataUpdate {
	sdu.mutation.AddMetadatumIDs(ids...)
	return sdu
}

// AddMetadata adds the "metadata" edges to the Metadata entity.
func (sdu *SourceDataUpdate) AddMetadata(m ...*Metadata) *SourceDataUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return sdu.AddMetadatumIDs(ids...)
}

// Mutation returns the SourceDataMutation object of the builder.
func (sdu *SourceDataUpdate) Mutation() *SourceDataMutation {
	return sdu.mutation
}

// ClearHashes clears all "hashes" edges to the HashesEntry entity.
func (sdu *SourceDataUpdate) ClearHashes() *SourceDataUpdate {
	sdu.mutation.ClearHashes()
	return sdu
}

// RemoveHashIDs removes the "hashes" edge to HashesEntry entities by IDs.
func (sdu *SourceDataUpdate) RemoveHashIDs(ids ...uuid.UUID) *SourceDataUpdate {
	sdu.mutation.RemoveHashIDs(ids...)
	return sdu
}

// RemoveHashes removes "hashes" edges to HashesEntry entities.
func (sdu *SourceDataUpdate) RemoveHashes(h ...*HashesEntry) *SourceDataUpdate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return sdu.RemoveHashIDs(ids...)
}

// ClearMetadata clears all "metadata" edges to the Metadata entity.
func (sdu *SourceDataUpdate) ClearMetadata() *SourceDataUpdate {
	sdu.mutation.ClearMetadata()
	return sdu
}

// RemoveMetadatumIDs removes the "metadata" edge to Metadata entities by IDs.
func (sdu *SourceDataUpdate) RemoveMetadatumIDs(ids ...uuid.UUID) *SourceDataUpdate {
	sdu.mutation.RemoveMetadatumIDs(ids...)
	return sdu
}

// RemoveMetadata removes "metadata" edges to Metadata entities.
func (sdu *SourceDataUpdate) RemoveMetadata(m ...*Metadata) *SourceDataUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return sdu.RemoveMetadatumIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sdu *SourceDataUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, sdu.sqlSave, sdu.mutation, sdu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sdu *SourceDataUpdate) SaveX(ctx context.Context) int {
	affected, err := sdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sdu *SourceDataUpdate) Exec(ctx context.Context) error {
	_, err := sdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdu *SourceDataUpdate) ExecX(ctx context.Context) {
	if err := sdu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sdu *SourceDataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(sourcedata.Table, sourcedata.Columns, sqlgraph.NewFieldSpec(sourcedata.FieldID, field.TypeUUID))
	if ps := sdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sdu.mutation.Format(); ok {
		_spec.SetField(sourcedata.FieldFormat, field.TypeString, value)
	}
	if value, ok := sdu.mutation.Size(); ok {
		_spec.SetField(sourcedata.FieldSize, field.TypeInt64, value)
	}
	if value, ok := sdu.mutation.AddedSize(); ok {
		_spec.AddField(sourcedata.FieldSize, field.TypeInt64, value)
	}
	if value, ok := sdu.mutation.URI(); ok {
		_spec.SetField(sourcedata.FieldURI, field.TypeString, value)
	}
	if sdu.mutation.URICleared() {
		_spec.ClearField(sourcedata.FieldURI, field.TypeString)
	}
	if sdu.mutation.HashesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sourcedata.HashesTable,
			Columns: sourcedata.HashesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hashesentry.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sdu.mutation.RemovedHashesIDs(); len(nodes) > 0 && !sdu.mutation.HashesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sourcedata.HashesTable,
			Columns: sourcedata.HashesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hashesentry.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sdu.mutation.HashesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sourcedata.HashesTable,
			Columns: sourcedata.HashesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hashesentry.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sdu.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   sourcedata.MetadataTable,
			Columns: []string{sourcedata.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sdu.mutation.RemovedMetadataIDs(); len(nodes) > 0 && !sdu.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   sourcedata.MetadataTable,
			Columns: []string{sourcedata.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sdu.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   sourcedata.MetadataTable,
			Columns: []string{sourcedata.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sourcedata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	sdu.mutation.done = true
	return n, nil
}

// SourceDataUpdateOne is the builder for updating a single SourceData entity.
type SourceDataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SourceDataMutation
}

// SetFormat sets the "format" field.
func (sduo *SourceDataUpdateOne) SetFormat(s string) *SourceDataUpdateOne {
	sduo.mutation.SetFormat(s)
	return sduo
}

// SetNillableFormat sets the "format" field if the given value is not nil.
func (sduo *SourceDataUpdateOne) SetNillableFormat(s *string) *SourceDataUpdateOne {
	if s != nil {
		sduo.SetFormat(*s)
	}
	return sduo
}

// SetSize sets the "size" field.
func (sduo *SourceDataUpdateOne) SetSize(i int64) *SourceDataUpdateOne {
	sduo.mutation.ResetSize()
	sduo.mutation.SetSize(i)
	return sduo
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (sduo *SourceDataUpdateOne) SetNillableSize(i *int64) *SourceDataUpdateOne {
	if i != nil {
		sduo.SetSize(*i)
	}
	return sduo
}

// AddSize adds i to the "size" field.
func (sduo *SourceDataUpdateOne) AddSize(i int64) *SourceDataUpdateOne {
	sduo.mutation.AddSize(i)
	return sduo
}

// SetURI sets the "uri" field.
func (sduo *SourceDataUpdateOne) SetURI(s string) *SourceDataUpdateOne {
	sduo.mutation.SetURI(s)
	return sduo
}

// SetNillableURI sets the "uri" field if the given value is not nil.
func (sduo *SourceDataUpdateOne) SetNillableURI(s *string) *SourceDataUpdateOne {
	if s != nil {
		sduo.SetURI(*s)
	}
	return sduo
}

// ClearURI clears the value of the "uri" field.
func (sduo *SourceDataUpdateOne) ClearURI() *SourceDataUpdateOne {
	sduo.mutation.ClearURI()
	return sduo
}

// AddHashIDs adds the "hashes" edge to the HashesEntry entity by IDs.
func (sduo *SourceDataUpdateOne) AddHashIDs(ids ...uuid.UUID) *SourceDataUpdateOne {
	sduo.mutation.AddHashIDs(ids...)
	return sduo
}

// AddHashes adds the "hashes" edges to the HashesEntry entity.
func (sduo *SourceDataUpdateOne) AddHashes(h ...*HashesEntry) *SourceDataUpdateOne {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return sduo.AddHashIDs(ids...)
}

// AddMetadatumIDs adds the "metadata" edge to the Metadata entity by IDs.
func (sduo *SourceDataUpdateOne) AddMetadatumIDs(ids ...uuid.UUID) *SourceDataUpdateOne {
	sduo.mutation.AddMetadatumIDs(ids...)
	return sduo
}

// AddMetadata adds the "metadata" edges to the Metadata entity.
func (sduo *SourceDataUpdateOne) AddMetadata(m ...*Metadata) *SourceDataUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return sduo.AddMetadatumIDs(ids...)
}

// Mutation returns the SourceDataMutation object of the builder.
func (sduo *SourceDataUpdateOne) Mutation() *SourceDataMutation {
	return sduo.mutation
}

// ClearHashes clears all "hashes" edges to the HashesEntry entity.
func (sduo *SourceDataUpdateOne) ClearHashes() *SourceDataUpdateOne {
	sduo.mutation.ClearHashes()
	return sduo
}

// RemoveHashIDs removes the "hashes" edge to HashesEntry entities by IDs.
func (sduo *SourceDataUpdateOne) RemoveHashIDs(ids ...uuid.UUID) *SourceDataUpdateOne {
	sduo.mutation.RemoveHashIDs(ids...)
	return sduo
}

// RemoveHashes removes "hashes" edges to HashesEntry entities.
func (sduo *SourceDataUpdateOne) RemoveHashes(h ...*HashesEntry) *SourceDataUpdateOne {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return sduo.RemoveHashIDs(ids...)
}

// ClearMetadata clears all "metadata" edges to the Metadata entity.
func (sduo *SourceDataUpdateOne) ClearMetadata() *SourceDataUpdateOne {
	sduo.mutation.ClearMetadata()
	return sduo
}

// RemoveMetadatumIDs removes the "metadata" edge to Metadata entities by IDs.
func (sduo *SourceDataUpdateOne) RemoveMetadatumIDs(ids ...uuid.UUID) *SourceDataUpdateOne {
	sduo.mutation.RemoveMetadatumIDs(ids...)
	return sduo
}

// RemoveMetadata removes "metadata" edges to Metadata entities.
func (sduo *SourceDataUpdateOne) RemoveMetadata(m ...*Metadata) *SourceDataUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return sduo.RemoveMetadatumIDs(ids...)
}

// Where appends a list predicates to the SourceDataUpdate builder.
func (sduo *SourceDataUpdateOne) Where(ps ...predicate.SourceData) *SourceDataUpdateOne {
	sduo.mutation.Where(ps...)
	return sduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sduo *SourceDataUpdateOne) Select(field string, fields ...string) *SourceDataUpdateOne {
	sduo.fields = append([]string{field}, fields...)
	return sduo
}

// Save executes the query and returns the updated SourceData entity.
func (sduo *SourceDataUpdateOne) Save(ctx context.Context) (*SourceData, error) {
	return withHooks(ctx, sduo.sqlSave, sduo.mutation, sduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sduo *SourceDataUpdateOne) SaveX(ctx context.Context) *SourceData {
	node, err := sduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sduo *SourceDataUpdateOne) Exec(ctx context.Context) error {
	_, err := sduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sduo *SourceDataUpdateOne) ExecX(ctx context.Context) {
	if err := sduo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sduo *SourceDataUpdateOne) sqlSave(ctx context.Context) (_node *SourceData, err error) {
	_spec := sqlgraph.NewUpdateSpec(sourcedata.Table, sourcedata.Columns, sqlgraph.NewFieldSpec(sourcedata.FieldID, field.TypeUUID))
	id, ok := sduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SourceData.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sourcedata.FieldID)
		for _, f := range fields {
			if !sourcedata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sourcedata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sduo.mutation.Format(); ok {
		_spec.SetField(sourcedata.FieldFormat, field.TypeString, value)
	}
	if value, ok := sduo.mutation.Size(); ok {
		_spec.SetField(sourcedata.FieldSize, field.TypeInt64, value)
	}
	if value, ok := sduo.mutation.AddedSize(); ok {
		_spec.AddField(sourcedata.FieldSize, field.TypeInt64, value)
	}
	if value, ok := sduo.mutation.URI(); ok {
		_spec.SetField(sourcedata.FieldURI, field.TypeString, value)
	}
	if sduo.mutation.URICleared() {
		_spec.ClearField(sourcedata.FieldURI, field.TypeString)
	}
	if sduo.mutation.HashesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sourcedata.HashesTable,
			Columns: sourcedata.HashesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hashesentry.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sduo.mutation.RemovedHashesIDs(); len(nodes) > 0 && !sduo.mutation.HashesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sourcedata.HashesTable,
			Columns: sourcedata.HashesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hashesentry.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sduo.mutation.HashesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sourcedata.HashesTable,
			Columns: sourcedata.HashesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hashesentry.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sduo.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   sourcedata.MetadataTable,
			Columns: []string{sourcedata.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sduo.mutation.RemovedMetadataIDs(); len(nodes) > 0 && !sduo.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   sourcedata.MetadataTable,
			Columns: []string{sourcedata.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sduo.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   sourcedata.MetadataTable,
			Columns: []string{sourcedata.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SourceData{config: sduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sourcedata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sduo.mutation.done = true
	return _node, nil
}
