package server

import (
	"chess/server/game"
	"fmt"
	"log"

	"github.com/buger/jsonparser"
)

type Room struct {
	player1 *Player
	player2 *Player
	game    *game.Game
}

func NewRoom() *Room {
	return &Room{game: game.NewGame()}
}

func (r *Room) AddPlayer(p *Player) error {
	if r.player1 == nil {
		r.player1 = p
	} else if r.player2 == nil {
		r.player2 = p
	} else {
		return fmt.Errorf("Room is full")
	}
	return nil
}

func (r *Room) RemovePlayer(p *Player) error {
	if r.player1 == p {
		r.player1 = nil
		return nil
	} else if r.player2 == p {
		r.player2 = nil
		return nil
	}
	return fmt.Errorf("Player not found")
}

func (r *Room) GetRoomSize() int {
	size := 0
	if r.player1 != nil {
		size++
	}
	if r.player2 != nil {
		size++
	}
	return size
}

func (r *Room) HandleGame() {
	log.Println("Room activated")
	for {
		if r.player1 == nil {
			continue
		}
		messageType, message, err := r.player1.ws.ReadMessage()
		log.Println(messageType)
		if err != nil {
			log.Println("Some error:", err)
			r.player1 = nil
			return
		}
		log.Println("Get message.")

		reqAction, _ := jsonparser.GetString(message, "action")
		reqBody, _, _, _ := jsonparser.Get(message, "body")

		switch reqAction {
		case "request-moves":
			log.Println("Request moves")
			r.handleRequestMoves(reqBody, r.player1.ws)
		case "move-piece":
			log.Println("Move piece")
			r.handleMovePiece(reqBody, r.player1.ws)
		}
	}
}
