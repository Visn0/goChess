package infrastructure

import (
	"chess/server/application"
	"chess/server/domain"
)

type RequestDrawWsController struct {
	uc *application.RequestDrawAction
}

func NewRequestDrawWsController(c domain.ConnectionRepository) *RequestDrawWsController {
	return &RequestDrawWsController{
		uc: application.NewRequestDrawAction(c),
	}
}

func (c *RequestDrawWsController) Invoke(_ []byte) error {
	return c.uc.Invoke()
}
