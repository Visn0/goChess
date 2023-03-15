package application

import (
	"chess/server/domain"
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
	c      domain.ConnectionRepository
	cEnemy domain.ConnectionRepository
	game   *domain.Game
}

func NewMovePieceAction(c domain.ConnectionRepository, cEnemy domain.ConnectionRepository,
	game *domain.Game) *MovePieceAction {
	return &MovePieceAction{
		c:      c,
		cEnemy: cEnemy,
		game:   game,
	}
}

func (uc *MovePieceAction) setGameStatus(enemyColor domain.Color, output *MovePieceOutput) {
	enemyKingPos := uc.game.Board.GetKingPos(enemyColor)
	if enemyKingPos == nil {
		log.Println("##> Enemy King not found: ", enemyColor)
		panic("")
	}
	if uc.game.Board.PositionIsUnderAttack(enemyKingPos, !enemyColor) {
		output.KingCheck = enemyKingPos
		log.Println("##> Enemy King is under attack: ", enemyColor, enemyKingPos)
	}
	enemyHasMoves := uc.game.CalculateValidMoves(enemyColor)
	if !enemyHasMoves {
		log.Println("##> Enemy has no valid moves: ", enemyColor)
		if output.KingCheck != nil {
			output.EndGame = "checkmate"
		} else {
			output.EndGame = "draw"
		}
	}
}

func (uc *MovePieceAction) Invoke(p *MovePieceParams) error {
	// log.Println("==> Move piece params: ", shared.ToJSONString(p))
	move := &domain.Move{
		From: p.Src,
		To:   p.Dst,
	}

	playerColor := uc.game.ColorToMove
	enemyColor := !playerColor

	uc.game.Move(move, p.PromoteTo)
	fmt.Println("Game EnpassantPieces: ", uc.game.EnPassantPieces)
	uc.game.ColorToMove = enemyColor

	output := newMovePieceOutput(p.Src, p.Dst, p.PromoteTo)
	uc.setGameStatus(enemyColor, output)

	// log.Println("##> Move piece output: ", shared.ToJSONString(output))
	log.Println("##> Player to move: ", enemyColor)
	err := uc.c.SendWebSocketMessage(output)
	if err != nil {
		return nil
	}

	return uc.cEnemy.SendWebSocketMessage(output)
}
