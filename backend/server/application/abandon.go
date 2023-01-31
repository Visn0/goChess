package application

import (
	"chess/server/domain"
)

type AbandonOutput struct {
	Action     string `json:"action"`
}

type AbandonAction struct {
	c      domain.ConnectionRepository
}

func NewAbandonAction(c domain.ConnectionRepository) *AbandonAction {
	return &AbandonAction{
		c:      c,
	}
}

func (uc *AbandonAction) Invoke() error {
	output := AbandonOutput{
		Action:     "abandon",
	}

	return uc.c.SendWebSocketMessage(output)
}
