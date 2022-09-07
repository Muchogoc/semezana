// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Muchogoc/semezana/ent/channel"
	"github.com/google/uuid"
)

// Channel is the model entity for the Channel schema.
type Channel struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
	// StateAt holds the value of the "state_at" field.
	StateAt time.Time `json:"state_at,omitempty"`
	// sequential number of the latest message sent through the Channel
	Sequence int `json:"sequence,omitempty"`
	// timestamp of the last message sent to the Channel
	Touched time.Time `json:"touched,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ChannelQuery when eager-loading is set.
	Edges ChannelEdges `json:"edges"`
}

// ChannelEdges holds the relations/edges for other nodes in the graph.
type ChannelEdges struct {
	// Messages holds the value of the messages edge.
	Messages []*Message `json:"messages,omitempty"`
	// Members holds the value of the members edge.
	Members []*User `json:"members,omitempty"`
	// Subscriptions holds the value of the subscriptions edge.
	Subscriptions []*Subscription `json:"subscriptions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// MessagesOrErr returns the Messages value or an error if the edge
// was not loaded in eager-loading.
func (e ChannelEdges) MessagesOrErr() ([]*Message, error) {
	if e.loadedTypes[0] {
		return e.Messages, nil
	}
	return nil, &NotLoadedError{edge: "messages"}
}

// MembersOrErr returns the Members value or an error if the edge
// was not loaded in eager-loading.
func (e ChannelEdges) MembersOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Members, nil
	}
	return nil, &NotLoadedError{edge: "members"}
}

// SubscriptionsOrErr returns the Subscriptions value or an error if the edge
// was not loaded in eager-loading.
func (e ChannelEdges) SubscriptionsOrErr() ([]*Subscription, error) {
	if e.loadedTypes[2] {
		return e.Subscriptions, nil
	}
	return nil, &NotLoadedError{edge: "subscriptions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Channel) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case channel.FieldSequence:
			values[i] = new(sql.NullInt64)
		case channel.FieldName, channel.FieldType, channel.FieldState:
			values[i] = new(sql.NullString)
		case channel.FieldCreatedAt, channel.FieldUpdatedAt, channel.FieldStateAt, channel.FieldTouched:
			values[i] = new(sql.NullTime)
		case channel.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Channel", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Channel fields.
func (c *Channel) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case channel.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case channel.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case channel.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case channel.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case channel.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				c.Type = value.String
			}
		case channel.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				c.State = value.String
			}
		case channel.FieldStateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field state_at", values[i])
			} else if value.Valid {
				c.StateAt = value.Time
			}
		case channel.FieldSequence:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sequence", values[i])
			} else if value.Valid {
				c.Sequence = int(value.Int64)
			}
		case channel.FieldTouched:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field touched", values[i])
			} else if value.Valid {
				c.Touched = value.Time
			}
		}
	}
	return nil
}

// QueryMessages queries the "messages" edge of the Channel entity.
func (c *Channel) QueryMessages() *MessageQuery {
	return (&ChannelClient{config: c.config}).QueryMessages(c)
}

// QueryMembers queries the "members" edge of the Channel entity.
func (c *Channel) QueryMembers() *UserQuery {
	return (&ChannelClient{config: c.config}).QueryMembers(c)
}

// QuerySubscriptions queries the "subscriptions" edge of the Channel entity.
func (c *Channel) QuerySubscriptions() *SubscriptionQuery {
	return (&ChannelClient{config: c.config}).QuerySubscriptions(c)
}

// Update returns a builder for updating this Channel.
// Note that you need to call Channel.Unwrap() before calling this method if this Channel
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Channel) Update() *ChannelUpdateOne {
	return (&ChannelClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Channel entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Channel) Unwrap() *Channel {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Channel is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Channel) String() string {
	var builder strings.Builder
	builder.WriteString("Channel(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(c.Type)
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(c.State)
	builder.WriteString(", ")
	builder.WriteString("state_at=")
	builder.WriteString(c.StateAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("sequence=")
	builder.WriteString(fmt.Sprintf("%v", c.Sequence))
	builder.WriteString(", ")
	builder.WriteString("touched=")
	builder.WriteString(c.Touched.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Channels is a parsable slice of Channel.
type Channels []*Channel

func (c Channels) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}