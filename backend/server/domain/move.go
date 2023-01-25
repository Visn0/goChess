package domain

import (
	"fmt"
)

type Position struct {
	Rank `json:"rank"`
	File `json:"file"`
}

func (p *Position) String() string {
	return string(rune('A'+p.File)) + string(rune('1'+p.Rank))
}

func (p *Position) Valid() bool {
	return p.Rank >= 0 && p.Rank < 8 && p.File >= 0 && p.File < 8
}

func (p *Position) Add(d Direction) {
	p.Rank = Rank(int(p.Rank) + d.x)
	p.File = File(int(p.File) + d.y)
}

type Move struct {
	From *Position `json:"from"`
	To   *Position `json:"to"`
}

func (g *Game) Move(m *Move, promoteTo *PieceType) {
	p := g.Board.GetPiece(m.From.Rank, m.From.File)
	if p == nil {
		fmt.Println("No piece at", m.From)
		panic("No piece at " + m.From.String())
	}
	if p.GetColor() != g.ColotToMove {
		fmt.Println("Wrong color to move")
		panic("Wrong color to move")
	}
	g.removeEnPassantStatesIfNotThisPiece(p)

	if p.GetPieceType() == PAWN {
		g.checkPawnMove(m, p.(*Pawn), promoteTo)
	} else if p.GetPieceType() == KING {
		g.checkCastleMove(m, p.(*King))
		if g.ColotToMove == WHITE {
			g.Board.whiteKingPos = m.To
		} else {
			g.Board.blackKingPos = m.To
		}
	} else {
		g.Board.SetPiece(m.To.Rank, m.To.File, p)
		g.Board.RemovePiece(m.From.Rank, m.From.File)
	}
	g.ColotToMove = Color(!bool(g.ColotToMove))
}

func (g *Game) removeEnPassantStatesIfNotThisPiece(p IPiece) {
	var currentPieceIdx int = -1
	for i, pawn := range g.EnPassantPieces {
		if pawn == p {
			currentPieceIdx = i
		}
	}
	if currentPieceIdx == -1 {
		for _, pawn := range g.EnPassantPieces {
			pawn.EnPassantNeighbourPos = nil
		}
		g.EnPassantPieces = make([]*Pawn, 0, 2)
	} else {
		g.EnPassantPieces = append(g.EnPassantPieces[:currentPieceIdx], g.EnPassantPieces[currentPieceIdx+1:]...)
	}
}

func (g *Game) checkPawnMove(m *Move, pawn *Pawn, promoteTo *PieceType) {
	dirRank := Rank(pawn.GetValidDirections()[0].x)
	if g.checkPawnPromotion(m, pawn, promoteTo) {
		return
	}
	if m.From.File == m.To.File {
		// Move forward
		if m.From.Rank+dirRank == m.To.Rank {
			// 1 space forward
			g.Board.SetPiece(m.To.Rank, m.To.File, pawn)
			g.Board.RemovePiece(m.From.Rank, m.From.File)
		} else if pawn.FirstMove && m.From.Rank+2*dirRank == m.To.Rank && g.Board.GetPiece(m.To.Rank, m.To.File) == nil {
			// 2 spaces forward
			g.Board.SetPiece(m.To.Rank, m.To.File, pawn)
			g.Board.RemovePiece(m.From.Rank, m.From.File)
			// Set en passant left neighbour
			leftPos := &Position{m.To.Rank, m.To.File - 1}
			if leftPos.Valid() {
				leftPiece := g.Board.GetPiece(leftPos.Rank, leftPos.File)
				if leftPiece != nil && leftPiece.GetPieceType() == PAWN && leftPiece.GetColor() != pawn.GetColor() {
					leftPiece.(*Pawn).EnPassantNeighbourPos = &Position{m.To.Rank, m.To.File}
					g.EnPassantPieces = append(g.EnPassantPieces, leftPiece.(*Pawn))
				}
			}
			// Set en passant right neighbour
			rightPos := &Position{m.To.Rank, m.To.File + 1}
			if rightPos.Valid() {
				rightPiece := g.Board.GetPiece(rightPos.Rank, rightPos.File)
				if rightPiece != nil && rightPiece.GetPieceType() == PAWN && rightPiece.GetColor() != pawn.GetColor() {
					rightPiece.(*Pawn).EnPassantNeighbourPos = &Position{m.To.Rank, m.To.File}
					g.EnPassantPieces = append(g.EnPassantPieces, rightPiece.(*Pawn))
				}
			}
		} else {
			fmt.Println(pawn.FirstMove, m.From.Rank+2*dirRank, m.To.Rank)
			panic("Invalid forward pawn move")
		}
	} else {
		// Move diagonally
		if m.To.Rank == m.From.Rank+dirRank {
			dstPiece := g.Board.GetPiece(m.To.Rank, m.To.File)
			if dstPiece != nil && dstPiece.GetColor() != pawn.GetColor() {
				// Capture piece
				g.Board.SetPiece(m.To.Rank, m.To.File, pawn)
				g.Board.RemovePiece(m.From.Rank, m.From.File)
			} else if dstPiece == nil {
				// Check if en passant
				if pawn.EnPassantNeighbourPos != nil {
					neighbourPiece := g.Board.GetPiece(pawn.EnPassantNeighbourPos.Rank, pawn.EnPassantNeighbourPos.File)
					if neighbourPiece != nil && neighbourPiece.GetPieceType() == PAWN && neighbourPiece.GetColor() != pawn.GetColor() {
						g.Board.SetPiece(m.To.Rank, m.To.File, pawn)
						g.Board.RemovePiece(m.From.Rank, m.From.File)
						g.Board.RemovePiece(pawn.EnPassantNeighbourPos.Rank, pawn.EnPassantNeighbourPos.File)
					}
				}
			} else {
				panic("Invalid capture pawn move")
			}
		} else {
			panic("Invalid diagonal pawn move")
		}
	}
	pawn.EnPassantNeighbourPos = nil
	pawn.FirstMove = false
}

func (g *Game) checkPawnPromotion(m *Move, pawn *Pawn, promoteTo *PieceType) bool {

	var piece IPiece
	if promoteTo != nil {
		fmt.Println("Promoting pawn", m.To, promoteTo)
		if m.To.Rank != 0 && m.To.Rank != 7 {
			panic("Invalid promotion rank")
		}
		switch *promoteTo {
		case QUEEN:
			piece = &Queen{}
			setNewPiece(piece, QUEEN, pawn.GetColor())
		case ROOK:
			piece = &Rook{}
			setNewPiece(piece, ROOK, pawn.GetColor())
		case BISHOP:
			piece = &Bishop{}
			setNewPiece(piece, BISHOP, pawn.GetColor())
		case KNIGHT:
			piece = &Knight{}
			setNewPiece(piece, KNIGHT, pawn.GetColor())
		default:
			panic(fmt.Sprintf("Invalid promotion piece type: %v", promoteTo))
		}
		g.Board.SetPiece(m.To.Rank, m.To.File, piece)
		g.Board.RemovePiece(m.From.Rank, m.From.File)
		return true
	}
	return false
}

func (g *Game) checkCastleMove(m *Move, king *King) {

	distance := m.To.File - m.From.File
	// Check if king is moving 2 spaces
	if king.FirstMove && distance == 2 || distance == -2 {
		var rook IPiece
		if distance == 2 {
			rook = g.Board.GetPiece(m.From.Rank, H)
		} else {
			rook = g.Board.GetPiece(m.From.Rank, A)
		}
		// Check if rook is in correct position
		if rook == nil || rook.GetPieceType() != ROOK || !rook.(*Rook).FirstMove {
			panic("Invalid castle")
		}
		// Move the correct rook
		if distance == 2 {
			g.Board.SetPiece(m.From.Rank, F, rook)
			g.Board.RemovePiece(m.From.Rank, H)
		} else {
			g.Board.SetPiece(m.From.Rank, D, rook)
			g.Board.RemovePiece(m.From.Rank, A)
		}
		rook.(*Rook).FirstMove = false
	}
	g.Board.SetPiece(m.To.Rank, m.To.File, king)
	g.Board.RemovePiece(m.From.Rank, m.From.File)
	king.FirstMove = false
}
