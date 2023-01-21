package infrastructure

import (
	"chess/server/application"
	"chess/server/domain"
	"encoding/json"
	"log"
)

type JoinRoomWsController struct {
	uc *application.JoinRoomAction
}

func NewJoinRoomWsController(rm *domain.RoomManager, r domain.ConnectionRepository) *JoinRoomWsController {
	return &JoinRoomWsController{
		uc: application.NewJoinRoomAction(rm, r),
	}
}

func (c *JoinRoomWsController) Invoke(body []byte) (*domain.Room, error) {
	var p application.JoinRoomParams
	err := json.Unmarshal(body, &p)
	if err != nil {
		log.Println("Error unmarshalling request create room:", err)
		return nil, err
	}

	return c.uc.Invoke(&p)
}
