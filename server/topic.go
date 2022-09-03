package main

import "github.com/Muchogoc/semezana/server/dto"

// Topic is an isolated communication channel
type Topic struct {
	// Еxpanded/unique name of the topic.
	name string

	// Channel for receiving client messages from sessions or other topics, buffered = 256.
	clientMsg chan *dto.ClientComMessage
	// Channel for receiving server messages generated on the server or received from other cluster nodes, buffered = 64.
	serverMsg chan *dto.ServerComMessage
}
