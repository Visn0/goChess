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
