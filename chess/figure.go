package chess

type Figure struct {
	Name  FigureName
	Color Color
	Moves int
}

func (f Figure) Glyph() string {
	return f.Name.Glyph(f.Color)
}
