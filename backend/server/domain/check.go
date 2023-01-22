package domain

import "fmt"

func (g *Game) PositionIsUnderAttack(pos *Position, enemyColor Color) bool {
	//TODO: check dir for each piece because far away pawn cant attack
	queenDir := getPieceDirection(QUEEN, enemyColor)
	if g.positionIsUnderAttackUsingDirections(pos, enemyColor, queenDir) {
		return true
	}
	knightDir := getPieceDirection(ROOK, enemyColor)
	return g.positionIsUnderAttackUsingDirections(pos, enemyColor, knightDir)
}

func (g *Game) positionIsUnderAttackUsingDirections(pos *Position, enemyColor Color, directions []Direction) bool {
	for _, d := range directions {
		fmt.Println("check dir: ", d, " pos: ", pos)
		dCum := &Direction{0, 0}
		for {
			dCum.x += d.x
			dCum.y += d.y
			newPos := &Position{Rank: pos.Rank + Rank(dCum.x), File: pos.File + File(dCum.y)}
			if !newPos.Valid() {
				break
			}
			piece := g.Board.GetPiece(newPos.Rank, newPos.File)
			if piece != nil {
				if piece.GetColor() == enemyColor {
					fmt.Println("check piece: ", piece, " pos: ", newPos)
					return true
				}
				break
			}
		}
	}
	return false
}
