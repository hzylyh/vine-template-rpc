// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"vine-template-rpc/internal/system/data/ent/predicate"
	"vine-template-rpc/internal/system/data/ent/resource"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ResourceUpdate is the builder for updating Resource entities.
type ResourceUpdate struct {
	config
	hooks    []Hook
	mutation *ResourceMutation
}

// Where appends a list predicates to the ResourceUpdate builder.
func (ru *ResourceUpdate) Where(ps ...predicate.Resource) *ResourceUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetName sets the "name" field.
func (ru *ResourceUpdate) SetName(s string) *ResourceUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetCode sets the "code" field.
func (ru *ResourceUpdate) SetCode(s string) *ResourceUpdate {
	ru.mutation.SetCode(s)
	return ru
}

// SetPath sets the "path" field.
func (ru *ResourceUpdate) SetPath(s string) *ResourceUpdate {
	ru.mutation.SetPath(s)
	return ru
}

// SetType sets the "type" field.
func (ru *ResourceUpdate) SetType(s string) *ResourceUpdate {
	ru.mutation.SetType(s)
	return ru
}

// SetRemark sets the "remark" field.
func (ru *ResourceUpdate) SetRemark(s string) *ResourceUpdate {
	ru.mutation.SetRemark(s)
	return ru
}

// SetStatus sets the "status" field.
func (ru *ResourceUpdate) SetStatus(i int) *ResourceUpdate {
	ru.mutation.ResetStatus()
	ru.mutation.SetStatus(i)
	return ru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ru *ResourceUpdate) SetNillableStatus(i *int) *ResourceUpdate {
	if i != nil {
		ru.SetStatus(*i)
	}
	return ru
}

// AddStatus adds i to the "status" field.
func (ru *ResourceUpdate) AddStatus(i int) *ResourceUpdate {
	ru.mutation.AddStatus(i)
	return ru
}

// Mutation returns the ResourceMutation object of the builder.
func (ru *ResourceUpdate) Mutation() *ResourceMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ResourceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ResourceUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ResourceUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ResourceUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *ResourceUpdate) check() error {
	if v, ok := ru.mutation.Name(); ok {
		if err := resource.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Resource.name": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Code(); ok {
		if err := resource.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Resource.code": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Path(); ok {
		if err := resource.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "Resource.path": %w`, err)}
		}
	}
	if v, ok := ru.mutation.GetType(); ok {
		if err := resource.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Resource.type": %w`, err)}
		}
	}
	return nil
}

func (ru *ResourceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(resource.Table, resource.Columns, sqlgraph.NewFieldSpec(resource.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(resource.FieldName, field.TypeString, value)
	}
	if value, ok := ru.mutation.Code(); ok {
		_spec.SetField(resource.FieldCode, field.TypeString, value)
	}
	if value, ok := ru.mutation.Path(); ok {
		_spec.SetField(resource.FieldPath, field.TypeString, value)
	}
	if value, ok := ru.mutation.GetType(); ok {
		_spec.SetField(resource.FieldType, field.TypeString, value)
	}
	if value, ok := ru.mutation.Remark(); ok {
		_spec.SetField(resource.FieldRemark, field.TypeString, value)
	}
	if value, ok := ru.mutation.Status(); ok {
		_spec.SetField(resource.FieldStatus, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedStatus(); ok {
		_spec.AddField(resource.FieldStatus, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// ResourceUpdateOne is the builder for updating a single Resource entity.
type ResourceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ResourceMutation
}

// SetName sets the "name" field.
func (ruo *ResourceUpdateOne) SetName(s string) *ResourceUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetCode sets the "code" field.
func (ruo *ResourceUpdateOne) SetCode(s string) *ResourceUpdateOne {
	ruo.mutation.SetCode(s)
	return ruo
}

// SetPath sets the "path" field.
func (ruo *ResourceUpdateOne) SetPath(s string) *ResourceUpdateOne {
	ruo.mutation.SetPath(s)
	return ruo
}

// SetType sets the "type" field.
func (ruo *ResourceUpdateOne) SetType(s string) *ResourceUpdateOne {
	ruo.mutation.SetType(s)
	return ruo
}

// SetRemark sets the "remark" field.
func (ruo *ResourceUpdateOne) SetRemark(s string) *ResourceUpdateOne {
	ruo.mutation.SetRemark(s)
	return ruo
}

// SetStatus sets the "status" field.
func (ruo *ResourceUpdateOne) SetStatus(i int) *ResourceUpdateOne {
	ruo.mutation.ResetStatus()
	ruo.mutation.SetStatus(i)
	return ruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ruo *ResourceUpdateOne) SetNillableStatus(i *int) *ResourceUpdateOne {
	if i != nil {
		ruo.SetStatus(*i)
	}
	return ruo
}

// AddStatus adds i to the "status" field.
func (ruo *ResourceUpdateOne) AddStatus(i int) *ResourceUpdateOne {
	ruo.mutation.AddStatus(i)
	return ruo
}

// Mutation returns the ResourceMutation object of the builder.
func (ruo *ResourceUpdateOne) Mutation() *ResourceMutation {
	return ruo.mutation
}

// Where appends a list predicates to the ResourceUpdate builder.
func (ruo *ResourceUpdateOne) Where(ps ...predicate.Resource) *ResourceUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ResourceUpdateOne) Select(field string, fields ...string) *ResourceUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Resource entity.
func (ruo *ResourceUpdateOne) Save(ctx context.Context) (*Resource, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ResourceUpdateOne) SaveX(ctx context.Context) *Resource {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ResourceUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ResourceUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *ResourceUpdateOne) check() error {
	if v, ok := ruo.mutation.Name(); ok {
		if err := resource.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Resource.name": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Code(); ok {
		if err := resource.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Resource.code": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Path(); ok {
		if err := resource.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "Resource.path": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.GetType(); ok {
		if err := resource.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Resource.type": %w`, err)}
		}
	}
	return nil
}

func (ruo *ResourceUpdateOne) sqlSave(ctx context.Context) (_node *Resource, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(resource.Table, resource.Columns, sqlgraph.NewFieldSpec(resource.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Resource.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, resource.FieldID)
		for _, f := range fields {
			if !resource.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != resource.FieldID {
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
		_spec.SetField(resource.FieldName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Code(); ok {
		_spec.SetField(resource.FieldCode, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Path(); ok {
		_spec.SetField(resource.FieldPath, field.TypeString, value)
	}
	if value, ok := ruo.mutation.GetType(); ok {
		_spec.SetField(resource.FieldType, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Remark(); ok {
		_spec.SetField(resource.FieldRemark, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Status(); ok {
		_spec.SetField(resource.FieldStatus, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedStatus(); ok {
		_spec.AddField(resource.FieldStatus, field.TypeInt, value)
	}
	_node = &Resource{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}