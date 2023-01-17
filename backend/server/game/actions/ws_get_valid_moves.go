package actions

import (
	"chess/server/game"
	"chess/server/shared"

	"encoding/json"
	"fmt"
	"log"
)

type RequestMoves struct {
	*game.Position
	Piece string `json:"piece"`
}

type ResponseMoves struct {
	Action     string           `json:"action"`
	ValidMoves []*game.Position `json:"validMoves"`
}

func WsGetValidMoves(g *game.Game, body []byte, c *shared.WsConn) {
	log.Println("handle request moves")
	req := RequestMoves{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		log.Println("Error unmarshalling request moves:", err)
		return
	}
	validMoves := g.GetValidPositions(req.Rank, req.File)
	fmt.Printf("Found %d valid moves", len(validMoves))
	resp := ResponseMoves{
		Action:     "request-moves",
		ValidMoves: validMoves,
	}
	j, _ := json.Marshal(resp)
	log.Println(string(j))
	c.WriteJSON(resp)
}
