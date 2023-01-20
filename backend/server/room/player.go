package room

import (
	"chess/server/game"
	"time"

	websocket "github.com/gofiber/websocket/v2"
)

type wsConn = websocket.Conn

type Player struct {
	Ws             *wsConn
	ID             string
	Color          game.Color
	TimeConsumedMS int
	LastClockTime  time.Time
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
	p.TimeConsumedMS += int(time.Since(p.LastClockTime).Milliseconds())
	p.LastClockTime = time.Now()
}

func (p *Player) TimeLeft() int {
	if p.LastClockTime.IsZero() {
		return 10 * 60 * 1000
	}
	return 10*60*1000 - (p.TimeConsumedMS + int(time.Since(p.LastClockTime).Milliseconds()))
}
