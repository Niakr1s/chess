package chess

type FigureName string

const (
	FigureKing   FigureName = "king"   // король
	FigureRook   FigureName = "rook"   // ладья
	FigureBishop FigureName = "bishop" // слон
	FigureQueen  FigureName = "queen"  // ферзь
	FigureKnight FigureName = "knight" // конь
	FigurePawn   FigureName = "pawn"   // пешка
)

func (n FigureName) Glyph(color Color) string {
	if color == White {
		return n.glyphWhite()
	} else {
		return n.glyphBlack()
	}
}

func (n FigureName) glyphWhite() string {
	switch n {
	case FigureKing:
		return "♔"
	case FigureRook:
		return "♖"
	case FigureBishop:
		return "♗"
	case FigureQueen:
		return "♕"
	case FigureKnight:
		return "♘"
	case FigurePawn:
		return "♙"
	}
	return ""
}

func (n FigureName) glyphBlack() string {
	switch n {
	case FigureKing:
		return "♚"
	case FigureRook:
		return "♜"
	case FigureBishop:
		return "♝"
	case FigureQueen:
		return "♛"
	case FigureKnight:
		return "♞"
	case FigurePawn:
		return "♟"
	}
	return ""
}
