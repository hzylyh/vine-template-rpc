// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"vine-template-rpc/internal/system/data/ent/predicate"
	"vine-template-rpc/internal/system/data/ent/role"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoleUpdate is the builder for updating Role entities.
type RoleUpdate struct {
	config
	hooks    []Hook
	mutation *RoleMutation
}

// Where appends a list predicates to the RoleUpdate builder.
func (ru *RoleUpdate) Where(ps ...predicate.Role) *RoleUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetName sets the "name" field.
func (ru *RoleUpdate) SetName(s string) *RoleUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetCode sets the "code" field.
func (ru *RoleUpdate) SetCode(s string) *RoleUpdate {
	ru.mutation.SetCode(s)
	return ru
}

// SetRemark sets the "remark" field.
func (ru *RoleUpdate) SetRemark(s string) *RoleUpdate {
	ru.mutation.SetRemark(s)
	return ru
}

// SetStatus sets the "status" field.
func (ru *RoleUpdate) SetStatus(i int) *RoleUpdate {
	ru.mutation.ResetStatus()
	ru.mutation.SetStatus(i)
	return ru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableStatus(i *int) *RoleUpdate {
	if i != nil {
		ru.SetStatus(*i)
	}
	return ru
}

// AddStatus adds i to the "status" field.
func (ru *RoleUpdate) AddStatus(i int) *RoleUpdate {
	ru.mutation.AddStatus(i)
	return ru
}

// Mutation returns the RoleMutation object of the builder.
func (ru *RoleUpdate) Mutation() *RoleMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RoleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoleUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoleUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoleUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RoleUpdate) check() error {
	if v, ok := ru.mutation.Name(); ok {
		if err := role.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Role.name": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Code(); ok {
		if err := role.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Role.code": %w`, err)}
		}
	}
	return nil
}

func (ru *RoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(role.Table, role.Columns, sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(role.FieldName, field.TypeString, value)
	}
	if value, ok := ru.mutation.Code(); ok {
		_spec.SetField(role.FieldCode, field.TypeString, value)
	}
	if value, ok := ru.mutation.Remark(); ok {
		_spec.SetField(role.FieldRemark, field.TypeString, value)
	}
	if value, ok := ru.mutation.Status(); ok {
		_spec.SetField(role.FieldStatus, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedStatus(); ok {
		_spec.AddField(role.FieldStatus, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RoleUpdateOne is the builder for updating a single Role entity.
type RoleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoleMutation
}

// SetName sets the "name" field.
func (ruo *RoleUpdateOne) SetName(s string) *RoleUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetCode sets the "code" field.
func (ruo *RoleUpdateOne) SetCode(s string) *RoleUpdateOne {
	ruo.mutation.SetCode(s)
	return ruo
}

// SetRemark sets the "remark" field.
func (ruo *RoleUpdateOne) SetRemark(s string) *RoleUpdateOne {
	ruo.mutation.SetRemark(s)
	return ruo
}

// SetStatus sets the "status" field.
func (ruo *RoleUpdateOne) SetStatus(i int) *RoleUpdateOne {
	ruo.mutation.ResetStatus()
	ruo.mutation.SetStatus(i)
	return ruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableStatus(i *int) *RoleUpdateOne {
	if i != nil {
		ruo.SetStatus(*i)
	}
	return ruo
}

// AddStatus adds i to the "status" field.
func (ruo *RoleUpdateOne) AddStatus(i int) *RoleUpdateOne {
	ruo.mutation.AddStatus(i)
	return ruo
}

// Mutation returns the RoleMutation object of the builder.
func (ruo *RoleUpdateOne) Mutation() *RoleMutation {
	return ruo.mutation
}

// Where appends a list predicates to the RoleUpdate builder.
func (ruo *RoleUpdateOne) Where(ps ...predicate.Role) *RoleUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RoleUpdateOne) Select(field string, fields ...string) *RoleUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Role entity.
func (ruo *RoleUpdateOne) Save(ctx context.Context) (*Role, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoleUpdateOne) SaveX(ctx context.Context) *Role {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RoleUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoleUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RoleUpdateOne) check() error {
	if v, ok := ruo.mutation.Name(); ok {
		if err := role.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Role.name": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Code(); ok {
		if err := role.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Role.code": %w`, err)}
		}
	}
	return nil
}

func (ruo *RoleUpdateOne) sqlSave(ctx context.Context) (_node *Role, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(role.Table, role.Columns, sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Role.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, role.FieldID)
		for _, f := range fields {
			if !role.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != role.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(role.FieldName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Code(); ok {
		_spec.SetField(role.FieldCode, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Remark(); ok {
		_spec.SetField(role.FieldRemark, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Status(); ok {
		_spec.SetField(role.FieldStatus, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedStatus(); ok {
		_spec.AddField(role.FieldStatus, field.TypeInt, value)
	}
	_node = &Role{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
