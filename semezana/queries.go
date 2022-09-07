package semezana

import (
	"context"
	"fmt"

	"github.com/Muchogoc/semezana/ent"
	"github.com/google/uuid"
)

func _() {
	ctx := context.Background()

	// query a user's channel
	u, _ := globals.client.User.Get(ctx, uuid.New())
	u.QueryChannels().All(ctx)

	// query channel messages
	channel, _ := globals.client.Channel.Get(ctx, uuid.New())
	messages, _ := channel.QueryMessages().
		Limit(10).
		Offset(10).
		Order(ent.Desc()).
		All(ctx)

	for _, message := range messages {

		// should have message metadata
		recipientQuery := message.QueryRecipients()
		recipients, _ := recipientQuery.All(ctx)
		for _, recipient := range recipients {
			fmt.Println(recipient.Status.String())
			fmt.Println(recipient.Edges.User)
		}

		fmt.Println(message.ID)
		fmt.Println(message.Sequence)
		fmt.Println(message.CreatedAt)
		fmt.Println(message.Edges.Author)
		fmt.Println(message.Header)
		fmt.Println(message.Content)
		fmt.Println(message.Edges.Author)

	}
}
