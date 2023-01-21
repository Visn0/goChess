package application

import (
	"chess/server/domain"
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
	uc.addTestRooms(uc.rm)
	rooms := uc.rm.GetRooms()
	return newGetRoomsOutput(rooms)
}

func (uc *GetRoomsAction) addTestRooms(rm *domain.RoomManager) {
	roomTest1 := &domain.Room{
		ID: "roomTest1",
		Player1: &domain.Player{
			Ws: nil,
			ID: "player1",
		},
		Player2: &domain.Player{
			Ws: nil,
			ID: "player2",
		},
	}
	roomTest2 := &domain.Room{
		ID: "roomTest2",
		Player1: &domain.Player{
			Ws: nil,
			ID: "player3",
		},
		Player2: &domain.Player{
			Ws: nil,
			ID: "player4",
		},
	}
	rm.AddRoom(roomTest1)
	rm.AddRoom(roomTest2)
}
