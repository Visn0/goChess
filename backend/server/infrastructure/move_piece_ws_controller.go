package infrastructure

import (
	"chess/server/application"
	"chess/server/domain"
	"encoding/json"
	"log"
)

type MovePieceWsController struct {
	uc *application.MovePieceAction
}

func NewMovePieceWsController(ownRepository domain.ConnectionRepository, enemyRepository domain.ConnectionRepository,
	game *domain.Game) *MovePieceWsController {
	return &MovePieceWsController{
		uc: application.NewMovePieceAction(ownRepository, enemyRepository, game),
	}
}

func (c *MovePieceWsController) Invoke(body []byte) error {
	var p application.MovePieceParams
	err := json.Unmarshal(body, &p)
	if err != nil {
		log.Println("Error unmarshalling request create room:", err)
		return err
	}

	return c.uc.Invoke(&p)
}
