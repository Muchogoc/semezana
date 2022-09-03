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
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now()).UpdateDefault(time.Now),
		field.String("state"),
		field.Time("state_at"),
		field.Time("last_seen").Optional(),
		field.JSON("access", map[string]interface{}{}).Optional(),
		field.JSON("public", map[string]interface{}{}).Optional(),
		field.JSON("trusted", map[string]interface{}{}).Optional(),
		field.JSON("tags", map[string]interface{}{}).Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subscriptions", Subscription.Type),
		edge.To("messages", Message.Type),
		edge.To("topics", Topic.Type),
		edge.To("devices", Device.Type),
	}
}
