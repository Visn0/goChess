package infrastructure

import (
	"chess/server/application"
	"chess/server/shared/wsrouter"
)

func StartGameActionWsController(ctx *wsrouter.Context, durationMs int) error {
	outputPlayer1, outputPlayer2 := application.StartGameAction(ctx.Player, ctx.Enemy, durationMs)

	err := ctx.Player.SendWebSocketMessage(outputPlayer1)
	if err != nil {
		return err
	}

	return ctx.Enemy.SendWebSocketMessage(outputPlayer2)
}
