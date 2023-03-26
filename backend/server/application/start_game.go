package application

import (
	"chess/server/domain"
)

type StartGameOutput struct {
	Action       string `json:"action"`
	PlayerColor  string `json:"playerColor"`
	OpponentName string `json:"opponentName"`
	Duration     int    `json:"duration"`
}

func newStartGameOutput(playerColor domain.Color, opponentName string, durationMs int) *StartGameOutput {
	return &StartGameOutput{
		Action:       "start-game",
		PlayerColor:  domain.ColorToString(playerColor),
		OpponentName: opponentName,
		Duration:     durationMs,
	}
}

func StartGameAction(p1, p2 *domain.Player, durationMs int) (op1, op2 *StartGameOutput) {
	p1.Color = domain.GetRandomColor()
	if p1.Color == domain.WHITE {
		p2.Color = domain.BLACK
		defer p1.StartTimer()
	} else {
		p2.Color = domain.WHITE
		defer p2.StartTimer()
	}

	outputPlayer1 := newStartGameOutput(p1.Color, p2.ID, durationMs)
	outputPlayer2 := newStartGameOutput(p2.Color, p1.ID, durationMs)

	return outputPlayer1, outputPlayer2
}
