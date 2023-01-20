package infrastructure

import (
	"chess/server/application"
	"chess/server/domain"
	"encoding/json"
	"log"
)

type CreateRoomWsController struct {
	uc *application.CreateRoomAction
}

func NewCreateRoomWsController(rm *domain.RoomManager, r domain.ConnectionRepository) *CreateRoomWsController {
	return &CreateRoomWsController{
		uc: application.NewCreateRoomAction(rm, r),
	}
}

func (c *CreateRoomWsController) Invoke(body []byte) (*domain.Room, error) {
	var p application.CreateRoomParams
	err := json.Unmarshal(body, &p)
	if err != nil {
		log.Println("Error unmarshalling request create room:", err)
		return nil, err
	}

	return c.uc.Invoke(&p)
}
