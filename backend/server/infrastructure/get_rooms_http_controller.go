package infrastructure

import (
	"chess/server/application"
	"chess/server/domain"

	fiber "github.com/gofiber/fiber/v2"
)

type GetRoomsHTTPController struct {
	uc *application.GetRoomsAction
}

func NewGetRoomsHTTPController(rm *domain.RoomManager) *GetRoomsHTTPController {
	return &GetRoomsHTTPController{
		uc: application.NewGetRoomsAction(rm),
	}
}

func (c *GetRoomsHTTPController) Invoke(ctx *fiber.Ctx) error {
	return ctx.JSON(c.uc.Invoke())
}
