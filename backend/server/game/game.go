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

func (g *Game) GetValidPositions(rank Rank, file File) []*Position {
	fmt.Println("Getting valid moves for", rank, file)
	p := g.Board.GetPiece(rank, file)
	if p == nil {
		fmt.Println("No piece at", rank, file)
		return nil
	}
	fmt.Printf("Piece at %d %d is %+v\n", rank, file, p.GetName())
	positions := g.GetPieceValidMovesHandler(p.GetName())(rank, file, p)
	return positions
}

func (g *Game) GetPieceValidMovesHandler(pieceType string) func(rank Rank, file File, p IPiece) []*Position {
	switch pieceType {
	case "pawn":
		return g.GetPawnValidMoves
	case "rook":
		return g.GetRookValidMoves
	case "knight":
		return g.GetKnightValidMoves
	case "bishop":
		return g.GetBishopValidMoves
	case "queen":
		return g.GetQueenValidMoves
	case "king":
		return g.GetKingValidMoves
	default:
		panic("Invalid piece type")
	}
}

func (g *Game) GetShortDistanceMoves(rank Rank, file File, p IPiece) []*Position {
	positions := []*Position{}
	for _, d := range p.GetValidDirections() {
		newPos := &Position{Rank: rank + Rank(d.x), File: file + File(d.y)}
		if newPos.Valid() {
			if g.Board.GetPiece(newPos.Rank, newPos.File) == nil {
				positions = append(positions, newPos)
			} else if g.Board.GetPiece(newPos.Rank, newPos.File).GetColor() != p.GetColor() &&
				p.GetName() != "pawn" {
				positions = append(positions, newPos)
			}
		}
	}
	return positions
}

func (g *Game) GetLongDistanceMoves(rank Rank, file File, p IPiece) []*Position {
	positions := []*Position{}
	for _, d := range p.GetValidDirections() {
		dCum := &Direction{0, 0}
		for {
			dCum.x += d.x
			dCum.y += d.y
			newPos := &Position{Rank: rank + Rank(dCum.x), File: file + File(dCum.y)}
			if newPos.Valid() {
				if g.Board.GetPiece(newPos.Rank, newPos.File) == nil {
					positions = append(positions, newPos)
				} else if g.Board.GetPiece(newPos.Rank, newPos.File).GetColor() != p.GetColor() {
					positions = append(positions, newPos)
					break
				} else {
					break
				}
			} else {
				break
			}
		}
	}
	return positions
}

func (g *Game) GetPawnValidMoves(rank Rank, file File, p IPiece) []*Position {
	positions := g.GetShortDistanceMoves(rank, file, p)
	pawn := p.(*Pawn)
	// Check if pawn can move two spaces forward
	if pawn.FirstMove {
		positions = append(positions, &Position{Rank: rank + Rank(p.GetValidDirections()[0].x*2), File: file})
	}
	// Check if pawn can move diagonally
	fmt.Println("Checking diagonals")
	pawnRankDir := Rank(p.GetValidDirections()[0].x)
	topPos := Position{Rank: rank + pawnRankDir, File: file}
	if topPos.Valid() {
		// Check if there is a piece in front left of the pawn
		topLeftPos := Position{Rank: topPos.Rank, File: topPos.File - 1}
		if topLeftPos.Valid() {
			topLeftPiece := g.Board.GetPiece(topLeftPos.Rank, topLeftPos.File)
			// Check if the piece is an enemy piece
			if topLeftPiece != nil && pawn.IsEnemy(topLeftPiece) {
				positions = append(positions, &topLeftPos)
			}
		}
		// Check if there is a piece in front right of the pawn
		topRightPos := Position{Rank: topPos.Rank, File: topPos.File + 1}
		if topRightPos.Valid() {
			topRightPiece := g.Board.GetPiece(topRightPos.Rank, topRightPos.File)
			// Check if the piece is an enemy piece
			if topRightPiece != nil && pawn.IsEnemy(topRightPiece) {
				positions = append(positions, &topRightPos)
			}
		}
		// Check if pawn has En Passant move
		if pawn.EnPassantNeighbourPos != nil {
			dstPos := &Position{Rank: rank + pawnRankDir, File: pawn.EnPassantNeighbourPos.File}
			if dstPos.Valid() {
				positions = append(positions, dstPos)
			}
		}
	}
	fmt.Println("Done checking diagonals")
	return positions
}

func (g *Game) GetRookValidMoves(rank Rank, file File, p IPiece) []*Position {
	positions := g.GetLongDistanceMoves(rank, file, p)
	return positions
}

func (g *Game) GetKnightValidMoves(rank Rank, file File, p IPiece) []*Position {
	positions := g.GetShortDistanceMoves(rank, file, p)
	return positions
}

func (g *Game) GetBishopValidMoves(rank Rank, file File, p IPiece) []*Position {
	positions := g.GetLongDistanceMoves(rank, file, p)
	return positions
}

func (g *Game) GetQueenValidMoves(rank Rank, file File, p IPiece) []*Position {
	positions := g.GetLongDistanceMoves(rank, file, p)
	return positions
}

func (g *Game) GetKingValidMoves(rank Rank, file File, p IPiece) []*Position {
	positions := g.GetShortDistanceMoves(rank, file, p)
	g.getKingCastlePositions(rank, file, p, &positions)
	return positions
}

func (g *Game) getKingCastlePositions(rank Rank, file File, p IPiece, positions *([]*Position)) {
	// Check if king can castle
	if p.(*King).FirstMove {
		// Check if king can castle kingside
		if g.Board.GetPiece(rank, file+1) == nil && g.Board.GetPiece(rank, file+2) == nil {
			rook := g.Board.GetPiece(rank, file+3)
			if rook != nil && rook.GetName() == "rook" && rook.(*Rook).FirstMove {
				*positions = append(*positions, &Position{Rank: rank, File: file + 2})
			}
		}
		// Check if king can castle queenside
		if g.Board.GetPiece(rank, file-1) == nil && g.Board.GetPiece(rank, file-2) == nil && g.Board.GetPiece(rank, file-3) == nil {
			rook := g.Board.GetPiece(rank, file-4)
			if rook != nil && rook.GetName() == "rook" && rook.(*Rook).FirstMove {
				*positions = append(*positions, &Position{Rank: rank, File: file - 2})
			}
		}
	}
}
