package server

import (
	"encoding/json"
	"log"
)

type RequestMoves struct {
	*Coordinate
	Piece string `json:"piece"`
}

type ResponseMoves struct {
	Action     string        `json:"action"`
	ValidMoves []*Coordinate `json:"validMoves"`
}

type Coordinate struct {
	File int `json:"file"` // col
	Rank int `json:"rank"` // row
}

func (s *Server) handleRequestMoves(body []byte, c *wsConn) {
	log.Println("handle request moves")
	req := RequestMoves{}
	json.Unmarshal(body, &req)

	// game.legalMoves(req.piece, req.Coordinate)
	resp := ResponseMoves{
		Action: "request-moves",
		ValidMoves: []*Coordinate{
			{File: 1, Rank: 1},
			{File: 2, Rank: 2},
		},
	}
	c.WriteJSON(resp)
}

type RequestMovePiece struct {
	Src *Coordinate `json:"src"`
	Dst *Coordinate `json:"dst"`
}

type ResponseMovePiece struct {
	Action string      `json:"action"`
	Src    *Coordinate `json:"src"`
	Dst    *Coordinate `json:"dst"`
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
