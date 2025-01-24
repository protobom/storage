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
	"github.com/protobom/storage/internal/backends/ent/externalreference"
	"github.com/protobom/storage/internal/backends/ent/hashesentry"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/predicate"
	"github.com/protobom/storage/internal/backends/ent/sourcedata"
)

// HashesEntryUpdate is the builder for updating HashesEntry entities.
type HashesEntryUpdate struct {
	config
	hooks    []Hook
	mutation *HashesEntryMutation
}

// Where appends a list predicates to the HashesEntryUpdate builder.
func (heu *HashesEntryUpdate) Where(ps ...predicate.HashesEntry) *HashesEntryUpdate {
	heu.mutation.Where(ps...)
	return heu
}

// SetHashAlgorithm sets the "hash_algorithm" field.
func (heu *HashesEntryUpdate) SetHashAlgorithm(ha hashesentry.HashAlgorithm) *HashesEntryUpdate {
	heu.mutation.SetHashAlgorithm(ha)
	return heu
}

// SetNillableHashAlgorithm sets the "hash_algorithm" field if the given value is not nil.
func (heu *HashesEntryUpdate) SetNillableHashAlgorithm(ha *hashesentry.HashAlgorithm) *HashesEntryUpdate {
	if ha != nil {
		heu.SetHashAlgorithm(*ha)
	}
	return heu
}

// SetHashData sets the "hash_data" field.
func (heu *HashesEntryUpdate) SetHashData(s string) *HashesEntryUpdate {
	heu.mutation.SetHashData(s)
	return heu
}

// SetNillableHashData sets the "hash_data" field if the given value is not nil.
func (heu *HashesEntryUpdate) SetNillableHashData(s *string) *HashesEntryUpdate {
	if s != nil {
		heu.SetHashData(*s)
	}
	return heu
}

// AddExternalReferenceIDs adds the "external_references" edge to the ExternalReference entity by IDs.
func (heu *HashesEntryUpdate) AddExternalReferenceIDs(ids ...uuid.UUID) *HashesEntryUpdate {
	heu.mutation.AddExternalReferenceIDs(ids...)
	return heu
}

// AddExternalReferences adds the "external_references" edges to the ExternalReference entity.
func (heu *HashesEntryUpdate) AddExternalReferences(e ...*ExternalReference) *HashesEntryUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return heu.AddExternalReferenceIDs(ids...)
}

// AddNodeIDs adds the "nodes" edge to the Node entity by IDs.
func (heu *HashesEntryUpdate) AddNodeIDs(ids ...uuid.UUID) *HashesEntryUpdate {
	heu.mutation.AddNodeIDs(ids...)
	return heu
}

// AddNodes adds the "nodes" edges to the Node entity.
func (heu *HashesEntryUpdate) AddNodes(n ...*Node) *HashesEntryUpdate {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return heu.AddNodeIDs(ids...)
}

// AddSourceDatumIDs adds the "source_data" edge to the SourceData entity by IDs.
func (heu *HashesEntryUpdate) AddSourceDatumIDs(ids ...uuid.UUID) *HashesEntryUpdate {
	heu.mutation.AddSourceDatumIDs(ids...)
	return heu
}

// AddSourceData adds the "source_data" edges to the SourceData entity.
func (heu *HashesEntryUpdate) AddSourceData(s ...*SourceData) *HashesEntryUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return heu.AddSourceDatumIDs(ids...)
}

// Mutation returns the HashesEntryMutation object of the builder.
func (heu *HashesEntryUpdate) Mutation() *HashesEntryMutation {
	return heu.mutation
}

// ClearExternalReferences clears all "external_references" edges to the ExternalReference entity.
func (heu *HashesEntryUpdate) ClearExternalReferences() *HashesEntryUpdate {
	heu.mutation.ClearExternalReferences()
	return heu
}

// RemoveExternalReferenceIDs removes the "external_references" edge to ExternalReference entities by IDs.
func (heu *HashesEntryUpdate) RemoveExternalReferenceIDs(ids ...uuid.UUID) *HashesEntryUpdate {
	heu.mutation.RemoveExternalReferenceIDs(ids...)
	return heu
}

// RemoveExternalReferences removes "external_references" edges to ExternalReference entities.
func (heu *HashesEntryUpdate) RemoveExternalReferences(e ...*ExternalReference) *HashesEntryUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return heu.RemoveExternalReferenceIDs(ids...)
}

// ClearNodes clears all "nodes" edges to the Node entity.
func (heu *HashesEntryUpdate) ClearNodes() *HashesEntryUpdate {
	heu.mutation.ClearNodes()
	return heu
}

// RemoveNodeIDs removes the "nodes" edge to Node entities by IDs.
func (heu *HashesEntryUpdate) RemoveNodeIDs(ids ...uuid.UUID) *HashesEntryUpdate {
	heu.mutation.RemoveNodeIDs(ids...)
	return heu
}

// RemoveNodes removes "nodes" edges to Node entities.
func (heu *HashesEntryUpdate) RemoveNodes(n ...*Node) *HashesEntryUpdate {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return heu.RemoveNodeIDs(ids...)
}

// ClearSourceData clears all "source_data" edges to the SourceData entity.
func (heu *HashesEntryUpdate) ClearSourceData() *HashesEntryUpdate {
	heu.mutation.ClearSourceData()
	return heu
}

// RemoveSourceDatumIDs removes the "source_data" edge to SourceData entities by IDs.
func (heu *HashesEntryUpdate) RemoveSourceDatumIDs(ids ...uuid.UUID) *HashesEntryUpdate {
	heu.mutation.RemoveSourceDatumIDs(ids...)
	return heu
}

// RemoveSourceData removes "source_data" edges to SourceData entities.
func (heu *HashesEntryUpdate) RemoveSourceData(s ...*SourceData) *HashesEntryUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return heu.RemoveSourceDatumIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (heu *HashesEntryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, heu.sqlSave, heu.mutation, heu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (heu *HashesEntryUpdate) SaveX(ctx context.Context) int {
	affected, err := heu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (heu *HashesEntryUpdate) Exec(ctx context.Context) error {
	_, err := heu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (heu *HashesEntryUpdate) ExecX(ctx context.Context) {
	if err := heu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (heu *HashesEntryUpdate) check() error {
	if v, ok := heu.mutation.HashAlgorithm(); ok {
		if err := hashesentry.HashAlgorithmValidator(v); err != nil {
			return &ValidationError{Name: "hash_algorithm", err: fmt.Errorf(`ent: validator failed for field "HashesEntry.hash_algorithm": %w`, err)}
		}
	}
	return nil
}

func (heu *HashesEntryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := heu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(hashesentry.Table, hashesentry.Columns, sqlgraph.NewFieldSpec(hashesentry.FieldID, field.TypeUUID))
	if ps := heu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := heu.mutation.HashAlgorithm(); ok {
		_spec.SetField(hashesentry.FieldHashAlgorithm, field.TypeEnum, value)
	}
	if value, ok := heu.mutation.HashData(); ok {
		_spec.SetField(hashesentry.FieldHashData, field.TypeString, value)
	}
	if heu.mutation.ExternalReferencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hashesentry.ExternalReferencesTable,
			Columns: hashesentry.ExternalReferencesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(externalreference.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heu.mutation.RemovedExternalReferencesIDs(); len(nodes) > 0 && !heu.mutation.ExternalReferencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hashesentry.ExternalReferencesTable,
			Columns: hashesentry.ExternalReferencesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(externalreference.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heu.mutation.ExternalReferencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hashesentry.ExternalReferencesTable,
			Columns: hashesentry.ExternalReferencesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(externalreference.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if heu.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hashesentry.NodesTable,
			Columns: hashesentry.NodesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heu.mutation.RemovedNodesIDs(); len(nodes) > 0 && !heu.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hashesentry.NodesTable,
			Columns: hashesentry.NodesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heu.mutation.NodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hashesentry.NodesTable,
			Columns: hashesentry.NodesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if heu.mutation.SourceDataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hashesentry.SourceDataTable,
			Columns: hashesentry.SourceDataPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcedata.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heu.mutation.RemovedSourceDataIDs(); len(nodes) > 0 && !heu.mutation.SourceDataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hashesentry.SourceDataTable,
			Columns: hashesentry.SourceDataPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcedata.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heu.mutation.SourceDataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hashesentry.SourceDataTable,
			Columns: hashesentry.SourceDataPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcedata.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, heu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hashesentry.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	heu.mutation.done = true
	return n, nil
}

// HashesEntryUpdateOne is the builder for updating a single HashesEntry entity.
type HashesEntryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HashesEntryMutation
}

// SetHashAlgorithm sets the "hash_algorithm" field.
func (heuo *HashesEntryUpdateOne) SetHashAlgorithm(ha hashesentry.HashAlgorithm) *HashesEntryUpdateOne {
	heuo.mutation.SetHashAlgorithm(ha)
	return heuo
}

// SetNillableHashAlgorithm sets the "hash_algorithm" field if the given value is not nil.
func (heuo *HashesEntryUpdateOne) SetNillableHashAlgorithm(ha *hashesentry.HashAlgorithm) *HashesEntryUpdateOne {
	if ha != nil {
		heuo.SetHashAlgorithm(*ha)
	}
	return heuo
}

// SetHashData sets the "hash_data" field.
func (heuo *HashesEntryUpdateOne) SetHashData(s string) *HashesEntryUpdateOne {
	heuo.mutation.SetHashData(s)
	return heuo
}

// SetNillableHashData sets the "hash_data" field if the given value is not nil.
func (heuo *HashesEntryUpdateOne) SetNillableHashData(s *string) *HashesEntryUpdateOne {
	if s != nil {
		heuo.SetHashData(*s)
	}
	return heuo
}

// AddExternalReferenceIDs adds the "external_references" edge to the ExternalReference entity by IDs.
func (heuo *HashesEntryUpdateOne) AddExternalReferenceIDs(ids ...uuid.UUID) *HashesEntryUpdateOne {
	heuo.mutation.AddExternalReferenceIDs(ids...)
	return heuo
}

// AddExternalReferences adds the "external_references" edges to the ExternalReference entity.
func (heuo *HashesEntryUpdateOne) AddExternalReferences(e ...*ExternalReference) *HashesEntryUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return heuo.AddExternalReferenceIDs(ids...)
}

// AddNodeIDs adds the "nodes" edge to the Node entity by IDs.
func (heuo *HashesEntryUpdateOne) AddNodeIDs(ids ...uuid.UUID) *HashesEntryUpdateOne {
	heuo.mutation.AddNodeIDs(ids...)
	return heuo
}

// AddNodes adds the "nodes" edges to the Node entity.
func (heuo *HashesEntryUpdateOne) AddNodes(n ...*Node) *HashesEntryUpdateOne {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return heuo.AddNodeIDs(ids...)
}

// AddSourceDatumIDs adds the "source_data" edge to the SourceData entity by IDs.
func (heuo *HashesEntryUpdateOne) AddSourceDatumIDs(ids ...uuid.UUID) *HashesEntryUpdateOne {
	heuo.mutation.AddSourceDatumIDs(ids...)
	return heuo
}

// AddSourceData adds the "source_data" edges to the SourceData entity.
func (heuo *HashesEntryUpdateOne) AddSourceData(s ...*SourceData) *HashesEntryUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return heuo.AddSourceDatumIDs(ids...)
}

// Mutation returns the HashesEntryMutation object of the builder.
func (heuo *HashesEntryUpdateOne) Mutation() *HashesEntryMutation {
	return heuo.mutation
}

// ClearExternalReferences clears all "external_references" edges to the ExternalReference entity.
func (heuo *HashesEntryUpdateOne) ClearExternalReferences() *HashesEntryUpdateOne {
	heuo.mutation.ClearExternalReferences()
	return heuo
}

// RemoveExternalReferenceIDs removes the "external_references" edge to ExternalReference entities by IDs.
func (heuo *HashesEntryUpdateOne) RemoveExternalReferenceIDs(ids ...uuid.UUID) *HashesEntryUpdateOne {
	heuo.mutation.RemoveExternalReferenceIDs(ids...)
	return heuo
}

// RemoveExternalReferences removes "external_references" edges to ExternalReference entities.
func (heuo *HashesEntryUpdateOne) RemoveExternalReferences(e ...*ExternalReference) *HashesEntryUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return heuo.RemoveExternalReferenceIDs(ids...)
}

// ClearNodes clears all "nodes" edges to the Node entity.
func (heuo *HashesEntryUpdateOne) ClearNodes() *HashesEntryUpdateOne {
	heuo.mutation.ClearNodes()
	return heuo
}

// RemoveNodeIDs removes the "nodes" edge to Node entities by IDs.
func (heuo *HashesEntryUpdateOne) RemoveNodeIDs(ids ...uuid.UUID) *HashesEntryUpdateOne {
	heuo.mutation.RemoveNodeIDs(ids...)
	return heuo
}

// RemoveNodes removes "nodes" edges to Node entities.
func (heuo *HashesEntryUpdateOne) RemoveNodes(n ...*Node) *HashesEntryUpdateOne {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return heuo.RemoveNodeIDs(ids...)
}

// ClearSourceData clears all "source_data" edges to the SourceData entity.
func (heuo *HashesEntryUpdateOne) ClearSourceData() *HashesEntryUpdateOne {
	heuo.mutation.ClearSourceData()
	return heuo
}

// RemoveSourceDatumIDs removes the "source_data" edge to SourceData entities by IDs.
func (heuo *HashesEntryUpdateOne) RemoveSourceDatumIDs(ids ...uuid.UUID) *HashesEntryUpdateOne {
	heuo.mutation.RemoveSourceDatumIDs(ids...)
	return heuo
}

// RemoveSourceData removes "source_data" edges to SourceData entities.
func (heuo *HashesEntryUpdateOne) RemoveSourceData(s ...*SourceData) *HashesEntryUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return heuo.RemoveSourceDatumIDs(ids...)
}

// Where appends a list predicates to the HashesEntryUpdate builder.
func (heuo *HashesEntryUpdateOne) Where(ps ...predicate.HashesEntry) *HashesEntryUpdateOne {
	heuo.mutation.Where(ps...)
	return heuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (heuo *HashesEntryUpdateOne) Select(field string, fields ...string) *HashesEntryUpdateOne {
	heuo.fields = append([]string{field}, fields...)
	return heuo
}

// Save executes the query and returns the updated HashesEntry entity.
func (heuo *HashesEntryUpdateOne) Save(ctx context.Context) (*HashesEntry, error) {
	return withHooks(ctx, heuo.sqlSave, heuo.mutation, heuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (heuo *HashesEntryUpdateOne) SaveX(ctx context.Context) *HashesEntry {
	node, err := heuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (heuo *HashesEntryUpdateOne) Exec(ctx context.Context) error {
	_, err := heuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (heuo *HashesEntryUpdateOne) ExecX(ctx context.Context) {
	if err := heuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (heuo *HashesEntryUpdateOne) check() error {
	if v, ok := heuo.mutation.HashAlgorithm(); ok {
		if err := hashesentry.HashAlgorithmValidator(v); err != nil {
			return &ValidationError{Name: "hash_algorithm", err: fmt.Errorf(`ent: validator failed for field "HashesEntry.hash_algorithm": %w`, err)}
		}
	}
	return nil
}

func (heuo *HashesEntryUpdateOne) sqlSave(ctx context.Context) (_node *HashesEntry, err error) {
	if err := heuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(hashesentry.Table, hashesentry.Columns, sqlgraph.NewFieldSpec(hashesentry.FieldID, field.TypeUUID))
	id, ok := heuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "HashesEntry.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := heuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hashesentry.FieldID)
		for _, f := range fields {
			if !hashesentry.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != hashesentry.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := heuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := heuo.mutation.HashAlgorithm(); ok {
		_spec.SetField(hashesentry.FieldHashAlgorithm, field.TypeEnum, value)
	}
	if value, ok := heuo.mutation.HashData(); ok {
		_spec.SetField(hashesentry.FieldHashData, field.TypeString, value)
	}
	if heuo.mutation.ExternalReferencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hashesentry.ExternalReferencesTable,
			Columns: hashesentry.ExternalReferencesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(externalreference.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heuo.mutation.RemovedExternalReferencesIDs(); len(nodes) > 0 && !heuo.mutation.ExternalReferencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hashesentry.ExternalReferencesTable,
			Columns: hashesentry.ExternalReferencesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(externalreference.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heuo.mutation.ExternalReferencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hashesentry.ExternalReferencesTable,
			Columns: hashesentry.ExternalReferencesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(externalreference.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if heuo.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hashesentry.NodesTable,
			Columns: hashesentry.NodesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heuo.mutation.RemovedNodesIDs(); len(nodes) > 0 && !heuo.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hashesentry.NodesTable,
			Columns: hashesentry.NodesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heuo.mutation.NodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hashesentry.NodesTable,
			Columns: hashesentry.NodesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if heuo.mutation.SourceDataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hashesentry.SourceDataTable,
			Columns: hashesentry.SourceDataPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcedata.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heuo.mutation.RemovedSourceDataIDs(); len(nodes) > 0 && !heuo.mutation.SourceDataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hashesentry.SourceDataTable,
			Columns: hashesentry.SourceDataPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcedata.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := heuo.mutation.SourceDataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hashesentry.SourceDataTable,
			Columns: hashesentry.SourceDataPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcedata.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &HashesEntry{config: heuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, heuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hashesentry.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	heuo.mutation.done = true
	return _node, nil
}
