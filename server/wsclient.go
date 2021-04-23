package server

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type WsClient struct {
	Username string
	from     chan interface{}
	to       chan interface{}
	conn     *websocket.Conn
}

func newWsClient(conn *websocket.Conn) *WsClient {
	res := &WsClient{
		from: make(chan interface{}),
		to:   make(chan interface{}),
		conn: conn,
	}
	go res.wireConnToFromChannel()
	go res.wireToChannelToConn()
	return res
}

func (c *WsClient) From() <-chan interface{} {
	return c.from
}

func (c *WsClient) To() chan<- interface{} {
	return c.to
}

func (c *WsClient) Close() error {
	return c.conn.Close()
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
			if nm, ok := e.(AuthUsernameEvent); ok {
				log.Printf("got username for client: %s", nm.Username)
				c.Username = nm.Username
				return true
			}
		}
	}
}

func (c *WsClient) wireToChannelToConn() {
	for e := range c.to {
		raw, err := c.eventToJson(e)
		if err != nil {
			log.Printf("couldn't convert event to json: %v", err)
			continue
		}
		log.Printf("%s", string(raw))
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
	}()

	for {
		messageType, body, err := c.conn.ReadMessage()
		if err != nil || messageType == websocket.CloseMessage {
			log.Printf("connection closed")
			return
		}
		e, err := c.jsonToEvent(body)
		if err != nil {
			log.Printf("unknown event")
			continue
		}
		c.from <- e
	}
}
