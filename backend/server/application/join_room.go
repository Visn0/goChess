package application

import (
	"chess/server/domain"
	"chess/server/shared"
	"chess/server/shared/chesserror"
	"fmt"
	"log"
)

type JoinRoomParams struct {
	RoomID   string `json:"roomID"`
	PlayerID string `json:"playerID"`
	Password string `json:"password"`
}

type JoinRoomOutput struct {
	HTTPCode int                 `json:"httpCode"`
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

func newJoinRoomOutput(httpCode int, roomID string, player1, player2 *domain.Player) *JoinRoomOutput {
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
		HTTPCode: httpCode,
		Action:   "join-room",
		Room: &JoinRoomInfoOutput{
			ID:      roomID,
			Players: players,
		},
	}
}

type JoinRoomAction struct {
	rm *domain.RoomManager
	c  domain.ConnectionRepository
}

func NewJoinRoomAction(rm *domain.RoomManager, c domain.ConnectionRepository) *JoinRoomAction {
	return &JoinRoomAction{rm: rm, c: c}
}

func (uc *JoinRoomAction) Invoke(p *JoinRoomParams) (*domain.Room, error) {
	log.Println("==> Join room params: ", shared.ToJSONString(p))
	room, ok := uc.rm.GetRoom(p.RoomID)
	if !ok {
		err := chesserror.NewError(chesserror.ResourceNotFound,
			fmt.Sprintf("Room with id '%s' not found", p.RoomID))
		log.Println(err)
		return nil, err
	}

	player := domain.NewPlayer(uc.c, p.PlayerID)
	err := room.AddPlayer(player)
	if err != nil {
		return nil, err
	}

	output := newJoinRoomOutput(200, p.RoomID, room.Player1, room.Player2)
	log.Println("##> Join room output: ", shared.ToJSONString(output))

	err = uc.c.SendWebSocketMessage(output)
	if err != nil {
		return nil, err
	}

	return room, nil
}
