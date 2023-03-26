package infrastructure

import (
	"chess/server/application"
	"chess/server/shared/wsrouter"
)

type MovePieceWsController struct {
	uc *application.MovePieceAction
}

func NewMovePieceWsController() *MovePieceWsController {
	return &MovePieceWsController{
		uc: application.NewMovePieceAction(),
	}
}

func (c *MovePieceWsController) Invoke(ctx *wsrouter.Context) error {
	var p application.MovePieceParams
	err := ctx.Bind(&p)
	if err != nil {
		return err
	}

	output, err := c.uc.Invoke(ctx, &p)
	if err != nil {
		return nil
	}

	if err := ctx.Player.SendWebSocketMessage(output); err != nil {
		return err
	}

	return ctx.Enemy.SendWebSocketMessage(output)
}
