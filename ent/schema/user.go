package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.String("name"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now()).UpdateDefault(time.Now),
		field.String("state"),
		field.Time("state_at"),
		field.Time("last_seen").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("messages", Message.Type),
		edge.To("recipient_messages", Message.Type).
			Through("recipients", Recipient.Type),
		edge.To("devices", Device.Type),
		edge.To("channels", Channel.Type).
			Through("subscriptions", Subscription.Type),
	}
}
