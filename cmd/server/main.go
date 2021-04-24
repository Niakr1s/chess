package main

import (
	"log"
	"net/http"

	"github.com/niakr1s/chess/chat"
	"github.com/niakr1s/chess/events"
	"github.com/niakr1s/chess/lobby"
	"github.com/niakr1s/chess/server"
)

func main() {
	s := server.NewServer()

	Chat := chat.NewChat()
	Lobby := &lobby.Lobby{}

	go func() {
		for c := range s.Clients {
			chatCh, chessCh := events.SplitChannel(c.From())

			Chat.AddClient(chat.Client{Username: c.Username, Input: chatCh, Output: c.To()})
			Lobby.AddPlayer(lobby.Player{Username: c.Username, Input: chessCh, Output: c.To()})
		}
	}()

	addr := ":3333"
	log.Printf("starting server at %s", addr)
	log.Fatal(http.ListenAndServe(addr, s))
}
