package application

import (
	"chess/server/domain"
	"chess/server/shared"
	"chess/server/shared/chesserror"
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
	c  domain.ConnectionRepository
}

func NewCreateRoomAction(rm *domain.RoomManager, c domain.ConnectionRepository) *CreateRoomAction {
	return &CreateRoomAction{rm: rm, c: c}
}

func (uc *CreateRoomAction) Invoke(p *CreateRoomParams) (*domain.Room, error) {
	log.Println("==> Create room params: ", shared.ToJSONString(p))
	_, ok := uc.rm.GetRoom(p.RoomID)
	if ok {
		err := chesserror.NewError(chesserror.ResourceAlreadyExists,
			fmt.Sprintf("There is already a room with id '%s'", p.RoomID))
		log.Println(err)
		return nil, err
	}

	player := domain.NewPlayer(uc.c.GetWebSocketConnection(), p.PlayerID)

	r := domain.NewRoom(p.RoomID)
	err := r.AddPlayer(player)
	if err != nil {
		return nil, err
	}

	uc.rm.AddRoom(r)
	output := newCreateRoomOutput(p.PlayerID)
	log.Println("##> Create room output: ", shared.ToJSONString(output))

	err = uc.c.SendWebSocketMessage(output)
	if err != nil {
		return nil, err
	}

	return r, nil
}
