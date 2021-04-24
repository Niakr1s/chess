package lobby_test

import (
	"testing"
	"time"

	"github.com/niakr1s/chess/events"
	"github.com/niakr1s/chess/lobby"
	"github.com/stretchr/testify/assert"
)

func Test_Lobby(t *testing.T) {
	l := lobby.Lobby{}

	p1In := make(chan events.Event)
	p1Out := make(chan events.Event)
	p1 := lobby.Player{Input: p1In, Output: p1Out, Username: "player1"}

	p2In := make(chan events.Event)
	p2Out := make(chan events.Event)
	p2 := lobby.Player{Input: p2In, Output: p2Out, Username: "player2"}

	l.AddPlayer(p1)
	l.AddPlayer(p2)

	go func() {
		p1In <- events.ChessMoveEvent{From: "e2", To: "e4"}
		p1In <- events.ChessMoveEvent{From: "e4", To: "e5"}

		<-time.After(time.Millisecond * 100)
		p2In <- events.ChessMoveEvent{From: "c7", To: "c5"}
		p2In <- events.ChessMoveEvent{From: "c5", To: "c4"}
	}()

	// intial new turn event
	assert.IsType(t, events.ChessNewTurnEvent{}, <-p1Out)
	assert.IsType(t, events.ChessNewTurnEvent{}, <-p2Out)

	// after e2-e4
	assert.IsType(t, events.ChessNewTurnEvent{}, <-p1Out)
	assert.IsType(t, events.ChessNewTurnEvent{}, <-p2Out)

	// after e4-e5
	assert.IsType(t, events.ErrorEvent{}, <-p1Out)

	// after c7-c5
	assert.IsType(t, events.ChessNewTurnEvent{}, <-p1Out)
	assert.IsType(t, events.ChessNewTurnEvent{}, <-p2Out)

	// after c5-c4
	assert.IsType(t, events.ErrorEvent{}, <-p2Out)
}
