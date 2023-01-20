package game

type Game struct {
	Board           *Board
	LastMoves       []*Move
	EnPassantPieces []*Pawn
}

func NewGame() *Game {
	return &Game{
		Board:           NewBoard(),
		EnPassantPieces: make([]*Pawn, 0, 2),
	}
}
