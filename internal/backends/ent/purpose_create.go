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
	"github.com/protobom/storage/internal/backends/ent/document"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/purpose"
)

// PurposeCreate is the builder for creating a Purpose entity.
type PurposeCreate struct {
	config
	mutation *PurposeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetPrimaryPurpose sets the "primary_purpose" field.
func (pc *PurposeCreate) SetPrimaryPurpose(pp purpose.PrimaryPurpose) *PurposeCreate {
	pc.mutation.SetPrimaryPurpose(pp)
	return pc
}

// AddDocumentIDs adds the "documents" edge to the Document entity by IDs.
func (pc *PurposeCreate) AddDocumentIDs(ids ...uuid.UUID) *PurposeCreate {
	pc.mutation.AddDocumentIDs(ids...)
	return pc
}

// AddDocuments adds the "documents" edges to the Document entity.
func (pc *PurposeCreate) AddDocuments(d ...*Document) *PurposeCreate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pc.AddDocumentIDs(ids...)
}

// AddNodeIDs adds the "nodes" edge to the Node entity by IDs.
func (pc *PurposeCreate) AddNodeIDs(ids ...uuid.UUID) *PurposeCreate {
	pc.mutation.AddNodeIDs(ids...)
	return pc
}

// AddNodes adds the "nodes" edges to the Node entity.
func (pc *PurposeCreate) AddNodes(n ...*Node) *PurposeCreate {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return pc.AddNodeIDs(ids...)
}

// Mutation returns the PurposeMutation object of the builder.
func (pc *PurposeCreate) Mutation() *PurposeMutation {
	return pc.mutation
}

// Save creates the Purpose in the database.
func (pc *PurposeCreate) Save(ctx context.Context) (*Purpose, error) {
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PurposeCreate) SaveX(ctx context.Context) *Purpose {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PurposeCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PurposeCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PurposeCreate) check() error {
	if _, ok := pc.mutation.PrimaryPurpose(); !ok {
		return &ValidationError{Name: "primary_purpose", err: errors.New(`ent: missing required field "Purpose.primary_purpose"`)}
	}
	if v, ok := pc.mutation.PrimaryPurpose(); ok {
		if err := purpose.PrimaryPurposeValidator(v); err != nil {
			return &ValidationError{Name: "primary_purpose", err: fmt.Errorf(`ent: validator failed for field "Purpose.primary_purpose": %w`, err)}
		}
	}
	if len(pc.mutation.DocumentsIDs()) == 0 {
		return &ValidationError{Name: "documents", err: errors.New(`ent: missing required edge "Purpose.documents"`)}
	}
	if len(pc.mutation.NodesIDs()) == 0 {
		return &ValidationError{Name: "nodes", err: errors.New(`ent: missing required edge "Purpose.nodes"`)}
	}
	return nil
}

func (pc *PurposeCreate) sqlSave(ctx context.Context) (*Purpose, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PurposeCreate) createSpec() (*Purpose, *sqlgraph.CreateSpec) {
	var (
		_node = &Purpose{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(purpose.Table, sqlgraph.NewFieldSpec(purpose.FieldID, field.TypeInt))
	)
	_spec.OnConflict = pc.conflict
	if value, ok := pc.mutation.PrimaryPurpose(); ok {
		_spec.SetField(purpose.FieldPrimaryPurpose, field.TypeEnum, value)
		_node.PrimaryPurpose = value
	}
	if nodes := pc.mutation.DocumentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   purpose.DocumentsTable,
			Columns: purpose.DocumentsPrimaryKey,
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
	if nodes := pc.mutation.NodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   purpose.NodesTable,
			Columns: purpose.NodesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeUUID),
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
//	client.Purpose.Create().
//		SetPrimaryPurpose(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PurposeUpsert) {
//			SetPrimaryPurpose(v+v).
//		}).
//		Exec(ctx)
func (pc *PurposeCreate) OnConflict(opts ...sql.ConflictOption) *PurposeUpsertOne {
	pc.conflict = opts
	return &PurposeUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Purpose.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *PurposeCreate) OnConflictColumns(columns ...string) *PurposeUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PurposeUpsertOne{
		create: pc,
	}
}

type (
	// PurposeUpsertOne is the builder for "upsert"-ing
	//  one Purpose node.
	PurposeUpsertOne struct {
		create *PurposeCreate
	}

	// PurposeUpsert is the "OnConflict" setter.
	PurposeUpsert struct {
		*sql.UpdateSet
	}
)

// SetPrimaryPurpose sets the "primary_purpose" field.
func (u *PurposeUpsert) SetPrimaryPurpose(v purpose.PrimaryPurpose) *PurposeUpsert {
	u.Set(purpose.FieldPrimaryPurpose, v)
	return u
}

// UpdatePrimaryPurpose sets the "primary_purpose" field to the value that was provided on create.
func (u *PurposeUpsert) UpdatePrimaryPurpose() *PurposeUpsert {
	u.SetExcluded(purpose.FieldPrimaryPurpose)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Purpose.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PurposeUpsertOne) UpdateNewValues() *PurposeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Purpose.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PurposeUpsertOne) Ignore() *PurposeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PurposeUpsertOne) DoNothing() *PurposeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PurposeCreate.OnConflict
// documentation for more info.
func (u *PurposeUpsertOne) Update(set func(*PurposeUpsert)) *PurposeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PurposeUpsert{UpdateSet: update})
	}))
	return u
}

// SetPrimaryPurpose sets the "primary_purpose" field.
func (u *PurposeUpsertOne) SetPrimaryPurpose(v purpose.PrimaryPurpose) *PurposeUpsertOne {
	return u.Update(func(s *PurposeUpsert) {
		s.SetPrimaryPurpose(v)
	})
}

// UpdatePrimaryPurpose sets the "primary_purpose" field to the value that was provided on create.
func (u *PurposeUpsertOne) UpdatePrimaryPurpose() *PurposeUpsertOne {
	return u.Update(func(s *PurposeUpsert) {
		s.UpdatePrimaryPurpose()
	})
}

// Exec executes the query.
func (u *PurposeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PurposeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PurposeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PurposeUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PurposeUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PurposeCreateBulk is the builder for creating many Purpose entities in bulk.
type PurposeCreateBulk struct {
	config
	err      error
	builders []*PurposeCreate
	conflict []sql.ConflictOption
}

// Save creates the Purpose entities in the database.
func (pcb *PurposeCreateBulk) Save(ctx context.Context) ([]*Purpose, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Purpose, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PurposeMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PurposeCreateBulk) SaveX(ctx context.Context) []*Purpose {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PurposeCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PurposeCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Purpose.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PurposeUpsert) {
//			SetPrimaryPurpose(v+v).
//		}).
//		Exec(ctx)
func (pcb *PurposeCreateBulk) OnConflict(opts ...sql.ConflictOption) *PurposeUpsertBulk {
	pcb.conflict = opts
	return &PurposeUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Purpose.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *PurposeCreateBulk) OnConflictColumns(columns ...string) *PurposeUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PurposeUpsertBulk{
		create: pcb,
	}
}

// PurposeUpsertBulk is the builder for "upsert"-ing
// a bulk of Purpose nodes.
type PurposeUpsertBulk struct {
	create *PurposeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Purpose.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PurposeUpsertBulk) UpdateNewValues() *PurposeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Purpose.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PurposeUpsertBulk) Ignore() *PurposeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PurposeUpsertBulk) DoNothing() *PurposeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PurposeCreateBulk.OnConflict
// documentation for more info.
func (u *PurposeUpsertBulk) Update(set func(*PurposeUpsert)) *PurposeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PurposeUpsert{UpdateSet: update})
	}))
	return u
}

// SetPrimaryPurpose sets the "primary_purpose" field.
func (u *PurposeUpsertBulk) SetPrimaryPurpose(v purpose.PrimaryPurpose) *PurposeUpsertBulk {
	return u.Update(func(s *PurposeUpsert) {
		s.SetPrimaryPurpose(v)
	})
}

// UpdatePrimaryPurpose sets the "primary_purpose" field to the value that was provided on create.
func (u *PurposeUpsertBulk) UpdatePrimaryPurpose() *PurposeUpsertBulk {
	return u.Update(func(s *PurposeUpsert) {
		s.UpdatePrimaryPurpose()
	})
}

// Exec executes the query.
func (u *PurposeUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PurposeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PurposeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PurposeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
