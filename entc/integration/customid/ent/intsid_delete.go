// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/intsid"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/schema/field"
)

// IntSIDDelete is the builder for deleting a IntSID entity.
type IntSIDDelete struct {
	config
	hooks    []Hook
	mutation *IntSIDMutation
}

// Where appends a list predicates to the IntSIDDelete builder.
func (isd *IntSIDDelete) Where(ps ...predicate.IntSID) *IntSIDDelete {
	isd.mutation.Where(ps...)
	return isd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (isd *IntSIDDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(isd.hooks) == 0 {
		affected, err = isd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IntSIDMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			isd.mutation = mutation
			affected, err = isd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(isd.hooks) - 1; i >= 0; i-- {
			if isd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = isd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, isd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (isd *IntSIDDelete) ExecX(ctx context.Context) int {
	n, err := isd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (isd *IntSIDDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: intsid.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: intsid.FieldID,
			},
		},
	}
	if ps := isd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, isd.driver, _spec)
}

// IntSIDDeleteOne is the builder for deleting a single IntSID entity.
type IntSIDDeleteOne struct {
	isd *IntSIDDelete
}

// Exec executes the deletion query.
func (isdo *IntSIDDeleteOne) Exec(ctx context.Context) error {
	n, err := isdo.isd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{intsid.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (isdo *IntSIDDeleteOne) ExecX(ctx context.Context) {
	isdo.isd.ExecX(ctx)
}
