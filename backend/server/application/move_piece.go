package application

import (
	"chess/server/domain"
	"chess/server/shared"
	"chess/server/shared/wsrouter"
	"fmt"
	"log"
)

type MovePieceParams struct {
	Src       *domain.Position  `json:"src"`
	Dst       *domain.Position  `json:"dst"`
	PromoteTo *domain.PieceType `json:"promoteTo"`
}

type MovePieceOutput struct {
	Action    string            `json:"action"`
	Src       *domain.Position  `json:"src"`
	Dst       *domain.Position  `json:"dst"`
	PromoteTo *domain.PieceType `json:"promoteTo"`
	EndGame   string            `json:"endGame"`
	KingCheck *domain.Position  `json:"kingCheck"`
}

func newMovePieceOutput(src, dst *domain.Position, promoteTo *domain.PieceType) *MovePieceOutput {
	return &MovePieceOutput{
		Action:    "move-piece",
		Src:       src,
		Dst:       dst,
		PromoteTo: promoteTo,
		KingCheck: nil,
	}
}

type MovePieceAction struct {
}

func NewMovePieceAction() *MovePieceAction {
	return &MovePieceAction{}
}

func (uc *MovePieceAction) setGameStatus(ctx *wsrouter.Context, enemyColor domain.Color, output *MovePieceOutput) {
	enemyKingPos := ctx.Game.Board.GetKingPos(enemyColor)
	if enemyKingPos == nil {
		log.Println("##> Enemy King not found: ", enemyColor)
		panic("")
	}
	if ctx.Game.Board.PositionIsUnderAttack(enemyKingPos, !enemyColor) {
		output.KingCheck = enemyKingPos
		log.Println("##> Enemy King is under attack: ", enemyColor, enemyKingPos)
	}
	enemyHasMoves := ctx.Game.CalculateValidMoves(enemyColor)
	if !enemyHasMoves {
		log.Println("##> Enemy has no valid moves: ", enemyColor)
		if output.KingCheck != nil {
			output.EndGame = "checkmate"
		} else {
			output.EndGame = "draw"
		}
	}
}

func (uc *MovePieceAction) Invoke(ctx *wsrouter.Context, p *MovePieceParams) (*MovePieceOutput, error) {
	log.Println("==> Move piece params: ", shared.ToJSONString(p))
	move := &domain.Move{
		From: p.Src,
		To:   p.Dst,
	}

	playerColor := ctx.Game.ColorToMove
	enemyColor := !playerColor

	ctx.Game.Move(move, p.PromoteTo)
	fmt.Println("Game EnpassantPieces: ", ctx.Game.EnPassantPieces)
	ctx.Game.ColorToMove = enemyColor

	output := newMovePieceOutput(p.Src, p.Dst, p.PromoteTo)
	uc.setGameStatus(ctx, enemyColor, output)

	ctx.Player.StopTimer()
	ctx.Enemy.StartTimer()

	log.Println("##> Move piece output: ", shared.ToJSONString(output))
	return output, nil
}
