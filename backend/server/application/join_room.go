package application

import (
	"chess/server/domain"
	"fmt"
	"log"
)

type JoinRoomParams struct {
	RoomID   string `json:"roomID"`
	PlayerID string `json:"playerID"`
	Password string `json:"password"`
}

type JoinRoomOutput struct {
	HttpCode int                 `json:"httpCode"`
	Action   string              `json:"action"`
	Room     *JoinRoomInfoOutput `json:"room"`
}
type JoinRoomInfoOutput struct {
	ID      string                  `json:"id"`
	Players []*JoinRoomPlayerOutput `json:"players"`
}

type JoinRoomPlayerOutput struct {
	ID string `json:"id"`
}

func newJoinRoomOutput(httpCode int, roomID string, player1 *domain.Player, player2 *domain.Player) *JoinRoomOutput {
	players := make([]*JoinRoomPlayerOutput, 0, 2)
	if player1 != nil {
		players = append(players, &JoinRoomPlayerOutput{
			ID: player1.ID,
		})
	}

	if player2 != nil {
		players = append(players, &JoinRoomPlayerOutput{
			ID: player2.ID,
		})
	}

	return &JoinRoomOutput{
		HttpCode: httpCode,
		Action:   "join-room",
		Room: &JoinRoomInfoOutput{
			ID:      roomID,
			Players: players,
		},
	}

}

type JoinRoomAction struct {
	rm *domain.RoomManager
	r  domain.ConnectionRepository
}

func NewJoinRoomAction(rm *domain.RoomManager, r domain.ConnectionRepository) *JoinRoomAction {
	return &JoinRoomAction{rm: rm, r: r}
}

func (uc *JoinRoomAction) Invoke(p *JoinRoomParams) (*domain.Room, error) {
	room, ok := uc.rm.GetRoom(p.RoomID)
	if !ok {
		err := fmt.Errorf("There is already a room with id %q", p.RoomID)
		log.Println(err)
		return nil, err
	}
	player := domain.NewPlayer(uc.r.GetWebSocketConnection(), p.PlayerID, domain.BLACK)
	room.AddPlayer(player)
	log.Println("Player joined room", player.ID, room.ID)

	output := newJoinRoomOutput(200, p.RoomID, room.Player1, room.Player2)
	err := uc.r.SendWebSocketMessage(output)
	if err != nil {
		return nil, err
	}

	return room, nil
}
