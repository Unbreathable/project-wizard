package lobby_routes

import (
	"github.com/Liphium/project-wizard/backend/integration"
	"github.com/Liphium/project-wizard/backend/service"
	"github.com/gofiber/fiber/v2"
)

type LobbyInfoRequest struct {
	LobbyId string `json:"lobby_id" validate:"required"`
}

type LobbyInfoResponse struct {
	Success bool   `json:"success"`
	Player1 string `json:"player_1"`
	Player2 string `json:"player_2"`
}

// Route: /lobby/info
func infoLobby(c *fiber.Ctx) error {
	var req LobbyInfoRequest
	if err := c.BodyParser(&req); err != nil {
		return integration.InvalidRequest(c, "request is invalid")
	}
	if err := service.Validate.Struct(req); err != nil {
		return integration.InvalidRequest(c, "request format is invalid")
	}

	lobby, ok := service.GetLobby(req.LobbyId)
	if !ok {
		return integration.InvalidRequest(c, "invalid id")
	}

	return c.JSON(LobbyInfoResponse{
		Success: true,
		Player1: lobby.GetPlayer1().Name,
		Player2: lobby.GetPlayer2().Name,
	})
}
