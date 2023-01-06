package server

import (
	"chess/server/game"
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

type RequestCreateRoom struct {
	Name     string `json:"name"`
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
	}
	_, ok := s.rooms[req.Name]
	if ok {
		resp.HttpCode = 400
		err = c.WriteJSON(resp)
		if err != nil {
			log.Println("Error sending error response room exists:", err)
		}
	}
	room := NewRoom()
	player := &Player{
		ws: c,
		id: "player1",
	}
	room.AddPlayer(player)
	s.rooms[req.Name] = room
	log.Println("Room created")

	roomWG := &sync.WaitGroup{}
	go room.HandleGame(true, roomWG)

	resp.Room = &ResponseRoom{
		ID: req.Name,
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
	Name     string `json:"name"`
	Password string `json:"password"`
}

type ResponseJoinRoom struct {
	Action   string `json:"action"`
	HttpCode int    `json:"httpCode"`
}

func (s *Server) handleJoinRoom(body []byte, c *wsConn) {
	req := RequestJoinRoom{}
	resp := ResponseJoinRoom{
		Action:   "join-room",
		HttpCode: 200,
	}
	err := json.Unmarshal(body, &req)
	if err != nil {
		log.Println("Error unmarshalling request join room:", err)
		resp.HttpCode = 400
		err = c.WriteJSON(resp)
		if err != nil {
			log.Println("Error sending error response unmarshalling:", err)
		}
	}
	room, ok := s.rooms[req.Name]
	if !ok {
		resp.HttpCode = 400
		err = c.WriteJSON(resp)
		if err != nil {
			log.Println("Error sending error response room exists:", err)
		}
	}
	player := &Player{
		ws: c,
		id: "player2",
	}
	room.AddPlayer(player)
	log.Println("Player joined room")
	roomWG := &sync.WaitGroup{}
	go room.HandleGame(false, roomWG)

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
