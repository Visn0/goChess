package application

import (
	"chess/server/domain"
	"chess/server/shared"

	"encoding/json"
	"fmt"
	"log"
)

type RequestMoves struct {
	Rank domain.Rank `json:"rank"`
	File domain.File `json:"file"`
}

type ResponseMoves struct {
	Action     string             `json:"action"`
	ValidMoves []*domain.Position `json:"validMoves"`
}

func WsGetValidMoves(g *domain.Game, body []byte, c *shared.WsConn) {
	log.Println("Handle request moves")
	req := RequestMoves{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		log.Println("Error unmarshalling request moves:", err)
		return
	}
	validMoves := g.GetValidMoves(req.Rank, req.File)
	if validMoves == nil {
		fmt.Println("No valid moves found")
		return
	}
	// fmt.Printf("Found %d valid moves", len(validMoves))
	resp := ResponseMoves{
		Action:     "request-moves",
		ValidMoves: validMoves,
	}
	// j, _ := json.Marshal(resp)
	// log.Println(string(j))
	c.WriteJSON(resp)
}
