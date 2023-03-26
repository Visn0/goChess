package infrastructure

import (
	"chess/server/application"
	"chess/server/shared/wsrouter"
	"encoding/json"
	"log"
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
	err := json.Unmarshal(ctx.Body, &p)
	if err != nil {
		log.Println("Error unmarshalling draw response:", err)
		return err
	}

	output := c.uc.Invoke(ctx, &p)
	return ctx.EnemyRepository.SendWebSocketMessage(output)
}
