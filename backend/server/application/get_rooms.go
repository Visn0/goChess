package application

import (
	"chess/server/domain"
	"chess/server/shared"
	"log"
)

type GetRoomsOutput struct {
	Rooms []*GetRoomInfoOutput `json:"rooms"`
}

type GetRoomInfoOutput struct {
	ID      string                  `json:"id"`
	Players []*GetRoomsPlayerOutput `json:"players"`
}

type GetRoomsPlayerOutput struct {
	ID string `json:"id"`
}

func newGetRoomsOutput(rooms []*domain.Room) *GetRoomsOutput {
	outputRooms := make([]*GetRoomInfoOutput, 0, len(rooms))
	for i := range rooms {
		r := rooms[i]
		players := make([]*GetRoomsPlayerOutput, 0, 2)

		if r.Player1 != nil {
			players = append(players, &GetRoomsPlayerOutput{
				ID: r.Player1.ID,
			})
		}

		if r.Player2 != nil {
			players = append(players, &GetRoomsPlayerOutput{
				ID: r.Player2.ID,
			})
		}

		outputRooms = append(outputRooms, &GetRoomInfoOutput{
			ID:      r.ID,
			Players: players,
		})
	}

	return &GetRoomsOutput{
		Rooms: outputRooms,
	}
}

type GetRoomsAction struct {
	rm *domain.RoomManager
}

func NewGetRoomsAction(rm *domain.RoomManager) *GetRoomsAction {
	return &GetRoomsAction{rm: rm}
}

func (uc *GetRoomsAction) Invoke() *GetRoomsOutput {
	rooms := uc.rm.GetRooms()
	output := newGetRoomsOutput(rooms)
	log.Println("==> Get rooms output: ", shared.ToJSONString(output))

	return output
}
