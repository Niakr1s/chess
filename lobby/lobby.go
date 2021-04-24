package lobby

import (
	"github.com/niakr1s/chess/chess"
	"github.com/niakr1s/chess/events"
)

type Player struct {
	Username string
	Input    <-chan events.Event
	Output   chan<- events.Event
	color    chess.Color
}

type Lobby struct {
	waitingClient *Player
}

func (l *Lobby) AddPlayer(player Player) error {
	if l.waitingClient != nil {
		l.startGame([2]Player{*l.waitingClient, player})
		l.waitingClient = nil
		return nil
	}
	l.waitingClient = &player
	return nil
}

type players [2]Player

func (p players) notifyNewTurn(game chess.Game) {
	for _, player := range p {
		player.Output <- events.ChessNewTurnEvent{Game: game}
	}
}

func (l *Lobby) startGame(players players) {
	players[0].color = chess.White
	players[1].color = chess.Black

	game := chess.NewGame()

	go players.notifyNewTurn(*game)
	for _, player := range players {
		player := player
		go func() {
			for e := range player.Input {
				switch t := e.(type) {
				case events.ChessMoveEvent:
					if game.CurrentPlayer != player.color {
						player.Output <- events.ErrorEvent{Message: "not your turn"}
						continue
					}
					err := game.MoveStr(t.From, t.To)
					if err != nil {
						player.Output <- events.ErrorEvent{Message: err.Error()}
						continue
					}
					players.notifyNewTurn(*game)
				}
			}
		}()
	}
}
