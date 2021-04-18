package chess

type Moves []Pos

func filterInvalidMoves(moves Moves) Moves {
	res := Moves{}
	for _, p := range moves {
		if p.IsValid() {
			res = append(res, p)
		}
	}
	return res
}

func (m Moves) Contains(pos Pos) bool {
	for _, p := range m {
		if p == pos {
			return true
		}
	}
	return false
}
