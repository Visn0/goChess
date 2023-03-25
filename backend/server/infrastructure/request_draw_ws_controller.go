package infrastructure

import (
	"chess/server/application"
	"chess/server/shared/wsrouter"
)

type RequestDrawWsController struct {
	uc *application.RequestDrawAction
}

func NewRequestDrawWsController() *RequestDrawWsController {
	return &RequestDrawWsController{
		uc: application.NewRequestDrawAction(),
	}
}

func (c *RequestDrawWsController) Invoke(ctx *wsrouter.Context) error {
	return c.uc.Invoke(ctx)
}
