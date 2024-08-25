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
	"github.com/protobom/storage/internal/backends/ent/metadata"
	"github.com/protobom/storage/internal/backends/ent/tool"
)

// ToolCreate is the builder for creating a Tool entity.
type ToolCreate struct {
	config
	mutation *ToolMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDocumentID sets the "document_id" field.
func (tc *ToolCreate) SetDocumentID(u uuid.UUID) *ToolCreate {
	tc.mutation.SetDocumentID(u)
	return tc
}

// SetNillableDocumentID sets the "document_id" field if the given value is not nil.
func (tc *ToolCreate) SetNillableDocumentID(u *uuid.UUID) *ToolCreate {
	if u != nil {
		tc.SetDocumentID(*u)
	}
	return tc
}

// SetProtoMessage sets the "proto_message" field.
func (tc *ToolCreate) SetProtoMessage(s *sbom.Tool) *ToolCreate {
	tc.mutation.SetProtoMessage(s)
	return tc
}

// SetMetadataID sets the "metadata_id" field.
func (tc *ToolCreate) SetMetadataID(s string) *ToolCreate {
	tc.mutation.SetMetadataID(s)
	return tc
}

// SetNillableMetadataID sets the "metadata_id" field if the given value is not nil.
func (tc *ToolCreate) SetNillableMetadataID(s *string) *ToolCreate {
	if s != nil {
		tc.SetMetadataID(*s)
	}
	return tc
}

// SetName sets the "name" field.
func (tc *ToolCreate) SetName(s string) *ToolCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetVersion sets the "version" field.
func (tc *ToolCreate) SetVersion(s string) *ToolCreate {
	tc.mutation.SetVersion(s)
	return tc
}

// SetVendor sets the "vendor" field.
func (tc *ToolCreate) SetVendor(s string) *ToolCreate {
	tc.mutation.SetVendor(s)
	return tc
}

// SetID sets the "id" field.
func (tc *ToolCreate) SetID(u uuid.UUID) *ToolCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetDocument sets the "document" edge to the Document entity.
func (tc *ToolCreate) SetDocument(d *Document) *ToolCreate {
	return tc.SetDocumentID(d.ID)
}

// SetMetadata sets the "metadata" edge to the Metadata entity.
func (tc *ToolCreate) SetMetadata(m *Metadata) *ToolCreate {
	return tc.SetMetadataID(m.ID)
}

// Mutation returns the ToolMutation object of the builder.
func (tc *ToolCreate) Mutation() *ToolMutation {
	return tc.mutation
}

// Save creates the Tool in the database.
func (tc *ToolCreate) Save(ctx context.Context) (*Tool, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *ToolCreate) SaveX(ctx context.Context) *Tool {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *ToolCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *ToolCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *ToolCreate) defaults() {
	if _, ok := tc.mutation.DocumentID(); !ok {
		v := tool.DefaultDocumentID()
		tc.mutation.SetDocumentID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *ToolCreate) check() error {
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Tool.name"`)}
	}
	if _, ok := tc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "Tool.version"`)}
	}
	if _, ok := tc.mutation.Vendor(); !ok {
		return &ValidationError{Name: "vendor", err: errors.New(`ent: missing required field "Tool.vendor"`)}
	}
	return nil
}

func (tc *ToolCreate) sqlSave(ctx context.Context) (*Tool, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
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
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *ToolCreate) createSpec() (*Tool, *sqlgraph.CreateSpec) {
	var (
		_node = &Tool{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(tool.Table, sqlgraph.NewFieldSpec(tool.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = tc.conflict
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.ProtoMessage(); ok {
		_spec.SetField(tool.FieldProtoMessage, field.TypeJSON, value)
		_node.ProtoMessage = value
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(tool.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tc.mutation.Version(); ok {
		_spec.SetField(tool.FieldVersion, field.TypeString, value)
		_node.Version = value
	}
	if value, ok := tc.mutation.Vendor(); ok {
		_spec.SetField(tool.FieldVendor, field.TypeString, value)
		_node.Vendor = value
	}
	if nodes := tc.mutation.DocumentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tool.DocumentTable,
			Columns: []string{tool.DocumentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(document.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DocumentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tool.MetadataTable,
			Columns: []string{tool.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.MetadataID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Tool.Create().
//		SetDocumentID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ToolUpsert) {
//			SetDocumentID(v+v).
//		}).
//		Exec(ctx)
func (tc *ToolCreate) OnConflict(opts ...sql.ConflictOption) *ToolUpsertOne {
	tc.conflict = opts
	return &ToolUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Tool.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *ToolCreate) OnConflictColumns(columns ...string) *ToolUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &ToolUpsertOne{
		create: tc,
	}
}

type (
	// ToolUpsertOne is the builder for "upsert"-ing
	//  one Tool node.
	ToolUpsertOne struct {
		create *ToolCreate
	}

	// ToolUpsert is the "OnConflict" setter.
	ToolUpsert struct {
		*sql.UpdateSet
	}
)

// SetProtoMessage sets the "proto_message" field.
func (u *ToolUpsert) SetProtoMessage(v *sbom.Tool) *ToolUpsert {
	u.Set(tool.FieldProtoMessage, v)
	return u
}

// UpdateProtoMessage sets the "proto_message" field to the value that was provided on create.
func (u *ToolUpsert) UpdateProtoMessage() *ToolUpsert {
	u.SetExcluded(tool.FieldProtoMessage)
	return u
}

// ClearProtoMessage clears the value of the "proto_message" field.
func (u *ToolUpsert) ClearProtoMessage() *ToolUpsert {
	u.SetNull(tool.FieldProtoMessage)
	return u
}

// SetMetadataID sets the "metadata_id" field.
func (u *ToolUpsert) SetMetadataID(v string) *ToolUpsert {
	u.Set(tool.FieldMetadataID, v)
	return u
}

// UpdateMetadataID sets the "metadata_id" field to the value that was provided on create.
func (u *ToolUpsert) UpdateMetadataID() *ToolUpsert {
	u.SetExcluded(tool.FieldMetadataID)
	return u
}

// ClearMetadataID clears the value of the "metadata_id" field.
func (u *ToolUpsert) ClearMetadataID() *ToolUpsert {
	u.SetNull(tool.FieldMetadataID)
	return u
}

// SetName sets the "name" field.
func (u *ToolUpsert) SetName(v string) *ToolUpsert {
	u.Set(tool.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ToolUpsert) UpdateName() *ToolUpsert {
	u.SetExcluded(tool.FieldName)
	return u
}

// SetVersion sets the "version" field.
func (u *ToolUpsert) SetVersion(v string) *ToolUpsert {
	u.Set(tool.FieldVersion, v)
	return u
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *ToolUpsert) UpdateVersion() *ToolUpsert {
	u.SetExcluded(tool.FieldVersion)
	return u
}

// SetVendor sets the "vendor" field.
func (u *ToolUpsert) SetVendor(v string) *ToolUpsert {
	u.Set(tool.FieldVendor, v)
	return u
}

// UpdateVendor sets the "vendor" field to the value that was provided on create.
func (u *ToolUpsert) UpdateVendor() *ToolUpsert {
	u.SetExcluded(tool.FieldVendor)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Tool.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(tool.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ToolUpsertOne) UpdateNewValues() *ToolUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(tool.FieldID)
		}
		if _, exists := u.create.mutation.DocumentID(); exists {
			s.SetIgnore(tool.FieldDocumentID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Tool.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ToolUpsertOne) Ignore() *ToolUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ToolUpsertOne) DoNothing() *ToolUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ToolCreate.OnConflict
// documentation for more info.
func (u *ToolUpsertOne) Update(set func(*ToolUpsert)) *ToolUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ToolUpsert{UpdateSet: update})
	}))
	return u
}

// SetProtoMessage sets the "proto_message" field.
func (u *ToolUpsertOne) SetProtoMessage(v *sbom.Tool) *ToolUpsertOne {
	return u.Update(func(s *ToolUpsert) {
		s.SetProtoMessage(v)
	})
}

// UpdateProtoMessage sets the "proto_message" field to the value that was provided on create.
func (u *ToolUpsertOne) UpdateProtoMessage() *ToolUpsertOne {
	return u.Update(func(s *ToolUpsert) {
		s.UpdateProtoMessage()
	})
}

// ClearProtoMessage clears the value of the "proto_message" field.
func (u *ToolUpsertOne) ClearProtoMessage() *ToolUpsertOne {
	return u.Update(func(s *ToolUpsert) {
		s.ClearProtoMessage()
	})
}

// SetMetadataID sets the "metadata_id" field.
func (u *ToolUpsertOne) SetMetadataID(v string) *ToolUpsertOne {
	return u.Update(func(s *ToolUpsert) {
		s.SetMetadataID(v)
	})
}

// UpdateMetadataID sets the "metadata_id" field to the value that was provided on create.
func (u *ToolUpsertOne) UpdateMetadataID() *ToolUpsertOne {
	return u.Update(func(s *ToolUpsert) {
		s.UpdateMetadataID()
	})
}

// ClearMetadataID clears the value of the "metadata_id" field.
func (u *ToolUpsertOne) ClearMetadataID() *ToolUpsertOne {
	return u.Update(func(s *ToolUpsert) {
		s.ClearMetadataID()
	})
}

// SetName sets the "name" field.
func (u *ToolUpsertOne) SetName(v string) *ToolUpsertOne {
	return u.Update(func(s *ToolUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ToolUpsertOne) UpdateName() *ToolUpsertOne {
	return u.Update(func(s *ToolUpsert) {
		s.UpdateName()
	})
}

// SetVersion sets the "version" field.
func (u *ToolUpsertOne) SetVersion(v string) *ToolUpsertOne {
	return u.Update(func(s *ToolUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *ToolUpsertOne) UpdateVersion() *ToolUpsertOne {
	return u.Update(func(s *ToolUpsert) {
		s.UpdateVersion()
	})
}

// SetVendor sets the "vendor" field.
func (u *ToolUpsertOne) SetVendor(v string) *ToolUpsertOne {
	return u.Update(func(s *ToolUpsert) {
		s.SetVendor(v)
	})
}

// UpdateVendor sets the "vendor" field to the value that was provided on create.
func (u *ToolUpsertOne) UpdateVendor() *ToolUpsertOne {
	return u.Update(func(s *ToolUpsert) {
		s.UpdateVendor()
	})
}

// Exec executes the query.
func (u *ToolUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ToolCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ToolUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ToolUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ToolUpsertOne.ID is not supported by MySQL driver. Use ToolUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ToolUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ToolCreateBulk is the builder for creating many Tool entities in bulk.
type ToolCreateBulk struct {
	config
	err      error
	builders []*ToolCreate
	conflict []sql.ConflictOption
}

// Save creates the Tool entities in the database.
func (tcb *ToolCreateBulk) Save(ctx context.Context) ([]*Tool, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tool, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ToolMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *ToolCreateBulk) SaveX(ctx context.Context) []*Tool {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *ToolCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *ToolCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Tool.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ToolUpsert) {
//			SetDocumentID(v+v).
//		}).
//		Exec(ctx)
func (tcb *ToolCreateBulk) OnConflict(opts ...sql.ConflictOption) *ToolUpsertBulk {
	tcb.conflict = opts
	return &ToolUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Tool.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *ToolCreateBulk) OnConflictColumns(columns ...string) *ToolUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &ToolUpsertBulk{
		create: tcb,
	}
}

// ToolUpsertBulk is the builder for "upsert"-ing
// a bulk of Tool nodes.
type ToolUpsertBulk struct {
	create *ToolCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Tool.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(tool.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ToolUpsertBulk) UpdateNewValues() *ToolUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(tool.FieldID)
			}
			if _, exists := b.mutation.DocumentID(); exists {
				s.SetIgnore(tool.FieldDocumentID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Tool.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ToolUpsertBulk) Ignore() *ToolUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ToolUpsertBulk) DoNothing() *ToolUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ToolCreateBulk.OnConflict
// documentation for more info.
func (u *ToolUpsertBulk) Update(set func(*ToolUpsert)) *ToolUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ToolUpsert{UpdateSet: update})
	}))
	return u
}

// SetProtoMessage sets the "proto_message" field.
func (u *ToolUpsertBulk) SetProtoMessage(v *sbom.Tool) *ToolUpsertBulk {
	return u.Update(func(s *ToolUpsert) {
		s.SetProtoMessage(v)
	})
}

// UpdateProtoMessage sets the "proto_message" field to the value that was provided on create.
func (u *ToolUpsertBulk) UpdateProtoMessage() *ToolUpsertBulk {
	return u.Update(func(s *ToolUpsert) {
		s.UpdateProtoMessage()
	})
}

// ClearProtoMessage clears the value of the "proto_message" field.
func (u *ToolUpsertBulk) ClearProtoMessage() *ToolUpsertBulk {
	return u.Update(func(s *ToolUpsert) {
		s.ClearProtoMessage()
	})
}

// SetMetadataID sets the "metadata_id" field.
func (u *ToolUpsertBulk) SetMetadataID(v string) *ToolUpsertBulk {
	return u.Update(func(s *ToolUpsert) {
		s.SetMetadataID(v)
	})
}

// UpdateMetadataID sets the "metadata_id" field to the value that was provided on create.
func (u *ToolUpsertBulk) UpdateMetadataID() *ToolUpsertBulk {
	return u.Update(func(s *ToolUpsert) {
		s.UpdateMetadataID()
	})
}

// ClearMetadataID clears the value of the "metadata_id" field.
func (u *ToolUpsertBulk) ClearMetadataID() *ToolUpsertBulk {
	return u.Update(func(s *ToolUpsert) {
		s.ClearMetadataID()
	})
}

// SetName sets the "name" field.
func (u *ToolUpsertBulk) SetName(v string) *ToolUpsertBulk {
	return u.Update(func(s *ToolUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ToolUpsertBulk) UpdateName() *ToolUpsertBulk {
	return u.Update(func(s *ToolUpsert) {
		s.UpdateName()
	})
}

// SetVersion sets the "version" field.
func (u *ToolUpsertBulk) SetVersion(v string) *ToolUpsertBulk {
	return u.Update(func(s *ToolUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *ToolUpsertBulk) UpdateVersion() *ToolUpsertBulk {
	return u.Update(func(s *ToolUpsert) {
		s.UpdateVersion()
	})
}

// SetVendor sets the "vendor" field.
func (u *ToolUpsertBulk) SetVendor(v string) *ToolUpsertBulk {
	return u.Update(func(s *ToolUpsert) {
		s.SetVendor(v)
	})
}

// UpdateVendor sets the "vendor" field to the value that was provided on create.
func (u *ToolUpsertBulk) UpdateVendor() *ToolUpsertBulk {
	return u.Update(func(s *ToolUpsert) {
		s.UpdateVendor()
	})
}

// Exec executes the query.
func (u *ToolUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ToolCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ToolCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ToolUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
