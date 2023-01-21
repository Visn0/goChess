package domain

type Game struct {
	Board           *Board
	ColotToMove     Color
	LastMoves       []*Move
	EnPassantPieces []*Pawn
}

func NewGame() *Game {
	return &Game{
		Board:           NewBoard(),
		ColotToMove:     WHITE,
		EnPassantPieces: make([]*Pawn, 0, 2),
	}
}
