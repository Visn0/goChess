package actions

import (
	"chess/server/room"

	fiber "github.com/gofiber/fiber/v2"
)

type PlayerResponse struct {
	ID string `json:"id"`
}

type RoomResponse struct {
	ID      string            `json:"id"`
	Players []*PlayerResponse `json:"players"`
}

type HttpRoomsResponse struct {
	Rooms []*RoomResponse `json:"rooms"`
}

func HttpGetRooms(ctx *fiber.Ctx, rm *room.RoomManager) error {
	// TEST ROOMS
	addTestRooms(rm) //TODO: remove

	rooms := rm.GetRooms()
	roomsResponse := make([]*RoomResponse, 0, len(rooms))
	for _, room := range rooms {
		info := room.GetPublicInfo()
		roomResponse := &RoomResponse{
			ID:      info.ID,
			Players: make([]*PlayerResponse, 0, 2),
		}
		for _, player := range info.Players {
			roomResponse.Players = append(roomResponse.Players, &PlayerResponse{
				ID: player.ID,
			})
		}
		roomsResponse = append(roomsResponse, roomResponse)
	}
	resp := HttpRoomsResponse{
		Rooms: roomsResponse,
	}
	return ctx.JSON(resp)
}

func addTestRooms(rm *room.RoomManager) {
	roomTest1 := &room.Room{
		ID: "roomTest1",
		Player1: &room.Player{
			Ws: nil,
			ID: "player1",
		},
		Player2: &room.Player{
			Ws: nil,
			ID: "player2",
		},
	}
	roomTest2 := &room.Room{
		ID: "roomTest2",
		Player1: &room.Player{
			Ws: nil,
			ID: "player3",
		},
		Player2: &room.Player{
			Ws: nil,
			ID: "player4",
		},
	}
	rm.AddRoom(roomTest1)
	rm.AddRoom(roomTest2)
}