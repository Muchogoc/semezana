// Code generated by ent, DO NOT EDIT.

package message

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the message type in the database.
	Label = "message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTopicID holds the string denoting the topic_id field in the database.
	FieldTopicID = "topic_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldSequence holds the string denoting the sequence field in the database.
	FieldSequence = "sequence"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldHeader holds the string denoting the header field in the database.
	FieldHeader = "header"
	// EdgeTopic holds the string denoting the topic edge name in mutations.
	EdgeTopic = "topic"
	// EdgeSender holds the string denoting the sender edge name in mutations.
	EdgeSender = "sender"
	// Table holds the table name of the message in the database.
	Table = "messages"
	// TopicTable is the table that holds the topic relation/edge.
	TopicTable = "messages"
	// TopicInverseTable is the table name for the Topic entity.
	// It exists in this package in order to avoid circular dependency with the "topic" package.
	TopicInverseTable = "topics"
	// TopicColumn is the table column denoting the topic relation/edge.
	TopicColumn = "topic_id"
	// SenderTable is the table that holds the sender relation/edge.
	SenderTable = "messages"
	// SenderInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	SenderInverseTable = "users"
	// SenderColumn is the table column denoting the sender relation/edge.
	SenderColumn = "user_messages"
)

// Columns holds all SQL columns for message fields.
var Columns = []string{
	FieldID,
	FieldTopicID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldSequence,
	FieldContent,
	FieldHeader,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "messages"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_messages",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
