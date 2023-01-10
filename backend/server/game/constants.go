package game

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

const (
	INIT_BOARD  = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
	EMPTY_BOARD = "8/8/8/8/8/8/8/8"
)
