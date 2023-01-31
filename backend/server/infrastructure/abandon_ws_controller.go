package infrastructure

import (
	"chess/server/application"
	"chess/server/domain"
)

type AbandonWsController struct {
	uc *application.AbandonAction
}

func NewAbandonWsController(c domain.ConnectionRepository) *AbandonWsController {
	return &AbandonWsController{
		uc: application.NewAbandonAction(c),
	}
}

func (c *AbandonWsController) Invoke() error {
	return c.uc.Invoke()
}
