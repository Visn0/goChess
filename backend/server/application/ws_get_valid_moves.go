package application

import (
	"chess/server/domain"

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
	r    domain.ConnectionRepository
	game *domain.Game
}

func NewGetValidMovesAction(r domain.ConnectionRepository, game *domain.Game) *GetValidMovesAction {
	return &GetValidMovesAction{
		r:    r,
		game: game,
	}
}

func (uc *GetValidMovesAction) Invoke(p *GetValidMovesParams) error {
	validMoves := uc.game.GetValidMoves(p.Rank, p.File)
	if validMoves == nil {
		fmt.Println("No valid moves found")
		return nil
	}

	output := GetValidMovesOutput{
		Action:     "request-moves",
		ValidMoves: validMoves,
	}

	return uc.r.SendWebSocketMessage(output)
}
