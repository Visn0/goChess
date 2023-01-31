package application

import (
	"chess/server/domain"
	"chess/server/shared"
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

func (uc *MovePieceAction) Invoke(p *MovePieceParams) error {
	log.Println("==> Move piece params: ", shared.ToJSONString(p))
	move := &domain.Move{
		From: p.Src,
		To:   p.Dst,
	}

	uc.game.Move(move, p.PromoteTo)
	output := newMovePieceOutput(p.Src, p.Dst, p.PromoteTo)

	nextColorToMove := uc.game.ColotToMove
	enemyKingPos := uc.game.Board.GetKingPos(nextColorToMove)
	if enemyKingPos == nil {
		log.Println("##> King not found: ", nextColorToMove)
		panic("King not found")
	}
	if uc.game.Board.PositionIsUnderAttack(enemyKingPos, !nextColorToMove) {
		output.KingCheck = enemyKingPos
		log.Println("##> King is under attack: ", nextColorToMove, enemyKingPos)
	}
	log.Println("##> Move piece output: ", shared.ToJSONString(output))
	log.Println("##> Player to move: ", nextColorToMove)
	err := uc.c.SendWebSocketMessage(output)
	if err != nil {
		return nil
	}

	return uc.cEnemy.SendWebSocketMessage(output)
}
