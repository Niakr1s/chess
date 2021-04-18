package chess_test

import (
	"testing"

	"github.com/niakr1s/chess/chess"
	"github.com/stretchr/testify/assert"
)

func Test_NewChessBoard(t *testing.T) {
	c := chess.NewChessBoard()

	for _, row := range []int{0, 1, 6, 7} {
		for col := 0; col < 8; col++ {
			f := c.GetFigure(chess.Pos{Col: col, Row: row})
			assert.NotNil(t, f)
		}
	}

	for row := 2; row < 6; row++ {
		for col := 0; col < 8; col++ {
			f := c.GetFigure(chess.Pos{Col: col, Row: row})
			assert.Nil(t, f)
		}
	}
}

func Test_Move(t *testing.T) {
	type TestCase struct {
		From          string
		To            string
		ErrorExpected bool
	}

	doTest := func(cb *chess.ChessBoard, tc TestCase) {
		from, err := chess.NewPosFromStr(tc.From)
		assert.NoError(t, err)

		to, err := chess.NewPosFromStr(tc.To)
		assert.NoError(t, err)

		err = cb.Move(from, to)
		if err != nil {
			assert.True(t, tc.ErrorExpected)
			return
		}
		assert.False(t, tc.ErrorExpected)
	}

	t.Run(string(chess.FigureKing), func(t *testing.T) {
		cb := chess.NewChessBoard()
		doTest(cb, TestCase{From: "e1", To: "e2", ErrorExpected: true})
		doTest(cb, TestCase{From: "e2", To: "e4", ErrorExpected: false})
		doTest(cb, TestCase{From: "e1", To: "e2", ErrorExpected: false})
		doTest(cb, TestCase{From: "e2", To: "e3", ErrorExpected: false})

		doTest(cb, TestCase{From: "e3", To: "e4", ErrorExpected: true})
		doTest(cb, TestCase{From: "e3", To: "c3", ErrorExpected: true})
		doTest(cb, TestCase{From: "e3", To: "g3", ErrorExpected: true})
		doTest(cb, TestCase{From: "e3", To: "e1", ErrorExpected: true})

		doTest(cb, TestCase{From: "e3", To: "d3", ErrorExpected: false})
		doTest(cb, TestCase{From: "d3", To: "c4", ErrorExpected: false})
		doTest(cb, TestCase{From: "c4", To: "a6", ErrorExpected: true})
	})
	t.Run(string(chess.FigureQueen), func(t *testing.T) {
		cb := chess.NewChessBoard()
		doTest(cb, TestCase{From: "d1", To: "d2", ErrorExpected: true})
		doTest(cb, TestCase{From: "d2", To: "d4", ErrorExpected: false})
		doTest(cb, TestCase{From: "d1", To: "d2", ErrorExpected: false})
		doTest(cb, TestCase{From: "d2", To: "d3", ErrorExpected: false})

		doTest(cb, TestCase{From: "d3", To: "d4", ErrorExpected: true})
		doTest(cb, TestCase{From: "d3", To: "d5", ErrorExpected: true})
		doTest(cb, TestCase{From: "d3", To: "d6", ErrorExpected: true})

		doTest(cb, TestCase{From: "d3", To: "c3", ErrorExpected: false})
		doTest(cb, TestCase{From: "c3", To: "e5", ErrorExpected: true})
		doTest(cb, TestCase{From: "c3", To: "a5", ErrorExpected: false})
		doTest(cb, TestCase{From: "a5", To: "h5", ErrorExpected: false})
		doTest(cb, TestCase{From: "h5", To: "h8", ErrorExpected: true})
		doTest(cb, TestCase{From: "h5", To: "h7", ErrorExpected: false})
	})
	t.Run(string(chess.FigureBishop), func(t *testing.T) {
		cb := chess.NewChessBoard()
		doTest(cb, TestCase{From: "c1", To: "d2", ErrorExpected: true})
		doTest(cb, TestCase{From: "d2", To: "d4", ErrorExpected: false})
		doTest(cb, TestCase{From: "c1", To: "d2", ErrorExpected: false})

		doTest(cb, TestCase{From: "d2", To: "d3", ErrorExpected: true})
		doTest(cb, TestCase{From: "d2", To: "d4", ErrorExpected: true})
		doTest(cb, TestCase{From: "d2", To: "d5", ErrorExpected: true})
		doTest(cb, TestCase{From: "d2", To: "e5", ErrorExpected: true})
		doTest(cb, TestCase{From: "d2", To: "e1", ErrorExpected: true})

		doTest(cb, TestCase{From: "d2", To: "f4", ErrorExpected: false})
		doTest(cb, TestCase{From: "f4", To: "b8", ErrorExpected: true})
		doTest(cb, TestCase{From: "f4", To: "c7", ErrorExpected: false})
	})
	t.Run(string(chess.FigureKnight), func(t *testing.T) {
		cb := chess.NewChessBoard()
		doTest(cb, TestCase{From: "b1", To: "a2", ErrorExpected: true})
		doTest(cb, TestCase{From: "b1", To: "b2", ErrorExpected: true})
		doTest(cb, TestCase{From: "b1", To: "c2", ErrorExpected: true})
		doTest(cb, TestCase{From: "b1", To: "d2", ErrorExpected: true})

		doTest(cb, TestCase{From: "b1", To: "a3", ErrorExpected: false})
		doTest(cb, TestCase{From: "a3", To: "b4", ErrorExpected: true})
		doTest(cb, TestCase{From: "a3", To: "b5", ErrorExpected: false})
		doTest(cb, TestCase{From: "b5", To: "c7", ErrorExpected: false})
	})
	t.Run(string(chess.FigureRook), func(t *testing.T) {
		cb := chess.NewChessBoard()
		doTest(cb, TestCase{From: "a1", To: "a2", ErrorExpected: true})
		doTest(cb, TestCase{From: "a1", To: "b1", ErrorExpected: true})
		doTest(cb, TestCase{From: "a2", To: "a4", ErrorExpected: false})
		doTest(cb, TestCase{From: "a1", To: "a3", ErrorExpected: false})
		doTest(cb, TestCase{From: "a3", To: "g3", ErrorExpected: false})
		doTest(cb, TestCase{From: "g3", To: "f4", ErrorExpected: true})
		doTest(cb, TestCase{From: "g3", To: "g7", ErrorExpected: false})
	})
	t.Run(string(chess.FigurePawn), func(t *testing.T) {
		cb := chess.NewChessBoard()
		doTest(cb, TestCase{From: "a2", To: "a4", ErrorExpected: false})
		doTest(cb, TestCase{From: "a4", To: "a3", ErrorExpected: true})
		doTest(cb, TestCase{From: "a4", To: "b3", ErrorExpected: true})
		doTest(cb, TestCase{From: "a4", To: "b4", ErrorExpected: true})
		doTest(cb, TestCase{From: "a4", To: "b5", ErrorExpected: true})

		doTest(cb, TestCase{From: "a4", To: "a5", ErrorExpected: false})
		doTest(cb, TestCase{From: "a5", To: "a6", ErrorExpected: false})
		doTest(cb, TestCase{From: "a6", To: "a7", ErrorExpected: true})
		doTest(cb, TestCase{From: "a6", To: "b6", ErrorExpected: true})
		doTest(cb, TestCase{From: "a6", To: "b7", ErrorExpected: false})
		doTest(cb, TestCase{From: "b7", To: "a8", ErrorExpected: false})
	})
}

func Test_GetPossibleMoves(t *testing.T) {
	type TestCase struct {
		Pos           string
		ExpectedMoves []chess.Pos
	}

	doTest := func(cb *chess.ChessBoard, tc TestCase) {
		pos, err := chess.NewPosFromStr(tc.Pos)
		assert.NoError(t, err)

		moves, err := cb.GetPossibleMoves(pos)
		assert.NoError(t, err)

		assert.Len(t, moves, len(tc.ExpectedMoves))
		for _, m := range tc.ExpectedMoves {
			assert.Contains(t, moves, m)
		}
	}

	t.Run("King", func(t *testing.T) {
		cb := chess.NewChessBoard()
		doTest(cb, TestCase{Pos: "d1", ExpectedMoves: []chess.Pos{}})
	})
}
