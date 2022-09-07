// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChannelsColumns holds the columns for the "channels" table.
	ChannelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "state", Type: field.TypeString},
		{Name: "state_at", Type: field.TypeTime},
		{Name: "sequence", Type: field.TypeInt, Default: 0},
		{Name: "touched", Type: field.TypeTime},
	}
	// ChannelsTable holds the schema information for the "channels" table.
	ChannelsTable = &schema.Table{
		Name:       "channels",
		Columns:    ChannelsColumns,
		PrimaryKey: []*schema.Column{ChannelsColumns[0]},
	}
	// DevicesColumns holds the columns for the "devices" table.
	DevicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "hash", Type: field.TypeString},
		{Name: "device_id", Type: field.TypeString},
		{Name: "platform", Type: field.TypeString},
		{Name: "last_seen", Type: field.TypeTime},
		{Name: "language", Type: field.TypeString},
		{Name: "user_devices", Type: field.TypeUUID},
	}
	// DevicesTable holds the schema information for the "devices" table.
	DevicesTable = &schema.Table{
		Name:       "devices",
		Columns:    DevicesColumns,
		PrimaryKey: []*schema.Column{DevicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "devices_users_devices",
				Columns:    []*schema.Column{DevicesColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "device_hash",
				Unique:  true,
				Columns: []*schema.Column{DevicesColumns[3]},
			},
		},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "sequence", Type: field.TypeInt},
		{Name: "header", Type: field.TypeJSON},
		{Name: "content", Type: field.TypeJSON},
		{Name: "channel_id", Type: field.TypeUUID},
		{Name: "user_messages", Type: field.TypeUUID},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "messages_channels_messages",
				Columns:    []*schema.Column{MessagesColumns[6]},
				RefColumns: []*schema.Column{ChannelsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "messages_users_messages",
				Columns:    []*schema.Column{MessagesColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "message_channel_id_sequence",
				Unique:  true,
				Columns: []*schema.Column{MessagesColumns[6], MessagesColumns[3]},
			},
		},
	}
	// RecipientsColumns holds the columns for the "recipients" table.
	RecipientsColumns = []*schema.Column{
		{Name: "status", Type: field.TypeEnum, Enums: []string{"DELIVERED", "UNREAD", "READ"}},
		{Name: "delivered_at", Type: field.TypeTime, Nullable: true},
		{Name: "read_at", Type: field.TypeTime, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "message_id", Type: field.TypeUUID},
	}
	// RecipientsTable holds the schema information for the "recipients" table.
	RecipientsTable = &schema.Table{
		Name:       "recipients",
		Columns:    RecipientsColumns,
		PrimaryKey: []*schema.Column{RecipientsColumns[3], RecipientsColumns[4]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "recipients_users_user",
				Columns:    []*schema.Column{RecipientsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "recipients_messages_message",
				Columns:    []*schema.Column{RecipientsColumns[4]},
				RefColumns: []*schema.Column{MessagesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SubscriptionsColumns holds the columns for the "subscriptions" table.
	SubscriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "role", Type: field.TypeString},
		{Name: "status", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "channel_id", Type: field.TypeUUID},
	}
	// SubscriptionsTable holds the schema information for the "subscriptions" table.
	SubscriptionsTable = &schema.Table{
		Name:       "subscriptions",
		Columns:    SubscriptionsColumns,
		PrimaryKey: []*schema.Column{SubscriptionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subscriptions_users_user",
				Columns:    []*schema.Column{SubscriptionsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "subscriptions_channels_channel",
				Columns:    []*schema.Column{SubscriptionsColumns[4]},
				RefColumns: []*schema.Column{ChannelsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "subscription_user_id_channel_id",
				Unique:  true,
				Columns: []*schema.Column{SubscriptionsColumns[3], SubscriptionsColumns[4]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "state", Type: field.TypeString},
		{Name: "state_at", Type: field.TypeTime},
		{Name: "last_seen", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChannelsTable,
		DevicesTable,
		MessagesTable,
		RecipientsTable,
		SubscriptionsTable,
		UsersTable,
	}
)

func init() {
	DevicesTable.ForeignKeys[0].RefTable = UsersTable
	MessagesTable.ForeignKeys[0].RefTable = ChannelsTable
	MessagesTable.ForeignKeys[1].RefTable = UsersTable
	RecipientsTable.ForeignKeys[0].RefTable = UsersTable
	RecipientsTable.ForeignKeys[1].RefTable = MessagesTable
	SubscriptionsTable.ForeignKeys[0].RefTable = UsersTable
	SubscriptionsTable.ForeignKeys[1].RefTable = ChannelsTable
}
