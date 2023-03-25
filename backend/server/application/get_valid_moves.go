package application

import (
	"chess/server/domain"
	"chess/server/shared/chesserror"
	"chess/server/shared/wsrouter"
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
}

func NewGetValidMovesAction() *GetValidMovesAction {
	return &GetValidMovesAction{}
}

func (uc *GetValidMovesAction) Invoke(ctx *wsrouter.Context, p *GetValidMovesParams) (*GetValidMovesOutput, error) {
	piece := ctx.Game.Board.GetPiece(&domain.Position{Rank: p.Rank, File: p.File})
	if piece == nil {
		return nil, chesserror.NewError(chesserror.ResourceNotFound, "No piece found at given position")
	}

	if piece.GetColor() != ctx.Game.ColorToMove {
		return nil, chesserror.NewError(chesserror.GenericError,
			"Movement not valid: it is not the turn of the given player.")
	}

	validMoves := piece.GetValidMoves()
	output := &GetValidMovesOutput{
		Action:     "request-moves",
		ValidMoves: validMoves,
	}

	return output, nil
}
