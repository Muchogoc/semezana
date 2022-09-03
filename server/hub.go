package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/Muchogoc/semezana/server/dto"
)

// Hub is the core structure which holds topics.
type Hub struct {

	// Topics must be indexed by name
	topics *sync.Map

	// Current number of loaded topics
	numTopics int

	// Channel for routing client-side messages, buffered at 4096
	routeClient chan *dto.ClientComMessage

	// Channel for routing server-generated messages, buffered at 4096
	routeServer chan *dto.ServerComMessage

	// subscribe session to topic, possibly creating a new topic, buffered at 256
	join chan *dto.ClientComMessage

	// Request to shutdown, unbuffered
	shutdown chan chan<- bool
}

func newHub() *Hub {
	h := &Hub{
		topics:      &sync.Map{},
		routeClient: make(chan *dto.ClientComMessage, 4096),
		routeServer: make(chan *dto.ServerComMessage, 4096),
		join:        make(chan *dto.ClientComMessage, 256),
		shutdown:    make(chan chan<- bool),
	}

	go h.run()

	return h
}

func (h *Hub) topicGet(name string) *Topic {
	if t, ok := h.topics.Load(name); ok {
		return t.(*Topic)
	}
	return nil
}

func (h *Hub) topicPut(name string, t *Topic) {
	h.numTopics++
	h.topics.Store(name, t)
}

func (h *Hub) topicDel(name string) {
	h.numTopics--
	h.topics.Delete(name)
}

func (h *Hub) run() {
	for {
		select {
		case join := <-h.join:
			fmt.Print(join)
		case msg := <-h.routeClient:
			if destination := h.topicGet(msg.Receiver); destination != nil {
				if destination.clientMsg != nil {
					select {
					case destination.clientMsg <- msg:
					default:
						log.Println("hub: topic's broadcast queue is full", destination.name)
					}
				}
			} else {
				log.Println("hub: invalid topic category for broadcast", destination.name)
			}
		case msg := <-h.routeServer:
			fmt.Print(msg)
		case hubdone := <-h.shutdown:
			fmt.Print(hubdone)
		}
	}
}
