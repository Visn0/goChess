package server

import (
	"chess/server/game"
	"encoding/json"
	"fmt"
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

func (s *Server) handleCreateRoom(body []byte, c *wsConn) {
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
	_, ok := s.rooms[req.RoomID]
	if ok {
		resp.HttpCode = 400
		err = c.WriteJSON(resp)
		if err != nil {
			log.Println("Error sending error response room exists:", err)
		}
		return
	}
	room := NewRoom()
	player := &Player{
		ws: c,
		id: req.PlayerID,
	}
	room.AddPlayer(player)
	s.rooms[req.RoomID] = room
	log.Println("Room created")

	roomWG := &sync.WaitGroup{}
	go room.HandleGame(true, roomWG)

	resp.Room = &ResponseRoom{
		ID: req.RoomID,
		Players: []*ResponsePlayer{
			{
				ID: player.id,
			},
		},
	}
	c.WriteJSON(resp)
	roomWG.Wait()
}

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

func (s *Server) handleJoinRoom(body []byte, c *wsConn) {
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
	room, ok := s.rooms[req.RoomID]
	if !ok {
		resp.HttpCode = 400
		err = c.WriteJSON(resp)
		if err != nil {
			log.Println("Error sending error response room exists:", err)
		}
		return
	}
	player := &Player{
		ws: c,
		id: req.PlayerID,
	}
	room.AddPlayer(player)
	log.Println("Player joined room")
	roomWG := &sync.WaitGroup{}
	go room.HandleGame(false, roomWG)

	playersInfo := []*ResponsePlayer{}
	if room.player1 != nil {
		playersInfo = append(playersInfo, &ResponsePlayer{
			ID: room.player1.id,
		})
	}
	if room.player2 != nil {
		playersInfo = append(playersInfo, &ResponsePlayer{
			ID: room.player2.id,
		})
	}
	resp.Room = &ResponseRoom{
		ID:      req.RoomID,
		Players: playersInfo,
	}
	j, _ := json.Marshal(resp)
	log.Println(string(j))
	c.WriteJSON(resp)
	roomWG.Wait()
}

type RequestMoves struct {
	*game.Move
	Piece string `json:"piece"`
}

type ResponseMoves struct {
	Action     string       `json:"action"`
	ValidMoves []*game.Move `json:"validMoves"`
}

func (r *Room) handleRequestMoves(body []byte, c *wsConn) {
	log.Println("handle request moves")
	req := RequestMoves{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		log.Println("Error unmarshalling request moves:", err)
		return
	}

	validMoves := r.game.GetValidMoves(req.Rank, req.File)
	for _, move := range validMoves {
		fmt.Println(move)
	}
	fmt.Printf("Found %d valid moves", len(validMoves))
	resp := ResponseMoves{
		Action:     "request-moves",
		ValidMoves: validMoves,
	}
	j, _ := json.Marshal(resp)
	log.Println(string(j))
	c.WriteJSON(resp)
}

type RequestMovePiece struct {
	Src *game.Move `json:"src"`
	Dst *game.Move `json:"dst"`
}

type ResponseMovePiece struct {
	Action string     `json:"action"`
	Src    *game.Move `json:"src"`
	Dst    *game.Move `json:"dst"`
}

func (r *Room) handleMovePiece(body []byte, c *wsConn) {
	log.Println("handle move piece")
	req := RequestMovePiece{}
	json.Unmarshal(body, &req)

	r.game.Move(req.Src, req.Dst)
	resp := ResponseMovePiece{
		Action: "move-piece",
		Src:    req.Src,
		Dst:    req.Dst,
	}

	c.WriteJSON(resp)
}
