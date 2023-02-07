package domain

import (
	"fmt"
	"log"
)

// CalculateValidMoves calculates valid moves for all pieces of a given color
// and returns true if the player has valid moves
func (g *Game) CalculateValidMoves(color Color) bool {
	playerHasValidMoves := false
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			p := g.Board.GetPiece(Rank(rank), File(file))
			if p != nil && p.GetColor() == color {
				validMoves := g.GetValidMoves(Rank(rank), File(file))
				if len(validMoves) > 0 {
					p.SetValidMoves(validMoves)
					playerHasValidMoves = true
				}
			}
		}
	}
	return playerHasValidMoves
}

// GetValidMoves returns a list of valid moves for a piece at a given position
func (g *Game) GetValidMoves(rank Rank, file File) []*Position {
	fmt.Println("Getting valid moves for", rank, file)
	p := g.Board.GetPiece(rank, file)
	if p == nil {
		fmt.Println("No piece at", rank, file)
		return nil
	}
	if p.GetColor() != g.ColotToMove {
		fmt.Println("Wrong color to move")
		return nil
	}
	fmt.Printf("Piece at %d %d is %+v\n", rank, file, p.GetPieceType())
	positions := g.getPieceValidMovesHandler(p.GetPieceType())(rank, file, p)
	positions = g.filterMovesIfCheck(rank, file, positions)
	precalculatedPositions := p.GetValidMoves()
	if len(precalculatedPositions) > 0 && len(positions) != len(precalculatedPositions) {
		log.Println("WARNING: Precalculated moves and calculated moves do not match")
	}
	return positions
}

func (g *Game) getPieceValidMovesHandler(pieceType PieceType) func(rank Rank, file File, p IPiece) []*Position {
	switch pieceType {
	case PAWN:
		return g.getPawnValidMoves
	case ROOK:
		return g.getRookValidMoves
	case KNIGHT:
		return g.getKnightValidMoves
	case BISHOP:
		return g.getBishopValidMoves
	case QUEEN:
		return g.getQueenValidMoves
	case KING:
		return g.getKingValidMoves
	default:
		panic("Invalid piece type")
	}
}

func (g *Game) getShortDistanceMoves(rank Rank, file File, p IPiece) []*Position {
	positions := []*Position{}
	for _, d := range p.GetValidDirections() {
		newPos := &Position{Rank: rank + Rank(d.x), File: file + File(d.y)}
		if newPos.Valid() {
			if g.Board.GetPiece(newPos.Rank, newPos.File) == nil {
				positions = append(positions, newPos)
			} else if g.Board.GetPiece(newPos.Rank, newPos.File).GetColor() != p.GetColor() &&
				p.GetPieceType() != PAWN {
				// Pawns can not take pieces in front of them
				positions = append(positions, newPos)
			}
		}
	}
	return positions
}

func (g *Game) getLongDistanceMoves(rank Rank, file File, p IPiece) []*Position {
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

func (g *Game) getPawnValidMoves(rank Rank, file File, p IPiece) []*Position {
	positions := g.getShortDistanceMoves(rank, file, p)
	pawn := p.(*Pawn)
	// Check if pawn can move two spaces forward
	if pawn.FirstMove {
		dstPos := &Position{Rank: rank + Rank(p.GetValidDirections()[0].x*2), File: file}
		if dstPos.Valid() && g.Board.GetPiece(dstPos.Rank, dstPos.File) == nil {
			positions = append(positions, dstPos)
		}
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

func (g *Game) getRookValidMoves(rank Rank, file File, p IPiece) []*Position {
	positions := g.getLongDistanceMoves(rank, file, p)
	return positions
}

func (g *Game) getKnightValidMoves(rank Rank, file File, p IPiece) []*Position {
	positions := g.getShortDistanceMoves(rank, file, p)
	return positions
}

func (g *Game) getBishopValidMoves(rank Rank, file File, p IPiece) []*Position {
	positions := g.getLongDistanceMoves(rank, file, p)
	return positions
}

func (g *Game) getQueenValidMoves(rank Rank, file File, p IPiece) []*Position {
	positions := g.getLongDistanceMoves(rank, file, p)
	return positions
}

func (g *Game) getKingValidMoves(rank Rank, file File, p IPiece) []*Position {
	positions := g.getShortDistanceMoves(rank, file, p)
	g.getKingCastlePositions(rank, file, p, &positions)
	return positions
}

func (g *Game) getKingCastlePositions(rank Rank, file File, p IPiece, positions *[]*Position) {
	// Check if king can castle
	if p.(*King).FirstMove {
		// Check if king can castle kingside
		if g.Board.GetPiece(rank, file+1) == nil && g.Board.GetPiece(rank, file+2) == nil {
			rook := g.Board.GetPiece(rank, file+3)
			if rook != nil && rook.GetPieceType() == ROOK && rook.(*Rook).FirstMove {
				*positions = append(*positions, &Position{Rank: rank, File: file + 2})
			}
		}
		// Check if king can castle queenside
		if g.Board.GetPiece(rank, file-1) == nil &&
			g.Board.GetPiece(rank, file-2) == nil &&
			g.Board.GetPiece(rank, file-3) == nil {
			rook := g.Board.GetPiece(rank, file-4)
			if rook != nil && rook.GetPieceType() == ROOK && rook.(*Rook).FirstMove {
				*positions = append(*positions, &Position{Rank: rank, File: file - 2})
			}
		}
	}
}

// Remove all positions that will put the king in check
// TODO: has some bugs during castling and en passant
// TODO: not the most efficient way to do this (implement it with a copy of the board)
func (g *Game) filterMovesIfCheck(rank Rank, file File, positions []*Position) []*Position {
	filteredPositions := []*Position{}
	for _, pos := range positions {
		if !g.isCheckAfterMove(&Position{Rank: rank, File: file}, pos) {
			filteredPositions = append(filteredPositions, pos)
		}
	}
	return filteredPositions
}

// Check if the move will put the king in check
func (g *Game) isCheckAfterMove(from, to *Position) bool {
	boardCopy := g.Board.Copy()
	boardCopy.MovePiece(from, to) //TODO: doesn't work perfectly for castling or en passant yet
	// Check if the king is in check
	enemyKingPos := g.Board.GetKingPos(g.ColotToMove)
	// If the king is moved, update the king position
	if enemyKingPos.Rank == from.Rank && enemyKingPos.File == from.File {
		enemyKingPos = to
	}

	if enemyKingPos == nil {
		log.Println("##> King not found: ", g.ColotToMove)
		panic("King not found")
	}
	if boardCopy.PositionIsUnderAttack(enemyKingPos, !g.ColotToMove) {
		return true
	}
	return false
}
