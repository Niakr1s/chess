package chess

import (
	"fmt"
	"strings"
)

type ChessBoard struct {
	Board     [8][8]*Figure `json:"board"`
	graveYard []*Figure
}

func NewCleanChessBoard() *ChessBoard {
	return &ChessBoard{
		graveYard: make([]*Figure, 0),
	}
}

func NewChessBoard() *ChessBoard {
	board := NewCleanChessBoard()

	fillFigureRow(&board.Board[0], White)
	fillPawnRow(&board.Board[1], White)
	fillPawnRow(&board.Board[6], Black)
	fillFigureRow(&board.Board[7], Black)

	return board
}

func fillPawnRow(row *[8]*Figure, color Color) {
	for i := 0; i < len(row); i++ {
		row[i] = &Figure{
			Name:  FigurePawn,
			Color: color,
		}
	}
}

func fillFigureRow(row *[8]*Figure, color Color) {
	for _, col := range []int{0, 7} {
		row[col] = &Figure{
			Name:  FigureRook,
			Color: color,
		}
	}
	for _, col := range []int{1, 6} {
		row[col] = &Figure{
			Name:  FigureKnight,
			Color: color,
		}
	}
	for _, col := range []int{2, 5} {
		row[col] = &Figure{
			Name:  FigureBishop,
			Color: color,
		}
	}
	row[3] = &Figure{
		Name:  FigureQueen,
		Color: color,
	}
	row[4] = &Figure{
		Name:  FigureKing,
		Color: color,
	}
}

func (b ChessBoard) GetFigure(pos Pos) *Figure {
	return b.Board[pos.Row][pos.Col]
}

func (b *ChessBoard) MoveStr(from, to string) error {
	f, err := NewPosFromStr(from)
	if err != nil {
		return err
	}

	t, err := NewPosFromStr(to)
	if err != nil {
		return err
	}

	return b.Move(f, t)
}

func (b *ChessBoard) Move(from, to Pos) error {
	moves, err := b.GetPossibleMoves(from)
	if err != nil {
		return err
	}
	if !moves.Contains(to) {
		return fmt.Errorf("not possible move")
	}
	if toFig := b.GetFigure(to); toFig != nil {
		b.graveYard = append(b.graveYard, toFig)
	}
	removedFig := b.removeFigure(from)
	b.setFigure(removedFig, to)
	removedFig.Moves++
	return nil
}

func (b *ChessBoard) removeFigure(pos Pos) *Figure {
	f := b.Board[pos.Row][pos.Col]
	b.Board[pos.Row][pos.Col] = nil
	return f
}

func (b *ChessBoard) setFigure(f *Figure, pos Pos) {
	b.Board[pos.Row][pos.Col] = f
}

func (b ChessBoard) GetPossibleMoves(pos Pos) (Moves, error) {
	f := b.GetFigure(pos)
	if f == nil {
		return nil, fmt.Errorf("no figure")
	}
	res := Moves{}

	switch f.Name {
	case FigureKing:
		moves := pos.adjacent()
		for _, movePos := range moves {
			if movePosFigure := b.GetFigure(movePos); movePosFigure == nil || !SameColor(f.Color, movePosFigure.Color) {
				res = append(res, movePos)
			}
		}

	case FigureQueen:
		for rowOffset := -1; rowOffset <= 1; rowOffset++ {
			for colOffset := -1; colOffset <= 1; colOffset++ {
				line := pos.line(rowOffset, colOffset)
				res = append(res, b.filterLine(f.Color, line)...)
			}
		}

	case FigureBishop:
		for _, offsets := range [][]int{{-1, -1}, {1, -1}, {-1, 1}, {1, 1}} {
			line := pos.line(offsets[0], offsets[1])
			res = append(res, b.filterLine(f.Color, line)...)
		}

	case FigureKnight:
		moves := pos.knightMoves()
		res = append(res, b.filterAdjacent(f.Color, moves)...)

	case FigurePawn:
		rowOffset := rowOffsetForColor(f.Color)
		line := pos.line(rowOffset, 0)
		if len(line) > 0 {
			if b.GetFigure(line[0]) == nil {
				res = append(res, line[0])
				if f.Moves == 0 && len(line) > 1 && b.GetFigure(line[1]) == nil {
					res = append(res, line[1])
				}
			}
		}
		diags := Moves{
			{Col: pos.Col - 1, Row: pos.Row + rowOffset},
			{Col: pos.Col + 1, Row: pos.Row + rowOffset},
		}
		diags = filterInvalidMoves(diags)
		for _, diag := range diags {
			if diagFigure := b.GetFigure(diag); diagFigure != nil && !SameColor(diagFigure.Color, f.Color) {
				res = append(res, diag)
			}
		}

	case FigureRook:
		for _, offsets := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			line := pos.line(offsets[0], offsets[1])
			res = append(res, b.filterLine(f.Color, line)...)
		}

	default:
		return nil, fmt.Errorf("unknown figure")
	}

	return res, nil
}

// filterAdjacent returns slice, that contains nil figures or figures with another color.
func (b ChessBoard) filterAdjacent(color Color, line Moves) Moves {
	res := Moves{}
	for _, pos := range line {
		f := b.GetFigure(pos)
		if f == nil || !SameColor(f.Color, color) {
			res = append(res, pos)
		}
	}
	return res
}

// filterLine returns line, that contains nil figures and first figure with different color included (if exists).
func (b ChessBoard) filterLine(color Color, line Moves) Moves {
	res := Moves{}
	for _, pos := range line {
		f := b.GetFigure(pos)
		if f == nil {
			res = append(res, pos)
		} else {
			if !SameColor(color, f.Color) {
				res = append(res, pos)
			}
			break
		}
	}
	return res
}

const header = "   a b c d e f g h"

func (b ChessBoard) String() string {
	builder := strings.Builder{}
	builder.WriteString(header + "\n")

	for r := len(b.Board) - 1; r >= 0; r-- {
		row := b.Board[r]
		builder.WriteString(fmt.Sprintf("%d %s\n", r+1, b.rowToString(row)))
	}
	return builder.String()
}

func (b ChessBoard) rowToString(row [8]*Figure) string {
	glyphs := []string{}
	for _, f := range row {
		glyph := " "
		if f != nil {
			glyph = f.Glyph()
		}
		glyphs = append(glyphs, glyph)
	}
	return strings.Join(glyphs, " ")
}
