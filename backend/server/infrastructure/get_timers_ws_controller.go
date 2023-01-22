package infrastructure

import (
	"chess/server/application"
	"chess/server/domain"
)

type GetTimersWsController struct {
	uc *application.GetTimersAction
}

func NewGetTimersWsController(c domain.ConnectionRepository, player, enemy *domain.Player) *GetTimersWsController {
	return &GetTimersWsController{
		uc: application.NewGetTimersAction(c, player, enemy),
	}
}

func (c *GetTimersWsController) Invoke() error {
	return c.uc.Invoke()
}
