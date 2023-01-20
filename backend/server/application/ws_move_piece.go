package application

import (
	"chess/server/domain"
	"chess/server/shared"
	"encoding/json"
	"log"
	"time"
)

type RequestMovePiece struct {
	Src       *domain.Position  `json:"src"`
	Dst       *domain.Position  `json:"dst"`
	PromoteTo *domain.PieceType `json:"promoteTo"`
}

type ResponseMovePiece struct {
	Action    string            `json:"action"`
	Src       *domain.Position  `json:"src"`
	Dst       *domain.Position  `json:"dst"`
	PromoteTo *domain.PieceType `json:"promoteTo"`
}

func WsMovePiece(g *domain.Game, body []byte, player, enemy *shared.WsConn) {
	log.Println("handle move piece")
	req := RequestMovePiece{}
	json.Unmarshal(body, &req)

	move := &domain.Move{
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
