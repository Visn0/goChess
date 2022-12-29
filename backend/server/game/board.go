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
		return &Piece{Name: "pawn", Black: false}
	case "p":
		return &Piece{Name: "pawn", Black: true}
	case "R":
		return &Piece{Name: "rook", Black: false}
	case "r":
		return &Piece{Name: "rook", Black: true}
	case "N":
		return &Piece{Name: "knight", Black: false}
	case "n":
		return &Piece{Name: "knight", Black: true}
	case "B":
		return &Piece{Name: "bishop", Black: false}
	case "b":
		return &Piece{Name: "bishop", Black: true}
	case "Q":
		return &Piece{Name: "queen", Black: false}
	case "q":
		return &Piece{Name: "queen", Black: true}
	case "K":
		return &Piece{Name: "king", Black: false}
	case "k":
		return &Piece{Name: "king", Black: true}
	default:
		return nil
	}
}
