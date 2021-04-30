package chess_test

import (
	"fmt"
	"testing"

	"github.com/niakr1s/chess/chess"
	"github.com/stretchr/testify/assert"
)

func Test_FigureEncode(t *testing.T) {
	testCases := []struct {
		F        chess.Figure
		Expected byte
	}{
		{chess.Figure{Name: chess.FigureKing, Color: chess.White}, 0b00000001},
		{chess.Figure{Name: chess.FigureKing, Color: chess.Black}, 0b01000001},
		{chess.Figure{Name: chess.FigureKing, Color: chess.Black, Moves: 2}, 0b11000001},

		{chess.Figure{Name: chess.FigureKing, Color: chess.White}, 0b00000001},
		{chess.Figure{Name: chess.FigureQueen, Color: chess.White}, 0b00000010},
		{chess.Figure{Name: chess.FigureBishop, Color: chess.White}, 0b00000100},
		{chess.Figure{Name: chess.FigureKnight, Color: chess.White}, 0b00001000},
		{chess.Figure{Name: chess.FigureRook, Color: chess.White}, 0b00010000},
		{chess.Figure{Name: chess.FigurePawn, Color: chess.White}, 0b00100000},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test #%d", i), func(t *testing.T) {
			enc := tc.F.Encode()
			assert.Equal(t, tc.Expected, enc)
		})
	}
}
