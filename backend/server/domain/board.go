package domain

type Board struct {
	board        [][]IPiece
	whiteKingPos *Position
	blackKingPos *Position
}

func NewBoard() *Board {
	return NewBoardFromFEN(InitBoardFen)
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
			pos := &Position{rank, file}
			if p.GetPieceType() == KING {
				if p.GetColor() == WHITE {
					b.whiteKingPos = pos
				} else {
					b.blackKingPos = pos
				}
			}
			b.SetPiece(pos, p)
			file++
		}
	}
	return b
}

func (b *Board) GetPiece(pos *Position) IPiece {
	return b.board[pos.Rank][pos.File]
}

func (b *Board) SetPiece(pos *Position, p IPiece) {
	b.board[pos.Rank][pos.File] = p
}

func (b *Board) RemovePiece(pos *Position) {
	b.board[pos.Rank][pos.File] = nil
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

func (b *Board) ExecuteMove(m *Move) {
	p := b.GetPiece(m.From)
	if p == nil {
		panic("No piece at from position")
	}
	b.RemovePiece(m.From)
	b.SetPiece(m.To, p)
}
