package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type WsHandler struct {
	mux *http.ServeMux

	// Clients is channel, from wich new clients can be received.
	// Client already has Name.
	Clients chan *WsClient
}

func NewServer() *WsHandler {
	s := &WsHandler{
		mux:     http.NewServeMux(),
		Clients: make(chan *WsClient),
	}

	s.mux.HandleFunc("/ws", s.wsHandler)
	return s
}

func (h *WsHandler) wsHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := newWsClient(conn)
	if gotName := client.waitForUsername(time.Second * 5); !gotName {
		log.Printf("client didn't provided name")
		return
	}
	h.Clients <- client
}

func (h *WsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}
