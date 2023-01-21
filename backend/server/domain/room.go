package domain

import (
	// "chess/server/application"
	"fmt"
)

type Room struct {
	ID      string
	Player1 *Player
	Player2 *Player
	Game    *Game
}

func NewRoom(id string) *Room {
	return &Room{ID: id, Game: NewGame()}
}

func (r *Room) AddPlayer(p *Player) error {
	if r.Player1 == nil {
		r.Player1 = p
	} else if r.Player2 == nil {
		r.Player2 = p
		fmt.Println("Set player 2")
		r.Player1.StartTimer()
	} else {
		return fmt.Errorf("Room is full")
	}
	return nil
}

func (r *Room) RemovePlayer(p *Player) error {
	if r.Player1 == p {
		r.Player1 = nil
		return nil
	} else if r.Player2 == p {
		r.Player2 = nil
		return nil
	}
	return fmt.Errorf("Player not found")
}

func (r *Room) GetRoomSize() int {
	size := 0
	if r.Player1 != nil {
		size++
	}
	if r.Player2 != nil {
		size++
	}
	return size
}
