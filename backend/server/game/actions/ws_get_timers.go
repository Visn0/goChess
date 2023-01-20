package actions

import "chess/server/shared"

type ResponseTimers struct {
	Action     string `json:"action"`
	PlayerTime int    `json:"player-time"`
	EnemyTime  int    `json:"enemy-time"`
}

func WsGetTimers(p1, p2 *shared.WsConn, t1, t2 int) {
	p1.WriteJSON(ResponseTimers{
		Action:     "get-timers",
		PlayerTime: t1,
		EnemyTime:  t2,
	})
	p2.WriteJSON(ResponseTimers{
		Action:     "get-timers",
		PlayerTime: t2,
		EnemyTime:  t1,
	})
}
