package domain

import "fmt"

func (g *Game) PositionIsUnderAttack(pos *Position, enemyColor Color) bool {
	for i := 0; i < N_PIECE_TYPES; i++ {
		pieceType := PieceType(i)
		directions := getPieceDirection(pieceType, enemyColor)
		if g.positionIsUnderAttackUsingDirections(pos, pieceType, enemyColor, directions) {
			return true
		}
	}
	return false
}

func (g *Game) positionIsUnderAttackUsingDirections(pos *Position, pieceType PieceType, enemyColor Color, directions []Direction) bool {
	if pieceType == PAWN {
		directions = []Direction{{-directions[0].x, -1}, {-directions[0].x, 1}}
	}
	for _, d := range directions {
		fmt.Println("pos: ", pos, "check ", pieceType.String(), " dir: ", d)
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
				if piece.GetColor() == enemyColor && piece.GetPieceType() == pieceType {
					fmt.Println("check piece: ", piece, " pos: ", newPos)
					return true
				}
				break
			}
			// TODO: improve this s**t
			if pieceType == PAWN || pieceType == KING || pieceType == KNIGHT {
				break
			}
		}
	}
	return false
}
