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
	ValidPositions  []*Position
	FirstMove       bool
}

func NewPawn(isBlack bool) *Piece {
	if isBlack {
		return &Piece{
			Name:            "pawn",
			Black:           isBlack,
			ValidDirections: []Direction{{-1, 0}},
			FirstMove:       true,
		}
	}
	return &Piece{
		Name:            "pawn",
		Black:           isBlack,
		ValidDirections: []Direction{{1, 0}},
		FirstMove:       true,
	}
}

func NewRook(black bool) *Piece {
	return &Piece{
		Name:            "rook",
		Black:           black,
		ValidDirections: []Direction{{1, 0}, {0, 1}, {-1, 0}, {0, -1}},
		FirstMove:       true,
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
		FirstMove: true,
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
		FirstMove: true,
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
		FirstMove: true,
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
		FirstMove: true,
	}
}

func (p *Piece) String() string {
	j, _ := json.MarshalIndent(p, "", " ")
	return string(j)
}
