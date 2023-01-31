package domain

import (
	"crypto/rand"
	"math/big"
)

type File int

const (
	A File = iota
	B
	C
	D
	E
	F
	G
	H
)

type Rank int

const (
	_1 Rank = iota
	_2
	_3
	_4
	_5
	_6
	_7
	_8
)

type PieceType int

const NPieceTypes = 6

const (
	PAWN PieceType = iota
	ROOK
	KNIGHT
	BISHOP
	QUEEN
	KING
)

func (p PieceType) String() string {
	return [...]string{"PAWN", "ROOK", "KNIGHT", "BISHOP", "QUEEN", "KING"}[p]
}

type Color bool

const (
	WHITE Color = false
	BLACK Color = true
)

func GetRandomColor() Color {
	bigInt, _ := rand.Int(rand.Reader, big.NewInt(100))
	n := bigInt.Int64()

	if n%2 == 0 {
		return WHITE
	}
	return BLACK
}

func ColorToString(color Color) string {
	if color == WHITE {
		return "white"
	}
	return "black"
}

const (
	InitBoardFen  = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
	EmptyBoardFen = "8/8/8/8/8/8/8/8"
)
