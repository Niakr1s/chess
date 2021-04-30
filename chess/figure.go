package chess

type Figure struct {
	Name  FigureName
	Color Color
	Moves int
}

func (f Figure) Glyph() string {
	return f.Name.Glyph(f.Color)
}

// first 6 bytes
const (
	FigureKingByteOffset = 1 << iota
	FigureQueenByteOffset
	FigureBishopByteOffset
	FigureKnightByteOffset
	FigureRookByteOffset
	FigurePawnByteOffset
)

func (f Figure) Encode() byte {
	var res byte = 0

	switch f.Name {
	case FigureKing:
		res |= FigureKingByteOffset
	case FigureQueen:
		res |= FigureQueenByteOffset
	case FigureBishop:
		res |= FigureBishopByteOffset
	case FigureKnight:
		res |= FigureKnightByteOffset
	case FigureRook:
		res |= FigureRookByteOffset
	case FigurePawn:
		res |= FigurePawnByteOffset
	default:
		return 0
	}

	// writing color to 7 byte
	res |= byte(f.Color) << 6

	var wasMoved byte = 0
	if f.Moves != 0 {
		wasMoved = 1
	}

	res |= wasMoved << 7

	return res
}
