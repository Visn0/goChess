package server

import (
	"chess/server/game"
	"encoding/json"
	"fmt"
	"log"
)

type RequestMoves struct {
	*game.Move
	Piece string `json:"piece"`
}

type ResponseMoves struct {
	Action     string       `json:"action"`
	ValidMoves []*game.Move `json:"validMoves"`
}

func (s *Server) handleRequestMoves(body []byte, c *wsConn) {
	log.Println("handle request moves")
	req := RequestMoves{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		log.Println("Error unmarshalling request moves:", err)
		return
	}

	fmt.Println("Searching room")
	room, ok := s.rooms["room1"]
	if !ok {
		fmt.Println("Room not found")
		return
	}
	validMoves := room.game.GetValidMoves(req.Rank, req.File)
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

func (s *Server) handleMovePiece(body []byte, c *wsConn) {
	log.Println("handle move piece")
	req := RequestMovePiece{}
	json.Unmarshal(body, &req)

	// game.movePiece(req.Src, req.Dst)
	resp := ResponseMovePiece{
		Action: "move-piece",
		Src:    req.Src,
		Dst:    req.Dst,
	}
	c.WriteJSON(resp)
}
