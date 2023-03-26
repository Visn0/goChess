package application

import "chess/server/shared/wsrouter"

type RequestDrawOutput struct {
	Action string `json:"action"`
}

type RequestDrawAction struct {
}

func NewRequestDrawAction() *RequestDrawAction {
	return &RequestDrawAction{}
}

func (uc *RequestDrawAction) Invoke(ctx *wsrouter.Context) *RequestDrawOutput {
	output := &RequestDrawOutput{
		Action: "draw-request",
	}

	return output
}
