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
	// if p.ValidMoves != nil {
	// 	fmt.Println("Returning cached moves")
	// 	return p.ValidMoves
	// }
	positions := []*Position{}
	fromPos := &Position{Rank: rank, File: file}
	for _, d := range p.ValidDirections {
		dCum := Direction{0, 0}
		for {
			dCum.x += d.x
			dCum.y += d.y
			pos := &Position{Rank: fromPos.Rank + Rank(dCum.x), File: fromPos.File + File(dCum.y)}
			if !pos.Valid() {
				fmt.Println("Invalid move", pos.Rank, pos.File, d)
				break
			}
			fmt.Println("Checking", pos.Rank, pos.File, d)
			if g.Board.GetPiece(pos.Rank, pos.File) == nil {
				positions = append(positions, pos)
			} else {
				break
			}
			if p.Name == "king" || p.Name == "knight" {
				break
			}
		}
	}
	p.ValidPositions = positions
	return positions
}
