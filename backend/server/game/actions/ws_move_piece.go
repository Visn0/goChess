package actions

import (
	"chess/server/game"
	"chess/server/shared"
	"encoding/json"
	"log"
	"time"
)

type RequestMovePiece struct {
	Src       *game.Position  `json:"src"`
	Dst       *game.Position  `json:"dst"`
	PromoteTo *game.PieceType `json:"promoteTo"`
}

type ResponseMovePiece struct {
	Action    string          `json:"action"`
	Src       *game.Position  `json:"src"`
	Dst       *game.Position  `json:"dst"`
	PromoteTo *game.PieceType `json:"promoteTo"`
}

func WsMovePiece(g *game.Game, body []byte, player, enemy *shared.WsConn) {
	log.Println("handle move piece")
	req := RequestMovePiece{}
	json.Unmarshal(body, &req)

	move := &game.Move{
		From: req.Src,
		To:   req.Dst,
	}
	g.Move(move, req.PromoteTo)
	resp := ResponseMovePiece{
		Action:    "move-piece",
		Src:       req.Src,
		Dst:       req.Dst,
		PromoteTo: req.PromoteTo,
	}

	player.WriteJSON(resp)
	time.Sleep(100 * time.Millisecond)
	enemy.WriteJSON(resp)
}
