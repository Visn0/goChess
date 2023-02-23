package infrastructure

import (
	"chess/server/application"
	"chess/server/domain"
	"encoding/json"
	"log"
)

type ResponseDrawWsController struct {
	uc *application.ResponseDrawAction
}

func NewResponseDrawWsController(c domain.ConnectionRepository) *ResponseDrawWsController {
	return &ResponseDrawWsController{
		uc: application.NewResponseDrawAction(c),
	}
}

func (c *ResponseDrawWsController) Invoke(body []byte) error {
	var p application.ResponseDrawParam
	err := json.Unmarshal(body, &p)
	if err != nil {
		log.Println("Error unmarshalling draw response:", err)
		return err
	}

	return c.uc.Invoke(p.DrawResponse)
}
