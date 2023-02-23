package domain

import "fmt"

// CalculateValidMoves calculates valid moves for all pieces of a given color
// and returns true if the player has valid moves
func (g *Game) CalculateValidMoves(color Color) bool {
	playerHasValidMoves := false
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			pos := &Position{Rank(rank), File(file)}
			piece := g.Board.GetPiece(pos)
			if piece != nil && piece.GetColor() == color {
				validMoves := g.getPieceValidMoves(pos, piece)
				if len(validMoves) > 0 {
					piece.SetValidMoves(validMoves)
					playerHasValidMoves = true
				}
			}
		}
	}
	return playerHasValidMoves
}

// getPieceValidMoves returns a list of valid moves for a piece at a given position
func (g *Game) getPieceValidMoves(pos *Position, piece IPiece) []*Position {
	pieceValidMovesHandler := g.getPieceValidMovesHandler(piece.GetPieceType())
	positions := pieceValidMovesHandler(pos, piece)
	positions = g.filterMovesIfCheck(pos, positions)
	return positions
}

func (g *Game) getPieceValidMovesHandler(pieceType PieceType) func(pos *Position, p IPiece) []*Position {
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

func (g *Game) getShortDistanceMoves(pos *Position, p IPiece) []*Position {
	positions := []*Position{}
	for _, d := range p.GetValidDirections() {
		dstPos := &Position{Rank: pos.Rank + Rank(d.x), File: pos.File + File(d.y)}
		if dstPos.Valid() {
			dstPiece := g.Board.GetPiece(dstPos)
			if dstPiece == nil {
				positions = append(positions, dstPos)
			} else if p.GetPieceType() != PAWN && p.IsEnemy(dstPiece) {
				// Pawns can not take pieces in front of them
				positions = append(positions, dstPos)
			}
		}
	}
	return positions
}

func (g *Game) getLongDistanceMoves(pos *Position, p IPiece) []*Position {
	positions := []*Position{}
	for _, d := range p.GetValidDirections() {
		dCum := &Direction{0, 0}
		for {
			dCum.x += d.x
			dCum.y += d.y
			dstPos := &Position{Rank: pos.Rank + Rank(dCum.x), File: pos.File + File(dCum.y)}
			if dstPos.Valid() {
				dstPiece := g.Board.GetPiece(dstPos)
				if dstPiece == nil {
					positions = append(positions, dstPos)
				} else if p.IsEnemy(dstPiece) {
					positions = append(positions, dstPos)
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

func (g *Game) getPawnValidMoves(pos *Position, p IPiece) []*Position {
	positions := g.getShortDistanceMoves(pos, p)
	pawn := p.(*Pawn)
	// Check if pawn can move two spaces forward
	if pawn.FirstMove {
		dstPos := &Position{Rank: pos.Rank + Rank(p.GetValidDirections()[0].x*2), File: pos.File}
		if dstPos.Valid() && g.Board.GetPiece(dstPos) == nil {
			positions = append(positions, dstPos)
		}
	}
	// Check if pawn can move diagonally
	pawnRankDir := Rank(p.GetValidDirections()[0].x)
	topPos := Position{Rank: pos.Rank + pawnRankDir, File: pos.File}
	if topPos.Valid() {
		// Check if there is a piece in front left of the pawn
		topLeftPos := &Position{Rank: topPos.Rank, File: topPos.File - 1}
		if topLeftPos.Valid() {
			topLeftPiece := g.Board.GetPiece(topLeftPos)
			// Check if the piece is an enemy piece
			if topLeftPiece != nil && pawn.IsEnemy(topLeftPiece) {
				positions = append(positions, topLeftPos)
			}
		}
		// Check if there is a piece in front right of the pawn
		topRightPos := &Position{Rank: topPos.Rank, File: topPos.File + 1}
		if topRightPos.Valid() {
			topRightPiece := g.Board.GetPiece(topRightPos)
			// Check if the piece is an enemy piece
			if topRightPiece != nil && pawn.IsEnemy(topRightPiece) {
				positions = append(positions, topRightPos)
			}
		}
		// Check if pawn has En Passant move
		if pawn.EnPassantNeighbourPos != nil {
			dstPos := &Position{Rank: pos.Rank + pawnRankDir, File: pawn.EnPassantNeighbourPos.File}
			if dstPos.Valid() {
				positions = append(positions, dstPos)
			}
		}
	}
	return positions
}

func (g *Game) getRookValidMoves(pos *Position, p IPiece) []*Position {
	positions := g.getLongDistanceMoves(pos, p)
	return positions
}

func (g *Game) getKnightValidMoves(pos *Position, p IPiece) []*Position {
	positions := g.getShortDistanceMoves(pos, p)
	return positions
}

func (g *Game) getBishopValidMoves(pos *Position, p IPiece) []*Position {
	positions := g.getLongDistanceMoves(pos, p)
	return positions
}

func (g *Game) getQueenValidMoves(pos *Position, p IPiece) []*Position {
	positions := g.getLongDistanceMoves(pos, p)
	return positions
}

func (g *Game) getKingValidMoves(pos *Position, p IPiece) []*Position {
	positions := g.getShortDistanceMoves(pos, p)
	g.getKingCastlePositions(pos, p, &positions)
	return positions
}

func (g *Game) getKingCastlePositions(pos *Position, p IPiece, positions *[]*Position) {
	// Check if king can castle
	if !p.(*King).FirstMove {
		return
	}
	// Check if king can castle kingside
	oneRightPos := &Position{Rank: pos.Rank, File: pos.File + 1}
	twoRightPos := &Position{Rank: pos.Rank, File: pos.File + 2}
	if g.Board.GetPiece(oneRightPos) == nil && g.Board.GetPiece(twoRightPos) == nil {
		rookPos := &Position{Rank: pos.Rank, File: pos.File + 3}
		rook := g.Board.GetPiece(rookPos)
		if rook != nil && rook.GetPieceType() == ROOK && rook.(*Rook).FirstMove {
			*positions = append(*positions, twoRightPos)
		}
	}
	// Check if king can castle queenside
	oneLeftPos := &Position{Rank: pos.Rank, File: pos.File - 1}
	twoLeftPos := &Position{Rank: pos.Rank, File: pos.File - 2}
	threeLeftPos := &Position{Rank: pos.Rank, File: pos.File - 3}
	if g.Board.GetPiece(oneLeftPos) == nil &&
		g.Board.GetPiece(twoLeftPos) == nil &&
		g.Board.GetPiece(threeLeftPos) == nil {
		rookPos := &Position{Rank: pos.Rank, File: pos.File - 4}
		rook := g.Board.GetPiece(rookPos)
		if rook != nil && rook.GetPieceType() == ROOK && rook.(*Rook).FirstMove {
			*positions = append(*positions, twoLeftPos)
		}
	}
}

// Remove all positions that will put the king in check
// TODO: has some bugs during castling and en passant
// TODO: not the most efficient way to do this (implement it with a copy of the board)
func (g *Game) filterMovesIfCheck(from *Position, positions []*Position) []*Position {
	filteredPositions := []*Position{}
	for _, pos := range positions {
		move := &Move{From: from, To: pos}
		if !g.isCheckAfterMove(move) {
			filteredPositions = append(filteredPositions, pos)
		}
	}
	return filteredPositions
}

// Check if the move will put the king in check
func (g *Game) isCheckAfterMove(move *Move) bool {
	boardCopy := g.Board.Copy()
	boardCopy.ExecuteMove(move) //TODO: doesn't work perfectly for castling or en passant yet
	// Check if the king is in check
	enemyKingPos := g.Board.GetKingPos(g.ColorToMove)
	// If the king is moved, update the king position
	if enemyKingPos.Rank == move.From.Rank && enemyKingPos.File == move.From.File {
		enemyKingPos = move.To
	}
	if enemyKingPos == nil {
		fmt.Println("King not found: " + g.ColorToMove.String())
	}
	if boardCopy.PositionIsUnderAttack(enemyKingPos, !g.ColorToMove) {
		return true
	}
	return false
}
