package game

type Board struct {
	board [][]*Piece
}

func NewBoard() *Board {
	return NewBoardFromFEN(INIT_BOARD)
}

// Fill board with pieces using FEN
func NewBoardFromFEN(fen string) *Board {
	b := &Board{
		board: make([][]*Piece, 8),
	}
	for i := range b.board {
		b.board[i] = make([]*Piece, 8)
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
			b.SetPiece(rank, file, p)
			file++
		}
	}
	return b
}

func (b *Board) GetPiece(rank Rank, file File) *Piece {
	return b.board[rank][file]
}

func (b *Board) SetPiece(rank Rank, file File, p *Piece) {
	b.board[rank][file] = p
}

func (b *Board) RemovePiece(rank Rank, file File) {
	b.board[rank][file] = nil
}

func fenCharToPiece(fen string) *Piece {
	switch fen {
	case "P":
		return NewPawn(false)
	case "p":
		return NewPawn(true)
	case "N":
		return NewKnight(false)
	case "n":
		return NewKnight(true)
	case "B":
		return NewBishop(false)
	case "b":
		return NewBishop(true)
	case "R":
		return NewRook(false)
	case "r":
		return NewRook(true)
	case "Q":
		return NewQueen(false)
	case "q":
		return NewQueen(true)
	case "K":
		return NewKing(false)
	case "k":
		return NewKing(true)
	default:
		return nil
	}
}
