package chat

import (
	"sync"
	"time"

	"github.com/niakr1s/chess/events"
)

type Client struct {
	Username string
	Input    <-chan events.Event
	Output   chan<- events.Event
}

type Chat struct {
	mu      sync.Mutex
	clients map[string]Client
}

func NewChat() *Chat {
	return &Chat{
		clients: make(map[string]Client),
	}
}

func (c *Chat) AddClient(client Client) {
	c.RemoveClient(client.Username)
	c.mu.Lock()
	c.clients[client.Username] = client
	c.mu.Unlock()

	c.broadcast(events.ChatUserJoinEvent{Username: client.Username, Time: time.Now()})
	go func() {
		for e := range client.Input {
			switch t := e.(type) {
			case events.ChatMessageEvent:
				c.broadcast(t)
			}
		}
		c.RemoveClient(client.Username)
	}()
}

func (c *Chat) RemoveClient(username string) {
	c.mu.Lock()
	delete(c.clients, username)
	c.mu.Unlock()

	c.broadcast(events.ChatUserLeaveEvent{Username: username, Time: time.Now()})
}

func (c *Chat) broadcast(e events.Event) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, c := range c.clients {
		c := c
		go func() {
			c.Output <- e
		}()
	}
}
