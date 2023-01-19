package actions

import (
	"chess/server/game"
	"chess/server/shared"
	"encoding/json"
	"log"
)

type RequestMovePiece struct {
	Src       *game.Position `json:"src"`
	Dst       *game.Position `json:"dst"`
	Promotion string         `json:"promotion"`
}

type ResponseMovePiece struct {
	Action    string         `json:"action"`
	Src       *game.Position `json:"src"`
	Dst       *game.Position `json:"dst"`
	Promotion string         `json:"promotion"`
}

func WsMovePiece(g *game.Game, body []byte, c *shared.WsConn) {
	log.Println("handle move piece")
	req := RequestMovePiece{}
	json.Unmarshal(body, &req)

	move := &game.Move{
		From: req.Src,
		To:   req.Dst,
	}
	g.Move(move)
	resp := ResponseMovePiece{
		Action:    "move-piece",
		Src:       req.Src,
		Dst:       req.Dst,
		Promotion: "",
	}

	c.WriteJSON(resp)
}
