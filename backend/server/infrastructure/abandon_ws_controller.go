package infrastructure

import (
	"chess/server/application"
	"chess/server/shared/wsrouter"
)

type AbandonWsController struct {
	uc *application.AbandonAction
}

func NewAbandonWsController() *AbandonWsController {
	return &AbandonWsController{
		uc: application.NewAbandonAction(),
	}
}

func (c *AbandonWsController) Invoke(ctx *wsrouter.Context) error {
	return c.uc.Invoke(ctx)
}
