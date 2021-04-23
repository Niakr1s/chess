package main

import (
	"log"
	"net/http"

	"github.com/niakr1s/chess/server"
)

func main() {
	s := server.NewServer()

	go func() {
		for c := range s.Clients {
			for m := range c.From() {
				log.Println(m)
				if m, ok := m.(server.ChatMessageEvent); ok {
					c.To() <- m
				}
			}
		}
	}()

	addr := ":3333"
	log.Printf("starting server at %s", addr)
	log.Fatal(http.ListenAndServe(addr, s))
}
