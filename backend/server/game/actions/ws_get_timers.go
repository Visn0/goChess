package actions

import (
	"chess/server/shared"
	"log"
)

type ResponseTimers struct {
	Action     string `json:"action"`
	PlayerTime int    `json:"playerTime"`
	EnemyTime  int    `json:"enemyTime"`
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

	log.Println("=>> timers =>> p1: ", t1/1000, "| p2: ", t2)
}
