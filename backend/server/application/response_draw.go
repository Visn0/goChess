package application

import (
	"chess/server/domain"
)

type ResponseDrawParam struct {
	DrawResponse bool `json:"drawResponse"`
}

type ResponseDrawOutput struct {
	Action       string `json:"action"`
	DrawResponse bool   `json:"drawResponse"`
}

type ResponseDrawAction struct {
	c domain.ConnectionRepository
}

func NewResponseDrawAction(c domain.ConnectionRepository) *ResponseDrawAction {
	return &ResponseDrawAction{
		c: c,
	}
}

func (uc *ResponseDrawAction) Invoke(drawResponse bool) error {
	output := ResponseDrawOutput{
		Action:       "draw-response",
		DrawResponse: drawResponse,
	}

	return uc.c.SendWebSocketMessage(output)
}
