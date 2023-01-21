package domain

import (
	"time"

	websocket "github.com/gofiber/websocket/v2"
)

type wsConn = websocket.Conn

type Player struct {
	Ws             *wsConn
	ID             string
	Color          Color
	TimeConsumedMS int
	LastClockTime  time.Time
}

func NewPlayer(ws *wsConn, id string, color Color) *Player {
	return &Player{
		Ws:             ws,
		ID:             id,
		Color:          color,
		TimeConsumedMS: 0,
		LastClockTime:  time.Time{},
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
