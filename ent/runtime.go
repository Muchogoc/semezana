// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/Muchogoc/semezana/ent/device"
	"github.com/Muchogoc/semezana/ent/message"
	"github.com/Muchogoc/semezana/ent/schema"
	"github.com/Muchogoc/semezana/ent/subscription"
	"github.com/Muchogoc/semezana/ent/topic"
	"github.com/Muchogoc/semezana/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	deviceFields := schema.Device{}.Fields()
	_ = deviceFields
	// deviceDescCreatedAt is the schema descriptor for created_at field.
	deviceDescCreatedAt := deviceFields[1].Descriptor()
	// device.DefaultCreatedAt holds the default value on creation for the created_at field.
	device.DefaultCreatedAt = deviceDescCreatedAt.Default.(func() time.Time)
	// deviceDescUpdatedAt is the schema descriptor for updated_at field.
	deviceDescUpdatedAt := deviceFields[2].Descriptor()
	// device.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	device.DefaultUpdatedAt = deviceDescUpdatedAt.Default.(time.Time)
	// device.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	device.UpdateDefaultUpdatedAt = deviceDescUpdatedAt.UpdateDefault.(func() time.Time)
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescCreatedAt is the schema descriptor for created_at field.
	messageDescCreatedAt := messageFields[2].Descriptor()
	// message.DefaultCreatedAt holds the default value on creation for the created_at field.
	message.DefaultCreatedAt = messageDescCreatedAt.Default.(func() time.Time)
	// messageDescUpdatedAt is the schema descriptor for updated_at field.
	messageDescUpdatedAt := messageFields[3].Descriptor()
	// message.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	message.DefaultUpdatedAt = messageDescUpdatedAt.Default.(time.Time)
	// message.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	message.UpdateDefaultUpdatedAt = messageDescUpdatedAt.UpdateDefault.(func() time.Time)
	// messageDescSequenceID is the schema descriptor for sequence_id field.
	messageDescSequenceID := messageFields[4].Descriptor()
	// message.SequenceIDValidator is a validator for the "sequence_id" field. It is called by the builders before save.
	message.SequenceIDValidator = messageDescSequenceID.Validators[0].(func(int) error)
	subscriptionFields := schema.Subscription{}.Fields()
	_ = subscriptionFields
	// subscriptionDescCreatedAt is the schema descriptor for created_at field.
	subscriptionDescCreatedAt := subscriptionFields[1].Descriptor()
	// subscription.DefaultCreatedAt holds the default value on creation for the created_at field.
	subscription.DefaultCreatedAt = subscriptionDescCreatedAt.Default.(func() time.Time)
	// subscriptionDescUpdatedAt is the schema descriptor for updated_at field.
	subscriptionDescUpdatedAt := subscriptionFields[2].Descriptor()
	// subscription.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	subscription.DefaultUpdatedAt = subscriptionDescUpdatedAt.Default.(time.Time)
	// subscription.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	subscription.UpdateDefaultUpdatedAt = subscriptionDescUpdatedAt.UpdateDefault.(func() time.Time)
	topicFields := schema.Topic{}.Fields()
	_ = topicFields
	// topicDescCreatedAt is the schema descriptor for created_at field.
	topicDescCreatedAt := topicFields[1].Descriptor()
	// topic.DefaultCreatedAt holds the default value on creation for the created_at field.
	topic.DefaultCreatedAt = topicDescCreatedAt.Default.(func() time.Time)
	// topicDescUpdatedAt is the schema descriptor for updated_at field.
	topicDescUpdatedAt := topicFields[2].Descriptor()
	// topic.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	topic.DefaultUpdatedAt = topicDescUpdatedAt.Default.(time.Time)
	// topic.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	topic.UpdateDefaultUpdatedAt = topicDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[1].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[2].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
