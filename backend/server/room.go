package server

import (
	"chess/server/game"
	"fmt"
)

type Room struct {
	player1 *Player
	player2 *Player
	game    *game.Game
}

func NewRoom() *Room {
	return &Room{game: game.NewGame()}
}

func (r *Room) AddPlayer(p *Player) error {
	if r.player1 == nil {
		r.player1 = p
	} else if r.player2 == nil {
		r.player2 = p
	} else {
		return fmt.Errorf("Room is full")
	}
	return nil
}

func (r *Room) RemovePlayer(p *Player) error {
	if r.player1 == p {
		r.player1 = nil
		return nil
	} else if r.player2 == p {
		r.player2 = nil
		return nil
	}
	return fmt.Errorf("Player not found")
}

func (r *Room) GetRoomSize() int {
	size := 0
	if r.player1 != nil {
		size++
	}
	if r.player2 != nil {
		size++
	}
	return size
}
