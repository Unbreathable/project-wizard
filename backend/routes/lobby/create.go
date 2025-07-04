package lobby_routes

import (
	"github.com/Liphium/project-wizard/backend/integration"
	"github.com/Liphium/project-wizard/backend/service"
	"github.com/gofiber/fiber/v2"
)

type LobbyCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type LobbyCreateResponse struct {
	Success  bool   `json:"success"`
	LobbyId  string `json:"lobby_id"`
	PlayerId string `json:"player_id"`
}

// Route: /lobby/create
func createLobby(c *fiber.Ctx) error {
	var req LobbyCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return integration.InvalidRequest(c, "request is invalid")
	}
	if err := service.Validate.Struct(req); err != nil {
		return integration.InvalidRequest(c, "request format is invalid")
	}

	lobbyId, playerId := service.CreateLobby(req.Name)

	return c.JSON(LobbyCreateResponse{
		Success:  true,
		LobbyId:  lobbyId.String(),
		PlayerId: playerId.String(),
	})
}
