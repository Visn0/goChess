package application

import (
	"chess/server/domain"
)

type StartGameOutput struct {
	Action      string `json:"action"`
	PlayerColor string `json:"playerColor"`
	Duration    int    `json:"duration"`
}

func newStartGameOutput(playerColor domain.Color, duration int) *StartGameOutput {
	return &StartGameOutput{
		Action:      "start-game",
		PlayerColor: domain.ColorToString(playerColor),
		Duration:    duration,
	}
}

func StartGameAction(p1Color, p2Color domain.Color, p1Repo, p2Repo domain.ConnectionRepository, duration int) error {
	output := newStartGameOutput(p1Color, duration)
	err := p1Repo.SendWebSocketMessage(output)
	if err != nil {
		return err
	}

	output = newStartGameOutput(p2Color, duration)
	return p2Repo.SendWebSocketMessage(output)
}
