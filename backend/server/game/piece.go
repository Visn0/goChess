package game

import "encoding/json"

type Direction struct {
	x int
	y int
}

type Piece struct {
	Name            string
	Black           bool
	ValidDirections []Direction
	ValidMoves      []*Move
}

func NewPawn(black bool) *Piece {
	if black {
		return &Piece{
			Name:            "pawn",
			Black:           black,
			ValidDirections: []Direction{{-1, 0}},
		}
	}
	return &Piece{
		Name:            "pawn",
		Black:           black,
		ValidDirections: []Direction{{1, 0}},
	}
}

func NewRook(black bool) *Piece {
	return &Piece{
		Name:            "rook",
		Black:           black,
		ValidDirections: []Direction{{1, 0}, {0, 1}, {-1, 0}, {0, -1}},
	}
}

func NewKnight(black bool) *Piece {
	return &Piece{
		Name:  "knight",
		Black: black,
		ValidDirections: []Direction{
			{1, -2}, {1, 2},
			{-1, -2}, {-1, 2},
			{2, -1}, {2, 1},
			{-2, -1}, {-2, 1},
		},
	}
}

func NewBishop(black bool) *Piece {
	return &Piece{
		Name:  "bishop",
		Black: black,
		ValidDirections: []Direction{
			{-1, -1},
			{-1, 1},
			{1, 1},
			{1, -1},
		},
	}
}

func NewQueen(black bool) *Piece {
	return &Piece{
		Name:  "queen",
		Black: black,
		ValidDirections: []Direction{
			{-1, -1},
			{-1, 1},
			{1, 1},
			{1, -1},
			{1, 0},
			{0, 1},
			{-1, 0},
			{0, -1},
		},
	}
}

func NewKing(black bool) *Piece {
	return &Piece{
		Name:  "king",
		Black: black,
		ValidDirections: []Direction{
			{-1, -1},
			{-1, 1},
			{1, 1},
			{1, -1},
			{1, 0},
			{0, 1},
			{-1, 0},
			{0, -1},
		},
	}
}

func (p *Piece) String() string {
	j, _ := json.MarshalIndent(p, "", " ")
	return string(j)
}
