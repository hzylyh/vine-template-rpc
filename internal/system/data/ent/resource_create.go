// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"vine-template-rpc/internal/system/data/ent/resource"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ResourceCreate is the builder for creating a Resource entity.
type ResourceCreate struct {
	config
	mutation *ResourceMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (rc *ResourceCreate) SetName(s string) *ResourceCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetCode sets the "code" field.
func (rc *ResourceCreate) SetCode(s string) *ResourceCreate {
	rc.mutation.SetCode(s)
	return rc
}

// SetPath sets the "path" field.
func (rc *ResourceCreate) SetPath(s string) *ResourceCreate {
	rc.mutation.SetPath(s)
	return rc
}

// SetType sets the "type" field.
func (rc *ResourceCreate) SetType(s string) *ResourceCreate {
	rc.mutation.SetType(s)
	return rc
}

// SetRemark sets the "remark" field.
func (rc *ResourceCreate) SetRemark(s string) *ResourceCreate {
	rc.mutation.SetRemark(s)
	return rc
}

// SetStatus sets the "status" field.
func (rc *ResourceCreate) SetStatus(i int) *ResourceCreate {
	rc.mutation.SetStatus(i)
	return rc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rc *ResourceCreate) SetNillableStatus(i *int) *ResourceCreate {
	if i != nil {
		rc.SetStatus(*i)
	}
	return rc
}

// Mutation returns the ResourceMutation object of the builder.
func (rc *ResourceCreate) Mutation() *ResourceMutation {
	return rc.mutation
}

// Save creates the Resource in the database.
func (rc *ResourceCreate) Save(ctx context.Context) (*Resource, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ResourceCreate) SaveX(ctx context.Context) *Resource {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ResourceCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ResourceCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *ResourceCreate) defaults() {
	if _, ok := rc.mutation.Status(); !ok {
		v := resource.DefaultStatus
		rc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ResourceCreate) check() error {
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Resource.name"`)}
	}
	if v, ok := rc.mutation.Name(); ok {
		if err := resource.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Resource.name": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Resource.code"`)}
	}
	if v, ok := rc.mutation.Code(); ok {
		if err := resource.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Resource.code": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "Resource.path"`)}
	}
	if v, ok := rc.mutation.Path(); ok {
		if err := resource.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "Resource.path": %w`, err)}
		}
	}
	if _, ok := rc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Resource.type"`)}
	}
	if v, ok := rc.mutation.GetType(); ok {
		if err := resource.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Resource.type": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Remark(); !ok {
		return &ValidationError{Name: "remark", err: errors.New(`ent: missing required field "Resource.remark"`)}
	}
	if _, ok := rc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Resource.status"`)}
	}
	return nil
}

func (rc *ResourceCreate) sqlSave(ctx context.Context) (*Resource, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *ResourceCreate) createSpec() (*Resource, *sqlgraph.CreateSpec) {
	var (
		_node = &Resource{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(resource.Table, sqlgraph.NewFieldSpec(resource.FieldID, field.TypeInt))
	)
	if value, ok := rc.mutation.Name(); ok {
		_spec.SetField(resource.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := rc.mutation.Code(); ok {
		_spec.SetField(resource.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := rc.mutation.Path(); ok {
		_spec.SetField(resource.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := rc.mutation.GetType(); ok {
		_spec.SetField(resource.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := rc.mutation.Remark(); ok {
		_spec.SetField(resource.FieldRemark, field.TypeString, value)
		_node.Remark = value
	}
	if value, ok := rc.mutation.Status(); ok {
		_spec.SetField(resource.FieldStatus, field.TypeInt, value)
		_node.Status = value
	}
	return _node, _spec
}

// ResourceCreateBulk is the builder for creating many Resource entities in bulk.
type ResourceCreateBulk struct {
	config
	err      error
	builders []*ResourceCreate
}

// Save creates the Resource entities in the database.
func (rcb *ResourceCreateBulk) Save(ctx context.Context) ([]*Resource, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Resource, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ResourceMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ResourceCreateBulk) SaveX(ctx context.Context) []*Resource {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ResourceCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ResourceCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
