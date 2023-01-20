package room

import (
	"chess/server/game"
	"time"

	websocket "github.com/gofiber/websocket/v2"
)

type wsConn = websocket.Conn

type Player struct {
	Ws                    *wsConn
	ID                    string
	Color                 game.Color
	TimeConsumedInSeconds int
	LastClockTime         time.Time
}

type PlayerPublicInfo struct {
	ID string `json:"id"`
}

func (p *Player) GetPublicInfo() *PlayerPublicInfo {
	return &PlayerPublicInfo{
		ID: p.ID,
	}
}

func (p *Player) StartTimer() {
	p.LastClockTime = time.Now()
}

func (p *Player) StopTimer() {
	p.TimeConsumedInSeconds += int(time.Since(p.LastClockTime).Seconds())
	p.LastClockTime = time.Now()
}

func (p *Player) TimeLeft() int {
	return 10*60 - p.TimeConsumedInSeconds
}
