package server

import (
	"chess/server/game"
	"encoding/json"
	"fmt"
	"log"
)

type RequestCreateRoom struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (s *Server) handleCreateRoom(body []byte, c *wsConn) error {
	req := RequestCreateRoom{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		log.Println("Error unmarshalling request create room:", err)
		return err
	}
	_, ok := s.rooms[req.Name]
	if ok {
		return fmt.Errorf("room already exists")
	}
	room := NewRoom()
	player := &Player{
		ws: c,
		id: "player1",
	}
	room.AddPlayer(player)
	s.rooms[req.Name] = room
	log.Println("Room created")
	go room.HandleGame(true)
	return nil
}

type RequestJoinRoom struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (s *Server) handleJoinRoom(body []byte, c *wsConn) error {
	req := RequestJoinRoom{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		log.Println("Error unmarshalling request join room:", err)
		return err
	}
	room, ok := s.rooms[req.Name]
	if !ok {
		return fmt.Errorf("room does not exist")
	}
	player := &Player{
		ws: c,
		id: "player2",
	}
	room.AddPlayer(player)
	log.Println("Player joined room")
	go room.HandleGame(false)
	return nil
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
