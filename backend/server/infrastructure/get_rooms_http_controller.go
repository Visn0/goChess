package infrastructure

import (
	"chess/server/application"
	"chess/server/domain"

	"github.com/gofiber/fiber/v2"
)

type GetRoomsHttpController struct {
	uc *application.GetRoomsAction
}

func NewGetRoomsHttpController(rm *domain.RoomManager) *GetRoomsHttpController {
	return &GetRoomsHttpController{
		uc: application.NewGetRoomsAction(rm),
	}
}

func (c *GetRoomsHttpController) Invoke(ctx *fiber.Ctx) error {
	return ctx.JSON(c.uc.Invoke())
}
