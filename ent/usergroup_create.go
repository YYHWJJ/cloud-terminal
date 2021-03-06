// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/willie-lin/cloud-terminal/ent/usergroup"
)

// UserGroupCreate is the builder for creating a UserGroup entity.
type UserGroupCreate struct {
	config
	mutation *UserGroupMutation
	hooks    []Hook
}

// Mutation returns the UserGroupMutation object of the builder.
func (ugc *UserGroupCreate) Mutation() *UserGroupMutation {
	return ugc.mutation
}

// Save creates the UserGroup in the database.
func (ugc *UserGroupCreate) Save(ctx context.Context) (*UserGroup, error) {
	var (
		err  error
		node *UserGroup
	)
	if len(ugc.hooks) == 0 {
		if err = ugc.check(); err != nil {
			return nil, err
		}
		node, err = ugc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ugc.check(); err != nil {
				return nil, err
			}
			ugc.mutation = mutation
			node, err = ugc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ugc.hooks) - 1; i >= 0; i-- {
			mut = ugc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ugc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ugc *UserGroupCreate) SaveX(ctx context.Context) *UserGroup {
	v, err := ugc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (ugc *UserGroupCreate) check() error {
	return nil
}

func (ugc *UserGroupCreate) sqlSave(ctx context.Context) (*UserGroup, error) {
	_node, _spec := ugc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ugc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ugc *UserGroupCreate) createSpec() (*UserGroup, *sqlgraph.CreateSpec) {
	var (
		_node = &UserGroup{config: ugc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: usergroup.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usergroup.FieldID,
			},
		}
	)
	return _node, _spec
}

// UserGroupCreateBulk is the builder for creating many UserGroup entities in bulk.
type UserGroupCreateBulk struct {
	config
	builders []*UserGroupCreate
}

// Save creates the UserGroup entities in the database.
func (ugcb *UserGroupCreateBulk) Save(ctx context.Context) ([]*UserGroup, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ugcb.builders))
	nodes := make([]*UserGroup, len(ugcb.builders))
	mutators := make([]Mutator, len(ugcb.builders))
	for i := range ugcb.builders {
		func(i int, root context.Context) {
			builder := ugcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserGroupMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ugcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ugcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ugcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ugcb *UserGroupCreateBulk) SaveX(ctx context.Context) []*UserGroup {
	v, err := ugcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
