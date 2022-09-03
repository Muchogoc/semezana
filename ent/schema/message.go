package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.UUID("topic_id", uuid.New()),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now()).UpdateDefault(time.Now),
		field.Int("sequence_id").Positive(),
		field.UUID("from", uuid.New()).Optional().Nillable().Comment("Who sent the message"),
		field.JSON("content", map[string]interface{}{}).Comment("The message data"),
		field.JSON("header", map[string]interface{}{}).Comment("The message header"),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("topic", Topic.Type).Ref("messages").Field("topic_id").Unique().Required(),
		edge.From("sender", User.Type).Ref("messages").Unique(),
	}
}

// Indexes of the Subscription.
func (Message) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("topic_id", "sequence_id").Unique(),
	}
}
