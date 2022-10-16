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
	// field.UUID("from", uuid.New()).Optional().Nillable().Comment("Who sent the message"),

	return []ent.Field{
		field.UUID("id", uuid.New()).
			Default(uuid.New),
		field.UUID("channel_id", uuid.New()),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now()).
			UpdateDefault(time.Now),
		field.Int("sequence"),
		field.JSON("header", MessageHeaders{}).
			Comment("The message header"),
		field.JSON("content", MessageContent{}).
			Comment("The message data"),
	}
}

type MessageContent struct {
	Text string `json:"text"`
}

type MessageHeaders struct{}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).
			Ref("messages").
			Unique().
			Required(),
		edge.From("channel", Channel.Type).
			Ref("messages").
			Field("channel_id").
			Unique().
			Required(),
		edge.From("message_recipients", User.Type).
			Ref("recipient_messages").
			Through("recipients", Recipient.Type),
	}
}

// Indexes of the Subscription.
func (Message) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("channel_id", "sequence").Unique(),
	}
}
