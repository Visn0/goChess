package game

import (
	"encoding/json"
)

type Direction struct {
	x int
	y int
}

type IPiece interface {
	SetPiece(pieceType PieceType, color Color)
	GetPieceType() PieceType
	GetColor() Color
	GetValidDirections() []Direction
	String() string
	IsEnemy(piece IPiece) bool
}

type PieceBase struct {
	PieceType
	Color
	FirstMove bool
}

var PIECE_DIRECTION = map[PieceType][]Direction{
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

func NewPawn(color Color) *Pawn {
	return &Pawn{
		PieceBase: PieceBase{
			PieceType: PAWN,
			Color:     color,
			FirstMove: true,
		},
		EnPassantNeighbourPos: nil,
	}
}

func NewRook(color Color) *Rook {
	return &Rook{
		PieceBase: PieceBase{
			PieceType: ROOK,
			Color:     color,
			FirstMove: true,
		},
	}
}

func NewKnight(color Color) *Knight {
	return &Knight{
		PieceBase: PieceBase{
			PieceType: KNIGHT,
			Color:     color,
			FirstMove: true,
		},
	}
}

func NewBishop(color Color) *Bishop {
	return &Bishop{
		PieceBase: PieceBase{
			PieceType: BISHOP,
			Color:     color,
			FirstMove: true,
		},
	}
}

func NewQueen(color Color) *Queen {
	return &Queen{
		PieceBase: PieceBase{
			PieceType: QUEEN,
			Color:     color,
			FirstMove: true,
		},
	}
}

func NewKing(color Color) *King {
	return &King{
		PieceBase: PieceBase{
			PieceType: KING,
			Color:     color,
			FirstMove: true,
		},
	}
}

func (p *PieceBase) SetPiece(pieceType PieceType, color Color) {
	p.PieceType = pieceType
	p.Color = color
	p.FirstMove = true
}

func (p *Pawn) SetPiece(pieceType PieceType, color Color) {
	p.PieceBase.SetPiece(pieceType, color)
	p.EnPassantNeighbourPos = nil
}

func (p *PieceBase) GetPieceType() PieceType {
	return p.PieceType
}

func (p *PieceBase) String() string {
	j, _ := json.MarshalIndent(p, "", " ")
	return string(j)
}

func (p *PieceBase) GetValidDirections() []Direction {
	dir := PIECE_DIRECTION[p.PieceType]
	if p.PieceType == PAWN && p.Color == BLACK {
		return []Direction{{-1, 0}}
	}
	return dir
}

func (p *PieceBase) GetColor() Color {
	return p.Color
}

func (p *PieceBase) IsEnemy(piece IPiece) bool {
	return bool(p.GetColor()) != bool(piece.GetColor())
}
