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
	g.Board.SetPiece(m.To.Rank, m.To.File, p)
	g.Board.RemovePiece(m.From.Rank, m.From.File)
}

func (g *Game) GetValidPositions(rank Rank, file File) []*Position {
	fmt.Println("Getting valid moves for", rank, file)
	p := g.Board.GetPiece(rank, file)
	if p == nil {
		fmt.Println("No piece at", rank, file)
		return nil
	}
	fmt.Println("Piece at", rank, file, "is", p.String())
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
	if p.(*Pawn).FirstMove {
		positions = append(positions, &Position{Rank: rank + Rank(p.GetValidDirections()[0].x*2), File: file})
	}
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
	return positions
}
