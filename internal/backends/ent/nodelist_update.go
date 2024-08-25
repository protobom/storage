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
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/nodelist"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// NodeListUpdate is the builder for updating NodeList entities.
type NodeListUpdate struct {
	config
	hooks    []Hook
	mutation *NodeListMutation
}

// Where appends a list predicates to the NodeListUpdate builder.
func (nlu *NodeListUpdate) Where(ps ...predicate.NodeList) *NodeListUpdate {
	nlu.mutation.Where(ps...)
	return nlu
}

// SetProtoMessage sets the "proto_message" field.
func (nlu *NodeListUpdate) SetProtoMessage(sl *sbom.NodeList) *NodeListUpdate {
	nlu.mutation.SetProtoMessage(sl)
	return nlu
}

// ClearProtoMessage clears the value of the "proto_message" field.
func (nlu *NodeListUpdate) ClearProtoMessage() *NodeListUpdate {
	nlu.mutation.ClearProtoMessage()
	return nlu
}

// SetRootElements sets the "root_elements" field.
func (nlu *NodeListUpdate) SetRootElements(s []string) *NodeListUpdate {
	nlu.mutation.SetRootElements(s)
	return nlu
}

// AppendRootElements appends s to the "root_elements" field.
func (nlu *NodeListUpdate) AppendRootElements(s []string) *NodeListUpdate {
	nlu.mutation.AppendRootElements(s)
	return nlu
}

// AddNodeIDs adds the "nodes" edge to the Node entity by IDs.
func (nlu *NodeListUpdate) AddNodeIDs(ids ...string) *NodeListUpdate {
	nlu.mutation.AddNodeIDs(ids...)
	return nlu
}

// AddNodes adds the "nodes" edges to the Node entity.
func (nlu *NodeListUpdate) AddNodes(n ...*Node) *NodeListUpdate {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nlu.AddNodeIDs(ids...)
}

// Mutation returns the NodeListMutation object of the builder.
func (nlu *NodeListUpdate) Mutation() *NodeListMutation {
	return nlu.mutation
}

// ClearNodes clears all "nodes" edges to the Node entity.
func (nlu *NodeListUpdate) ClearNodes() *NodeListUpdate {
	nlu.mutation.ClearNodes()
	return nlu
}

// RemoveNodeIDs removes the "nodes" edge to Node entities by IDs.
func (nlu *NodeListUpdate) RemoveNodeIDs(ids ...string) *NodeListUpdate {
	nlu.mutation.RemoveNodeIDs(ids...)
	return nlu
}

// RemoveNodes removes "nodes" edges to Node entities.
func (nlu *NodeListUpdate) RemoveNodes(n ...*Node) *NodeListUpdate {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nlu.RemoveNodeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nlu *NodeListUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, nlu.sqlSave, nlu.mutation, nlu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nlu *NodeListUpdate) SaveX(ctx context.Context) int {
	affected, err := nlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nlu *NodeListUpdate) Exec(ctx context.Context) error {
	_, err := nlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nlu *NodeListUpdate) ExecX(ctx context.Context) {
	if err := nlu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nlu *NodeListUpdate) check() error {
	if nlu.mutation.DocumentCleared() && len(nlu.mutation.DocumentIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "NodeList.document"`)
	}
	return nil
}

func (nlu *NodeListUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := nlu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(nodelist.Table, nodelist.Columns, sqlgraph.NewFieldSpec(nodelist.FieldID, field.TypeUUID))
	if ps := nlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nlu.mutation.ProtoMessage(); ok {
		_spec.SetField(nodelist.FieldProtoMessage, field.TypeJSON, value)
	}
	if nlu.mutation.ProtoMessageCleared() {
		_spec.ClearField(nodelist.FieldProtoMessage, field.TypeJSON)
	}
	if value, ok := nlu.mutation.RootElements(); ok {
		_spec.SetField(nodelist.FieldRootElements, field.TypeJSON, value)
	}
	if value, ok := nlu.mutation.AppendedRootElements(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, nodelist.FieldRootElements, value)
		})
	}
	if nlu.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodelist.NodesTable,
			Columns: []string{nodelist.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nlu.mutation.RemovedNodesIDs(); len(nodes) > 0 && !nlu.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodelist.NodesTable,
			Columns: []string{nodelist.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nlu.mutation.NodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodelist.NodesTable,
			Columns: []string{nodelist.NodesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, nlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nodelist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nlu.mutation.done = true
	return n, nil
}

// NodeListUpdateOne is the builder for updating a single NodeList entity.
type NodeListUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NodeListMutation
}

// SetProtoMessage sets the "proto_message" field.
func (nluo *NodeListUpdateOne) SetProtoMessage(sl *sbom.NodeList) *NodeListUpdateOne {
	nluo.mutation.SetProtoMessage(sl)
	return nluo
}

// ClearProtoMessage clears the value of the "proto_message" field.
func (nluo *NodeListUpdateOne) ClearProtoMessage() *NodeListUpdateOne {
	nluo.mutation.ClearProtoMessage()
	return nluo
}

// SetRootElements sets the "root_elements" field.
func (nluo *NodeListUpdateOne) SetRootElements(s []string) *NodeListUpdateOne {
	nluo.mutation.SetRootElements(s)
	return nluo
}

// AppendRootElements appends s to the "root_elements" field.
func (nluo *NodeListUpdateOne) AppendRootElements(s []string) *NodeListUpdateOne {
	nluo.mutation.AppendRootElements(s)
	return nluo
}

// AddNodeIDs adds the "nodes" edge to the Node entity by IDs.
func (nluo *NodeListUpdateOne) AddNodeIDs(ids ...string) *NodeListUpdateOne {
	nluo.mutation.AddNodeIDs(ids...)
	return nluo
}

// AddNodes adds the "nodes" edges to the Node entity.
func (nluo *NodeListUpdateOne) AddNodes(n ...*Node) *NodeListUpdateOne {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nluo.AddNodeIDs(ids...)
}

// Mutation returns the NodeListMutation object of the builder.
func (nluo *NodeListUpdateOne) Mutation() *NodeListMutation {
	return nluo.mutation
}

// ClearNodes clears all "nodes" edges to the Node entity.
func (nluo *NodeListUpdateOne) ClearNodes() *NodeListUpdateOne {
	nluo.mutation.ClearNodes()
	return nluo
}

// RemoveNodeIDs removes the "nodes" edge to Node entities by IDs.
func (nluo *NodeListUpdateOne) RemoveNodeIDs(ids ...string) *NodeListUpdateOne {
	nluo.mutation.RemoveNodeIDs(ids...)
	return nluo
}

// RemoveNodes removes "nodes" edges to Node entities.
func (nluo *NodeListUpdateOne) RemoveNodes(n ...*Node) *NodeListUpdateOne {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nluo.RemoveNodeIDs(ids...)
}

// Where appends a list predicates to the NodeListUpdate builder.
func (nluo *NodeListUpdateOne) Where(ps ...predicate.NodeList) *NodeListUpdateOne {
	nluo.mutation.Where(ps...)
	return nluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nluo *NodeListUpdateOne) Select(field string, fields ...string) *NodeListUpdateOne {
	nluo.fields = append([]string{field}, fields...)
	return nluo
}

// Save executes the query and returns the updated NodeList entity.
func (nluo *NodeListUpdateOne) Save(ctx context.Context) (*NodeList, error) {
	return withHooks(ctx, nluo.sqlSave, nluo.mutation, nluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nluo *NodeListUpdateOne) SaveX(ctx context.Context) *NodeList {
	node, err := nluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nluo *NodeListUpdateOne) Exec(ctx context.Context) error {
	_, err := nluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nluo *NodeListUpdateOne) ExecX(ctx context.Context) {
	if err := nluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nluo *NodeListUpdateOne) check() error {
	if nluo.mutation.DocumentCleared() && len(nluo.mutation.DocumentIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "NodeList.document"`)
	}
	return nil
}

func (nluo *NodeListUpdateOne) sqlSave(ctx context.Context) (_node *NodeList, err error) {
	if err := nluo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(nodelist.Table, nodelist.Columns, sqlgraph.NewFieldSpec(nodelist.FieldID, field.TypeUUID))
	id, ok := nluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NodeList.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nodelist.FieldID)
		for _, f := range fields {
			if !nodelist.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != nodelist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nluo.mutation.ProtoMessage(); ok {
		_spec.SetField(nodelist.FieldProtoMessage, field.TypeJSON, value)
	}
	if nluo.mutation.ProtoMessageCleared() {
		_spec.ClearField(nodelist.FieldProtoMessage, field.TypeJSON)
	}
	if value, ok := nluo.mutation.RootElements(); ok {
		_spec.SetField(nodelist.FieldRootElements, field.TypeJSON, value)
	}
	if value, ok := nluo.mutation.AppendedRootElements(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, nodelist.FieldRootElements, value)
		})
	}
	if nluo.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodelist.NodesTable,
			Columns: []string{nodelist.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nluo.mutation.RemovedNodesIDs(); len(nodes) > 0 && !nluo.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodelist.NodesTable,
			Columns: []string{nodelist.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nluo.mutation.NodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   nodelist.NodesTable,
			Columns: []string{nodelist.NodesColumn},
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
	_node = &NodeList{config: nluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nodelist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nluo.mutation.done = true
	return _node, nil
}
