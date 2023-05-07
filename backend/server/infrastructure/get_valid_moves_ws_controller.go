package infrastructure

import (
	"chess/server/application"
	"chess/server/shared/wsrouter"
)

type GetValidMovesWsController struct {
	uc *application.GetValidMovesAction
}

func NewGetValidMovesWsController() *GetValidMovesWsController {
	return &GetValidMovesWsController{
		uc: application.NewGetValidMovesAction(),
	}
}

func (c *GetValidMovesWsController) Invoke(ctx *wsrouter.Context) error {
	var p application.GetValidMovesParams
	err := ctx.Bind(&p)
	if err != nil {
		return err
	}

	output, err := c.uc.Invoke(ctx, &p)
	if err != nil {
		return err
	}

	return ctx.Player.SendWebSocketMessage(output)
}
