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
	return c.uc.Invoke(ctx)
}
