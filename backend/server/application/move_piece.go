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
}

func newMovePieceOutput(src *domain.Position, dst *domain.Position, promoteTo *domain.PieceType) *MovePieceOutput {
	return &MovePieceOutput{
		Action:    "move-piece",
		Src:       src,
		Dst:       dst,
		PromoteTo: promoteTo,
	}
}

type MovePieceAction struct {
	ownRepository   domain.ConnectionRepository
	enemyRepository domain.ConnectionRepository
	game            *domain.Game
}

func NewMovePieceAction(ownRepository domain.ConnectionRepository, enemyRepository domain.ConnectionRepository,
	game *domain.Game) *MovePieceAction {
	return &MovePieceAction{
		ownRepository:   ownRepository,
		enemyRepository: enemyRepository,
		game:            game,
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
	log.Println("##> Move piece output: ", shared.ToJSONString(output))

	err := uc.ownRepository.SendWebSocketMessage(output)
	if err != nil {
		return nil
	}

	return uc.enemyRepository.SendWebSocketMessage(output)
}
