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
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/predicate"
	"github.com/protobom/storage/internal/backends/ent/property"
)

// PropertyUpdate is the builder for updating Property entities.
type PropertyUpdate struct {
	config
	hooks    []Hook
	mutation *PropertyMutation
}

// Where appends a list predicates to the PropertyUpdate builder.
func (pu *PropertyUpdate) Where(ps ...predicate.Property) *PropertyUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetProtoMessage sets the "proto_message" field.
func (pu *PropertyUpdate) SetProtoMessage(s *sbom.Property) *PropertyUpdate {
	pu.mutation.SetProtoMessage(s)
	return pu
}

// ClearProtoMessage clears the value of the "proto_message" field.
func (pu *PropertyUpdate) ClearProtoMessage() *PropertyUpdate {
	pu.mutation.ClearProtoMessage()
	return pu
}

// SetNodeID sets the "node_id" field.
func (pu *PropertyUpdate) SetNodeID(s string) *PropertyUpdate {
	pu.mutation.SetNodeID(s)
	return pu
}

// SetNillableNodeID sets the "node_id" field if the given value is not nil.
func (pu *PropertyUpdate) SetNillableNodeID(s *string) *PropertyUpdate {
	if s != nil {
		pu.SetNodeID(*s)
	}
	return pu
}

// ClearNodeID clears the value of the "node_id" field.
func (pu *PropertyUpdate) ClearNodeID() *PropertyUpdate {
	pu.mutation.ClearNodeID()
	return pu
}

// SetName sets the "name" field.
func (pu *PropertyUpdate) SetName(s string) *PropertyUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PropertyUpdate) SetNillableName(s *string) *PropertyUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetData sets the "data" field.
func (pu *PropertyUpdate) SetData(s string) *PropertyUpdate {
	pu.mutation.SetData(s)
	return pu
}

// SetNillableData sets the "data" field if the given value is not nil.
func (pu *PropertyUpdate) SetNillableData(s *string) *PropertyUpdate {
	if s != nil {
		pu.SetData(*s)
	}
	return pu
}

// SetNode sets the "node" edge to the Node entity.
func (pu *PropertyUpdate) SetNode(n *Node) *PropertyUpdate {
	return pu.SetNodeID(n.ID)
}

// Mutation returns the PropertyMutation object of the builder.
func (pu *PropertyUpdate) Mutation() *PropertyMutation {
	return pu.mutation
}

// ClearNode clears the "node" edge to the Node entity.
func (pu *PropertyUpdate) ClearNode() *PropertyUpdate {
	pu.mutation.ClearNode()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PropertyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PropertyUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PropertyUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PropertyUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PropertyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(property.Table, property.Columns, sqlgraph.NewFieldSpec(property.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.ProtoMessage(); ok {
		_spec.SetField(property.FieldProtoMessage, field.TypeBytes, value)
	}
	if pu.mutation.ProtoMessageCleared() {
		_spec.ClearField(property.FieldProtoMessage, field.TypeBytes)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(property.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Data(); ok {
		_spec.SetField(property.FieldData, field.TypeString, value)
	}
	if pu.mutation.NodeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   property.NodeTable,
			Columns: []string{property.NodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.NodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   property.NodeTable,
			Columns: []string{property.NodeColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{property.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PropertyUpdateOne is the builder for updating a single Property entity.
type PropertyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PropertyMutation
}

// SetProtoMessage sets the "proto_message" field.
func (puo *PropertyUpdateOne) SetProtoMessage(s *sbom.Property) *PropertyUpdateOne {
	puo.mutation.SetProtoMessage(s)
	return puo
}

// ClearProtoMessage clears the value of the "proto_message" field.
func (puo *PropertyUpdateOne) ClearProtoMessage() *PropertyUpdateOne {
	puo.mutation.ClearProtoMessage()
	return puo
}

// SetNodeID sets the "node_id" field.
func (puo *PropertyUpdateOne) SetNodeID(s string) *PropertyUpdateOne {
	puo.mutation.SetNodeID(s)
	return puo
}

// SetNillableNodeID sets the "node_id" field if the given value is not nil.
func (puo *PropertyUpdateOne) SetNillableNodeID(s *string) *PropertyUpdateOne {
	if s != nil {
		puo.SetNodeID(*s)
	}
	return puo
}

// ClearNodeID clears the value of the "node_id" field.
func (puo *PropertyUpdateOne) ClearNodeID() *PropertyUpdateOne {
	puo.mutation.ClearNodeID()
	return puo
}

// SetName sets the "name" field.
func (puo *PropertyUpdateOne) SetName(s string) *PropertyUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PropertyUpdateOne) SetNillableName(s *string) *PropertyUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetData sets the "data" field.
func (puo *PropertyUpdateOne) SetData(s string) *PropertyUpdateOne {
	puo.mutation.SetData(s)
	return puo
}

// SetNillableData sets the "data" field if the given value is not nil.
func (puo *PropertyUpdateOne) SetNillableData(s *string) *PropertyUpdateOne {
	if s != nil {
		puo.SetData(*s)
	}
	return puo
}

// SetNode sets the "node" edge to the Node entity.
func (puo *PropertyUpdateOne) SetNode(n *Node) *PropertyUpdateOne {
	return puo.SetNodeID(n.ID)
}

// Mutation returns the PropertyMutation object of the builder.
func (puo *PropertyUpdateOne) Mutation() *PropertyMutation {
	return puo.mutation
}

// ClearNode clears the "node" edge to the Node entity.
func (puo *PropertyUpdateOne) ClearNode() *PropertyUpdateOne {
	puo.mutation.ClearNode()
	return puo
}

// Where appends a list predicates to the PropertyUpdate builder.
func (puo *PropertyUpdateOne) Where(ps ...predicate.Property) *PropertyUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PropertyUpdateOne) Select(field string, fields ...string) *PropertyUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Property entity.
func (puo *PropertyUpdateOne) Save(ctx context.Context) (*Property, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PropertyUpdateOne) SaveX(ctx context.Context) *Property {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PropertyUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PropertyUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PropertyUpdateOne) sqlSave(ctx context.Context) (_node *Property, err error) {
	_spec := sqlgraph.NewUpdateSpec(property.Table, property.Columns, sqlgraph.NewFieldSpec(property.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Property.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, property.FieldID)
		for _, f := range fields {
			if !property.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != property.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.ProtoMessage(); ok {
		_spec.SetField(property.FieldProtoMessage, field.TypeBytes, value)
	}
	if puo.mutation.ProtoMessageCleared() {
		_spec.ClearField(property.FieldProtoMessage, field.TypeBytes)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(property.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Data(); ok {
		_spec.SetField(property.FieldData, field.TypeString, value)
	}
	if puo.mutation.NodeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   property.NodeTable,
			Columns: []string{property.NodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.NodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   property.NodeTable,
			Columns: []string{property.NodeColumn},
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
	_node = &Property{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{property.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
