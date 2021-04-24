package server

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/niakr1s/chess/events"
)

type WsHandler struct {
	mux *http.ServeMux

	// Clients is channel, from wich new clients can be received.
	// Client already has Name.
	Clients chan *WsClient

	mu          sync.Mutex
	clientNames map[string]struct{}
}

func (h *WsHandler) addClientName(username string) bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.clientNames[username]; ok {
		return false
	}
	h.clientNames[username] = struct{}{}
	return true
}

func (h *WsHandler) removeClientName(username string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	delete(h.clientNames, username)
}

func NewServer() *WsHandler {
	s := &WsHandler{
		mux:         http.NewServeMux(),
		Clients:     make(chan *WsClient),
		clientNames: make(map[string]struct{}),
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
	if added := h.addClientName(client.Username); !added || client.Username == "" {
		client.To() <- events.ErrorEvent{Message: "invalid username"}
		<-time.After(time.Millisecond * 100)
		client.Close()
		return
	}

	go func(username string) {
		<-client.Done()
		h.removeClientName(username)
	}(client.Username)

	h.Clients <- client
}

func (h *WsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}
