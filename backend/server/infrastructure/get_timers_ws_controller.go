package infrastructure

import (
	"chess/server/application"
	"chess/server/shared/wsrouter"
)

type GetTimersWsController struct {
	uc *application.GetTimersAction
}

func NewGetTimersWsController() *GetTimersWsController {
	return &GetTimersWsController{
		uc: application.NewGetTimersAction(),
	}
}

func (c *GetTimersWsController) Invoke(ctx *wsrouter.Context) error {
	output := c.uc.Invoke(ctx)
	return ctx.Player.SendWebSocketMessage(output)
}
