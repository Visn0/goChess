package infrastructure

import (
	"chess/server/application"
	"chess/server/domain"
	"encoding/json"
	"log"
)

type GetValidMovesWsController struct {
	uc *application.GetValidMovesAction
}

func NewGetValidMovesWsController(r domain.ConnectionRepository, game *domain.Game) *GetValidMovesWsController {
	return &GetValidMovesWsController{
		uc: application.NewGetValidMovesAction(r, game),
	}
}

func (c *GetValidMovesWsController) Invoke(body []byte) error {
	var p application.GetValidMovesParams
	err := json.Unmarshal(body, &p)
	if err != nil {
		log.Println("Error unmarshalling request create room:", err)
		return err
	}

	return c.uc.Invoke(&p)
}
