package server

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/niakr1s/chess/events"
)

type WsClient struct {
	Username string
	from     chan events.Event
	to       chan events.Event
	conn     *websocket.Conn
	done     chan struct{}
}

func newWsClient(conn *websocket.Conn) *WsClient {
	res := &WsClient{
		from: make(chan events.Event),
		to:   make(chan events.Event),
		done: make(chan struct{}),
		conn: conn,
	}
	go res.wireConnToFromChannel()
	go res.wireToChannelToConn()
	return res
}

func (c *WsClient) From() <-chan events.Event {
	return c.from
}

func (c *WsClient) To() chan<- events.Event {
	return c.to
}

func (c *WsClient) Close() error {
	return c.conn.Close()
}

func (c *WsClient) Done() <-chan struct{} {
	return c.done
}

func (c *WsClient) waitForUsername(timeout time.Duration) bool {
	timeOutCh := time.After(timeout)
	for {
		select {
		case <-timeOutCh:
			err := c.conn.Close()
			if err != nil {
				log.Printf("couldn't close connection: %v", err)
			}
			return false
		case e := <-c.from:
			if nm, ok := e.(events.AuthUsernameEvent); ok {
				log.Printf("got username for client: %s", nm.Username)
				c.Username = nm.Username
				return true
			}
		}
	}
}

func (c *WsClient) wireToChannelToConn() {
	for e := range c.to {
		raw, err := events.EventToJson(e)
		if err != nil {
			log.Printf("couldn't convert event to json: %v", err)
			continue
		}
		err = c.conn.WriteMessage(websocket.TextMessage, raw)
		if err != nil {
			log.Printf("couldn't write json to conn: %v", err)
			continue
		}
	}
}

func (c *WsClient) wireConnToFromChannel() {
	defer func() {
		close(c.from)
		close(c.to)
		close(c.done)
	}()

	for {
		messageType, body, err := c.conn.ReadMessage()
		if err != nil || messageType == websocket.CloseMessage {
			log.Printf("connection closed")
			return
		}
		e, err := events.JsonToEvent(body, c.Username)
		if err != nil {
			log.Printf("unknown event")
			continue
		}
		log.Printf("got event %v", e)
		c.from <- e
	}
}
