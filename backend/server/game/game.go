package game

import "fmt"

type Game struct {
	Board *Board
}

func NewGame() *Game {
	return &Game{Board: NewBoard()}
}

func (g *Game) Move(fromRank, fromFile, toRank, toFile int) {
	p := g.Board.GetPiece(fromRank, fromFile)
	g.Board.SetPiece(toRank, toFile, p)
	g.Board.SetPiece(fromRank, fromFile, nil)
}

func (g *Game) GetValidMoves(rank, file int) []*Move {
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
			move := &Move{Rank: rank + d.x, File: file + d.y}
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
