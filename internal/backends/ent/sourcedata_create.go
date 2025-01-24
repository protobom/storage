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
	"github.com/protobom/storage/internal/backends/ent/hashesentry"
	"github.com/protobom/storage/internal/backends/ent/metadata"
	"github.com/protobom/storage/internal/backends/ent/sourcedata"
)

// SourceDataCreate is the builder for creating a SourceData entity.
type SourceDataCreate struct {
	config
	mutation *SourceDataMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetProtoMessage sets the "proto_message" field.
func (sdc *SourceDataCreate) SetProtoMessage(sd *sbom.SourceData) *SourceDataCreate {
	sdc.mutation.SetProtoMessage(sd)
	return sdc
}

// SetFormat sets the "format" field.
func (sdc *SourceDataCreate) SetFormat(s string) *SourceDataCreate {
	sdc.mutation.SetFormat(s)
	return sdc
}

// SetSize sets the "size" field.
func (sdc *SourceDataCreate) SetSize(i int64) *SourceDataCreate {
	sdc.mutation.SetSize(i)
	return sdc
}

// SetURI sets the "uri" field.
func (sdc *SourceDataCreate) SetURI(s string) *SourceDataCreate {
	sdc.mutation.SetURI(s)
	return sdc
}

// SetNillableURI sets the "uri" field if the given value is not nil.
func (sdc *SourceDataCreate) SetNillableURI(s *string) *SourceDataCreate {
	if s != nil {
		sdc.SetURI(*s)
	}
	return sdc
}

// SetID sets the "id" field.
func (sdc *SourceDataCreate) SetID(u uuid.UUID) *SourceDataCreate {
	sdc.mutation.SetID(u)
	return sdc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sdc *SourceDataCreate) SetNillableID(u *uuid.UUID) *SourceDataCreate {
	if u != nil {
		sdc.SetID(*u)
	}
	return sdc
}

// AddHashIDs adds the "hashes" edge to the HashesEntry entity by IDs.
func (sdc *SourceDataCreate) AddHashIDs(ids ...uuid.UUID) *SourceDataCreate {
	sdc.mutation.AddHashIDs(ids...)
	return sdc
}

// AddHashes adds the "hashes" edges to the HashesEntry entity.
func (sdc *SourceDataCreate) AddHashes(h ...*HashesEntry) *SourceDataCreate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return sdc.AddHashIDs(ids...)
}

// AddDocumentIDs adds the "documents" edge to the Document entity by IDs.
func (sdc *SourceDataCreate) AddDocumentIDs(ids ...uuid.UUID) *SourceDataCreate {
	sdc.mutation.AddDocumentIDs(ids...)
	return sdc
}

// AddDocuments adds the "documents" edges to the Document entity.
func (sdc *SourceDataCreate) AddDocuments(d ...*Document) *SourceDataCreate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return sdc.AddDocumentIDs(ids...)
}

// AddMetadatumIDs adds the "metadata" edge to the Metadata entity by IDs.
func (sdc *SourceDataCreate) AddMetadatumIDs(ids ...uuid.UUID) *SourceDataCreate {
	sdc.mutation.AddMetadatumIDs(ids...)
	return sdc
}

// AddMetadata adds the "metadata" edges to the Metadata entity.
func (sdc *SourceDataCreate) AddMetadata(m ...*Metadata) *SourceDataCreate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return sdc.AddMetadatumIDs(ids...)
}

// Mutation returns the SourceDataMutation object of the builder.
func (sdc *SourceDataCreate) Mutation() *SourceDataMutation {
	return sdc.mutation
}

// Save creates the SourceData in the database.
func (sdc *SourceDataCreate) Save(ctx context.Context) (*SourceData, error) {
	if err := sdc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, sdc.sqlSave, sdc.mutation, sdc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sdc *SourceDataCreate) SaveX(ctx context.Context) *SourceData {
	v, err := sdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdc *SourceDataCreate) Exec(ctx context.Context) error {
	_, err := sdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdc *SourceDataCreate) ExecX(ctx context.Context) {
	if err := sdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sdc *SourceDataCreate) defaults() error {
	if _, ok := sdc.mutation.ID(); !ok {
		if sourcedata.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized sourcedata.DefaultID (forgotten import ent/runtime?)")
		}
		v := sourcedata.DefaultID()
		sdc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sdc *SourceDataCreate) check() error {
	if _, ok := sdc.mutation.ProtoMessage(); !ok {
		return &ValidationError{Name: "proto_message", err: errors.New(`ent: missing required field "SourceData.proto_message"`)}
	}
	if _, ok := sdc.mutation.Format(); !ok {
		return &ValidationError{Name: "format", err: errors.New(`ent: missing required field "SourceData.format"`)}
	}
	if _, ok := sdc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "SourceData.size"`)}
	}
	if len(sdc.mutation.DocumentsIDs()) == 0 {
		return &ValidationError{Name: "documents", err: errors.New(`ent: missing required edge "SourceData.documents"`)}
	}
	if len(sdc.mutation.MetadataIDs()) == 0 {
		return &ValidationError{Name: "metadata", err: errors.New(`ent: missing required edge "SourceData.metadata"`)}
	}
	return nil
}

func (sdc *SourceDataCreate) sqlSave(ctx context.Context) (*SourceData, error) {
	if err := sdc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sdc.driver, _spec); err != nil {
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
	sdc.mutation.id = &_node.ID
	sdc.mutation.done = true
	return _node, nil
}

func (sdc *SourceDataCreate) createSpec() (*SourceData, *sqlgraph.CreateSpec) {
	var (
		_node = &SourceData{config: sdc.config}
		_spec = sqlgraph.NewCreateSpec(sourcedata.Table, sqlgraph.NewFieldSpec(sourcedata.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = sdc.conflict
	if id, ok := sdc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sdc.mutation.ProtoMessage(); ok {
		_spec.SetField(sourcedata.FieldProtoMessage, field.TypeBytes, value)
		_node.ProtoMessage = value
	}
	if value, ok := sdc.mutation.Format(); ok {
		_spec.SetField(sourcedata.FieldFormat, field.TypeString, value)
		_node.Format = value
	}
	if value, ok := sdc.mutation.Size(); ok {
		_spec.SetField(sourcedata.FieldSize, field.TypeInt64, value)
		_node.Size = value
	}
	if value, ok := sdc.mutation.URI(); ok {
		_spec.SetField(sourcedata.FieldURI, field.TypeString, value)
		_node.URI = &value
	}
	if nodes := sdc.mutation.HashesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sdc.mutation.DocumentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   sourcedata.DocumentsTable,
			Columns: sourcedata.DocumentsPrimaryKey,
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
	if nodes := sdc.mutation.MetadataIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SourceData.Create().
//		SetProtoMessage(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SourceDataUpsert) {
//			SetProtoMessage(v+v).
//		}).
//		Exec(ctx)
func (sdc *SourceDataCreate) OnConflict(opts ...sql.ConflictOption) *SourceDataUpsertOne {
	sdc.conflict = opts
	return &SourceDataUpsertOne{
		create: sdc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SourceData.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sdc *SourceDataCreate) OnConflictColumns(columns ...string) *SourceDataUpsertOne {
	sdc.conflict = append(sdc.conflict, sql.ConflictColumns(columns...))
	return &SourceDataUpsertOne{
		create: sdc,
	}
}

type (
	// SourceDataUpsertOne is the builder for "upsert"-ing
	//  one SourceData node.
	SourceDataUpsertOne struct {
		create *SourceDataCreate
	}

	// SourceDataUpsert is the "OnConflict" setter.
	SourceDataUpsert struct {
		*sql.UpdateSet
	}
)

// SetFormat sets the "format" field.
func (u *SourceDataUpsert) SetFormat(v string) *SourceDataUpsert {
	u.Set(sourcedata.FieldFormat, v)
	return u
}

// UpdateFormat sets the "format" field to the value that was provided on create.
func (u *SourceDataUpsert) UpdateFormat() *SourceDataUpsert {
	u.SetExcluded(sourcedata.FieldFormat)
	return u
}

// SetSize sets the "size" field.
func (u *SourceDataUpsert) SetSize(v int64) *SourceDataUpsert {
	u.Set(sourcedata.FieldSize, v)
	return u
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *SourceDataUpsert) UpdateSize() *SourceDataUpsert {
	u.SetExcluded(sourcedata.FieldSize)
	return u
}

// AddSize adds v to the "size" field.
func (u *SourceDataUpsert) AddSize(v int64) *SourceDataUpsert {
	u.Add(sourcedata.FieldSize, v)
	return u
}

// SetURI sets the "uri" field.
func (u *SourceDataUpsert) SetURI(v string) *SourceDataUpsert {
	u.Set(sourcedata.FieldURI, v)
	return u
}

// UpdateURI sets the "uri" field to the value that was provided on create.
func (u *SourceDataUpsert) UpdateURI() *SourceDataUpsert {
	u.SetExcluded(sourcedata.FieldURI)
	return u
}

// ClearURI clears the value of the "uri" field.
func (u *SourceDataUpsert) ClearURI() *SourceDataUpsert {
	u.SetNull(sourcedata.FieldURI)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.SourceData.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sourcedata.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SourceDataUpsertOne) UpdateNewValues() *SourceDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(sourcedata.FieldID)
		}
		if _, exists := u.create.mutation.ProtoMessage(); exists {
			s.SetIgnore(sourcedata.FieldProtoMessage)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SourceData.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SourceDataUpsertOne) Ignore() *SourceDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SourceDataUpsertOne) DoNothing() *SourceDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SourceDataCreate.OnConflict
// documentation for more info.
func (u *SourceDataUpsertOne) Update(set func(*SourceDataUpsert)) *SourceDataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SourceDataUpsert{UpdateSet: update})
	}))
	return u
}

// SetFormat sets the "format" field.
func (u *SourceDataUpsertOne) SetFormat(v string) *SourceDataUpsertOne {
	return u.Update(func(s *SourceDataUpsert) {
		s.SetFormat(v)
	})
}

// UpdateFormat sets the "format" field to the value that was provided on create.
func (u *SourceDataUpsertOne) UpdateFormat() *SourceDataUpsertOne {
	return u.Update(func(s *SourceDataUpsert) {
		s.UpdateFormat()
	})
}

// SetSize sets the "size" field.
func (u *SourceDataUpsertOne) SetSize(v int64) *SourceDataUpsertOne {
	return u.Update(func(s *SourceDataUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *SourceDataUpsertOne) AddSize(v int64) *SourceDataUpsertOne {
	return u.Update(func(s *SourceDataUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *SourceDataUpsertOne) UpdateSize() *SourceDataUpsertOne {
	return u.Update(func(s *SourceDataUpsert) {
		s.UpdateSize()
	})
}

// SetURI sets the "uri" field.
func (u *SourceDataUpsertOne) SetURI(v string) *SourceDataUpsertOne {
	return u.Update(func(s *SourceDataUpsert) {
		s.SetURI(v)
	})
}

// UpdateURI sets the "uri" field to the value that was provided on create.
func (u *SourceDataUpsertOne) UpdateURI() *SourceDataUpsertOne {
	return u.Update(func(s *SourceDataUpsert) {
		s.UpdateURI()
	})
}

// ClearURI clears the value of the "uri" field.
func (u *SourceDataUpsertOne) ClearURI() *SourceDataUpsertOne {
	return u.Update(func(s *SourceDataUpsert) {
		s.ClearURI()
	})
}

// Exec executes the query.
func (u *SourceDataUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SourceDataCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SourceDataUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SourceDataUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: SourceDataUpsertOne.ID is not supported by MySQL driver. Use SourceDataUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SourceDataUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SourceDataCreateBulk is the builder for creating many SourceData entities in bulk.
type SourceDataCreateBulk struct {
	config
	err      error
	builders []*SourceDataCreate
	conflict []sql.ConflictOption
}

// Save creates the SourceData entities in the database.
func (sdcb *SourceDataCreateBulk) Save(ctx context.Context) ([]*SourceData, error) {
	if sdcb.err != nil {
		return nil, sdcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(sdcb.builders))
	nodes := make([]*SourceData, len(sdcb.builders))
	mutators := make([]Mutator, len(sdcb.builders))
	for i := range sdcb.builders {
		func(i int, root context.Context) {
			builder := sdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SourceDataMutation)
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
					_, err = mutators[i+1].Mutate(root, sdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = sdcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sdcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sdcb *SourceDataCreateBulk) SaveX(ctx context.Context) []*SourceData {
	v, err := sdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdcb *SourceDataCreateBulk) Exec(ctx context.Context) error {
	_, err := sdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdcb *SourceDataCreateBulk) ExecX(ctx context.Context) {
	if err := sdcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SourceData.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SourceDataUpsert) {
//			SetProtoMessage(v+v).
//		}).
//		Exec(ctx)
func (sdcb *SourceDataCreateBulk) OnConflict(opts ...sql.ConflictOption) *SourceDataUpsertBulk {
	sdcb.conflict = opts
	return &SourceDataUpsertBulk{
		create: sdcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SourceData.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sdcb *SourceDataCreateBulk) OnConflictColumns(columns ...string) *SourceDataUpsertBulk {
	sdcb.conflict = append(sdcb.conflict, sql.ConflictColumns(columns...))
	return &SourceDataUpsertBulk{
		create: sdcb,
	}
}

// SourceDataUpsertBulk is the builder for "upsert"-ing
// a bulk of SourceData nodes.
type SourceDataUpsertBulk struct {
	create *SourceDataCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SourceData.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(sourcedata.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SourceDataUpsertBulk) UpdateNewValues() *SourceDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(sourcedata.FieldID)
			}
			if _, exists := b.mutation.ProtoMessage(); exists {
				s.SetIgnore(sourcedata.FieldProtoMessage)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SourceData.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SourceDataUpsertBulk) Ignore() *SourceDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SourceDataUpsertBulk) DoNothing() *SourceDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SourceDataCreateBulk.OnConflict
// documentation for more info.
func (u *SourceDataUpsertBulk) Update(set func(*SourceDataUpsert)) *SourceDataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SourceDataUpsert{UpdateSet: update})
	}))
	return u
}

// SetFormat sets the "format" field.
func (u *SourceDataUpsertBulk) SetFormat(v string) *SourceDataUpsertBulk {
	return u.Update(func(s *SourceDataUpsert) {
		s.SetFormat(v)
	})
}

// UpdateFormat sets the "format" field to the value that was provided on create.
func (u *SourceDataUpsertBulk) UpdateFormat() *SourceDataUpsertBulk {
	return u.Update(func(s *SourceDataUpsert) {
		s.UpdateFormat()
	})
}

// SetSize sets the "size" field.
func (u *SourceDataUpsertBulk) SetSize(v int64) *SourceDataUpsertBulk {
	return u.Update(func(s *SourceDataUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *SourceDataUpsertBulk) AddSize(v int64) *SourceDataUpsertBulk {
	return u.Update(func(s *SourceDataUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *SourceDataUpsertBulk) UpdateSize() *SourceDataUpsertBulk {
	return u.Update(func(s *SourceDataUpsert) {
		s.UpdateSize()
	})
}

// SetURI sets the "uri" field.
func (u *SourceDataUpsertBulk) SetURI(v string) *SourceDataUpsertBulk {
	return u.Update(func(s *SourceDataUpsert) {
		s.SetURI(v)
	})
}

// UpdateURI sets the "uri" field to the value that was provided on create.
func (u *SourceDataUpsertBulk) UpdateURI() *SourceDataUpsertBulk {
	return u.Update(func(s *SourceDataUpsert) {
		s.UpdateURI()
	})
}

// ClearURI clears the value of the "uri" field.
func (u *SourceDataUpsertBulk) ClearURI() *SourceDataUpsertBulk {
	return u.Update(func(s *SourceDataUpsert) {
		s.ClearURI()
	})
}

// Exec executes the query.
func (u *SourceDataUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SourceDataCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SourceDataCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SourceDataUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
