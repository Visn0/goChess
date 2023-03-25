package application

import (
	"chess/server/shared/wsrouter"
)

type AbandonOutput struct {
	Action string `json:"action"`
}

type AbandonAction struct {
}

func NewAbandonAction() *AbandonAction {
	return &AbandonAction{}
}

func (uc *AbandonAction) Invoke(ctx *wsrouter.Context) error {
	output := AbandonOutput{
		Action: "abandon",
	}

	return ctx.EnemyRepository.SendWebSocketMessage(output)
}
