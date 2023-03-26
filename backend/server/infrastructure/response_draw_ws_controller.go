package infrastructure

import (
	"chess/server/application"
	"chess/server/shared/wsrouter"
)

type ResponseDrawWsController struct {
	uc *application.ResponseDrawAction
}

func NewResponseDrawWsController() *ResponseDrawWsController {
	return &ResponseDrawWsController{
		uc: application.NewResponseDrawAction(),
	}
}

func (c *ResponseDrawWsController) Invoke(ctx *wsrouter.Context) error {
	var p application.ResponseDrawParam
	err := ctx.Bind(&p)
	if err != nil {
		return err
	}

	output := c.uc.Invoke(ctx, &p)
	return ctx.EnemyRepository.SendWebSocketMessage(output)
}
