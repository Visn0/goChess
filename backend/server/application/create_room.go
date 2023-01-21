package application

import (
	"chess/server/domain"
	"fmt"
	"log"
)

type CreateRoomParams struct {
	RoomID   string `json:"roomID"`
	PlayerID string `json:"playerID"`
	Password string `json:"password"`
}

type CreateRoomOutput struct {
	ID      string                    `json:"id"`
	Players []*CreateRoomPlayerOutput `json:"players"`
}

type CreateRoomPlayerOutput struct {
	ID string `json:"id"`
}

func newCreateRoomOutput(playerID string) *CreateRoomOutput {
	return &CreateRoomOutput{
		ID: playerID,
		Players: []*CreateRoomPlayerOutput{{
			ID: playerID,
		}},
	}
}

type CreateRoomAction struct {
	rm *domain.RoomManager
	r  domain.ConnectionRepository
}

func NewCreateRoomAction(rm *domain.RoomManager, r domain.ConnectionRepository) *CreateRoomAction {
	return &CreateRoomAction{rm: rm, r: r}
}

func (uc *CreateRoomAction) Invoke(p *CreateRoomParams) (*domain.Room, error) {
	_, ok := uc.rm.GetRoom(p.RoomID)
	if ok {
		err := fmt.Errorf("There is already a room with id %q", p.RoomID)
		log.Println(err)
		return nil, err
	}

	player := domain.NewPlayer(uc.r.GetWebSocketConnection(), p.PlayerID, domain.WHITE)

	r := domain.NewRoom(p.RoomID)
	r.AddPlayer(player)

	uc.rm.AddRoom(r)
	log.Println("Room created", r.ID, "by player", player.ID)

	output := newCreateRoomOutput(p.PlayerID)
	err := uc.r.SendWebSocketMessage(output)
	if err != nil {
		return nil, err
	}

	return r, nil
}
