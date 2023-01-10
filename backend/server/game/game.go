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
	fmt.Println("Piece at", rank, file, "is", p.String())
	// if p.ValidMoves != nil {
	// 	fmt.Println("Returning cached moves")
	// 	return p.ValidMoves
	// }
	moves := []*Move{}
	fromMove := &Move{Rank: rank, File: file}
	for _, d := range p.ValidDirections {
		dCum := Direction{0, 0}
		for {
			dCum.x += d.x
			dCum.y += d.y
			move := &Move{Rank: fromMove.Rank + Rank(dCum.x), File: fromMove.File + File(dCum.y)}
			if !move.Valid() {
				fmt.Println("Invalid move", move.Rank, move.File, d)
				break
			}
			fmt.Println("Checking", move.Rank, move.File, d)
			if g.Board.GetPiece(move.Rank, move.File) == nil {
				moves = append(moves, move)
			} else {
				break
			}
			if p.Name == "king" || p.Name == "knight" {
				break
			}
		}
	}
	p.ValidMoves = moves
	return moves
}
