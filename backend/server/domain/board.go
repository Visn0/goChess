package domain

type Board struct {
	board        [][]IPiece
	whiteKingPos *Position
	blackKingPos *Position
}

func NewBoard() *Board {
	return NewBoardFromFEN(INIT_BOARD)
}

func (b *Board) Copy() *Board {
	newBoard := &Board{
		board: make([][]IPiece, 8),
		whiteKingPos: &Position{
			Rank: b.whiteKingPos.Rank,
			File: b.whiteKingPos.File,
		},
		blackKingPos: &Position{
			Rank: b.blackKingPos.Rank,
			File: b.blackKingPos.File,
		},
	}
	for i := range newBoard.board {
		newBoard.board[i] = make([]IPiece, 8)
	}
	for rank := range b.board {
		for file := range b.board[rank] {
			if b.board[rank][file] != nil {
				newBoard.board[rank][file] = b.board[rank][file].Copy()
			}
		}
	}
	return newBoard
}

// Fill board with pieces using FEN
func NewBoardFromFEN(fen string) *Board {
	b := &Board{
		board: make([][]IPiece, 8),
	}
	for i := range b.board {
		b.board[i] = make([]IPiece, 8)
	}
	rank, file := _8, A
	for _, piece := range fen {
		if piece == ' ' { // TODO: handle other parts of FEN
			break
		}
		if piece == '/' {
			rank--
			file = 0
			continue
		}
		if piece >= '1' && piece <= '8' {
			file += File(int(piece - '0'))
			continue
		}
		p := fenCharToPiece(string(piece))
		if p != nil {
			if p.GetPieceType() == KING {
				if p.GetColor() == WHITE {
					b.whiteKingPos = &Position{rank, file}
				} else {
					b.blackKingPos = &Position{rank, file}
				}
			}
			b.SetPiece(rank, file, p)
			file++
		}
	}
	return b
}

func (b *Board) GetPiece(rank Rank, file File) IPiece {
	return b.board[rank][file]
}

func (b *Board) SetPiece(rank Rank, file File, p IPiece) {
	b.board[rank][file] = p
}

func (b *Board) RemovePiece(rank Rank, file File) {
	b.board[rank][file] = nil
}

func fenCharToPiece(fen string) IPiece {
	var p IPiece

	switch fen {
	case "P":
		p = &Pawn{}
		setNewPiece(p, PAWN, false)
	case "p":
		p = &Pawn{}
		setNewPiece(p, PAWN, true)
	case "N":
		p = &Knight{}
		setNewPiece(p, KNIGHT, false)
	case "n":
		p = &Knight{}
		setNewPiece(p, KNIGHT, true)
	case "B":
		p = &Bishop{}
		setNewPiece(p, BISHOP, false)
	case "b":
		p = &Bishop{}
		setNewPiece(p, BISHOP, true)
	case "R":
		p = &Rook{}
		setNewPiece(p, ROOK, false)
	case "r":
		p = &Rook{}
		setNewPiece(p, ROOK, true)
	case "Q":
		p = &Queen{}
		setNewPiece(p, QUEEN, false)
	case "q":
		p = &Queen{}
		setNewPiece(p, QUEEN, true)
	case "K":
		p = &King{}
		setNewPiece(p, KING, false)
	case "k":
		p = &King{}
		setNewPiece(p, KING, true)
	default:
		return nil
	}
	return p
}

func (b *Board) GetKingPos(color Color) *Position {
	if color == WHITE {
		return b.whiteKingPos
	}
	return b.blackKingPos
}

func (b *Board) MovePiece(from, to *Position) {
	p := b.GetPiece(from.Rank, from.File)
	if p == nil {
		panic("No piece at from position")
	}
	b.RemovePiece(from.Rank, from.File)
	b.SetPiece(to.Rank, to.File, p)
}
