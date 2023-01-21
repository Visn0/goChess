package infrastructure

import (
	"chess/server/application"
	"chess/server/domain"
	"chess/server/shared/chesserror"
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

	room, err := c.uc.Invoke(&p)
	if err != nil {
		return nil, chesserror.WithAction(err, "create-room")
	}

	return room, nil
}
