package game

import "fmt"

type Game struct {
	Board *Board
}

func NewGame() *Game {
	return &Game{Board: NewBoard()}
}

func (g *Game) Move(src, dst *Move) {
	p := g.Board.GetPiece(src.Rank, src.File)
	g.Board.SetPiece(dst.Rank, dst.File, p)
	g.Board.RemovePiece(src.Rank, src.File)
}

func (g *Game) GetValidMoves(rank Rank, file File) []*Move {
	fmt.Println("Getting valid moves for", rank, file)
	p := g.Board.GetPiece(rank, file)
	if p == nil {
		fmt.Println("No piece at", rank, file)
		return nil
	}
	if p.ValidMoves != nil {
		fmt.Println("Returning cached moves")
		return p.ValidMoves
	}
	moves := []*Move{}
	for _, d := range p.ValidDirections {
		for {
			move := &Move{Rank: rank, File: file}
			move.Add(d)
			if !move.Valid() {
				break
			}
			fmt.Println("Checking", move)
			if g.Board.GetPiece(move.Rank, move.File) == nil {
				moves = append(moves, move)
			}
		}
	}
	p.ValidMoves = moves
	return moves
}
