package infrastructure

import (
	"chess/server/application"
	"chess/server/shared/wsrouter"
	"encoding/json"
	"log"
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
	err := json.Unmarshal(ctx.Body, &p)
	if err != nil {
		log.Println("Error unmarshalling request create room:", err)
		return err
	}

	output, err := c.uc.Invoke(ctx, &p)
	if err != nil {
		return err
	}

	return ctx.OwnRepository.SendWebSocketMessage(output)
}
