package chess

import (
	"fmt"
	"strings"
)

type Game struct {
	currentPlayer Color
	board         *ChessBoard
}

func NewGame() *Game {
	return &Game{
		currentPlayer: White,
		board:         NewChessBoard(),
	}
}

func (g Game) CurrentPlayer() Color {
	return g.currentPlayer
}

func (g *Game) MoveStr(from, to string) error {
	fromPos, err := NewPosFromStr(from)
	if err != nil {
		return err
	}
	fromFig := g.board.GetFigure(fromPos)
	if fromFig != nil && !SameColor(fromFig.Color, g.currentPlayer) {
		return fmt.Errorf("wrong color")
	}
	err = g.board.MoveStr(from, to)
	if err != nil {
		return err
	}
	g.currentPlayer = Opponent(g.currentPlayer)
	return nil
}

func (g Game) String() string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("      %s turn\n", g.currentPlayer))
	b.WriteString(g.board.String())
	b.WriteString("\n")
	return b.String()
}
