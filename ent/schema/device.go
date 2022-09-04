package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Device holds the schema definition for the Device entity.
type Device struct {
	ent.Schema
}

// Fields of the Device.
func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).
			Default(uuid.New),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now()).UpdateDefault(time.Now),
		field.String("hash"),
		field.String("device_id"),
		field.String("platform"),
		field.Time("last_seen"),
		field.String("language"),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("devices").Unique().Required(),
	}
}

// Indexes of the Device.
func (Device) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("hash").Unique(),
	}
}
