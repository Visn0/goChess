package application

import (
	"chess/server/domain"
	"time"
)

type MovePieceParams struct {
	Src       *domain.Position  `json:"src"`
	Dst       *domain.Position  `json:"dst"`
	PromoteTo *domain.PieceType `json:"promoteTo"`
}

type MovePieceOutput struct {
	Action    string            `json:"action"`
	Src       *domain.Position  `json:"src"`
	Dst       *domain.Position  `json:"dst"`
	PromoteTo *domain.PieceType `json:"promoteTo"`
}

func newMovePieceOutput(src *domain.Position, dst *domain.Position, promoteTo *domain.PieceType) *MovePieceOutput {
	return &MovePieceOutput{
		Action:    "move-piece",
		Src:       src,
		Dst:       dst,
		PromoteTo: promoteTo,
	}
}

type MovePieceAction struct {
	ownRepository   domain.ConnectionRepository
	enemyRepository domain.ConnectionRepository
	game            *domain.Game
}

func NewMovePieceAction(ownRepository domain.ConnectionRepository, enemyRepository domain.ConnectionRepository,
	game *domain.Game) *MovePieceAction {
	return &MovePieceAction{
		ownRepository:   ownRepository,
		enemyRepository: enemyRepository,
		game:            game,
	}
}

func (uc *MovePieceAction) Invoke(p *MovePieceParams) error {
	move := &domain.Move{
		From: p.Src,
		To:   p.Dst,
	}

	uc.game.Move(move, p.PromoteTo)
	output := newMovePieceOutput(p.Src, p.Dst, p.PromoteTo)

	err := uc.ownRepository.SendWebSocketMessage(output)
	if err != nil {
		return nil
	}

	time.Sleep(100 * time.Millisecond)
	return uc.enemyRepository.SendWebSocketMessage(output)
}

//
//
//
//
//
//
//
//
//
//
//
//
// type RequestMovePiece struct {
// 	Src       *domain.Position  `json:"src"`
// 	Dst       *domain.Position  `json:"dst"`
// 	PromoteTo *domain.PieceType `json:"promoteTo"`
// }

// type ResponseMovePiece struct {
// 	Action    string            `json:"action"`
// 	Src       *domain.Position  `json:"src"`
// 	Dst       *domain.Position  `json:"dst"`
// 	PromoteTo *domain.PieceType `json:"promoteTo"`
// }

// func WsMovePiece(g *domain.Game, body []byte, player, enemy *shared.WsConn) {
// 	log.Println("handle move piece")
// 	req := RequestMovePiece{}
// 	json.Unmarshal(body, &req)

// 	move := &domain.Move{
// 		From: req.Src,
// 		To:   req.Dst,
// 	}
// 	g.Move(move, req.PromoteTo)
// 	resp := ResponseMovePiece{
// 		Action:    "move-piece",
// 		Src:       req.Src,
// 		Dst:       req.Dst,
// 		PromoteTo: req.PromoteTo,
// 	}

// 	player.WriteJSON(resp)
// 	time.Sleep(100 * time.Millisecond)
// 	enemy.WriteJSON(resp)
// }
