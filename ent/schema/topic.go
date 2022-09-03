package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Topic holds the schema definition for the Topic entity.
type Topic struct {
	ent.Schema
}

// Fields of the Topic.
func (Topic) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now()).UpdateDefault(time.Now),
		field.String("name"),
		field.String("type"),
		field.String("state"),
		field.Time("state_at"),
		field.Int("sequence_id").
			Comment("sequential ID of the latest message sent through the topic"),
		field.Time("touched").
			Comment("timestamp of the last message sent to the topic"),
		field.JSON("access", map[string]interface{}{}),
		field.JSON("public", map[string]interface{}{}),
		field.JSON("trusted", map[string]interface{}{}),
		field.JSON("tags", map[string]interface{}{}),
	}
}

// Edges of the Topic.
func (Topic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("messages", Message.Type),
		edge.To("subscriptions", Subscription.Type),
		edge.From("owner", User.Type).Ref("topics"),
	}
}
