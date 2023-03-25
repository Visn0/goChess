package application

import (
	"chess/server/shared/wsrouter"
)

type ResponseDrawParam struct {
	DrawResponse bool `json:"drawResponse"`
}

type ResponseDrawOutput struct {
	Action       string `json:"action"`
	DrawResponse bool   `json:"drawResponse"`
}

type ResponseDrawAction struct {
}

func NewResponseDrawAction() *ResponseDrawAction {
	return &ResponseDrawAction{}
}

func (uc *ResponseDrawAction) Invoke(ctx *wsrouter.Context, p *ResponseDrawParam) error {
	output := ResponseDrawOutput{
		Action:       "draw-response",
		DrawResponse: p.DrawResponse,
	}

	return ctx.EnemyRepository.SendWebSocketMessage(output)
}
