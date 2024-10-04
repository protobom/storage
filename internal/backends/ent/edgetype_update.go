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
	"github.com/protobom/storage/internal/backends/ent/edgetype"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// EdgeTypeUpdate is the builder for updating EdgeType entities.
type EdgeTypeUpdate struct {
	config
	hooks    []Hook
	mutation *EdgeTypeMutation
}

// Where appends a list predicates to the EdgeTypeUpdate builder.
func (etu *EdgeTypeUpdate) Where(ps ...predicate.EdgeType) *EdgeTypeUpdate {
	etu.mutation.Where(ps...)
	return etu
}

// SetType sets the "type" field.
func (etu *EdgeTypeUpdate) SetType(e edgetype.Type) *EdgeTypeUpdate {
	etu.mutation.SetType(e)
	return etu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (etu *EdgeTypeUpdate) SetNillableType(e *edgetype.Type) *EdgeTypeUpdate {
	if e != nil {
		etu.SetType(*e)
	}
	return etu
}

// SetNodeID sets the "node_id" field.
func (etu *EdgeTypeUpdate) SetNodeID(s string) *EdgeTypeUpdate {
	etu.mutation.SetNodeID(s)
	return etu
}

// SetNillableNodeID sets the "node_id" field if the given value is not nil.
func (etu *EdgeTypeUpdate) SetNillableNodeID(s *string) *EdgeTypeUpdate {
	if s != nil {
		etu.SetNodeID(*s)
	}
	return etu
}

// SetToNodeID sets the "to_node_id" field.
func (etu *EdgeTypeUpdate) SetToNodeID(s string) *EdgeTypeUpdate {
	etu.mutation.SetToNodeID(s)
	return etu
}

// SetNillableToNodeID sets the "to_node_id" field if the given value is not nil.
func (etu *EdgeTypeUpdate) SetNillableToNodeID(s *string) *EdgeTypeUpdate {
	if s != nil {
		etu.SetToNodeID(*s)
	}
	return etu
}

// SetFromID sets the "from" edge to the Node entity by ID.
func (etu *EdgeTypeUpdate) SetFromID(id string) *EdgeTypeUpdate {
	etu.mutation.SetFromID(id)
	return etu
}

// SetFrom sets the "from" edge to the Node entity.
func (etu *EdgeTypeUpdate) SetFrom(n *Node) *EdgeTypeUpdate {
	return etu.SetFromID(n.ID)
}

// SetToID sets the "to" edge to the Node entity by ID.
func (etu *EdgeTypeUpdate) SetToID(id string) *EdgeTypeUpdate {
	etu.mutation.SetToID(id)
	return etu
}

// SetTo sets the "to" edge to the Node entity.
func (etu *EdgeTypeUpdate) SetTo(n *Node) *EdgeTypeUpdate {
	return etu.SetToID(n.ID)
}

// Mutation returns the EdgeTypeMutation object of the builder.
func (etu *EdgeTypeUpdate) Mutation() *EdgeTypeMutation {
	return etu.mutation
}

// ClearFrom clears the "from" edge to the Node entity.
func (etu *EdgeTypeUpdate) ClearFrom() *EdgeTypeUpdate {
	etu.mutation.ClearFrom()
	return etu
}

// ClearTo clears the "to" edge to the Node entity.
func (etu *EdgeTypeUpdate) ClearTo() *EdgeTypeUpdate {
	etu.mutation.ClearTo()
	return etu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (etu *EdgeTypeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, etu.sqlSave, etu.mutation, etu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (etu *EdgeTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := etu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (etu *EdgeTypeUpdate) Exec(ctx context.Context) error {
	_, err := etu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etu *EdgeTypeUpdate) ExecX(ctx context.Context) {
	if err := etu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (etu *EdgeTypeUpdate) check() error {
	if v, ok := etu.mutation.GetType(); ok {
		if err := edgetype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "EdgeType.type": %w`, err)}
		}
	}
	if etu.mutation.FromCleared() && len(etu.mutation.FromIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "EdgeType.from"`)
	}
	if etu.mutation.ToCleared() && len(etu.mutation.ToIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "EdgeType.to"`)
	}
	return nil
}

func (etu *EdgeTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := etu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(edgetype.Table, edgetype.Columns, sqlgraph.NewFieldSpec(edgetype.FieldID, field.TypeInt))
	if ps := etu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := etu.mutation.GetType(); ok {
		_spec.SetField(edgetype.FieldType, field.TypeEnum, value)
	}
	if etu.mutation.FromCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   edgetype.FromTable,
			Columns: []string{edgetype.FromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := etu.mutation.FromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   edgetype.FromTable,
			Columns: []string{edgetype.FromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if etu.mutation.ToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   edgetype.ToTable,
			Columns: []string{edgetype.ToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := etu.mutation.ToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   edgetype.ToTable,
			Columns: []string{edgetype.ToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, etu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{edgetype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	etu.mutation.done = true
	return n, nil
}

// EdgeTypeUpdateOne is the builder for updating a single EdgeType entity.
type EdgeTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EdgeTypeMutation
}

// SetType sets the "type" field.
func (etuo *EdgeTypeUpdateOne) SetType(e edgetype.Type) *EdgeTypeUpdateOne {
	etuo.mutation.SetType(e)
	return etuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (etuo *EdgeTypeUpdateOne) SetNillableType(e *edgetype.Type) *EdgeTypeUpdateOne {
	if e != nil {
		etuo.SetType(*e)
	}
	return etuo
}

// SetNodeID sets the "node_id" field.
func (etuo *EdgeTypeUpdateOne) SetNodeID(s string) *EdgeTypeUpdateOne {
	etuo.mutation.SetNodeID(s)
	return etuo
}

// SetNillableNodeID sets the "node_id" field if the given value is not nil.
func (etuo *EdgeTypeUpdateOne) SetNillableNodeID(s *string) *EdgeTypeUpdateOne {
	if s != nil {
		etuo.SetNodeID(*s)
	}
	return etuo
}

// SetToNodeID sets the "to_node_id" field.
func (etuo *EdgeTypeUpdateOne) SetToNodeID(s string) *EdgeTypeUpdateOne {
	etuo.mutation.SetToNodeID(s)
	return etuo
}

// SetNillableToNodeID sets the "to_node_id" field if the given value is not nil.
func (etuo *EdgeTypeUpdateOne) SetNillableToNodeID(s *string) *EdgeTypeUpdateOne {
	if s != nil {
		etuo.SetToNodeID(*s)
	}
	return etuo
}

// SetFromID sets the "from" edge to the Node entity by ID.
func (etuo *EdgeTypeUpdateOne) SetFromID(id string) *EdgeTypeUpdateOne {
	etuo.mutation.SetFromID(id)
	return etuo
}

// SetFrom sets the "from" edge to the Node entity.
func (etuo *EdgeTypeUpdateOne) SetFrom(n *Node) *EdgeTypeUpdateOne {
	return etuo.SetFromID(n.ID)
}

// SetToID sets the "to" edge to the Node entity by ID.
func (etuo *EdgeTypeUpdateOne) SetToID(id string) *EdgeTypeUpdateOne {
	etuo.mutation.SetToID(id)
	return etuo
}

// SetTo sets the "to" edge to the Node entity.
func (etuo *EdgeTypeUpdateOne) SetTo(n *Node) *EdgeTypeUpdateOne {
	return etuo.SetToID(n.ID)
}

// Mutation returns the EdgeTypeMutation object of the builder.
func (etuo *EdgeTypeUpdateOne) Mutation() *EdgeTypeMutation {
	return etuo.mutation
}

// ClearFrom clears the "from" edge to the Node entity.
func (etuo *EdgeTypeUpdateOne) ClearFrom() *EdgeTypeUpdateOne {
	etuo.mutation.ClearFrom()
	return etuo
}

// ClearTo clears the "to" edge to the Node entity.
func (etuo *EdgeTypeUpdateOne) ClearTo() *EdgeTypeUpdateOne {
	etuo.mutation.ClearTo()
	return etuo
}

// Where appends a list predicates to the EdgeTypeUpdate builder.
func (etuo *EdgeTypeUpdateOne) Where(ps ...predicate.EdgeType) *EdgeTypeUpdateOne {
	etuo.mutation.Where(ps...)
	return etuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (etuo *EdgeTypeUpdateOne) Select(field string, fields ...string) *EdgeTypeUpdateOne {
	etuo.fields = append([]string{field}, fields...)
	return etuo
}

// Save executes the query and returns the updated EdgeType entity.
func (etuo *EdgeTypeUpdateOne) Save(ctx context.Context) (*EdgeType, error) {
	return withHooks(ctx, etuo.sqlSave, etuo.mutation, etuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (etuo *EdgeTypeUpdateOne) SaveX(ctx context.Context) *EdgeType {
	node, err := etuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (etuo *EdgeTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := etuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etuo *EdgeTypeUpdateOne) ExecX(ctx context.Context) {
	if err := etuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (etuo *EdgeTypeUpdateOne) check() error {
	if v, ok := etuo.mutation.GetType(); ok {
		if err := edgetype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "EdgeType.type": %w`, err)}
		}
	}
	if etuo.mutation.FromCleared() && len(etuo.mutation.FromIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "EdgeType.from"`)
	}
	if etuo.mutation.ToCleared() && len(etuo.mutation.ToIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "EdgeType.to"`)
	}
	return nil
}

func (etuo *EdgeTypeUpdateOne) sqlSave(ctx context.Context) (_node *EdgeType, err error) {
	if err := etuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(edgetype.Table, edgetype.Columns, sqlgraph.NewFieldSpec(edgetype.FieldID, field.TypeInt))
	id, ok := etuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EdgeType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := etuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, edgetype.FieldID)
		for _, f := range fields {
			if !edgetype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != edgetype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := etuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := etuo.mutation.GetType(); ok {
		_spec.SetField(edgetype.FieldType, field.TypeEnum, value)
	}
	if etuo.mutation.FromCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   edgetype.FromTable,
			Columns: []string{edgetype.FromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := etuo.mutation.FromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   edgetype.FromTable,
			Columns: []string{edgetype.FromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if etuo.mutation.ToCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   edgetype.ToTable,
			Columns: []string{edgetype.ToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := etuo.mutation.ToIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   edgetype.ToTable,
			Columns: []string{edgetype.ToColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EdgeType{config: etuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, etuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{edgetype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	etuo.mutation.done = true
	return _node, nil
}
