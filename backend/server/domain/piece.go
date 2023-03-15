package domain

import (
	"fmt"
)

type Direction struct {
	x int
	y int
}

type IPiece interface {
	Copy() IPiece
	SetPiece(pieceType PieceType, color Color)
	GetPieceType() PieceType
	GetColor() Color
	GetValidDirections() []Direction
	String() string
	IsEnemy(piece IPiece) bool
	GetValidMoves() []*Position
	SetValidMoves([]*Position)
}

type PieceBase struct {
	PieceType
	Color
	FirstMove  bool
	validMoves []*Position
}

var PieceDirection = map[PieceType][]Direction{
	PAWN: {
		{1, 0},
	},
	ROOK: {
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	},
	KNIGHT: {
		{1, -2}, {1, 2},
		{-1, -2}, {-1, 2},
		{2, -1}, {2, 1},
		{-2, -1}, {-2, 1},
	},
	BISHOP: {
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	},
	QUEEN: {
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	},
	KING: {
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	},
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

func setNewPiece(piece IPiece, pieceType PieceType, color Color) {
	piece.SetPiece(pieceType, color)
	if pieceType == PAWN {
		piece.(*Pawn).EnPassantNeighbourPos = nil
	}
}

func (p *PieceBase) SetPiece(pieceType PieceType, color Color) {
	p.PieceType = pieceType
	p.Color = color
	p.FirstMove = true
}

func (p *PieceBase) Copy() IPiece {
	var piece IPiece
	switch p.PieceType {
	case PAWN:
		piece = &Pawn{}
	case ROOK:
		piece = &Rook{}
	case KNIGHT:
		piece = &Knight{}
	case BISHOP:
		piece = &Bishop{}
	case QUEEN:
		piece = &Queen{}
	case KING:
		piece = &King{}
	}
	piece.SetPiece(p.PieceType, p.Color)
	return piece
}

func (p *Pawn) SetPiece(pieceType PieceType, color Color) {
	p.PieceBase.SetPiece(pieceType, color)
	p.EnPassantNeighbourPos = nil
}

func (p *Pawn) Copy() IPiece {
	return &Pawn{
		PieceBase: PieceBase{
			PieceType: p.PieceType,
			Color:     p.Color,
			FirstMove: p.FirstMove,
		},
		EnPassantNeighbourPos: p.EnPassantNeighbourPos,
	}
}

func (p *PieceBase) GetPieceType() PieceType {
	return p.PieceType
}

func (p *PieceBase) String() string {
	return fmt.Sprintf("{Type: %s, Color: %s, FirstMove: %v}", p.PieceType, p.Color, p.FirstMove)
}

func (p *PieceBase) GetValidDirections() []Direction {
	return getPieceDirection(p.PieceType, p.Color)
}

func (p *PieceBase) GetColor() Color {
	return p.Color
}

func (p *PieceBase) IsEnemy(piece IPiece) bool {
	return bool(p.GetColor()) != bool(piece.GetColor())
}

func getPieceDirection(t PieceType, color Color) []Direction {
	dir := PieceDirection[t]
	if t == PAWN && color == BLACK {
		return []Direction{{-1, 0}}
	}
	return dir
}

func (p *PieceBase) GetValidMoves() []*Position {
	return p.validMoves
}

func (p *PieceBase) SetValidMoves(moves []*Position) {
	p.validMoves = moves
}
