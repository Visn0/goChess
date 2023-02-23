package application

import (
	"chess/server/domain"
)

type RequestDrawOutput struct {
	Action string `json:"action"`
}

type RequestDrawAction struct {
	c domain.ConnectionRepository
}

func NewRequestDrawAction(c domain.ConnectionRepository) *RequestDrawAction {
	return &RequestDrawAction{
		c: c,
	}
}

func (uc *RequestDrawAction) Invoke() error {
	output := RequestDrawOutput{
		Action: "receive-draw-request",
	}

	return uc.c.SendWebSocketMessage(output)
}
