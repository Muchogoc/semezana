package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Recipient holds the schema definition for the Recipient entity.
type Recipient struct {
	ent.Schema
}

// Annotations of the schema
func (Recipient) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "message_id"),
	}
}

// Fields of the Recipient.
func (Recipient) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("message_id", uuid.New()),
		field.UUID("user_id", uuid.New()),
		field.Enum("status").Values("DELIVERED", "UNREAD", "READ"),
		field.Time("status_at"),
	}
}

// Edges of the Recipient.
func (Recipient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
		edge.To("message", Message.Type).
			Unique().
			Required().
			Field("message_id"),
	}
}
