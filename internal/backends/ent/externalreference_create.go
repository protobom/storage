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
	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/externalreference"
	"github.com/protobom/storage/internal/backends/ent/hashesentry"
	"github.com/protobom/storage/internal/backends/ent/node"
)

// ExternalReferenceCreate is the builder for creating a ExternalReference entity.
type ExternalReferenceCreate struct {
	config
	mutation *ExternalReferenceMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetProtoMessage sets the "proto_message" field.
func (erc *ExternalReferenceCreate) SetProtoMessage(sr *sbom.ExternalReference) *ExternalReferenceCreate {
	erc.mutation.SetProtoMessage(sr)
	return erc
}

// SetNodeID sets the "node_id" field.
func (erc *ExternalReferenceCreate) SetNodeID(s string) *ExternalReferenceCreate {
	erc.mutation.SetNodeID(s)
	return erc
}

// SetNillableNodeID sets the "node_id" field if the given value is not nil.
func (erc *ExternalReferenceCreate) SetNillableNodeID(s *string) *ExternalReferenceCreate {
	if s != nil {
		erc.SetNodeID(*s)
	}
	return erc
}

// SetURL sets the "url" field.
func (erc *ExternalReferenceCreate) SetURL(s string) *ExternalReferenceCreate {
	erc.mutation.SetURL(s)
	return erc
}

// SetComment sets the "comment" field.
func (erc *ExternalReferenceCreate) SetComment(s string) *ExternalReferenceCreate {
	erc.mutation.SetComment(s)
	return erc
}

// SetAuthority sets the "authority" field.
func (erc *ExternalReferenceCreate) SetAuthority(s string) *ExternalReferenceCreate {
	erc.mutation.SetAuthority(s)
	return erc
}

// SetNillableAuthority sets the "authority" field if the given value is not nil.
func (erc *ExternalReferenceCreate) SetNillableAuthority(s *string) *ExternalReferenceCreate {
	if s != nil {
		erc.SetAuthority(*s)
	}
	return erc
}

// SetType sets the "type" field.
func (erc *ExternalReferenceCreate) SetType(e externalreference.Type) *ExternalReferenceCreate {
	erc.mutation.SetType(e)
	return erc
}

// SetDocumentID sets the "document" edge to the Document entity by ID.
func (erc *ExternalReferenceCreate) SetDocumentID(id string) *ExternalReferenceCreate {
	erc.mutation.SetDocumentID(id)
	return erc
}

// SetNillableDocumentID sets the "document" edge to the Document entity by ID if the given value is not nil.
func (erc *ExternalReferenceCreate) SetNillableDocumentID(id *string) *ExternalReferenceCreate {
	if id != nil {
		erc = erc.SetDocumentID(*id)
	}
	return erc
}

// SetDocument sets the "document" edge to the Document entity.
func (erc *ExternalReferenceCreate) SetDocument(d *Document) *ExternalReferenceCreate {
	return erc.SetDocumentID(d.ID)
}

// AddHashIDs adds the "hashes" edge to the HashesEntry entity by IDs.
func (erc *ExternalReferenceCreate) AddHashIDs(ids ...int) *ExternalReferenceCreate {
	erc.mutation.AddHashIDs(ids...)
	return erc
}

// AddHashes adds the "hashes" edges to the HashesEntry entity.
func (erc *ExternalReferenceCreate) AddHashes(h ...*HashesEntry) *ExternalReferenceCreate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return erc.AddHashIDs(ids...)
}

// SetNode sets the "node" edge to the Node entity.
func (erc *ExternalReferenceCreate) SetNode(n *Node) *ExternalReferenceCreate {
	return erc.SetNodeID(n.ID)
}

// Mutation returns the ExternalReferenceMutation object of the builder.
func (erc *ExternalReferenceCreate) Mutation() *ExternalReferenceMutation {
	return erc.mutation
}

// Save creates the ExternalReference in the database.
func (erc *ExternalReferenceCreate) Save(ctx context.Context) (*ExternalReference, error) {
	return withHooks(ctx, erc.sqlSave, erc.mutation, erc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (erc *ExternalReferenceCreate) SaveX(ctx context.Context) *ExternalReference {
	v, err := erc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (erc *ExternalReferenceCreate) Exec(ctx context.Context) error {
	_, err := erc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (erc *ExternalReferenceCreate) ExecX(ctx context.Context) {
	if err := erc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (erc *ExternalReferenceCreate) check() error {
	if _, ok := erc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "ExternalReference.url"`)}
	}
	if _, ok := erc.mutation.Comment(); !ok {
		return &ValidationError{Name: "comment", err: errors.New(`ent: missing required field "ExternalReference.comment"`)}
	}
	if _, ok := erc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "ExternalReference.type"`)}
	}
	if v, ok := erc.mutation.GetType(); ok {
		if err := externalreference.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "ExternalReference.type": %w`, err)}
		}
	}
	return nil
}

func (erc *ExternalReferenceCreate) sqlSave(ctx context.Context) (*ExternalReference, error) {
	if err := erc.check(); err != nil {
		return nil, err
	}
	_node, _spec := erc.createSpec()
	if err := sqlgraph.CreateNode(ctx, erc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	erc.mutation.id = &_node.ID
	erc.mutation.done = true
	return _node, nil
}

func (erc *ExternalReferenceCreate) createSpec() (*ExternalReference, *sqlgraph.CreateSpec) {
	var (
		_node = &ExternalReference{config: erc.config}
		_spec = sqlgraph.NewCreateSpec(externalreference.Table, sqlgraph.NewFieldSpec(externalreference.FieldID, field.TypeInt))
	)
	_spec.OnConflict = erc.conflict
	if value, ok := erc.mutation.ProtoMessage(); ok {
		_spec.SetField(externalreference.FieldProtoMessage, field.TypeJSON, value)
		_node.ProtoMessage = value
	}
	if value, ok := erc.mutation.URL(); ok {
		_spec.SetField(externalreference.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := erc.mutation.Comment(); ok {
		_spec.SetField(externalreference.FieldComment, field.TypeString, value)
		_node.Comment = value
	}
	if value, ok := erc.mutation.Authority(); ok {
		_spec.SetField(externalreference.FieldAuthority, field.TypeString, value)
		_node.Authority = value
	}
	if value, ok := erc.mutation.GetType(); ok {
		_spec.SetField(externalreference.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if nodes := erc.mutation.DocumentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   externalreference.DocumentTable,
			Columns: []string{externalreference.DocumentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(document.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.document_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := erc.mutation.HashesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   externalreference.HashesTable,
			Columns: []string{externalreference.HashesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hashesentry.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := erc.mutation.NodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   externalreference.NodeTable,
			Columns: []string{externalreference.NodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.NodeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ExternalReference.Create().
//		SetProtoMessage(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ExternalReferenceUpsert) {
//			SetProtoMessage(v+v).
//		}).
//		Exec(ctx)
func (erc *ExternalReferenceCreate) OnConflict(opts ...sql.ConflictOption) *ExternalReferenceUpsertOne {
	erc.conflict = opts
	return &ExternalReferenceUpsertOne{
		create: erc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ExternalReference.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (erc *ExternalReferenceCreate) OnConflictColumns(columns ...string) *ExternalReferenceUpsertOne {
	erc.conflict = append(erc.conflict, sql.ConflictColumns(columns...))
	return &ExternalReferenceUpsertOne{
		create: erc,
	}
}

type (
	// ExternalReferenceUpsertOne is the builder for "upsert"-ing
	//  one ExternalReference node.
	ExternalReferenceUpsertOne struct {
		create *ExternalReferenceCreate
	}

	// ExternalReferenceUpsert is the "OnConflict" setter.
	ExternalReferenceUpsert struct {
		*sql.UpdateSet
	}
)

// SetProtoMessage sets the "proto_message" field.
func (u *ExternalReferenceUpsert) SetProtoMessage(v *sbom.ExternalReference) *ExternalReferenceUpsert {
	u.Set(externalreference.FieldProtoMessage, v)
	return u
}

// UpdateProtoMessage sets the "proto_message" field to the value that was provided on create.
func (u *ExternalReferenceUpsert) UpdateProtoMessage() *ExternalReferenceUpsert {
	u.SetExcluded(externalreference.FieldProtoMessage)
	return u
}

// ClearProtoMessage clears the value of the "proto_message" field.
func (u *ExternalReferenceUpsert) ClearProtoMessage() *ExternalReferenceUpsert {
	u.SetNull(externalreference.FieldProtoMessage)
	return u
}

// SetNodeID sets the "node_id" field.
func (u *ExternalReferenceUpsert) SetNodeID(v string) *ExternalReferenceUpsert {
	u.Set(externalreference.FieldNodeID, v)
	return u
}

// UpdateNodeID sets the "node_id" field to the value that was provided on create.
func (u *ExternalReferenceUpsert) UpdateNodeID() *ExternalReferenceUpsert {
	u.SetExcluded(externalreference.FieldNodeID)
	return u
}

// ClearNodeID clears the value of the "node_id" field.
func (u *ExternalReferenceUpsert) ClearNodeID() *ExternalReferenceUpsert {
	u.SetNull(externalreference.FieldNodeID)
	return u
}

// SetURL sets the "url" field.
func (u *ExternalReferenceUpsert) SetURL(v string) *ExternalReferenceUpsert {
	u.Set(externalreference.FieldURL, v)
	return u
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *ExternalReferenceUpsert) UpdateURL() *ExternalReferenceUpsert {
	u.SetExcluded(externalreference.FieldURL)
	return u
}

// SetComment sets the "comment" field.
func (u *ExternalReferenceUpsert) SetComment(v string) *ExternalReferenceUpsert {
	u.Set(externalreference.FieldComment, v)
	return u
}

// UpdateComment sets the "comment" field to the value that was provided on create.
func (u *ExternalReferenceUpsert) UpdateComment() *ExternalReferenceUpsert {
	u.SetExcluded(externalreference.FieldComment)
	return u
}

// SetAuthority sets the "authority" field.
func (u *ExternalReferenceUpsert) SetAuthority(v string) *ExternalReferenceUpsert {
	u.Set(externalreference.FieldAuthority, v)
	return u
}

// UpdateAuthority sets the "authority" field to the value that was provided on create.
func (u *ExternalReferenceUpsert) UpdateAuthority() *ExternalReferenceUpsert {
	u.SetExcluded(externalreference.FieldAuthority)
	return u
}

// ClearAuthority clears the value of the "authority" field.
func (u *ExternalReferenceUpsert) ClearAuthority() *ExternalReferenceUpsert {
	u.SetNull(externalreference.FieldAuthority)
	return u
}

// SetType sets the "type" field.
func (u *ExternalReferenceUpsert) SetType(v externalreference.Type) *ExternalReferenceUpsert {
	u.Set(externalreference.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *ExternalReferenceUpsert) UpdateType() *ExternalReferenceUpsert {
	u.SetExcluded(externalreference.FieldType)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.ExternalReference.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ExternalReferenceUpsertOne) UpdateNewValues() *ExternalReferenceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ExternalReference.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ExternalReferenceUpsertOne) Ignore() *ExternalReferenceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ExternalReferenceUpsertOne) DoNothing() *ExternalReferenceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ExternalReferenceCreate.OnConflict
// documentation for more info.
func (u *ExternalReferenceUpsertOne) Update(set func(*ExternalReferenceUpsert)) *ExternalReferenceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ExternalReferenceUpsert{UpdateSet: update})
	}))
	return u
}

// SetProtoMessage sets the "proto_message" field.
func (u *ExternalReferenceUpsertOne) SetProtoMessage(v *sbom.ExternalReference) *ExternalReferenceUpsertOne {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.SetProtoMessage(v)
	})
}

// UpdateProtoMessage sets the "proto_message" field to the value that was provided on create.
func (u *ExternalReferenceUpsertOne) UpdateProtoMessage() *ExternalReferenceUpsertOne {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.UpdateProtoMessage()
	})
}

// ClearProtoMessage clears the value of the "proto_message" field.
func (u *ExternalReferenceUpsertOne) ClearProtoMessage() *ExternalReferenceUpsertOne {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.ClearProtoMessage()
	})
}

// SetNodeID sets the "node_id" field.
func (u *ExternalReferenceUpsertOne) SetNodeID(v string) *ExternalReferenceUpsertOne {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.SetNodeID(v)
	})
}

// UpdateNodeID sets the "node_id" field to the value that was provided on create.
func (u *ExternalReferenceUpsertOne) UpdateNodeID() *ExternalReferenceUpsertOne {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.UpdateNodeID()
	})
}

// ClearNodeID clears the value of the "node_id" field.
func (u *ExternalReferenceUpsertOne) ClearNodeID() *ExternalReferenceUpsertOne {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.ClearNodeID()
	})
}

// SetURL sets the "url" field.
func (u *ExternalReferenceUpsertOne) SetURL(v string) *ExternalReferenceUpsertOne {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *ExternalReferenceUpsertOne) UpdateURL() *ExternalReferenceUpsertOne {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.UpdateURL()
	})
}

// SetComment sets the "comment" field.
func (u *ExternalReferenceUpsertOne) SetComment(v string) *ExternalReferenceUpsertOne {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.SetComment(v)
	})
}

// UpdateComment sets the "comment" field to the value that was provided on create.
func (u *ExternalReferenceUpsertOne) UpdateComment() *ExternalReferenceUpsertOne {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.UpdateComment()
	})
}

// SetAuthority sets the "authority" field.
func (u *ExternalReferenceUpsertOne) SetAuthority(v string) *ExternalReferenceUpsertOne {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.SetAuthority(v)
	})
}

// UpdateAuthority sets the "authority" field to the value that was provided on create.
func (u *ExternalReferenceUpsertOne) UpdateAuthority() *ExternalReferenceUpsertOne {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.UpdateAuthority()
	})
}

// ClearAuthority clears the value of the "authority" field.
func (u *ExternalReferenceUpsertOne) ClearAuthority() *ExternalReferenceUpsertOne {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.ClearAuthority()
	})
}

// SetType sets the "type" field.
func (u *ExternalReferenceUpsertOne) SetType(v externalreference.Type) *ExternalReferenceUpsertOne {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *ExternalReferenceUpsertOne) UpdateType() *ExternalReferenceUpsertOne {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.UpdateType()
	})
}

// Exec executes the query.
func (u *ExternalReferenceUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ExternalReferenceCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ExternalReferenceUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ExternalReferenceUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ExternalReferenceUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ExternalReferenceCreateBulk is the builder for creating many ExternalReference entities in bulk.
type ExternalReferenceCreateBulk struct {
	config
	err      error
	builders []*ExternalReferenceCreate
	conflict []sql.ConflictOption
}

// Save creates the ExternalReference entities in the database.
func (ercb *ExternalReferenceCreateBulk) Save(ctx context.Context) ([]*ExternalReference, error) {
	if ercb.err != nil {
		return nil, ercb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ercb.builders))
	nodes := make([]*ExternalReference, len(ercb.builders))
	mutators := make([]Mutator, len(ercb.builders))
	for i := range ercb.builders {
		func(i int, root context.Context) {
			builder := ercb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ExternalReferenceMutation)
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
					_, err = mutators[i+1].Mutate(root, ercb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ercb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ercb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, ercb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ercb *ExternalReferenceCreateBulk) SaveX(ctx context.Context) []*ExternalReference {
	v, err := ercb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ercb *ExternalReferenceCreateBulk) Exec(ctx context.Context) error {
	_, err := ercb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ercb *ExternalReferenceCreateBulk) ExecX(ctx context.Context) {
	if err := ercb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ExternalReference.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ExternalReferenceUpsert) {
//			SetProtoMessage(v+v).
//		}).
//		Exec(ctx)
func (ercb *ExternalReferenceCreateBulk) OnConflict(opts ...sql.ConflictOption) *ExternalReferenceUpsertBulk {
	ercb.conflict = opts
	return &ExternalReferenceUpsertBulk{
		create: ercb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ExternalReference.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ercb *ExternalReferenceCreateBulk) OnConflictColumns(columns ...string) *ExternalReferenceUpsertBulk {
	ercb.conflict = append(ercb.conflict, sql.ConflictColumns(columns...))
	return &ExternalReferenceUpsertBulk{
		create: ercb,
	}
}

// ExternalReferenceUpsertBulk is the builder for "upsert"-ing
// a bulk of ExternalReference nodes.
type ExternalReferenceUpsertBulk struct {
	create *ExternalReferenceCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ExternalReference.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ExternalReferenceUpsertBulk) UpdateNewValues() *ExternalReferenceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ExternalReference.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ExternalReferenceUpsertBulk) Ignore() *ExternalReferenceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ExternalReferenceUpsertBulk) DoNothing() *ExternalReferenceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ExternalReferenceCreateBulk.OnConflict
// documentation for more info.
func (u *ExternalReferenceUpsertBulk) Update(set func(*ExternalReferenceUpsert)) *ExternalReferenceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ExternalReferenceUpsert{UpdateSet: update})
	}))
	return u
}

// SetProtoMessage sets the "proto_message" field.
func (u *ExternalReferenceUpsertBulk) SetProtoMessage(v *sbom.ExternalReference) *ExternalReferenceUpsertBulk {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.SetProtoMessage(v)
	})
}

// UpdateProtoMessage sets the "proto_message" field to the value that was provided on create.
func (u *ExternalReferenceUpsertBulk) UpdateProtoMessage() *ExternalReferenceUpsertBulk {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.UpdateProtoMessage()
	})
}

// ClearProtoMessage clears the value of the "proto_message" field.
func (u *ExternalReferenceUpsertBulk) ClearProtoMessage() *ExternalReferenceUpsertBulk {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.ClearProtoMessage()
	})
}

// SetNodeID sets the "node_id" field.
func (u *ExternalReferenceUpsertBulk) SetNodeID(v string) *ExternalReferenceUpsertBulk {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.SetNodeID(v)
	})
}

// UpdateNodeID sets the "node_id" field to the value that was provided on create.
func (u *ExternalReferenceUpsertBulk) UpdateNodeID() *ExternalReferenceUpsertBulk {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.UpdateNodeID()
	})
}

// ClearNodeID clears the value of the "node_id" field.
func (u *ExternalReferenceUpsertBulk) ClearNodeID() *ExternalReferenceUpsertBulk {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.ClearNodeID()
	})
}

// SetURL sets the "url" field.
func (u *ExternalReferenceUpsertBulk) SetURL(v string) *ExternalReferenceUpsertBulk {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *ExternalReferenceUpsertBulk) UpdateURL() *ExternalReferenceUpsertBulk {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.UpdateURL()
	})
}

// SetComment sets the "comment" field.
func (u *ExternalReferenceUpsertBulk) SetComment(v string) *ExternalReferenceUpsertBulk {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.SetComment(v)
	})
}

// UpdateComment sets the "comment" field to the value that was provided on create.
func (u *ExternalReferenceUpsertBulk) UpdateComment() *ExternalReferenceUpsertBulk {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.UpdateComment()
	})
}

// SetAuthority sets the "authority" field.
func (u *ExternalReferenceUpsertBulk) SetAuthority(v string) *ExternalReferenceUpsertBulk {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.SetAuthority(v)
	})
}

// UpdateAuthority sets the "authority" field to the value that was provided on create.
func (u *ExternalReferenceUpsertBulk) UpdateAuthority() *ExternalReferenceUpsertBulk {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.UpdateAuthority()
	})
}

// ClearAuthority clears the value of the "authority" field.
func (u *ExternalReferenceUpsertBulk) ClearAuthority() *ExternalReferenceUpsertBulk {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.ClearAuthority()
	})
}

// SetType sets the "type" field.
func (u *ExternalReferenceUpsertBulk) SetType(v externalreference.Type) *ExternalReferenceUpsertBulk {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *ExternalReferenceUpsertBulk) UpdateType() *ExternalReferenceUpsertBulk {
	return u.Update(func(s *ExternalReferenceUpsert) {
		s.UpdateType()
	})
}

// Exec executes the query.
func (u *ExternalReferenceUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ExternalReferenceCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ExternalReferenceCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ExternalReferenceUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
