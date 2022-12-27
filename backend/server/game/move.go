package game

type Move struct {
	Rank int `json:"rank"`
	File int `json:"file"`
}

func (m *Move) Equals(other Move) bool {
	return m.Rank == other.Rank && m.File == other.File
}

func (m *Move) String() string {
	return string(rune('a'+m.File)) + string(rune('1'+m.Rank))
}

func (m *Move) Valid() bool {
	return m.Rank >= 0 && m.Rank < 8 && m.File >= 0 && m.File < 8
}
