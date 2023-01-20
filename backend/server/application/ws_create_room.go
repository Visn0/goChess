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

	// var roomWG sync.WaitGroup
	// roomWG.Add(1)
	// go r.HandleGame(true, &roomWG)

	// roomWG.Wait()
	return r, nil
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
// type RequestCreateRoom struct {
// 	RoomID   string `json:"roomID"`
// 	PlayerID string `json:"playerID"`
// 	Password string `json:"password"`
// }

// type ResponseCreateRoom struct {
// 	HttpCode int           `json:"httpCode"`
// 	Action   string        `json:"action"`
// 	Room     *ResponseRoom `json:"room"`
// }

type ResponseRoom struct {
	ID      string            `json:"id"`
	Players []*ResponsePlayer `json:"players"`
}

type ResponsePlayer struct {
	ID string `json:"id"`
}

// func WsCreateRoom(rm *domain.RoomManager, body []byte, c *shared.WsConn) {
// 	req := RequestCreateRoom{}
// 	resp := ResponseCreateRoom{
// 		Action: "create-room",
// 	}
// 	err := json.Unmarshal(body, &req)
// 	if err != nil {
// 		log.Println("Error unmarshalling request create room:", err)
// 		resp.HttpCode = 400
// 		err = c.WriteJSON(resp)
// 		if err != nil {
// 			log.Println("Error sending error response unmarshalling:", err)
// 		}
// 		return
// 	}

// 	_, ok := rm.GetRoom(req.RoomID)
// 	if ok {
// 		resp.HttpCode = 400
// 		err = c.WriteJSON(resp)
// 		if err != nil {
// 			log.Println("Error sending error response room exists:", err)
// 		}
// 		return
// 	}

// 	r := domain.NewRoom(req.RoomID)
// 	player := &domain.Player{
// 		Ws:                    c,
// 		ID:                    req.PlayerID,
// 		TimeConsumedInSeconds: 0,
// 		Color:                 domain.WHITE,
// 	}

// 	r.AddPlayer(player)
// 	rm.AddRoom(r)
// 	log.Println("Room created", r.ID, "by player", player.ID)

// 	resp.Room = &ResponseRoom{
// 		ID: req.RoomID,
// 		Players: []*ResponsePlayer{
// 			{
// 				ID: player.ID,
// 			},
// 		},
// 	}

// 	c.WriteJSON(resp)

// 	var roomWG sync.WaitGroup
// 	roomWG.Add(1)
// 	go r.HandleGame(true, &roomWG)

// 	roomWG.Wait()
// }
