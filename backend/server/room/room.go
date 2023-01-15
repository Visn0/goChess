package room

import (
	"chess/server/game"
	gameActions "chess/server/game/actions"
	"fmt"
	"log"
	"sync"

	"github.com/buger/jsonparser"
)

type Room struct {
	ID      string
	Player1 *Player
	Player2 *Player
	game    *game.Game
}

func NewRoom() *Room {
	return &Room{game: game.NewGame()}
}

func (r *Room) AddPlayer(p *Player) error {
	if r.Player1 == nil {
		r.Player1 = p
	} else if r.Player2 == nil {
		r.Player2 = p
	} else {
		return fmt.Errorf("Room is full")
	}
	return nil
}

func (r *Room) RemovePlayer(p *Player) error {
	if r.Player1 == p {
		r.Player1 = nil
		return nil
	} else if r.Player2 == p {
		r.Player2 = nil
		return nil
	}
	return fmt.Errorf("Player not found")
}

func (r *Room) GetRoomSize() int {
	size := 0
	if r.Player1 != nil {
		size++
	}
	if r.Player2 != nil {
		size++
	}
	return size
}

func (r *Room) HandleGame(isHost bool, roomsWG *sync.WaitGroup) {
	defer roomsWG.Done()

	log.Println("Room activated")
	var player *Player
	if isHost {
		player = r.Player1
	} else {
		player = r.Player2
	}

	for {
		if player == nil {
			return
		}

		messageType, message, err := player.Ws.ReadMessage()
		log.Println(messageType)
		if err != nil {
			log.Println("Some error:", err)
			player = nil
			return
		}
		log.Println("Get message.")

		reqAction, _ := jsonparser.GetString(message, "action")
		reqBody, _, _, _ := jsonparser.Get(message, "body")

		switch reqAction {
		case "request-moves":
			log.Println("Request moves")
			gameActions.WsGetValidMoves(r.game, reqBody, player.Ws)
		case "move-piece":
			log.Println("Move piece")
			gameActions.WsMovePiece(r.game, reqBody, player.Ws)
		default:
			log.Println("Unknown action")
		}
	}
}

type RoomPublicInfo struct {
	ID      string              `json:"id"`
	Players []*PlayerPublicInfo `json:"players"`
}

func (r *Room) GetPublicInfo() *RoomPublicInfo {
	players := make([]*PlayerPublicInfo, 0, 2)
	if r.Player1 != nil {
		info := r.Player1.GetPublicInfo()
		players = append(players, &PlayerPublicInfo{
			ID: info.ID,
		})
	}
	if r.Player2 != nil {
		info := r.Player2.GetPublicInfo()
		players = append(players, &PlayerPublicInfo{
			ID: info.ID,
		})
	}
	return &RoomPublicInfo{
		ID:      r.ID,
		Players: players,
	}
}