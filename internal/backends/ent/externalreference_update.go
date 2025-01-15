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
)

// ExternalReferenceUpdate is the builder for updating ExternalReference entities.
type ExternalReferenceUpdate struct {
	config
	hooks    []Hook
	mutation *ExternalReferenceMutation
}

// Where appends a list predicates to the ExternalReferenceUpdate builder.
func (eru *ExternalReferenceUpdate) Where(ps ...predicate.ExternalReference) *ExternalReferenceUpdate {
	eru.mutation.Where(ps...)
	return eru
}

// SetURL sets the "url" field.
func (eru *ExternalReferenceUpdate) SetURL(s string) *ExternalReferenceUpdate {
	eru.mutation.SetURL(s)
	return eru
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (eru *ExternalReferenceUpdate) SetNillableURL(s *string) *ExternalReferenceUpdate {
	if s != nil {
		eru.SetURL(*s)
	}
	return eru
}

// SetComment sets the "comment" field.
func (eru *ExternalReferenceUpdate) SetComment(s string) *ExternalReferenceUpdate {
	eru.mutation.SetComment(s)
	return eru
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (eru *ExternalReferenceUpdate) SetNillableComment(s *string) *ExternalReferenceUpdate {
	if s != nil {
		eru.SetComment(*s)
	}
	return eru
}

// SetAuthority sets the "authority" field.
func (eru *ExternalReferenceUpdate) SetAuthority(s string) *ExternalReferenceUpdate {
	eru.mutation.SetAuthority(s)
	return eru
}

// SetNillableAuthority sets the "authority" field if the given value is not nil.
func (eru *ExternalReferenceUpdate) SetNillableAuthority(s *string) *ExternalReferenceUpdate {
	if s != nil {
		eru.SetAuthority(*s)
	}
	return eru
}

// ClearAuthority clears the value of the "authority" field.
func (eru *ExternalReferenceUpdate) ClearAuthority() *ExternalReferenceUpdate {
	eru.mutation.ClearAuthority()
	return eru
}

// SetType sets the "type" field.
func (eru *ExternalReferenceUpdate) SetType(e externalreference.Type) *ExternalReferenceUpdate {
	eru.mutation.SetType(e)
	return eru
}

// SetNillableType sets the "type" field if the given value is not nil.
func (eru *ExternalReferenceUpdate) SetNillableType(e *externalreference.Type) *ExternalReferenceUpdate {
	if e != nil {
		eru.SetType(*e)
	}
	return eru
}

// AddHashIDs adds the "hashes" edge to the HashesEntry entity by IDs.
func (eru *ExternalReferenceUpdate) AddHashIDs(ids ...uuid.UUID) *ExternalReferenceUpdate {
	eru.mutation.AddHashIDs(ids...)
	return eru
}

// AddHashes adds the "hashes" edges to the HashesEntry entity.
func (eru *ExternalReferenceUpdate) AddHashes(h ...*HashesEntry) *ExternalReferenceUpdate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return eru.AddHashIDs(ids...)
}

// AddNodeIDs adds the "nodes" edge to the Node entity by IDs.
func (eru *ExternalReferenceUpdate) AddNodeIDs(ids ...uuid.UUID) *ExternalReferenceUpdate {
	eru.mutation.AddNodeIDs(ids...)
	return eru
}

// AddNodes adds the "nodes" edges to the Node entity.
func (eru *ExternalReferenceUpdate) AddNodes(n ...*Node) *ExternalReferenceUpdate {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return eru.AddNodeIDs(ids...)
}

// Mutation returns the ExternalReferenceMutation object of the builder.
func (eru *ExternalReferenceUpdate) Mutation() *ExternalReferenceMutation {
	return eru.mutation
}

// ClearHashes clears all "hashes" edges to the HashesEntry entity.
func (eru *ExternalReferenceUpdate) ClearHashes() *ExternalReferenceUpdate {
	eru.mutation.ClearHashes()
	return eru
}

// RemoveHashIDs removes the "hashes" edge to HashesEntry entities by IDs.
func (eru *ExternalReferenceUpdate) RemoveHashIDs(ids ...uuid.UUID) *ExternalReferenceUpdate {
	eru.mutation.RemoveHashIDs(ids...)
	return eru
}

// RemoveHashes removes "hashes" edges to HashesEntry entities.
func (eru *ExternalReferenceUpdate) RemoveHashes(h ...*HashesEntry) *ExternalReferenceUpdate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return eru.RemoveHashIDs(ids...)
}

// ClearNodes clears all "nodes" edges to the Node entity.
func (eru *ExternalReferenceUpdate) ClearNodes() *ExternalReferenceUpdate {
	eru.mutation.ClearNodes()
	return eru
}

// RemoveNodeIDs removes the "nodes" edge to Node entities by IDs.
func (eru *ExternalReferenceUpdate) RemoveNodeIDs(ids ...uuid.UUID) *ExternalReferenceUpdate {
	eru.mutation.RemoveNodeIDs(ids...)
	return eru
}

// RemoveNodes removes "nodes" edges to Node entities.
func (eru *ExternalReferenceUpdate) RemoveNodes(n ...*Node) *ExternalReferenceUpdate {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return eru.RemoveNodeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eru *ExternalReferenceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eru.sqlSave, eru.mutation, eru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eru *ExternalReferenceUpdate) SaveX(ctx context.Context) int {
	affected, err := eru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eru *ExternalReferenceUpdate) Exec(ctx context.Context) error {
	_, err := eru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eru *ExternalReferenceUpdate) ExecX(ctx context.Context) {
	if err := eru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eru *ExternalReferenceUpdate) check() error {
	if v, ok := eru.mutation.GetType(); ok {
		if err := externalreference.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "ExternalReference.type": %w`, err)}
		}
	}
	return nil
}

func (eru *ExternalReferenceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(externalreference.Table, externalreference.Columns, sqlgraph.NewFieldSpec(externalreference.FieldID, field.TypeUUID))
	if ps := eru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eru.mutation.URL(); ok {
		_spec.SetField(externalreference.FieldURL, field.TypeString, value)
	}
	if value, ok := eru.mutation.Comment(); ok {
		_spec.SetField(externalreference.FieldComment, field.TypeString, value)
	}
	if value, ok := eru.mutation.Authority(); ok {
		_spec.SetField(externalreference.FieldAuthority, field.TypeString, value)
	}
	if eru.mutation.AuthorityCleared() {
		_spec.ClearField(externalreference.FieldAuthority, field.TypeString)
	}
	if value, ok := eru.mutation.GetType(); ok {
		_spec.SetField(externalreference.FieldType, field.TypeEnum, value)
	}
	if eru.mutation.HashesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   externalreference.HashesTable,
			Columns: externalreference.HashesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hashesentry.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eru.mutation.RemovedHashesIDs(); len(nodes) > 0 && !eru.mutation.HashesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   externalreference.HashesTable,
			Columns: externalreference.HashesPrimaryKey,
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
	if nodes := eru.mutation.HashesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   externalreference.HashesTable,
			Columns: externalreference.HashesPrimaryKey,
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
	if eru.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   externalreference.NodesTable,
			Columns: externalreference.NodesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eru.mutation.RemovedNodesIDs(); len(nodes) > 0 && !eru.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   externalreference.NodesTable,
			Columns: externalreference.NodesPrimaryKey,
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
	if nodes := eru.mutation.NodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   externalreference.NodesTable,
			Columns: externalreference.NodesPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, eru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{externalreference.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eru.mutation.done = true
	return n, nil
}

// ExternalReferenceUpdateOne is the builder for updating a single ExternalReference entity.
type ExternalReferenceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ExternalReferenceMutation
}

// SetURL sets the "url" field.
func (eruo *ExternalReferenceUpdateOne) SetURL(s string) *ExternalReferenceUpdateOne {
	eruo.mutation.SetURL(s)
	return eruo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (eruo *ExternalReferenceUpdateOne) SetNillableURL(s *string) *ExternalReferenceUpdateOne {
	if s != nil {
		eruo.SetURL(*s)
	}
	return eruo
}

// SetComment sets the "comment" field.
func (eruo *ExternalReferenceUpdateOne) SetComment(s string) *ExternalReferenceUpdateOne {
	eruo.mutation.SetComment(s)
	return eruo
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (eruo *ExternalReferenceUpdateOne) SetNillableComment(s *string) *ExternalReferenceUpdateOne {
	if s != nil {
		eruo.SetComment(*s)
	}
	return eruo
}

// SetAuthority sets the "authority" field.
func (eruo *ExternalReferenceUpdateOne) SetAuthority(s string) *ExternalReferenceUpdateOne {
	eruo.mutation.SetAuthority(s)
	return eruo
}

// SetNillableAuthority sets the "authority" field if the given value is not nil.
func (eruo *ExternalReferenceUpdateOne) SetNillableAuthority(s *string) *ExternalReferenceUpdateOne {
	if s != nil {
		eruo.SetAuthority(*s)
	}
	return eruo
}

// ClearAuthority clears the value of the "authority" field.
func (eruo *ExternalReferenceUpdateOne) ClearAuthority() *ExternalReferenceUpdateOne {
	eruo.mutation.ClearAuthority()
	return eruo
}

// SetType sets the "type" field.
func (eruo *ExternalReferenceUpdateOne) SetType(e externalreference.Type) *ExternalReferenceUpdateOne {
	eruo.mutation.SetType(e)
	return eruo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (eruo *ExternalReferenceUpdateOne) SetNillableType(e *externalreference.Type) *ExternalReferenceUpdateOne {
	if e != nil {
		eruo.SetType(*e)
	}
	return eruo
}

// AddHashIDs adds the "hashes" edge to the HashesEntry entity by IDs.
func (eruo *ExternalReferenceUpdateOne) AddHashIDs(ids ...uuid.UUID) *ExternalReferenceUpdateOne {
	eruo.mutation.AddHashIDs(ids...)
	return eruo
}

// AddHashes adds the "hashes" edges to the HashesEntry entity.
func (eruo *ExternalReferenceUpdateOne) AddHashes(h ...*HashesEntry) *ExternalReferenceUpdateOne {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return eruo.AddHashIDs(ids...)
}

// AddNodeIDs adds the "nodes" edge to the Node entity by IDs.
func (eruo *ExternalReferenceUpdateOne) AddNodeIDs(ids ...uuid.UUID) *ExternalReferenceUpdateOne {
	eruo.mutation.AddNodeIDs(ids...)
	return eruo
}

// AddNodes adds the "nodes" edges to the Node entity.
func (eruo *ExternalReferenceUpdateOne) AddNodes(n ...*Node) *ExternalReferenceUpdateOne {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return eruo.AddNodeIDs(ids...)
}

// Mutation returns the ExternalReferenceMutation object of the builder.
func (eruo *ExternalReferenceUpdateOne) Mutation() *ExternalReferenceMutation {
	return eruo.mutation
}

// ClearHashes clears all "hashes" edges to the HashesEntry entity.
func (eruo *ExternalReferenceUpdateOne) ClearHashes() *ExternalReferenceUpdateOne {
	eruo.mutation.ClearHashes()
	return eruo
}

// RemoveHashIDs removes the "hashes" edge to HashesEntry entities by IDs.
func (eruo *ExternalReferenceUpdateOne) RemoveHashIDs(ids ...uuid.UUID) *ExternalReferenceUpdateOne {
	eruo.mutation.RemoveHashIDs(ids...)
	return eruo
}

// RemoveHashes removes "hashes" edges to HashesEntry entities.
func (eruo *ExternalReferenceUpdateOne) RemoveHashes(h ...*HashesEntry) *ExternalReferenceUpdateOne {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return eruo.RemoveHashIDs(ids...)
}

// ClearNodes clears all "nodes" edges to the Node entity.
func (eruo *ExternalReferenceUpdateOne) ClearNodes() *ExternalReferenceUpdateOne {
	eruo.mutation.ClearNodes()
	return eruo
}

// RemoveNodeIDs removes the "nodes" edge to Node entities by IDs.
func (eruo *ExternalReferenceUpdateOne) RemoveNodeIDs(ids ...uuid.UUID) *ExternalReferenceUpdateOne {
	eruo.mutation.RemoveNodeIDs(ids...)
	return eruo
}

// RemoveNodes removes "nodes" edges to Node entities.
func (eruo *ExternalReferenceUpdateOne) RemoveNodes(n ...*Node) *ExternalReferenceUpdateOne {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return eruo.RemoveNodeIDs(ids...)
}

// Where appends a list predicates to the ExternalReferenceUpdate builder.
func (eruo *ExternalReferenceUpdateOne) Where(ps ...predicate.ExternalReference) *ExternalReferenceUpdateOne {
	eruo.mutation.Where(ps...)
	return eruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (eruo *ExternalReferenceUpdateOne) Select(field string, fields ...string) *ExternalReferenceUpdateOne {
	eruo.fields = append([]string{field}, fields...)
	return eruo
}

// Save executes the query and returns the updated ExternalReference entity.
func (eruo *ExternalReferenceUpdateOne) Save(ctx context.Context) (*ExternalReference, error) {
	return withHooks(ctx, eruo.sqlSave, eruo.mutation, eruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eruo *ExternalReferenceUpdateOne) SaveX(ctx context.Context) *ExternalReference {
	node, err := eruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (eruo *ExternalReferenceUpdateOne) Exec(ctx context.Context) error {
	_, err := eruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eruo *ExternalReferenceUpdateOne) ExecX(ctx context.Context) {
	if err := eruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eruo *ExternalReferenceUpdateOne) check() error {
	if v, ok := eruo.mutation.GetType(); ok {
		if err := externalreference.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "ExternalReference.type": %w`, err)}
		}
	}
	return nil
}

func (eruo *ExternalReferenceUpdateOne) sqlSave(ctx context.Context) (_node *ExternalReference, err error) {
	if err := eruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(externalreference.Table, externalreference.Columns, sqlgraph.NewFieldSpec(externalreference.FieldID, field.TypeUUID))
	id, ok := eruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ExternalReference.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := eruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, externalreference.FieldID)
		for _, f := range fields {
			if !externalreference.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != externalreference.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := eruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eruo.mutation.URL(); ok {
		_spec.SetField(externalreference.FieldURL, field.TypeString, value)
	}
	if value, ok := eruo.mutation.Comment(); ok {
		_spec.SetField(externalreference.FieldComment, field.TypeString, value)
	}
	if value, ok := eruo.mutation.Authority(); ok {
		_spec.SetField(externalreference.FieldAuthority, field.TypeString, value)
	}
	if eruo.mutation.AuthorityCleared() {
		_spec.ClearField(externalreference.FieldAuthority, field.TypeString)
	}
	if value, ok := eruo.mutation.GetType(); ok {
		_spec.SetField(externalreference.FieldType, field.TypeEnum, value)
	}
	if eruo.mutation.HashesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   externalreference.HashesTable,
			Columns: externalreference.HashesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hashesentry.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eruo.mutation.RemovedHashesIDs(); len(nodes) > 0 && !eruo.mutation.HashesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   externalreference.HashesTable,
			Columns: externalreference.HashesPrimaryKey,
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
	if nodes := eruo.mutation.HashesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   externalreference.HashesTable,
			Columns: externalreference.HashesPrimaryKey,
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
	if eruo.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   externalreference.NodesTable,
			Columns: externalreference.NodesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eruo.mutation.RemovedNodesIDs(); len(nodes) > 0 && !eruo.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   externalreference.NodesTable,
			Columns: externalreference.NodesPrimaryKey,
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
	if nodes := eruo.mutation.NodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   externalreference.NodesTable,
			Columns: externalreference.NodesPrimaryKey,
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
	_node = &ExternalReference{config: eruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, eruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{externalreference.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	eruo.mutation.done = true
	return _node, nil
}
