package game

type Position struct {
	Rank `json:"rank"`
	File `json:"file"`
}

func (p *Position) String() string {
	return string(rune('A'+p.File)) + string(rune('1'+p.Rank))
}

func (p *Position) Valid() bool {
	return p.Rank >= 0 && p.Rank < 8 && p.File >= 0 && p.File < 8
}

func (p *Position) Add(d Direction) {
	p.Rank = Rank(int(p.Rank) + d.x)
	p.File = File(int(p.File) + d.y)
}

type Move struct {
	From *Position `json:"from"`
	To   *Position `json:"to"`
}
