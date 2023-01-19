package game

import "fmt"

type Game struct {
	Board     *Board
	LastMoves []*Move
}

func NewGame() *Game {
	return &Game{Board: NewBoard()}
}

func (g *Game) Move(m *Move) {
	p := g.Board.GetPiece(m.From.Rank, m.From.File)

	if p.GetName() == "pawn" {
		g.checkPawnMove(m, p.(*Pawn))
	} else if p.GetName() == "king" {
		g.checkCastleMove(m, p.(*King))
	} else {
		g.Board.SetPiece(m.To.Rank, m.To.File, p)
		g.Board.RemovePiece(m.From.Rank, m.From.File)
	}
}

func (g *Game) checkPawnMove(m *Move, pawn *Pawn) {

	dirRank := Rank(pawn.GetValidDirections()[0].x)
	if m.From.File == m.To.File {
		// Move forward
		if m.From.Rank+dirRank == m.To.Rank {
			// 1 space forward
			g.Board.SetPiece(m.To.Rank, m.To.File, pawn)
			g.Board.RemovePiece(m.From.Rank, m.From.File)
		} else if pawn.FirstMove && m.From.Rank+2*dirRank == m.To.Rank {
			// 2 spaces forward
			g.Board.SetPiece(m.To.Rank, m.To.File, pawn)
			g.Board.RemovePiece(m.From.Rank, m.From.File)
			// Set en passant left neighbour
			leftPos := &Position{m.To.Rank, m.To.File - 1}
			if leftPos.Valid() {
				leftPiece := g.Board.GetPiece(leftPos.Rank, leftPos.File)
				if leftPiece != nil && leftPiece.GetName() == "pawn" && leftPiece.GetColor() != pawn.GetColor() {
					leftPiece.(*Pawn).EnPassantNeighbourPos = &Position{m.To.Rank, m.To.File}
				}
			}
			// Set en passant right neighbour
			rightPos := &Position{m.To.Rank, m.To.File + 1}
			if rightPos.Valid() {
				rightPiece := g.Board.GetPiece(rightPos.Rank, rightPos.File)
				if rightPiece != nil && rightPiece.GetName() == "pawn" && rightPiece.GetColor() != pawn.GetColor() {
					rightPiece.(*Pawn).EnPassantNeighbourPos = &Position{m.To.Rank, m.To.File}
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
					if neighbourPiece != nil && neighbourPiece.GetName() == "pawn" && neighbourPiece.GetColor() != pawn.GetColor() {
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
		if rook == nil || rook.GetName() != "rook" || !rook.(*Rook).FirstMove {
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
