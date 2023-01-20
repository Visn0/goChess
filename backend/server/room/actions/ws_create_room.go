package actions

import (
	"chess/server/game"
	"chess/server/room"
	"chess/server/shared"
	"encoding/json"
	"log"
	"sync"
)

type RequestCreateRoom struct {
	RoomID   string `json:"roomID"`
	PlayerID string `json:"playerID"`
	Password string `json:"password"`
}

type ResponseCreateRoom struct {
	HttpCode int           `json:"httpCode"`
	Action   string        `json:"action"`
	Room     *ResponseRoom `json:"room"`
}

type ResponseRoom struct {
	ID      string            `json:"id"`
	Players []*ResponsePlayer `json:"players"`
}

type ResponsePlayer struct {
	ID string `json:"id"`
}

func WsCreateRoom(rm *room.RoomManager, body []byte, c *shared.WsConn) {
	req := RequestCreateRoom{}
	resp := ResponseCreateRoom{
		Action: "create-room",
	}
	err := json.Unmarshal(body, &req)
	if err != nil {
		log.Println("Error unmarshalling request create room:", err)
		resp.HttpCode = 400
		err = c.WriteJSON(resp)
		if err != nil {
			log.Println("Error sending error response unmarshalling:", err)
		}
		return
	}

	_, ok := rm.GetRoom(req.RoomID)
	if ok {
		resp.HttpCode = 400
		err = c.WriteJSON(resp)
		if err != nil {
			log.Println("Error sending error response room exists:", err)
		}
		return
	}

	r := room.NewRoom(req.RoomID)
	player := &room.Player{
		Ws:                    c,
		ID:                    req.PlayerID,
		TimeConsumedInSeconds: 0,
		Color:                 game.WHITE,
	}

	r.AddPlayer(player)
	rm.AddRoom(r)
	log.Println("Room created", r.ID, "by player", player.ID)

	resp.Room = &ResponseRoom{
		ID: req.RoomID,
		Players: []*ResponsePlayer{
			{
				ID: player.ID,
			},
		},
	}

	c.WriteJSON(resp)

	var roomWG sync.WaitGroup
	roomWG.Add(1)
	go r.HandleGame(true, &roomWG)

	roomWG.Wait()
}
