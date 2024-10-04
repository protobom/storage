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

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/protobom/protobom/pkg/sbom"
	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/nodelist"
)

// NodeListCreate is the builder for creating a NodeList entity.
type NodeListCreate struct {
	config
	mutation *NodeListMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetProtoMessage sets the "proto_message" field.
func (nlc *NodeListCreate) SetProtoMessage(sl *sbom.NodeList) *NodeListCreate {
	nlc.mutation.SetProtoMessage(sl)
	return nlc
}

// SetRootElements sets the "root_elements" field.
func (nlc *NodeListCreate) SetRootElements(s []string) *NodeListCreate {
	nlc.mutation.SetRootElements(s)
	return nlc
}

// SetID sets the "id" field.
func (nlc *NodeListCreate) SetID(u uuid.UUID) *NodeListCreate {
	nlc.mutation.SetID(u)
	return nlc
}

// AddNodeIDs adds the "nodes" edge to the Node entity by IDs.
func (nlc *NodeListCreate) AddNodeIDs(ids ...string) *NodeListCreate {
	nlc.mutation.AddNodeIDs(ids...)
	return nlc
}

// AddNodes adds the "nodes" edges to the Node entity.
func (nlc *NodeListCreate) AddNodes(n ...*Node) *NodeListCreate {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nlc.AddNodeIDs(ids...)
}

// SetDocumentID sets the "document" edge to the Document entity by ID.
func (nlc *NodeListCreate) SetDocumentID(id uuid.UUID) *NodeListCreate {
	nlc.mutation.SetDocumentID(id)
	return nlc
}

// SetDocument sets the "document" edge to the Document entity.
func (nlc *NodeListCreate) SetDocument(d *Document) *NodeListCreate {
	return nlc.SetDocumentID(d.ID)
}

// Mutation returns the NodeListMutation object of the builder.
func (nlc *NodeListCreate) Mutation() *NodeListMutation {
	return nlc.mutation
}

// Save creates the NodeList in the database.
func (nlc *NodeListCreate) Save(ctx context.Context) (*NodeList, error) {
	return withHooks(ctx, nlc.sqlSave, nlc.mutation, nlc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nlc *NodeListCreate) SaveX(ctx context.Context) *NodeList {
	v, err := nlc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nlc *NodeListCreate) Exec(ctx context.Context) error {
	_, err := nlc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nlc *NodeListCreate) ExecX(ctx context.Context) {
	if err := nlc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nlc *NodeListCreate) check() error {
	if _, ok := nlc.mutation.RootElements(); !ok {
		return &ValidationError{Name: "root_elements", err: errors.New(`ent: missing required field "NodeList.root_elements"`)}
	}
	if len(nlc.mutation.DocumentIDs()) == 0 {
		return &ValidationError{Name: "document", err: errors.New(`ent: missing required edge "NodeList.document"`)}
	}
	return nil
}

func (nlc *NodeListCreate) sqlSave(ctx context.Context) (*NodeList, error) {
	if err := nlc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nlc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nlc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	nlc.mutation.id = &_node.ID
	nlc.mutation.done = true
	return _node, nil
}

func (nlc *NodeListCreate) createSpec() (*NodeList, *sqlgraph.CreateSpec) {
	var (
		_node = &NodeList{config: nlc.config}
		_spec = sqlgraph.NewCreateSpec(nodelist.Table, sqlgraph.NewFieldSpec(nodelist.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = nlc.conflict
	if id, ok := nlc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := nlc.mutation.ProtoMessage(); ok {
		_spec.SetField(nodelist.FieldProtoMessage, field.TypeJSON, value)
		_node.ProtoMessage = value
	}
	if value, ok := nlc.mutation.RootElements(); ok {
		_spec.SetField(nodelist.FieldRootElements, field.TypeJSON, value)
		_node.RootElements = value
	}
	if nodes := nlc.mutation.NodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   nodelist.NodesTable,
			Columns: nodelist.NodesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nlc.mutation.DocumentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   nodelist.DocumentTable,
			Columns: []string{nodelist.DocumentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(document.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.NodeList.Create().
//		SetProtoMessage(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NodeListUpsert) {
//			SetProtoMessage(v+v).
//		}).
//		Exec(ctx)
func (nlc *NodeListCreate) OnConflict(opts ...sql.ConflictOption) *NodeListUpsertOne {
	nlc.conflict = opts
	return &NodeListUpsertOne{
		create: nlc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.NodeList.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (nlc *NodeListCreate) OnConflictColumns(columns ...string) *NodeListUpsertOne {
	nlc.conflict = append(nlc.conflict, sql.ConflictColumns(columns...))
	return &NodeListUpsertOne{
		create: nlc,
	}
}

type (
	// NodeListUpsertOne is the builder for "upsert"-ing
	//  one NodeList node.
	NodeListUpsertOne struct {
		create *NodeListCreate
	}

	// NodeListUpsert is the "OnConflict" setter.
	NodeListUpsert struct {
		*sql.UpdateSet
	}
)

// SetProtoMessage sets the "proto_message" field.
func (u *NodeListUpsert) SetProtoMessage(v *sbom.NodeList) *NodeListUpsert {
	u.Set(nodelist.FieldProtoMessage, v)
	return u
}

// UpdateProtoMessage sets the "proto_message" field to the value that was provided on create.
func (u *NodeListUpsert) UpdateProtoMessage() *NodeListUpsert {
	u.SetExcluded(nodelist.FieldProtoMessage)
	return u
}

// ClearProtoMessage clears the value of the "proto_message" field.
func (u *NodeListUpsert) ClearProtoMessage() *NodeListUpsert {
	u.SetNull(nodelist.FieldProtoMessage)
	return u
}

// SetRootElements sets the "root_elements" field.
func (u *NodeListUpsert) SetRootElements(v []string) *NodeListUpsert {
	u.Set(nodelist.FieldRootElements, v)
	return u
}

// UpdateRootElements sets the "root_elements" field to the value that was provided on create.
func (u *NodeListUpsert) UpdateRootElements() *NodeListUpsert {
	u.SetExcluded(nodelist.FieldRootElements)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.NodeList.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(nodelist.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *NodeListUpsertOne) UpdateNewValues() *NodeListUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(nodelist.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.NodeList.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *NodeListUpsertOne) Ignore() *NodeListUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NodeListUpsertOne) DoNothing() *NodeListUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NodeListCreate.OnConflict
// documentation for more info.
func (u *NodeListUpsertOne) Update(set func(*NodeListUpsert)) *NodeListUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NodeListUpsert{UpdateSet: update})
	}))
	return u
}

// SetProtoMessage sets the "proto_message" field.
func (u *NodeListUpsertOne) SetProtoMessage(v *sbom.NodeList) *NodeListUpsertOne {
	return u.Update(func(s *NodeListUpsert) {
		s.SetProtoMessage(v)
	})
}

// UpdateProtoMessage sets the "proto_message" field to the value that was provided on create.
func (u *NodeListUpsertOne) UpdateProtoMessage() *NodeListUpsertOne {
	return u.Update(func(s *NodeListUpsert) {
		s.UpdateProtoMessage()
	})
}

// ClearProtoMessage clears the value of the "proto_message" field.
func (u *NodeListUpsertOne) ClearProtoMessage() *NodeListUpsertOne {
	return u.Update(func(s *NodeListUpsert) {
		s.ClearProtoMessage()
	})
}

// SetRootElements sets the "root_elements" field.
func (u *NodeListUpsertOne) SetRootElements(v []string) *NodeListUpsertOne {
	return u.Update(func(s *NodeListUpsert) {
		s.SetRootElements(v)
	})
}

// UpdateRootElements sets the "root_elements" field to the value that was provided on create.
func (u *NodeListUpsertOne) UpdateRootElements() *NodeListUpsertOne {
	return u.Update(func(s *NodeListUpsert) {
		s.UpdateRootElements()
	})
}

// Exec executes the query.
func (u *NodeListUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NodeListCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NodeListUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *NodeListUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: NodeListUpsertOne.ID is not supported by MySQL driver. Use NodeListUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *NodeListUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// NodeListCreateBulk is the builder for creating many NodeList entities in bulk.
type NodeListCreateBulk struct {
	config
	err      error
	builders []*NodeListCreate
	conflict []sql.ConflictOption
}

// Save creates the NodeList entities in the database.
func (nlcb *NodeListCreateBulk) Save(ctx context.Context) ([]*NodeList, error) {
	if nlcb.err != nil {
		return nil, nlcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(nlcb.builders))
	nodes := make([]*NodeList, len(nlcb.builders))
	mutators := make([]Mutator, len(nlcb.builders))
	for i := range nlcb.builders {
		func(i int, root context.Context) {
			builder := nlcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NodeListMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, nlcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = nlcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, nlcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, nlcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (nlcb *NodeListCreateBulk) SaveX(ctx context.Context) []*NodeList {
	v, err := nlcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nlcb *NodeListCreateBulk) Exec(ctx context.Context) error {
	_, err := nlcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nlcb *NodeListCreateBulk) ExecX(ctx context.Context) {
	if err := nlcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.NodeList.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NodeListUpsert) {
//			SetProtoMessage(v+v).
//		}).
//		Exec(ctx)
func (nlcb *NodeListCreateBulk) OnConflict(opts ...sql.ConflictOption) *NodeListUpsertBulk {
	nlcb.conflict = opts
	return &NodeListUpsertBulk{
		create: nlcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.NodeList.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (nlcb *NodeListCreateBulk) OnConflictColumns(columns ...string) *NodeListUpsertBulk {
	nlcb.conflict = append(nlcb.conflict, sql.ConflictColumns(columns...))
	return &NodeListUpsertBulk{
		create: nlcb,
	}
}

// NodeListUpsertBulk is the builder for "upsert"-ing
// a bulk of NodeList nodes.
type NodeListUpsertBulk struct {
	create *NodeListCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.NodeList.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(nodelist.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *NodeListUpsertBulk) UpdateNewValues() *NodeListUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(nodelist.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.NodeList.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *NodeListUpsertBulk) Ignore() *NodeListUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NodeListUpsertBulk) DoNothing() *NodeListUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NodeListCreateBulk.OnConflict
// documentation for more info.
func (u *NodeListUpsertBulk) Update(set func(*NodeListUpsert)) *NodeListUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NodeListUpsert{UpdateSet: update})
	}))
	return u
}

// SetProtoMessage sets the "proto_message" field.
func (u *NodeListUpsertBulk) SetProtoMessage(v *sbom.NodeList) *NodeListUpsertBulk {
	return u.Update(func(s *NodeListUpsert) {
		s.SetProtoMessage(v)
	})
}

// UpdateProtoMessage sets the "proto_message" field to the value that was provided on create.
func (u *NodeListUpsertBulk) UpdateProtoMessage() *NodeListUpsertBulk {
	return u.Update(func(s *NodeListUpsert) {
		s.UpdateProtoMessage()
	})
}

// ClearProtoMessage clears the value of the "proto_message" field.
func (u *NodeListUpsertBulk) ClearProtoMessage() *NodeListUpsertBulk {
	return u.Update(func(s *NodeListUpsert) {
		s.ClearProtoMessage()
	})
}

// SetRootElements sets the "root_elements" field.
func (u *NodeListUpsertBulk) SetRootElements(v []string) *NodeListUpsertBulk {
	return u.Update(func(s *NodeListUpsert) {
		s.SetRootElements(v)
	})
}

// UpdateRootElements sets the "root_elements" field to the value that was provided on create.
func (u *NodeListUpsertBulk) UpdateRootElements() *NodeListUpsertBulk {
	return u.Update(func(s *NodeListUpsert) {
		s.UpdateRootElements()
	})
}

// Exec executes the query.
func (u *NodeListUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the NodeListCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NodeListCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NodeListUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
