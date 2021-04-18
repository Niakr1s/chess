package chess_test

import (
	"fmt"
	"testing"

	"github.com/niakr1s/chess/chess"
	"github.com/stretchr/testify/assert"
)

func Test_NewPos(t *testing.T) {
	cases := []struct {
		col       int
		row       int
		expectErr bool
		expectPos chess.Pos
	}{
		{0, 0, false, chess.Pos{0, 0}},
		{2, 2, false, chess.Pos{2, 2}},
		{7, 7, false, chess.Pos{7, 7}},
		{7, 8, true, chess.Pos{}},
		{8, 8, true, chess.Pos{}},
		{-1, 0, true, chess.Pos{}},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			pos, err := chess.NewPos(c.col, c.row)
			if err != nil {
				assert.True(t, c.expectErr)
				return
			}
			assert.Equal(t, c.expectPos, pos)
		})
	}
}

func Test_NewPosFromStr(t *testing.T) {
	cases := []struct {
		posStr    string
		expectErr bool
		expectPos chess.Pos
	}{
		{"a1", false, chess.Pos{0, 0}},
		{"a8", false, chess.Pos{0, 7}},
		{"h1", false, chess.Pos{7, 0}},
		{"h8", false, chess.Pos{7, 7}},
		{"a9", true, chess.Pos{0, 0}},
		{"i8", true, chess.Pos{0, 0}},
		{"a11", true, chess.Pos{0, 0}},
		{"aa1", true, chess.Pos{0, 0}},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			pos, err := chess.NewPosFromStr(c.posStr)
			if err != nil {
				assert.True(t, c.expectErr)
				return
			}
			assert.Equal(t, c.expectPos, pos)
		})
	}
}
