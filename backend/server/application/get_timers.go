package application

import (
	"chess/server/domain"
)

type GetTimersOutput struct {
	Action     string `json:"action"`
	PlayerTime int    `json:"playerTime"`
	EnemyTime  int    `json:"enemyTime"`
}

type GetTimersAction struct {
	c      domain.ConnectionRepository
	player *domain.Player
	enemy  *domain.Player
}

func NewGetTimersAction(c domain.ConnectionRepository, p1, p2 *domain.Player) *GetTimersAction {
	return &GetTimersAction{
		c:      c,
		player: p1,
		enemy:  p2,
	}
}

func (uc *GetTimersAction) Invoke() error {
	// log.Println("==> Request timers: ")
	t1 := uc.player.TimeLeft()
	t2 := uc.enemy.TimeLeft()

	output := GetTimersOutput{
		Action:     "get-timers",
		PlayerTime: t1,
		EnemyTime:  t2,
	}

	// log.Println("##> Request timers output: ", shared.ToJSONString(output))
	return uc.c.SendWebSocketMessage(output)
}
