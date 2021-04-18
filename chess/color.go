package chess

type Color int

// Players
const (
	White Color = 0
	Black Color = 1
)

func Opponent(color Color) Color {
	if color == White {
		return Black
	}
	return White
}

func SameColor(c1, c2 Color) bool {
	return int(c1) == int(c2)
}

func (c Color) String() string {
	if c == White {
		return "White"
	}
	return "Black"
}
