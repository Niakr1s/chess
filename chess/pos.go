package chess

import (
	"fmt"
	"strings"
)

type Pos struct {
	Col int
	Row int
}

func NewPos(col, row int) (Pos, error) {
	res := Pos{Col: col, Row: row}
	if !res.IsValid() {
		return Pos{}, fmt.Errorf("invalid input")
	}
	return res, nil
}

func NewPosFromStr(str string) (Pos, error) {
	str = strings.ToLower(str)
	runes := []rune(str)
	if len(runes) != 2 {
		return Pos{}, fmt.Errorf("invalid str")
	}
	return NewPos(int(runes[0]-'a'), int(runes[1]-'1'))
}

func (p Pos) IsValid() bool {
	return p.Col >= 0 && p.Col <= 7 && p.Row >= 0 && p.Row <= 7
}

func (p Pos) withOffsets(rowOffset, colOffset int) Pos {
	return Pos{Col: p.Col + colOffset, Row: p.Row + rowOffset}
}

func (p Pos) adjacent() Moves {
	res := Moves{}
	for col := -1; col <= 1; col++ {
		for row := -1; row <= 1; row++ {
			pos := p.withOffsets(row, col)
			res = append(res, pos)
		}
	}
	return filterInvalidMoves(res)
}

func (p Pos) line(rowOffset, colOffset int) Moves {
	res := Moves{}

	rowOffset = sign(rowOffset)
	colOffset = sign(colOffset)

	if rowOffset == 0 && colOffset == 0 {
		return res
	}

	for {
		p = p.withOffsets(rowOffset, colOffset)
		if !p.IsValid() {
			break
		}
		res = append(res, p)
	}
	return res
}

func (p Pos) knightMoves() Moves {
	res := Moves{}

	for _, offset := range [][2]int{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}} {
		res = append(res, p.withOffsets(offset[0], offset[1]))
	}

	return filterInvalidMoves(res)
}

func sign(i int) int {
	if i < 0 {
		return -1
	}
	if i > 0 {
		return 1
	}
	return 0
}

func rowOffsetForColor(color Color) int {
	if color == White {
		return +1
	}
	return -1
}
