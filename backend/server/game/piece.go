package game

import "encoding/json"

type Direction struct {
	x int
	y int
}

type IPiece interface {
	GetName() string
	GetValidPositions() []*Position
	GetValidDirections() []Direction
	GetColor() bool
	String() string
}

type PieceBase struct {
	Name            string
	Black           bool
	ValidDirections []Direction
	ValidPositions  []*Position
	FirstMove       bool
}

type Pawn struct {
	PieceBase
	EnPassantNeighbourPos *Position
}

type Rook struct {
	PieceBase
}

type Knight struct {
	PieceBase
}

type Bishop struct {
	PieceBase
}

type Queen struct {
	PieceBase
}

type King struct {
	PieceBase
}

func NewPawn(isBlack bool) *Pawn {
	if isBlack {
		return &Pawn{
			PieceBase: PieceBase{
				Name:            "pawn",
				Black:           isBlack,
				ValidDirections: []Direction{{-1, 0}},
				FirstMove:       true,
			},
		}
	}
	return &Pawn{
		PieceBase: PieceBase{
			Name:            "pawn",
			Black:           isBlack,
			ValidDirections: []Direction{{1, 0}},
			FirstMove:       true,
		},
	}
}

func NewRook(black bool) *Rook {
	return &Rook{
		PieceBase: PieceBase{
			Name:            "rook",
			Black:           black,
			ValidDirections: []Direction{{1, 0}, {0, 1}, {-1, 0}, {0, -1}},
			FirstMove:       true,
		},
	}
}

func NewKnight(black bool) *Knight {
	return &Knight{
		PieceBase: PieceBase{
			Name:  "knight",
			Black: black,
			ValidDirections: []Direction{
				{1, -2}, {1, 2},
				{-1, -2}, {-1, 2},
				{2, -1}, {2, 1},
				{-2, -1}, {-2, 1},
			},
			FirstMove: true,
		},
	}
}

func NewBishop(black bool) *Bishop {
	return &Bishop{
		PieceBase: PieceBase{
			Name:  "bishop",
			Black: black,
			ValidDirections: []Direction{
				{-1, -1},
				{-1, 1},
				{1, 1},
				{1, -1},
			},
			FirstMove: true,
		},
	}
}

func NewQueen(black bool) *Queen {
	return &Queen{
		PieceBase: PieceBase{
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
		},
	}
}

func NewKing(black bool) *King {
	return &King{
		PieceBase: PieceBase{
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
		},
	}
}

func (p *PieceBase) GetName() string {
	return p.Name
}

func (p *PieceBase) String() string {
	j, _ := json.MarshalIndent(p, "", " ")
	return string(j)
}

func (p *PieceBase) GetValidPositions() []*Position {
	return p.ValidPositions
}

func (p *PieceBase) GetValidDirections() []Direction {
	return p.ValidDirections
}

func (p *PieceBase) GetColor() bool {
	return p.Black
}

func (p *PieceBase) IsEnemy(piece IPiece) bool {
	return p.Black != piece.GetColor()
}
