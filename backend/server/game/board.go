package game

type Board struct {
	board [][]IPiece
}

func NewBoard() *Board {
	return NewBoardFromFEN(INIT_BOARD)
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
	var p IPiece = NewPiece[Pawn](PAWN, false)
	return p
	// switch fen {
	// case "P":
	// 	p := NewPiece[Pawn](PAWN, false)
	// 	return &p
	// case "p":
	// 	return NewPiece[Pawn](PAWN, true)
	// case "N":
	// 	return NewPiece[Knight](KNIGHT, false)
	// case "n":
	// 	return NewPiece[Knight](KNIGHT, true)
	// case "B":
	// 	return NewPiece[Bishop](BISHOP, false)
	// case "b":
	// 	return NewPiece[Bishop](BISHOP, true)
	// case "R":
	// 	return NewPiece[Rook](ROOK, false)
	// case "r":
	// 	return NewPiece[Rook](ROOK, true)
	// case "Q":
	// 	return NewPiece[Queen](QUEEN, false)
	// case "q":
	// 	return NewPiece[Queen](QUEEN, true)
	// case "K":
	// 	return NewPiece[King](KING, false)
	// case "k":
	// 	return NewPiece[King](KING, true)
	// default:
	// 	return nil
	// }
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
