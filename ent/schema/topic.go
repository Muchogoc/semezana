package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Channel holds the schema definition for the Channel entity.
type Channel struct {
	ent.Schema
}

// Fields of the Channel.
func (Channel) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).
			Default(uuid.New),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now()).UpdateDefault(time.Now),
		field.String("name"),
		field.String("type"),
		field.String("state"),
		field.Time("state_at"),
		field.Int("sequence").
			Default(0).
			Comment("sequential number of the latest message sent through the Channel"),
		field.Time("touched").
			Default(time.Now()).
			Comment("timestamp of the last message sent to the Channel"),
	}
}

// Edges of the Channel.
func (Channel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("messages", Message.Type),
		edge.From("members", User.Type).Ref("channels").
			Through("subscriptions", Subscription.Type),
	}
}
