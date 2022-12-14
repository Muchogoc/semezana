// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Muchogoc/semezana/ent/device"
	"github.com/Muchogoc/semezana/ent/predicate"
	"github.com/Muchogoc/semezana/ent/user"
	"github.com/google/uuid"
)

// DeviceUpdate is the builder for updating Device entities.
type DeviceUpdate struct {
	config
	hooks    []Hook
	mutation *DeviceMutation
}

// Where appends a list predicates to the DeviceUpdate builder.
func (du *DeviceUpdate) Where(ps ...predicate.Device) *DeviceUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetCreatedAt sets the "created_at" field.
func (du *DeviceUpdate) SetCreatedAt(t time.Time) *DeviceUpdate {
	du.mutation.SetCreatedAt(t)
	return du
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (du *DeviceUpdate) SetNillableCreatedAt(t *time.Time) *DeviceUpdate {
	if t != nil {
		du.SetCreatedAt(*t)
	}
	return du
}

// SetUpdatedAt sets the "updated_at" field.
func (du *DeviceUpdate) SetUpdatedAt(t time.Time) *DeviceUpdate {
	du.mutation.SetUpdatedAt(t)
	return du
}

// SetHash sets the "hash" field.
func (du *DeviceUpdate) SetHash(s string) *DeviceUpdate {
	du.mutation.SetHash(s)
	return du
}

// SetDeviceID sets the "device_id" field.
func (du *DeviceUpdate) SetDeviceID(s string) *DeviceUpdate {
	du.mutation.SetDeviceID(s)
	return du
}

// SetPlatform sets the "platform" field.
func (du *DeviceUpdate) SetPlatform(s string) *DeviceUpdate {
	du.mutation.SetPlatform(s)
	return du
}

// SetLastSeen sets the "last_seen" field.
func (du *DeviceUpdate) SetLastSeen(t time.Time) *DeviceUpdate {
	du.mutation.SetLastSeen(t)
	return du
}

// SetLanguage sets the "language" field.
func (du *DeviceUpdate) SetLanguage(s string) *DeviceUpdate {
	du.mutation.SetLanguage(s)
	return du
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (du *DeviceUpdate) SetOwnerID(id uuid.UUID) *DeviceUpdate {
	du.mutation.SetOwnerID(id)
	return du
}

// SetOwner sets the "owner" edge to the User entity.
func (du *DeviceUpdate) SetOwner(u *User) *DeviceUpdate {
	return du.SetOwnerID(u.ID)
}

// Mutation returns the DeviceMutation object of the builder.
func (du *DeviceUpdate) Mutation() *DeviceMutation {
	return du.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (du *DeviceUpdate) ClearOwner() *DeviceUpdate {
	du.mutation.ClearOwner()
	return du
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DeviceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	du.defaults()
	if len(du.hooks) == 0 {
		if err = du.check(); err != nil {
			return 0, err
		}
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeviceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = du.check(); err != nil {
				return 0, err
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			if du.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DeviceUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DeviceUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DeviceUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DeviceUpdate) defaults() {
	if _, ok := du.mutation.UpdatedAt(); !ok {
		v := device.UpdateDefaultUpdatedAt()
		du.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DeviceUpdate) check() error {
	if _, ok := du.mutation.OwnerID(); du.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Device.owner"`)
	}
	return nil
}

func (du *DeviceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   device.Table,
			Columns: device.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: device.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: device.FieldCreatedAt,
		})
	}
	if value, ok := du.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: device.FieldUpdatedAt,
		})
	}
	if value, ok := du.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldHash,
		})
	}
	if value, ok := du.mutation.DeviceID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldDeviceID,
		})
	}
	if value, ok := du.mutation.Platform(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldPlatform,
		})
	}
	if value, ok := du.mutation.LastSeen(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: device.FieldLastSeen,
		})
	}
	if value, ok := du.mutation.Language(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldLanguage,
		})
	}
	if du.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   device.OwnerTable,
			Columns: []string{device.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   device.OwnerTable,
			Columns: []string{device.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{device.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// DeviceUpdateOne is the builder for updating a single Device entity.
type DeviceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeviceMutation
}

// SetCreatedAt sets the "created_at" field.
func (duo *DeviceUpdateOne) SetCreatedAt(t time.Time) *DeviceUpdateOne {
	duo.mutation.SetCreatedAt(t)
	return duo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableCreatedAt(t *time.Time) *DeviceUpdateOne {
	if t != nil {
		duo.SetCreatedAt(*t)
	}
	return duo
}

// SetUpdatedAt sets the "updated_at" field.
func (duo *DeviceUpdateOne) SetUpdatedAt(t time.Time) *DeviceUpdateOne {
	duo.mutation.SetUpdatedAt(t)
	return duo
}

// SetHash sets the "hash" field.
func (duo *DeviceUpdateOne) SetHash(s string) *DeviceUpdateOne {
	duo.mutation.SetHash(s)
	return duo
}

// SetDeviceID sets the "device_id" field.
func (duo *DeviceUpdateOne) SetDeviceID(s string) *DeviceUpdateOne {
	duo.mutation.SetDeviceID(s)
	return duo
}

// SetPlatform sets the "platform" field.
func (duo *DeviceUpdateOne) SetPlatform(s string) *DeviceUpdateOne {
	duo.mutation.SetPlatform(s)
	return duo
}

// SetLastSeen sets the "last_seen" field.
func (duo *DeviceUpdateOne) SetLastSeen(t time.Time) *DeviceUpdateOne {
	duo.mutation.SetLastSeen(t)
	return duo
}

// SetLanguage sets the "language" field.
func (duo *DeviceUpdateOne) SetLanguage(s string) *DeviceUpdateOne {
	duo.mutation.SetLanguage(s)
	return duo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (duo *DeviceUpdateOne) SetOwnerID(id uuid.UUID) *DeviceUpdateOne {
	duo.mutation.SetOwnerID(id)
	return duo
}

// SetOwner sets the "owner" edge to the User entity.
func (duo *DeviceUpdateOne) SetOwner(u *User) *DeviceUpdateOne {
	return duo.SetOwnerID(u.ID)
}

// Mutation returns the DeviceMutation object of the builder.
func (duo *DeviceUpdateOne) Mutation() *DeviceMutation {
	return duo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (duo *DeviceUpdateOne) ClearOwner() *DeviceUpdateOne {
	duo.mutation.ClearOwner()
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DeviceUpdateOne) Select(field string, fields ...string) *DeviceUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Device entity.
func (duo *DeviceUpdateOne) Save(ctx context.Context) (*Device, error) {
	var (
		err  error
		node *Device
	)
	duo.defaults()
	if len(duo.hooks) == 0 {
		if err = duo.check(); err != nil {
			return nil, err
		}
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeviceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = duo.check(); err != nil {
				return nil, err
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			if duo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, duo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Device)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DeviceMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DeviceUpdateOne) SaveX(ctx context.Context) *Device {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DeviceUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DeviceUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DeviceUpdateOne) defaults() {
	if _, ok := duo.mutation.UpdatedAt(); !ok {
		v := device.UpdateDefaultUpdatedAt()
		duo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DeviceUpdateOne) check() error {
	if _, ok := duo.mutation.OwnerID(); duo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Device.owner"`)
	}
	return nil
}

func (duo *DeviceUpdateOne) sqlSave(ctx context.Context) (_node *Device, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   device.Table,
			Columns: device.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: device.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Device.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, device.FieldID)
		for _, f := range fields {
			if !device.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != device.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: device.FieldCreatedAt,
		})
	}
	if value, ok := duo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: device.FieldUpdatedAt,
		})
	}
	if value, ok := duo.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldHash,
		})
	}
	if value, ok := duo.mutation.DeviceID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldDeviceID,
		})
	}
	if value, ok := duo.mutation.Platform(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldPlatform,
		})
	}
	if value, ok := duo.mutation.LastSeen(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: device.FieldLastSeen,
		})
	}
	if value, ok := duo.mutation.Language(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: device.FieldLanguage,
		})
	}
	if duo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   device.OwnerTable,
			Columns: []string{device.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   device.OwnerTable,
			Columns: []string{device.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Device{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{device.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
