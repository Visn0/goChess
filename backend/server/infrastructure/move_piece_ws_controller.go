package infrastructure

import (
	"chess/server/application"
	"chess/server/domain"
	"encoding/json"
	"log"
)

type MovePieceWsController struct {
	uc *application.MovePieceAction

	ownRepository   domain.ConnectionRepository
	enemyRepository domain.ConnectionRepository
}

func NewMovePieceWsController(
	player, enemy *domain.Player,
	ownRepository, enemyRepository domain.ConnectionRepository,
	game *domain.Game) *MovePieceWsController {
	return &MovePieceWsController{
		uc:              application.NewMovePieceAction(player, enemy, game),
		ownRepository:   ownRepository,
		enemyRepository: enemyRepository,
	}
}

func (c *MovePieceWsController) Invoke(body []byte) error {
	var p application.MovePieceParams
	err := json.Unmarshal(body, &p)
	if err != nil {
		log.Println("Error unmarshalling move piece params:", err)
		return err
	}

	output, err := c.uc.Invoke(&p)
	if err != nil {
		return nil
	}

	if err := c.ownRepository.SendWebSocketMessage(output); err != nil {
		return err
	}

	return c.enemyRepository.SendWebSocketMessage(output)
}
