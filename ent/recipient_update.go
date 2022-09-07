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
	"github.com/Muchogoc/semezana/ent/message"
	"github.com/Muchogoc/semezana/ent/predicate"
	"github.com/Muchogoc/semezana/ent/recipient"
	"github.com/Muchogoc/semezana/ent/user"
	"github.com/google/uuid"
)

// RecipientUpdate is the builder for updating Recipient entities.
type RecipientUpdate struct {
	config
	hooks    []Hook
	mutation *RecipientMutation
}

// Where appends a list predicates to the RecipientUpdate builder.
func (ru *RecipientUpdate) Where(ps ...predicate.Recipient) *RecipientUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetMessageID sets the "message_id" field.
func (ru *RecipientUpdate) SetMessageID(u uuid.UUID) *RecipientUpdate {
	ru.mutation.SetMessageID(u)
	return ru
}

// SetUserID sets the "user_id" field.
func (ru *RecipientUpdate) SetUserID(u uuid.UUID) *RecipientUpdate {
	ru.mutation.SetUserID(u)
	return ru
}

// SetStatus sets the "status" field.
func (ru *RecipientUpdate) SetStatus(r recipient.Status) *RecipientUpdate {
	ru.mutation.SetStatus(r)
	return ru
}

// SetDeliveredAt sets the "delivered_at" field.
func (ru *RecipientUpdate) SetDeliveredAt(t time.Time) *RecipientUpdate {
	ru.mutation.SetDeliveredAt(t)
	return ru
}

// SetNillableDeliveredAt sets the "delivered_at" field if the given value is not nil.
func (ru *RecipientUpdate) SetNillableDeliveredAt(t *time.Time) *RecipientUpdate {
	if t != nil {
		ru.SetDeliveredAt(*t)
	}
	return ru
}

// ClearDeliveredAt clears the value of the "delivered_at" field.
func (ru *RecipientUpdate) ClearDeliveredAt() *RecipientUpdate {
	ru.mutation.ClearDeliveredAt()
	return ru
}

// SetReadAt sets the "read_at" field.
func (ru *RecipientUpdate) SetReadAt(t time.Time) *RecipientUpdate {
	ru.mutation.SetReadAt(t)
	return ru
}

// SetNillableReadAt sets the "read_at" field if the given value is not nil.
func (ru *RecipientUpdate) SetNillableReadAt(t *time.Time) *RecipientUpdate {
	if t != nil {
		ru.SetReadAt(*t)
	}
	return ru
}

// ClearReadAt clears the value of the "read_at" field.
func (ru *RecipientUpdate) ClearReadAt() *RecipientUpdate {
	ru.mutation.ClearReadAt()
	return ru
}

// SetUser sets the "user" edge to the User entity.
func (ru *RecipientUpdate) SetUser(u *User) *RecipientUpdate {
	return ru.SetUserID(u.ID)
}

// SetMessage sets the "message" edge to the Message entity.
func (ru *RecipientUpdate) SetMessage(m *Message) *RecipientUpdate {
	return ru.SetMessageID(m.ID)
}

// Mutation returns the RecipientMutation object of the builder.
func (ru *RecipientUpdate) Mutation() *RecipientMutation {
	return ru.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ru *RecipientUpdate) ClearUser() *RecipientUpdate {
	ru.mutation.ClearUser()
	return ru
}

// ClearMessage clears the "message" edge to the Message entity.
func (ru *RecipientUpdate) ClearMessage() *RecipientUpdate {
	ru.mutation.ClearMessage()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RecipientUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecipientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RecipientUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RecipientUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RecipientUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RecipientUpdate) check() error {
	if v, ok := ru.mutation.Status(); ok {
		if err := recipient.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Recipient.status": %w`, err)}
		}
	}
	if _, ok := ru.mutation.UserID(); ru.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Recipient.user"`)
	}
	if _, ok := ru.mutation.MessageID(); ru.mutation.MessageCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Recipient.message"`)
	}
	return nil
}

func (ru *RecipientUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   recipient.Table,
			Columns: recipient.Columns,
			CompositeID: []*sqlgraph.FieldSpec{
				{
					Type:   field.TypeUUID,
					Column: recipient.FieldUserID,
				},
				{
					Type:   field.TypeUUID,
					Column: recipient.FieldMessageID,
				},
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: recipient.FieldStatus,
		})
	}
	if value, ok := ru.mutation.DeliveredAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: recipient.FieldDeliveredAt,
		})
	}
	if ru.mutation.DeliveredAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: recipient.FieldDeliveredAt,
		})
	}
	if value, ok := ru.mutation.ReadAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: recipient.FieldReadAt,
		})
	}
	if ru.mutation.ReadAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: recipient.FieldReadAt,
		})
	}
	if ru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   recipient.UserTable,
			Columns: []string{recipient.UserColumn},
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
	if nodes := ru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   recipient.UserTable,
			Columns: []string{recipient.UserColumn},
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
	if ru.mutation.MessageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   recipient.MessageTable,
			Columns: []string{recipient.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: message.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.MessageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   recipient.MessageTable,
			Columns: []string{recipient.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recipient.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// RecipientUpdateOne is the builder for updating a single Recipient entity.
type RecipientUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RecipientMutation
}

// SetMessageID sets the "message_id" field.
func (ruo *RecipientUpdateOne) SetMessageID(u uuid.UUID) *RecipientUpdateOne {
	ruo.mutation.SetMessageID(u)
	return ruo
}

// SetUserID sets the "user_id" field.
func (ruo *RecipientUpdateOne) SetUserID(u uuid.UUID) *RecipientUpdateOne {
	ruo.mutation.SetUserID(u)
	return ruo
}

// SetStatus sets the "status" field.
func (ruo *RecipientUpdateOne) SetStatus(r recipient.Status) *RecipientUpdateOne {
	ruo.mutation.SetStatus(r)
	return ruo
}

// SetDeliveredAt sets the "delivered_at" field.
func (ruo *RecipientUpdateOne) SetDeliveredAt(t time.Time) *RecipientUpdateOne {
	ruo.mutation.SetDeliveredAt(t)
	return ruo
}

// SetNillableDeliveredAt sets the "delivered_at" field if the given value is not nil.
func (ruo *RecipientUpdateOne) SetNillableDeliveredAt(t *time.Time) *RecipientUpdateOne {
	if t != nil {
		ruo.SetDeliveredAt(*t)
	}
	return ruo
}

// ClearDeliveredAt clears the value of the "delivered_at" field.
func (ruo *RecipientUpdateOne) ClearDeliveredAt() *RecipientUpdateOne {
	ruo.mutation.ClearDeliveredAt()
	return ruo
}

// SetReadAt sets the "read_at" field.
func (ruo *RecipientUpdateOne) SetReadAt(t time.Time) *RecipientUpdateOne {
	ruo.mutation.SetReadAt(t)
	return ruo
}

// SetNillableReadAt sets the "read_at" field if the given value is not nil.
func (ruo *RecipientUpdateOne) SetNillableReadAt(t *time.Time) *RecipientUpdateOne {
	if t != nil {
		ruo.SetReadAt(*t)
	}
	return ruo
}

// ClearReadAt clears the value of the "read_at" field.
func (ruo *RecipientUpdateOne) ClearReadAt() *RecipientUpdateOne {
	ruo.mutation.ClearReadAt()
	return ruo
}

// SetUser sets the "user" edge to the User entity.
func (ruo *RecipientUpdateOne) SetUser(u *User) *RecipientUpdateOne {
	return ruo.SetUserID(u.ID)
}

// SetMessage sets the "message" edge to the Message entity.
func (ruo *RecipientUpdateOne) SetMessage(m *Message) *RecipientUpdateOne {
	return ruo.SetMessageID(m.ID)
}

// Mutation returns the RecipientMutation object of the builder.
func (ruo *RecipientUpdateOne) Mutation() *RecipientMutation {
	return ruo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ruo *RecipientUpdateOne) ClearUser() *RecipientUpdateOne {
	ruo.mutation.ClearUser()
	return ruo
}

// ClearMessage clears the "message" edge to the Message entity.
func (ruo *RecipientUpdateOne) ClearMessage() *RecipientUpdateOne {
	ruo.mutation.ClearMessage()
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RecipientUpdateOne) Select(field string, fields ...string) *RecipientUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Recipient entity.
func (ruo *RecipientUpdateOne) Save(ctx context.Context) (*Recipient, error) {
	var (
		err  error
		node *Recipient
	)
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecipientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Recipient)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RecipientMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RecipientUpdateOne) SaveX(ctx context.Context) *Recipient {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RecipientUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RecipientUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RecipientUpdateOne) check() error {
	if v, ok := ruo.mutation.Status(); ok {
		if err := recipient.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Recipient.status": %w`, err)}
		}
	}
	if _, ok := ruo.mutation.UserID(); ruo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Recipient.user"`)
	}
	if _, ok := ruo.mutation.MessageID(); ruo.mutation.MessageCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Recipient.message"`)
	}
	return nil
}

func (ruo *RecipientUpdateOne) sqlSave(ctx context.Context) (_node *Recipient, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   recipient.Table,
			Columns: recipient.Columns,
			CompositeID: []*sqlgraph.FieldSpec{
				{
					Type:   field.TypeUUID,
					Column: recipient.FieldUserID,
				},
				{
					Type:   field.TypeUUID,
					Column: recipient.FieldMessageID,
				},
			},
		},
	}
	if id, ok := ruo.mutation.UserID(); !ok {
		return nil, &ValidationError{Name: "user_id", err: errors.New(`ent: missing "Recipient.user_id" for update`)}
	} else {
		_spec.Node.CompositeID[0].Value = id
	}
	if id, ok := ruo.mutation.MessageID(); !ok {
		return nil, &ValidationError{Name: "message_id", err: errors.New(`ent: missing "Recipient.message_id" for update`)}
	} else {
		_spec.Node.CompositeID[1].Value = id
	}
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, len(fields))
		for i, f := range fields {
			if !recipient.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			_spec.Node.Columns[i] = f
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: recipient.FieldStatus,
		})
	}
	if value, ok := ruo.mutation.DeliveredAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: recipient.FieldDeliveredAt,
		})
	}
	if ruo.mutation.DeliveredAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: recipient.FieldDeliveredAt,
		})
	}
	if value, ok := ruo.mutation.ReadAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: recipient.FieldReadAt,
		})
	}
	if ruo.mutation.ReadAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: recipient.FieldReadAt,
		})
	}
	if ruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   recipient.UserTable,
			Columns: []string{recipient.UserColumn},
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
	if nodes := ruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   recipient.UserTable,
			Columns: []string{recipient.UserColumn},
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
	if ruo.mutation.MessageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   recipient.MessageTable,
			Columns: []string{recipient.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: message.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.MessageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   recipient.MessageTable,
			Columns: []string{recipient.MessageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Recipient{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recipient.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
