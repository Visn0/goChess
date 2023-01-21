package domain

import (
	"math/rand"
	"time"
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

const N_PIECE_TYPES = 6

const (
	PAWN PieceType = iota
	ROOK
	KNIGHT
	BISHOP
	QUEEN
	KING
)

type Color bool

const (
	WHITE Color = false
	BLACK Color = true
)

func GetRandomColor() Color {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	n := r.Int()

	if n%2 == 0 {
		return WHITE
	}
	return BLACK
}

func ColorToString(color Color) string {
	if color == WHITE {
		return "white"
	} else {
		return "black"
	}
}

const (
	INIT_BOARD  = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
	EMPTY_BOARD = "8/8/8/8/8/8/8/8"
)
