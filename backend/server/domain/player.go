package domain

import (
	"chess/server/shared"
	"time"
)

type Player struct {
	Ws             *shared.WsConn
	ID             string
	Color          Color
	TimeConsumedMS int
	LastClockTime  time.Time
	paused         bool
}

func NewPlayer(ws *shared.WsConn, id string) *Player {
	return &Player{
		Ws:             ws,
		ID:             id,
		Color:          false,
		TimeConsumedMS: 0,
		LastClockTime:  time.Time{},
		paused:         true,
	}
}

func (p *Player) StartTimer() {
	p.LastClockTime = time.Now()
	p.paused = false
}

func (p *Player) StopTimer() {
	p.TimeConsumedMS += int(time.Since(p.LastClockTime).Milliseconds())
	p.LastClockTime = time.Now()
	p.paused = true
}

func (p *Player) TimeLeft() int {
	if p.LastClockTime.IsZero() {
		return 10 * 60 * 1000
	}

	if p.paused {
		return 10*60*1000 - p.TimeConsumedMS
	}

	return 10*60*1000 - (p.TimeConsumedMS + int(time.Since(p.LastClockTime).Milliseconds()))
}
