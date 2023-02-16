package domain

type Game struct {
	Board           *Board
	ColorToMove     Color
	LastMoves       []*Move
	EnPassantPieces []*Pawn
}

func NewGame() *Game {
	return &Game{
		Board:           NewBoard(),
		ColorToMove:     WHITE,
		EnPassantPieces: make([]*Pawn, 0, 2),
	}
}

func (g *Game) Copy() *Game {
	gCopy := &Game{
		Board:       g.Board.Copy(),
		ColorToMove: g.ColorToMove,
	}
	gCopy.LastMoves = make([]*Move, len(g.LastMoves))
	for i, m := range g.LastMoves {
		gCopy.LastMoves[i] = &Move{From: m.From, To: m.To}
	}
	gCopy.EnPassantPieces = make([]*Pawn, len(g.EnPassantPieces))
	for i, p := range g.EnPassantPieces {
		gCopy.EnPassantPieces[i] = &Pawn{
			PieceBase: PieceBase{
				PieceType:  p.PieceType,
				Color:      p.Color,
				FirstMove:  p.FirstMove,
				validMoves: []*Position{},
			},
			EnPassantNeighbourPos: &Position{p.EnPassantNeighbourPos.Rank, p.EnPassantNeighbourPos.File},
		}
	}
	return gCopy
}
