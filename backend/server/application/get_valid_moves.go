package application

import (
	"chess/server/domain"
	"chess/server/shared"
	"log"

	"fmt"
)

type GetValidMovesParams struct {
	Rank domain.Rank `json:"rank"`
	File domain.File `json:"file"`
}

type GetValidMovesOutput struct {
	Action     string             `json:"action"`
	ValidMoves []*domain.Position `json:"validMoves"`
}

type GetValidMovesAction struct {
	c    domain.ConnectionRepository
	game *domain.Game
}

func NewGetValidMovesAction(c domain.ConnectionRepository, game *domain.Game) *GetValidMovesAction {
	return &GetValidMovesAction{
		c:    c,
		game: game,
	}
}

func (uc *GetValidMovesAction) Invoke(p *GetValidMovesParams) error {
	// log.Println("==> Request moves params: ", shared.ToJSONString(p))
	validMoves := uc.game.GetValidMoves(p.Rank, p.File)
	if validMoves == nil {
		fmt.Println("No valid moves found")
		return nil
	}

	output := GetValidMovesOutput{
		Action:     "request-moves",
		ValidMoves: validMoves,
	}

	log.Println("##> Request moves output: ", shared.ToJSONString(output))
	return uc.c.SendWebSocketMessage(output)
}
