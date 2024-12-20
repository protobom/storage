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
	"github.com/protobom/storage/internal/backends/ent/metadata"
	"github.com/protobom/storage/internal/backends/ent/node"
	"github.com/protobom/storage/internal/backends/ent/person"
	"github.com/protobom/storage/internal/backends/ent/predicate"
)

// PersonUpdate is the builder for updating Person entities.
type PersonUpdate struct {
	config
	hooks    []Hook
	mutation *PersonMutation
}

// Where appends a list predicates to the PersonUpdate builder.
func (pu *PersonUpdate) Where(ps ...predicate.Person) *PersonUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetMetadataID sets the "metadata_id" field.
func (pu *PersonUpdate) SetMetadataID(u uuid.UUID) *PersonUpdate {
	pu.mutation.SetMetadataID(u)
	return pu
}

// SetNillableMetadataID sets the "metadata_id" field if the given value is not nil.
func (pu *PersonUpdate) SetNillableMetadataID(u *uuid.UUID) *PersonUpdate {
	if u != nil {
		pu.SetMetadataID(*u)
	}
	return pu
}

// ClearMetadataID clears the value of the "metadata_id" field.
func (pu *PersonUpdate) ClearMetadataID() *PersonUpdate {
	pu.mutation.ClearMetadataID()
	return pu
}

// SetNodeID sets the "node_id" field.
func (pu *PersonUpdate) SetNodeID(u uuid.UUID) *PersonUpdate {
	pu.mutation.SetNodeID(u)
	return pu
}

// SetNillableNodeID sets the "node_id" field if the given value is not nil.
func (pu *PersonUpdate) SetNillableNodeID(u *uuid.UUID) *PersonUpdate {
	if u != nil {
		pu.SetNodeID(*u)
	}
	return pu
}

// ClearNodeID clears the value of the "node_id" field.
func (pu *PersonUpdate) ClearNodeID() *PersonUpdate {
	pu.mutation.ClearNodeID()
	return pu
}

// SetName sets the "name" field.
func (pu *PersonUpdate) SetName(s string) *PersonUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PersonUpdate) SetNillableName(s *string) *PersonUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetIsOrg sets the "is_org" field.
func (pu *PersonUpdate) SetIsOrg(b bool) *PersonUpdate {
	pu.mutation.SetIsOrg(b)
	return pu
}

// SetNillableIsOrg sets the "is_org" field if the given value is not nil.
func (pu *PersonUpdate) SetNillableIsOrg(b *bool) *PersonUpdate {
	if b != nil {
		pu.SetIsOrg(*b)
	}
	return pu
}

// SetEmail sets the "email" field.
func (pu *PersonUpdate) SetEmail(s string) *PersonUpdate {
	pu.mutation.SetEmail(s)
	return pu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (pu *PersonUpdate) SetNillableEmail(s *string) *PersonUpdate {
	if s != nil {
		pu.SetEmail(*s)
	}
	return pu
}

// SetURL sets the "url" field.
func (pu *PersonUpdate) SetURL(s string) *PersonUpdate {
	pu.mutation.SetURL(s)
	return pu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (pu *PersonUpdate) SetNillableURL(s *string) *PersonUpdate {
	if s != nil {
		pu.SetURL(*s)
	}
	return pu
}

// SetPhone sets the "phone" field.
func (pu *PersonUpdate) SetPhone(s string) *PersonUpdate {
	pu.mutation.SetPhone(s)
	return pu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (pu *PersonUpdate) SetNillablePhone(s *string) *PersonUpdate {
	if s != nil {
		pu.SetPhone(*s)
	}
	return pu
}

// SetContactOwnerID sets the "contact_owner" edge to the Person entity by ID.
func (pu *PersonUpdate) SetContactOwnerID(id uuid.UUID) *PersonUpdate {
	pu.mutation.SetContactOwnerID(id)
	return pu
}

// SetNillableContactOwnerID sets the "contact_owner" edge to the Person entity by ID if the given value is not nil.
func (pu *PersonUpdate) SetNillableContactOwnerID(id *uuid.UUID) *PersonUpdate {
	if id != nil {
		pu = pu.SetContactOwnerID(*id)
	}
	return pu
}

// SetContactOwner sets the "contact_owner" edge to the Person entity.
func (pu *PersonUpdate) SetContactOwner(p *Person) *PersonUpdate {
	return pu.SetContactOwnerID(p.ID)
}

// AddContactIDs adds the "contacts" edge to the Person entity by IDs.
func (pu *PersonUpdate) AddContactIDs(ids ...uuid.UUID) *PersonUpdate {
	pu.mutation.AddContactIDs(ids...)
	return pu
}

// AddContacts adds the "contacts" edges to the Person entity.
func (pu *PersonUpdate) AddContacts(p ...*Person) *PersonUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddContactIDs(ids...)
}

// SetMetadata sets the "metadata" edge to the Metadata entity.
func (pu *PersonUpdate) SetMetadata(m *Metadata) *PersonUpdate {
	return pu.SetMetadataID(m.ID)
}

// SetNode sets the "node" edge to the Node entity.
func (pu *PersonUpdate) SetNode(n *Node) *PersonUpdate {
	return pu.SetNodeID(n.ID)
}

// Mutation returns the PersonMutation object of the builder.
func (pu *PersonUpdate) Mutation() *PersonMutation {
	return pu.mutation
}

// ClearContactOwner clears the "contact_owner" edge to the Person entity.
func (pu *PersonUpdate) ClearContactOwner() *PersonUpdate {
	pu.mutation.ClearContactOwner()
	return pu
}

// ClearContacts clears all "contacts" edges to the Person entity.
func (pu *PersonUpdate) ClearContacts() *PersonUpdate {
	pu.mutation.ClearContacts()
	return pu
}

// RemoveContactIDs removes the "contacts" edge to Person entities by IDs.
func (pu *PersonUpdate) RemoveContactIDs(ids ...uuid.UUID) *PersonUpdate {
	pu.mutation.RemoveContactIDs(ids...)
	return pu
}

// RemoveContacts removes "contacts" edges to Person entities.
func (pu *PersonUpdate) RemoveContacts(p ...*Person) *PersonUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveContactIDs(ids...)
}

// ClearMetadata clears the "metadata" edge to the Metadata entity.
func (pu *PersonUpdate) ClearMetadata() *PersonUpdate {
	pu.mutation.ClearMetadata()
	return pu
}

// ClearNode clears the "node" edge to the Node entity.
func (pu *PersonUpdate) ClearNode() *PersonUpdate {
	pu.mutation.ClearNode()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PersonUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PersonUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PersonUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PersonUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PersonUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(person.Table, person.Columns, sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(person.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.IsOrg(); ok {
		_spec.SetField(person.FieldIsOrg, field.TypeBool, value)
	}
	if value, ok := pu.mutation.Email(); ok {
		_spec.SetField(person.FieldEmail, field.TypeString, value)
	}
	if value, ok := pu.mutation.URL(); ok {
		_spec.SetField(person.FieldURL, field.TypeString, value)
	}
	if value, ok := pu.mutation.Phone(); ok {
		_spec.SetField(person.FieldPhone, field.TypeString, value)
	}
	if pu.mutation.ContactOwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   person.ContactOwnerTable,
			Columns: []string{person.ContactOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ContactOwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   person.ContactOwnerTable,
			Columns: []string{person.ContactOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ContactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   person.ContactsTable,
			Columns: []string{person.ContactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedContactsIDs(); len(nodes) > 0 && !pu.mutation.ContactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   person.ContactsTable,
			Columns: []string{person.ContactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ContactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   person.ContactsTable,
			Columns: []string{person.ContactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   person.MetadataTable,
			Columns: []string{person.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   person.MetadataTable,
			Columns: []string{person.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.NodeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   person.NodeTable,
			Columns: []string{person.NodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.NodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   person.NodeTable,
			Columns: []string{person.NodeColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{person.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PersonUpdateOne is the builder for updating a single Person entity.
type PersonUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PersonMutation
}

// SetMetadataID sets the "metadata_id" field.
func (puo *PersonUpdateOne) SetMetadataID(u uuid.UUID) *PersonUpdateOne {
	puo.mutation.SetMetadataID(u)
	return puo
}

// SetNillableMetadataID sets the "metadata_id" field if the given value is not nil.
func (puo *PersonUpdateOne) SetNillableMetadataID(u *uuid.UUID) *PersonUpdateOne {
	if u != nil {
		puo.SetMetadataID(*u)
	}
	return puo
}

// ClearMetadataID clears the value of the "metadata_id" field.
func (puo *PersonUpdateOne) ClearMetadataID() *PersonUpdateOne {
	puo.mutation.ClearMetadataID()
	return puo
}

// SetNodeID sets the "node_id" field.
func (puo *PersonUpdateOne) SetNodeID(u uuid.UUID) *PersonUpdateOne {
	puo.mutation.SetNodeID(u)
	return puo
}

// SetNillableNodeID sets the "node_id" field if the given value is not nil.
func (puo *PersonUpdateOne) SetNillableNodeID(u *uuid.UUID) *PersonUpdateOne {
	if u != nil {
		puo.SetNodeID(*u)
	}
	return puo
}

// ClearNodeID clears the value of the "node_id" field.
func (puo *PersonUpdateOne) ClearNodeID() *PersonUpdateOne {
	puo.mutation.ClearNodeID()
	return puo
}

// SetName sets the "name" field.
func (puo *PersonUpdateOne) SetName(s string) *PersonUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PersonUpdateOne) SetNillableName(s *string) *PersonUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetIsOrg sets the "is_org" field.
func (puo *PersonUpdateOne) SetIsOrg(b bool) *PersonUpdateOne {
	puo.mutation.SetIsOrg(b)
	return puo
}

// SetNillableIsOrg sets the "is_org" field if the given value is not nil.
func (puo *PersonUpdateOne) SetNillableIsOrg(b *bool) *PersonUpdateOne {
	if b != nil {
		puo.SetIsOrg(*b)
	}
	return puo
}

// SetEmail sets the "email" field.
func (puo *PersonUpdateOne) SetEmail(s string) *PersonUpdateOne {
	puo.mutation.SetEmail(s)
	return puo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (puo *PersonUpdateOne) SetNillableEmail(s *string) *PersonUpdateOne {
	if s != nil {
		puo.SetEmail(*s)
	}
	return puo
}

// SetURL sets the "url" field.
func (puo *PersonUpdateOne) SetURL(s string) *PersonUpdateOne {
	puo.mutation.SetURL(s)
	return puo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (puo *PersonUpdateOne) SetNillableURL(s *string) *PersonUpdateOne {
	if s != nil {
		puo.SetURL(*s)
	}
	return puo
}

// SetPhone sets the "phone" field.
func (puo *PersonUpdateOne) SetPhone(s string) *PersonUpdateOne {
	puo.mutation.SetPhone(s)
	return puo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (puo *PersonUpdateOne) SetNillablePhone(s *string) *PersonUpdateOne {
	if s != nil {
		puo.SetPhone(*s)
	}
	return puo
}

// SetContactOwnerID sets the "contact_owner" edge to the Person entity by ID.
func (puo *PersonUpdateOne) SetContactOwnerID(id uuid.UUID) *PersonUpdateOne {
	puo.mutation.SetContactOwnerID(id)
	return puo
}

// SetNillableContactOwnerID sets the "contact_owner" edge to the Person entity by ID if the given value is not nil.
func (puo *PersonUpdateOne) SetNillableContactOwnerID(id *uuid.UUID) *PersonUpdateOne {
	if id != nil {
		puo = puo.SetContactOwnerID(*id)
	}
	return puo
}

// SetContactOwner sets the "contact_owner" edge to the Person entity.
func (puo *PersonUpdateOne) SetContactOwner(p *Person) *PersonUpdateOne {
	return puo.SetContactOwnerID(p.ID)
}

// AddContactIDs adds the "contacts" edge to the Person entity by IDs.
func (puo *PersonUpdateOne) AddContactIDs(ids ...uuid.UUID) *PersonUpdateOne {
	puo.mutation.AddContactIDs(ids...)
	return puo
}

// AddContacts adds the "contacts" edges to the Person entity.
func (puo *PersonUpdateOne) AddContacts(p ...*Person) *PersonUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddContactIDs(ids...)
}

// SetMetadata sets the "metadata" edge to the Metadata entity.
func (puo *PersonUpdateOne) SetMetadata(m *Metadata) *PersonUpdateOne {
	return puo.SetMetadataID(m.ID)
}

// SetNode sets the "node" edge to the Node entity.
func (puo *PersonUpdateOne) SetNode(n *Node) *PersonUpdateOne {
	return puo.SetNodeID(n.ID)
}

// Mutation returns the PersonMutation object of the builder.
func (puo *PersonUpdateOne) Mutation() *PersonMutation {
	return puo.mutation
}

// ClearContactOwner clears the "contact_owner" edge to the Person entity.
func (puo *PersonUpdateOne) ClearContactOwner() *PersonUpdateOne {
	puo.mutation.ClearContactOwner()
	return puo
}

// ClearContacts clears all "contacts" edges to the Person entity.
func (puo *PersonUpdateOne) ClearContacts() *PersonUpdateOne {
	puo.mutation.ClearContacts()
	return puo
}

// RemoveContactIDs removes the "contacts" edge to Person entities by IDs.
func (puo *PersonUpdateOne) RemoveContactIDs(ids ...uuid.UUID) *PersonUpdateOne {
	puo.mutation.RemoveContactIDs(ids...)
	return puo
}

// RemoveContacts removes "contacts" edges to Person entities.
func (puo *PersonUpdateOne) RemoveContacts(p ...*Person) *PersonUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveContactIDs(ids...)
}

// ClearMetadata clears the "metadata" edge to the Metadata entity.
func (puo *PersonUpdateOne) ClearMetadata() *PersonUpdateOne {
	puo.mutation.ClearMetadata()
	return puo
}

// ClearNode clears the "node" edge to the Node entity.
func (puo *PersonUpdateOne) ClearNode() *PersonUpdateOne {
	puo.mutation.ClearNode()
	return puo
}

// Where appends a list predicates to the PersonUpdate builder.
func (puo *PersonUpdateOne) Where(ps ...predicate.Person) *PersonUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PersonUpdateOne) Select(field string, fields ...string) *PersonUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Person entity.
func (puo *PersonUpdateOne) Save(ctx context.Context) (*Person, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PersonUpdateOne) SaveX(ctx context.Context) *Person {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PersonUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PersonUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PersonUpdateOne) sqlSave(ctx context.Context) (_node *Person, err error) {
	_spec := sqlgraph.NewUpdateSpec(person.Table, person.Columns, sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Person.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, person.FieldID)
		for _, f := range fields {
			if !person.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != person.FieldID {
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
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(person.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.IsOrg(); ok {
		_spec.SetField(person.FieldIsOrg, field.TypeBool, value)
	}
	if value, ok := puo.mutation.Email(); ok {
		_spec.SetField(person.FieldEmail, field.TypeString, value)
	}
	if value, ok := puo.mutation.URL(); ok {
		_spec.SetField(person.FieldURL, field.TypeString, value)
	}
	if value, ok := puo.mutation.Phone(); ok {
		_spec.SetField(person.FieldPhone, field.TypeString, value)
	}
	if puo.mutation.ContactOwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   person.ContactOwnerTable,
			Columns: []string{person.ContactOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ContactOwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   person.ContactOwnerTable,
			Columns: []string{person.ContactOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ContactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   person.ContactsTable,
			Columns: []string{person.ContactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedContactsIDs(); len(nodes) > 0 && !puo.mutation.ContactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   person.ContactsTable,
			Columns: []string{person.ContactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ContactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   person.ContactsTable,
			Columns: []string{person.ContactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.MetadataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   person.MetadataTable,
			Columns: []string{person.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   person.MetadataTable,
			Columns: []string{person.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.NodeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   person.NodeTable,
			Columns: []string{person.NodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.NodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   person.NodeTable,
			Columns: []string{person.NodeColumn},
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
	_node = &Person{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{person.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
