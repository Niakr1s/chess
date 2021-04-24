package chess

import (
	"fmt"
	"strings"
)

type Game struct {
	CurrentPlayer Color       `json:"currentPlayer"`
	Board         *ChessBoard `json:"chessboard"`
}

func NewGame() *Game {
	return &Game{
		CurrentPlayer: White,
		Board:         NewChessBoard(),
	}
}

func (g *Game) MoveStr(from, to string) error {
	fromPos, err := NewPosFromStr(from)
	if err != nil {
		return err
	}
	fromFig := g.Board.GetFigure(fromPos)
	if fromFig != nil && !SameColor(fromFig.Color, g.CurrentPlayer) {
		return fmt.Errorf("wrong color")
	}
	err = g.Board.MoveStr(from, to)
	if err != nil {
		return err
	}
	g.CurrentPlayer = Opponent(g.CurrentPlayer)
	return nil
}

func (g Game) String() string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("      %s turn\n", g.CurrentPlayer))
	b.WriteString(g.Board.String())
	b.WriteString("\n")
	return b.String()
}
