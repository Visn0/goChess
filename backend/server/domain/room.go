package domain

import (
	"chess/server/shared/chesserror"
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
	} else {
		return chesserror.NewError(chesserror.WrongInputParameter, "Room is full")
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

	return chesserror.NewError(chesserror.ResourceNotFound, fmt.Sprintf("Player with id %q not found", p.ID))
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
