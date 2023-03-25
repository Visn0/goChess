package infrastructure

import (
	"chess/server/application"
	"chess/server/shared/wsrouter"
	"encoding/json"
	"log"
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
	err := json.Unmarshal(ctx.Body, &p)
	if err != nil {
		log.Println("Error unmarshalling move piece params:", err)
		return err
	}

	output, err := c.uc.Invoke(ctx, &p)
	if err != nil {
		return nil
	}

	if err := ctx.OwnRepository.SendWebSocketMessage(output); err != nil {
		return err
	}

	return ctx.EnemyRepository.SendWebSocketMessage(output)
}
