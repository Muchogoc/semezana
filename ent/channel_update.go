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
	"github.com/Muchogoc/semezana/ent/channel"
	"github.com/Muchogoc/semezana/ent/message"
	"github.com/Muchogoc/semezana/ent/predicate"
	"github.com/Muchogoc/semezana/ent/subscription"
	"github.com/Muchogoc/semezana/ent/user"
	"github.com/google/uuid"
)

// ChannelUpdate is the builder for updating Channel entities.
type ChannelUpdate struct {
	config
	hooks    []Hook
	mutation *ChannelMutation
}

// Where appends a list predicates to the ChannelUpdate builder.
func (cu *ChannelUpdate) Where(ps ...predicate.Channel) *ChannelUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *ChannelUpdate) SetCreatedAt(t time.Time) *ChannelUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *ChannelUpdate) SetNillableCreatedAt(t *time.Time) *ChannelUpdate {
	if t != nil {
		cu.SetCreatedAt(*t)
	}
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *ChannelUpdate) SetUpdatedAt(t time.Time) *ChannelUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetName sets the "name" field.
func (cu *ChannelUpdate) SetName(s string) *ChannelUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetDescription sets the "description" field.
func (cu *ChannelUpdate) SetDescription(s string) *ChannelUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetType sets the "type" field.
func (cu *ChannelUpdate) SetType(s string) *ChannelUpdate {
	cu.mutation.SetType(s)
	return cu
}

// SetState sets the "state" field.
func (cu *ChannelUpdate) SetState(s string) *ChannelUpdate {
	cu.mutation.SetState(s)
	return cu
}

// SetStateAt sets the "state_at" field.
func (cu *ChannelUpdate) SetStateAt(t time.Time) *ChannelUpdate {
	cu.mutation.SetStateAt(t)
	return cu
}

// SetSequence sets the "sequence" field.
func (cu *ChannelUpdate) SetSequence(i int) *ChannelUpdate {
	cu.mutation.ResetSequence()
	cu.mutation.SetSequence(i)
	return cu
}

// SetNillableSequence sets the "sequence" field if the given value is not nil.
func (cu *ChannelUpdate) SetNillableSequence(i *int) *ChannelUpdate {
	if i != nil {
		cu.SetSequence(*i)
	}
	return cu
}

// AddSequence adds i to the "sequence" field.
func (cu *ChannelUpdate) AddSequence(i int) *ChannelUpdate {
	cu.mutation.AddSequence(i)
	return cu
}

// SetTouched sets the "touched" field.
func (cu *ChannelUpdate) SetTouched(t time.Time) *ChannelUpdate {
	cu.mutation.SetTouched(t)
	return cu
}

// SetNillableTouched sets the "touched" field if the given value is not nil.
func (cu *ChannelUpdate) SetNillableTouched(t *time.Time) *ChannelUpdate {
	if t != nil {
		cu.SetTouched(*t)
	}
	return cu
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (cu *ChannelUpdate) AddMessageIDs(ids ...uuid.UUID) *ChannelUpdate {
	cu.mutation.AddMessageIDs(ids...)
	return cu
}

// AddMessages adds the "messages" edges to the Message entity.
func (cu *ChannelUpdate) AddMessages(m ...*Message) *ChannelUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cu.AddMessageIDs(ids...)
}

// AddMemberIDs adds the "members" edge to the User entity by IDs.
func (cu *ChannelUpdate) AddMemberIDs(ids ...uuid.UUID) *ChannelUpdate {
	cu.mutation.AddMemberIDs(ids...)
	return cu
}

// AddMembers adds the "members" edges to the User entity.
func (cu *ChannelUpdate) AddMembers(u ...*User) *ChannelUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.AddMemberIDs(ids...)
}

// AddSubscriptionIDs adds the "subscriptions" edge to the Subscription entity by IDs.
func (cu *ChannelUpdate) AddSubscriptionIDs(ids ...uuid.UUID) *ChannelUpdate {
	cu.mutation.AddSubscriptionIDs(ids...)
	return cu
}

// AddSubscriptions adds the "subscriptions" edges to the Subscription entity.
func (cu *ChannelUpdate) AddSubscriptions(s ...*Subscription) *ChannelUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.AddSubscriptionIDs(ids...)
}

// Mutation returns the ChannelMutation object of the builder.
func (cu *ChannelUpdate) Mutation() *ChannelMutation {
	return cu.mutation
}

// ClearMessages clears all "messages" edges to the Message entity.
func (cu *ChannelUpdate) ClearMessages() *ChannelUpdate {
	cu.mutation.ClearMessages()
	return cu
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (cu *ChannelUpdate) RemoveMessageIDs(ids ...uuid.UUID) *ChannelUpdate {
	cu.mutation.RemoveMessageIDs(ids...)
	return cu
}

// RemoveMessages removes "messages" edges to Message entities.
func (cu *ChannelUpdate) RemoveMessages(m ...*Message) *ChannelUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cu.RemoveMessageIDs(ids...)
}

// ClearMembers clears all "members" edges to the User entity.
func (cu *ChannelUpdate) ClearMembers() *ChannelUpdate {
	cu.mutation.ClearMembers()
	return cu
}

// RemoveMemberIDs removes the "members" edge to User entities by IDs.
func (cu *ChannelUpdate) RemoveMemberIDs(ids ...uuid.UUID) *ChannelUpdate {
	cu.mutation.RemoveMemberIDs(ids...)
	return cu
}

// RemoveMembers removes "members" edges to User entities.
func (cu *ChannelUpdate) RemoveMembers(u ...*User) *ChannelUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.RemoveMemberIDs(ids...)
}

// ClearSubscriptions clears all "subscriptions" edges to the Subscription entity.
func (cu *ChannelUpdate) ClearSubscriptions() *ChannelUpdate {
	cu.mutation.ClearSubscriptions()
	return cu
}

// RemoveSubscriptionIDs removes the "subscriptions" edge to Subscription entities by IDs.
func (cu *ChannelUpdate) RemoveSubscriptionIDs(ids ...uuid.UUID) *ChannelUpdate {
	cu.mutation.RemoveSubscriptionIDs(ids...)
	return cu
}

// RemoveSubscriptions removes "subscriptions" edges to Subscription entities.
func (cu *ChannelUpdate) RemoveSubscriptions(s ...*Subscription) *ChannelUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.RemoveSubscriptionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ChannelUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChannelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChannelUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChannelUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChannelUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ChannelUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := channel.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

func (cu *ChannelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   channel.Table,
			Columns: channel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: channel.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: channel.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: channel.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: channel.FieldName,
		})
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: channel.FieldDescription,
		})
	}
	if value, ok := cu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: channel.FieldType,
		})
	}
	if value, ok := cu.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: channel.FieldState,
		})
	}
	if value, ok := cu.mutation.StateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: channel.FieldStateAt,
		})
	}
	if value, ok := cu.mutation.Sequence(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: channel.FieldSequence,
		})
	}
	if value, ok := cu.mutation.AddedSequence(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: channel.FieldSequence,
		})
	}
	if value, ok := cu.mutation.Touched(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: channel.FieldTouched,
		})
	}
	if cu.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.MessagesTable,
			Columns: []string{channel.MessagesColumn},
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
	if nodes := cu.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !cu.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.MessagesTable,
			Columns: []string{channel.MessagesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.MessagesTable,
			Columns: []string{channel.MessagesColumn},
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
	if cu.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   channel.MembersTable,
			Columns: channel.MembersPrimaryKey,
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
	if nodes := cu.mutation.RemovedMembersIDs(); len(nodes) > 0 && !cu.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   channel.MembersTable,
			Columns: channel.MembersPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   channel.MembersTable,
			Columns: channel.MembersPrimaryKey,
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
	if cu.mutation.SubscriptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   channel.SubscriptionsTable,
			Columns: []string{channel.SubscriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: subscription.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedSubscriptionsIDs(); len(nodes) > 0 && !cu.mutation.SubscriptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   channel.SubscriptionsTable,
			Columns: []string{channel.SubscriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: subscription.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.SubscriptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   channel.SubscriptionsTable,
			Columns: []string{channel.SubscriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: subscription.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{channel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ChannelUpdateOne is the builder for updating a single Channel entity.
type ChannelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChannelMutation
}

// SetCreatedAt sets the "created_at" field.
func (cuo *ChannelUpdateOne) SetCreatedAt(t time.Time) *ChannelUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *ChannelUpdateOne) SetNillableCreatedAt(t *time.Time) *ChannelUpdateOne {
	if t != nil {
		cuo.SetCreatedAt(*t)
	}
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *ChannelUpdateOne) SetUpdatedAt(t time.Time) *ChannelUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetName sets the "name" field.
func (cuo *ChannelUpdateOne) SetName(s string) *ChannelUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *ChannelUpdateOne) SetDescription(s string) *ChannelUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetType sets the "type" field.
func (cuo *ChannelUpdateOne) SetType(s string) *ChannelUpdateOne {
	cuo.mutation.SetType(s)
	return cuo
}

// SetState sets the "state" field.
func (cuo *ChannelUpdateOne) SetState(s string) *ChannelUpdateOne {
	cuo.mutation.SetState(s)
	return cuo
}

// SetStateAt sets the "state_at" field.
func (cuo *ChannelUpdateOne) SetStateAt(t time.Time) *ChannelUpdateOne {
	cuo.mutation.SetStateAt(t)
	return cuo
}

// SetSequence sets the "sequence" field.
func (cuo *ChannelUpdateOne) SetSequence(i int) *ChannelUpdateOne {
	cuo.mutation.ResetSequence()
	cuo.mutation.SetSequence(i)
	return cuo
}

// SetNillableSequence sets the "sequence" field if the given value is not nil.
func (cuo *ChannelUpdateOne) SetNillableSequence(i *int) *ChannelUpdateOne {
	if i != nil {
		cuo.SetSequence(*i)
	}
	return cuo
}

// AddSequence adds i to the "sequence" field.
func (cuo *ChannelUpdateOne) AddSequence(i int) *ChannelUpdateOne {
	cuo.mutation.AddSequence(i)
	return cuo
}

// SetTouched sets the "touched" field.
func (cuo *ChannelUpdateOne) SetTouched(t time.Time) *ChannelUpdateOne {
	cuo.mutation.SetTouched(t)
	return cuo
}

// SetNillableTouched sets the "touched" field if the given value is not nil.
func (cuo *ChannelUpdateOne) SetNillableTouched(t *time.Time) *ChannelUpdateOne {
	if t != nil {
		cuo.SetTouched(*t)
	}
	return cuo
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (cuo *ChannelUpdateOne) AddMessageIDs(ids ...uuid.UUID) *ChannelUpdateOne {
	cuo.mutation.AddMessageIDs(ids...)
	return cuo
}

// AddMessages adds the "messages" edges to the Message entity.
func (cuo *ChannelUpdateOne) AddMessages(m ...*Message) *ChannelUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cuo.AddMessageIDs(ids...)
}

// AddMemberIDs adds the "members" edge to the User entity by IDs.
func (cuo *ChannelUpdateOne) AddMemberIDs(ids ...uuid.UUID) *ChannelUpdateOne {
	cuo.mutation.AddMemberIDs(ids...)
	return cuo
}

// AddMembers adds the "members" edges to the User entity.
func (cuo *ChannelUpdateOne) AddMembers(u ...*User) *ChannelUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.AddMemberIDs(ids...)
}

// AddSubscriptionIDs adds the "subscriptions" edge to the Subscription entity by IDs.
func (cuo *ChannelUpdateOne) AddSubscriptionIDs(ids ...uuid.UUID) *ChannelUpdateOne {
	cuo.mutation.AddSubscriptionIDs(ids...)
	return cuo
}

// AddSubscriptions adds the "subscriptions" edges to the Subscription entity.
func (cuo *ChannelUpdateOne) AddSubscriptions(s ...*Subscription) *ChannelUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.AddSubscriptionIDs(ids...)
}

// Mutation returns the ChannelMutation object of the builder.
func (cuo *ChannelUpdateOne) Mutation() *ChannelMutation {
	return cuo.mutation
}

// ClearMessages clears all "messages" edges to the Message entity.
func (cuo *ChannelUpdateOne) ClearMessages() *ChannelUpdateOne {
	cuo.mutation.ClearMessages()
	return cuo
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (cuo *ChannelUpdateOne) RemoveMessageIDs(ids ...uuid.UUID) *ChannelUpdateOne {
	cuo.mutation.RemoveMessageIDs(ids...)
	return cuo
}

// RemoveMessages removes "messages" edges to Message entities.
func (cuo *ChannelUpdateOne) RemoveMessages(m ...*Message) *ChannelUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cuo.RemoveMessageIDs(ids...)
}

// ClearMembers clears all "members" edges to the User entity.
func (cuo *ChannelUpdateOne) ClearMembers() *ChannelUpdateOne {
	cuo.mutation.ClearMembers()
	return cuo
}

// RemoveMemberIDs removes the "members" edge to User entities by IDs.
func (cuo *ChannelUpdateOne) RemoveMemberIDs(ids ...uuid.UUID) *ChannelUpdateOne {
	cuo.mutation.RemoveMemberIDs(ids...)
	return cuo
}

// RemoveMembers removes "members" edges to User entities.
func (cuo *ChannelUpdateOne) RemoveMembers(u ...*User) *ChannelUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.RemoveMemberIDs(ids...)
}

// ClearSubscriptions clears all "subscriptions" edges to the Subscription entity.
func (cuo *ChannelUpdateOne) ClearSubscriptions() *ChannelUpdateOne {
	cuo.mutation.ClearSubscriptions()
	return cuo
}

// RemoveSubscriptionIDs removes the "subscriptions" edge to Subscription entities by IDs.
func (cuo *ChannelUpdateOne) RemoveSubscriptionIDs(ids ...uuid.UUID) *ChannelUpdateOne {
	cuo.mutation.RemoveSubscriptionIDs(ids...)
	return cuo
}

// RemoveSubscriptions removes "subscriptions" edges to Subscription entities.
func (cuo *ChannelUpdateOne) RemoveSubscriptions(s ...*Subscription) *ChannelUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.RemoveSubscriptionIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ChannelUpdateOne) Select(field string, fields ...string) *ChannelUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Channel entity.
func (cuo *ChannelUpdateOne) Save(ctx context.Context) (*Channel, error) {
	var (
		err  error
		node *Channel
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChannelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Channel)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ChannelMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChannelUpdateOne) SaveX(ctx context.Context) *Channel {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ChannelUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChannelUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ChannelUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := channel.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

func (cuo *ChannelUpdateOne) sqlSave(ctx context.Context) (_node *Channel, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   channel.Table,
			Columns: channel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: channel.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Channel.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, channel.FieldID)
		for _, f := range fields {
			if !channel.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != channel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: channel.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: channel.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: channel.FieldName,
		})
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: channel.FieldDescription,
		})
	}
	if value, ok := cuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: channel.FieldType,
		})
	}
	if value, ok := cuo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: channel.FieldState,
		})
	}
	if value, ok := cuo.mutation.StateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: channel.FieldStateAt,
		})
	}
	if value, ok := cuo.mutation.Sequence(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: channel.FieldSequence,
		})
	}
	if value, ok := cuo.mutation.AddedSequence(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: channel.FieldSequence,
		})
	}
	if value, ok := cuo.mutation.Touched(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: channel.FieldTouched,
		})
	}
	if cuo.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.MessagesTable,
			Columns: []string{channel.MessagesColumn},
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
	if nodes := cuo.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !cuo.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.MessagesTable,
			Columns: []string{channel.MessagesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.MessagesTable,
			Columns: []string{channel.MessagesColumn},
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
	if cuo.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   channel.MembersTable,
			Columns: channel.MembersPrimaryKey,
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
	if nodes := cuo.mutation.RemovedMembersIDs(); len(nodes) > 0 && !cuo.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   channel.MembersTable,
			Columns: channel.MembersPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   channel.MembersTable,
			Columns: channel.MembersPrimaryKey,
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
	if cuo.mutation.SubscriptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   channel.SubscriptionsTable,
			Columns: []string{channel.SubscriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: subscription.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedSubscriptionsIDs(); len(nodes) > 0 && !cuo.mutation.SubscriptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   channel.SubscriptionsTable,
			Columns: []string{channel.SubscriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: subscription.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.SubscriptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   channel.SubscriptionsTable,
			Columns: []string{channel.SubscriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: subscription.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Channel{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{channel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
