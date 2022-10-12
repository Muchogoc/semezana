package semezana

import (
	"context"
	"fmt"
	"github.com/Muchogoc/semezana/ent"
	"github.com/Muchogoc/semezana/ent/message"
	"github.com/Muchogoc/semezana/ent/recipient"
	"github.com/Muchogoc/semezana/ent/user"
	"github.com/google/uuid"
)

func _() {
	ctx := context.Background()

	// query a user's channel
	usr, _ := globals.client.User.Get(ctx, uuid.New())
	usr.QueryChannels().All(ctx)

	// query messages for a user in a channel
	// - where the user is an audience
	// - where the user is an author
	channel, _ := globals.client.Channel.Get(ctx, uuid.New())
	messages, _ := channel.QueryMessages().
		Where(
			message.Or(
				message.HasAuthorWith(user.ID(usr.ID)),
				message.HasRecipientsWith(recipient.UserID(usr.ID)),
			),
		).
		Limit(10).
		Offset(10).
		Order(ent.Desc()).
		All(ctx)

	for _, msg := range messages {

		// should have message metadata
		recipientQuery := msg.QueryRecipients()
		recipients, _ := recipientQuery.All(ctx)
		for _, rcpnt := range recipients {
			fmt.Println(rcpnt.Status.String())
			fmt.Println(rcpnt.Edges.User)
		}

		fmt.Println(msg.ID)
		fmt.Println(msg.Sequence)
		fmt.Println(msg.CreatedAt)
		fmt.Println(msg.Edges.Author)
		fmt.Println(msg.Header)
		fmt.Println(msg.Content)
		fmt.Println(msg.Edges.Author)

	}
}
