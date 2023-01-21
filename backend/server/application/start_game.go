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

func StartGameAction(p1, p2 *domain.Player, p1Repo, p2Repo domain.ConnectionRepository, duration int) error {
	p1.Color = domain.GetRandomColor()
	if p1.Color == domain.WHITE {
		p2.Color = domain.BLACK
		defer p1.StartTimer()
	} else {
		p2.Color = domain.WHITE
		defer p2.StartTimer()
	}

	output := newStartGameOutput(p1.Color, duration)
	err := p1Repo.SendWebSocketMessage(output)
	if err != nil {
		return err
	}

	output = newStartGameOutput(p2.Color, duration)
	return p2Repo.SendWebSocketMessage(output)
}
