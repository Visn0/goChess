package application

import (
	"chess/server/shared/wsrouter"
)

type GetTimersOutput struct {
	Action     string `json:"action"`
	PlayerTime int    `json:"playerTime"`
	EnemyTime  int    `json:"enemyTime"`
}

type GetTimersAction struct {
}

func NewGetTimersAction() *GetTimersAction {
	return &GetTimersAction{}
}

func (uc *GetTimersAction) Invoke(ctx *wsrouter.Context) *GetTimersOutput {
	t1 := ctx.Player.TimeLeft()
	t2 := ctx.Enemy.TimeLeft()

	output := &GetTimersOutput{
		Action:     "get-timers",
		PlayerTime: t1,
		EnemyTime:  t2,
	}

	return output
}
