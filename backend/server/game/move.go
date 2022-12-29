package game

type Move struct {
	Rank `json:"rank"`
	File `json:"file"`
}

func (m *Move) Equals(other Move) bool {
	return m.Rank == other.Rank && m.File == other.File
}

func (m *Move) String() string {
	return string(rune('A'+m.File)) + string(rune('1'+m.Rank))
}

func (m *Move) Valid() bool {
	return m.Rank >= 0 && m.Rank < 8 && m.File >= 0 && m.File < 8
}

func (m *Move) Add(d Direction) {
	m.Rank = Rank(int(m.Rank) + d.x)
	m.File = File(int(m.File) + d.y)
}
