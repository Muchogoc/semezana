package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Subscription holds the schema definition for the Subscription entity.
type Subscription struct {
	ent.Schema
}

// Fields of the Subscription.
func (Subscription) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).
			Default(uuid.New),
		field.UUID("channel_id", uuid.New()),
		field.UUID("user_id", uuid.New()),
		field.String("role").Comment("Authorizations in channel i.e admin, moderator etc"),
		field.String("status").Comment("Access to channel i.e ok, banned"),
		field.Bool("pinned"),
		field.Time("pinned_at"),
	}
}

// Edges of the Subscription.
func (Subscription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
		edge.To("channel", Channel.Type).
			Unique().
			Required().
			Field("channel_id"),
	}
}
