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
	"github.com/protobom/storage/internal/backends/ent/annotation"
	"github.com/protobom/storage/internal/backends/ent/document"
)

// AnnotationCreate is the builder for creating a Annotation entity.
type AnnotationCreate struct {
	config
	mutation *AnnotationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDocumentID sets the "document_id" field.
func (ac *AnnotationCreate) SetDocumentID(s string) *AnnotationCreate {
	ac.mutation.SetDocumentID(s)
	return ac
}

// SetName sets the "name" field.
func (ac *AnnotationCreate) SetName(s string) *AnnotationCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetValue sets the "value" field.
func (ac *AnnotationCreate) SetValue(s string) *AnnotationCreate {
	ac.mutation.SetValue(s)
	return ac
}

// SetDocument sets the "document" edge to the Document entity.
func (ac *AnnotationCreate) SetDocument(d *Document) *AnnotationCreate {
	return ac.SetDocumentID(d.ID)
}

// Mutation returns the AnnotationMutation object of the builder.
func (ac *AnnotationCreate) Mutation() *AnnotationMutation {
	return ac.mutation
}

// Save creates the Annotation in the database.
func (ac *AnnotationCreate) Save(ctx context.Context) (*Annotation, error) {
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AnnotationCreate) SaveX(ctx context.Context) *Annotation {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AnnotationCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AnnotationCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AnnotationCreate) check() error {
	if _, ok := ac.mutation.DocumentID(); !ok {
		return &ValidationError{Name: "document_id", err: errors.New(`ent: missing required field "Annotation.document_id"`)}
	}
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Annotation.name"`)}
	}
	if _, ok := ac.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "Annotation.value"`)}
	}
	if len(ac.mutation.DocumentIDs()) == 0 {
		return &ValidationError{Name: "document", err: errors.New(`ent: missing required edge "Annotation.document"`)}
	}
	return nil
}

func (ac *AnnotationCreate) sqlSave(ctx context.Context) (*Annotation, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AnnotationCreate) createSpec() (*Annotation, *sqlgraph.CreateSpec) {
	var (
		_node = &Annotation{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(annotation.Table, sqlgraph.NewFieldSpec(annotation.FieldID, field.TypeInt))
	)
	_spec.OnConflict = ac.conflict
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(annotation.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.Value(); ok {
		_spec.SetField(annotation.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	if nodes := ac.mutation.DocumentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   annotation.DocumentTable,
			Columns: []string{annotation.DocumentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(document.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DocumentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Annotation.Create().
//		SetDocumentID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AnnotationUpsert) {
//			SetDocumentID(v+v).
//		}).
//		Exec(ctx)
func (ac *AnnotationCreate) OnConflict(opts ...sql.ConflictOption) *AnnotationUpsertOne {
	ac.conflict = opts
	return &AnnotationUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Annotation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AnnotationCreate) OnConflictColumns(columns ...string) *AnnotationUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AnnotationUpsertOne{
		create: ac,
	}
}

type (
	// AnnotationUpsertOne is the builder for "upsert"-ing
	//  one Annotation node.
	AnnotationUpsertOne struct {
		create *AnnotationCreate
	}

	// AnnotationUpsert is the "OnConflict" setter.
	AnnotationUpsert struct {
		*sql.UpdateSet
	}
)

// SetDocumentID sets the "document_id" field.
func (u *AnnotationUpsert) SetDocumentID(v string) *AnnotationUpsert {
	u.Set(annotation.FieldDocumentID, v)
	return u
}

// UpdateDocumentID sets the "document_id" field to the value that was provided on create.
func (u *AnnotationUpsert) UpdateDocumentID() *AnnotationUpsert {
	u.SetExcluded(annotation.FieldDocumentID)
	return u
}

// SetName sets the "name" field.
func (u *AnnotationUpsert) SetName(v string) *AnnotationUpsert {
	u.Set(annotation.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AnnotationUpsert) UpdateName() *AnnotationUpsert {
	u.SetExcluded(annotation.FieldName)
	return u
}

// SetValue sets the "value" field.
func (u *AnnotationUpsert) SetValue(v string) *AnnotationUpsert {
	u.Set(annotation.FieldValue, v)
	return u
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *AnnotationUpsert) UpdateValue() *AnnotationUpsert {
	u.SetExcluded(annotation.FieldValue)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Annotation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AnnotationUpsertOne) UpdateNewValues() *AnnotationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Annotation.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AnnotationUpsertOne) Ignore() *AnnotationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AnnotationUpsertOne) DoNothing() *AnnotationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AnnotationCreate.OnConflict
// documentation for more info.
func (u *AnnotationUpsertOne) Update(set func(*AnnotationUpsert)) *AnnotationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AnnotationUpsert{UpdateSet: update})
	}))
	return u
}

// SetDocumentID sets the "document_id" field.
func (u *AnnotationUpsertOne) SetDocumentID(v string) *AnnotationUpsertOne {
	return u.Update(func(s *AnnotationUpsert) {
		s.SetDocumentID(v)
	})
}

// UpdateDocumentID sets the "document_id" field to the value that was provided on create.
func (u *AnnotationUpsertOne) UpdateDocumentID() *AnnotationUpsertOne {
	return u.Update(func(s *AnnotationUpsert) {
		s.UpdateDocumentID()
	})
}

// SetName sets the "name" field.
func (u *AnnotationUpsertOne) SetName(v string) *AnnotationUpsertOne {
	return u.Update(func(s *AnnotationUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AnnotationUpsertOne) UpdateName() *AnnotationUpsertOne {
	return u.Update(func(s *AnnotationUpsert) {
		s.UpdateName()
	})
}

// SetValue sets the "value" field.
func (u *AnnotationUpsertOne) SetValue(v string) *AnnotationUpsertOne {
	return u.Update(func(s *AnnotationUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *AnnotationUpsertOne) UpdateValue() *AnnotationUpsertOne {
	return u.Update(func(s *AnnotationUpsert) {
		s.UpdateValue()
	})
}

// Exec executes the query.
func (u *AnnotationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AnnotationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AnnotationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AnnotationUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AnnotationUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AnnotationCreateBulk is the builder for creating many Annotation entities in bulk.
type AnnotationCreateBulk struct {
	config
	err      error
	builders []*AnnotationCreate
	conflict []sql.ConflictOption
}

// Save creates the Annotation entities in the database.
func (acb *AnnotationCreateBulk) Save(ctx context.Context) ([]*Annotation, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Annotation, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AnnotationMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AnnotationCreateBulk) SaveX(ctx context.Context) []*Annotation {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AnnotationCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AnnotationCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Annotation.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AnnotationUpsert) {
//			SetDocumentID(v+v).
//		}).
//		Exec(ctx)
func (acb *AnnotationCreateBulk) OnConflict(opts ...sql.ConflictOption) *AnnotationUpsertBulk {
	acb.conflict = opts
	return &AnnotationUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Annotation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AnnotationCreateBulk) OnConflictColumns(columns ...string) *AnnotationUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AnnotationUpsertBulk{
		create: acb,
	}
}

// AnnotationUpsertBulk is the builder for "upsert"-ing
// a bulk of Annotation nodes.
type AnnotationUpsertBulk struct {
	create *AnnotationCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Annotation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AnnotationUpsertBulk) UpdateNewValues() *AnnotationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Annotation.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AnnotationUpsertBulk) Ignore() *AnnotationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AnnotationUpsertBulk) DoNothing() *AnnotationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AnnotationCreateBulk.OnConflict
// documentation for more info.
func (u *AnnotationUpsertBulk) Update(set func(*AnnotationUpsert)) *AnnotationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AnnotationUpsert{UpdateSet: update})
	}))
	return u
}

// SetDocumentID sets the "document_id" field.
func (u *AnnotationUpsertBulk) SetDocumentID(v string) *AnnotationUpsertBulk {
	return u.Update(func(s *AnnotationUpsert) {
		s.SetDocumentID(v)
	})
}

// UpdateDocumentID sets the "document_id" field to the value that was provided on create.
func (u *AnnotationUpsertBulk) UpdateDocumentID() *AnnotationUpsertBulk {
	return u.Update(func(s *AnnotationUpsert) {
		s.UpdateDocumentID()
	})
}

// SetName sets the "name" field.
func (u *AnnotationUpsertBulk) SetName(v string) *AnnotationUpsertBulk {
	return u.Update(func(s *AnnotationUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AnnotationUpsertBulk) UpdateName() *AnnotationUpsertBulk {
	return u.Update(func(s *AnnotationUpsert) {
		s.UpdateName()
	})
}

// SetValue sets the "value" field.
func (u *AnnotationUpsertBulk) SetValue(v string) *AnnotationUpsertBulk {
	return u.Update(func(s *AnnotationUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *AnnotationUpsertBulk) UpdateValue() *AnnotationUpsertBulk {
	return u.Update(func(s *AnnotationUpsert) {
		s.UpdateValue()
	})
}

// Exec executes the query.
func (u *AnnotationUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AnnotationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AnnotationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AnnotationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
