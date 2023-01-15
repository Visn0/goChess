package actions

import (
	"chess/server/room"
	"chess/server/shared"
	"encoding/json"
	"log"
	"sync"
)

type RequestJoinRoom struct {
	RoomID   string `json:"roomID"`
	PlayerID string `json:"playerID"`
	Password string `json:"password"`
}

type ResponseJoinRoom struct {
	HttpCode int           `json:"httpCode"`
	Action   string        `json:"action"`
	Room     *ResponseRoom `json:"room"`
}

func WsJoinRoom(rm *room.RoomManager, body []byte, c *shared.WsConn) {
	req := RequestJoinRoom{}
	resp := ResponseJoinRoom{
		Action:   "join-room",
		HttpCode: 200,
		Room:     nil,
	}
	err := json.Unmarshal(body, &req)
	if err != nil {
		log.Println("Error unmarshalling request join room:", err)
		resp.HttpCode = 400
		err = c.WriteJSON(resp)
		if err != nil {
			log.Println("Error sending error response unmarshalling:", err)
		}
		return
	}
	r, ok := rm.GetRoom(req.RoomID)
	if !ok {
		resp.HttpCode = 400
		err = c.WriteJSON(resp)
		if err != nil {
			log.Println("Error sending error response room exists:", err)
		}
		return
	}
	player := &room.Player{
		Ws: c,
		ID: req.PlayerID,
	}
	r.AddPlayer(player)
	log.Println("Player joined room")

	playersInfo := []*ResponsePlayer{}
	if r.Player1 != nil {
		playersInfo = append(playersInfo, &ResponsePlayer{
			ID: r.Player1.ID,
		})
	}
	if r.Player2 != nil {
		playersInfo = append(playersInfo, &ResponsePlayer{
			ID: r.Player2.ID,
		})
	}
	resp.Room = &ResponseRoom{
		ID:      req.RoomID,
		Players: playersInfo,
	}
	j, _ := json.Marshal(resp)
	log.Println(string(j))
	c.WriteJSON(resp)

	var roomWG sync.WaitGroup
	roomWG.Add(1)
	go r.HandleGame(false, &roomWG)

	roomWG.Wait()
}
