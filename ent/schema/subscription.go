package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Subscription holds the schema definition for the Subscription entity.
type Subscription struct {
	ent.Schema
}

// Fields of the Subscription.
func (Subscription) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now()).UpdateDefault(time.Now),
		field.UUID("topic_id", uuid.New()),
		field.UUID("user_id", uuid.New()),
	}
}

// Edges of the Subscription.
func (Subscription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("subscriber", User.Type).Ref("subscriptions").Field("user_id").Unique().Required(),
		edge.From("topic", Topic.Type).Ref("subscriptions").Field("topic_id").Unique().Required(),
	}
}

// Indexes of the Subscription.
func (Subscription) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "topic_id").Unique(),
	}
}
