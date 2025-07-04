package lobby_routes

import (
	"github.com/Liphium/project-wizard/backend/integration"
	"github.com/Liphium/project-wizard/backend/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type LobbyInfoRequest struct {
	LobbyId string `json:"lobby_id"`
}

type LobbyInfoResponse struct {
	Success bool           `json:"success"`
	Player1 service.Player `json:"player_1"`
	Player2 service.Player `json:"player_2"`
}

// Route: /lobby/info
func infoLobby(c *fiber.Ctx) error {
	var req LobbyInfoRequest
	if err := c.BodyParser(&req); err != nil {
		return integration.InvalidRequest(c, "request is invalid")
	}

	lobbyId, err := uuid.Parse(req.LobbyId)
	if err != nil {
		return integration.InvalidRequest(c, "request invalid")
	}

	lobby, ok := service.GetLobby(lobbyId)
	if !ok {
		return integration.InvalidRequest(c, "invalid id")
	}

	return c.JSON(LobbyInfoResponse{
		Success: true,
		Player1: lobby.Player1,
		Player2: lobby.Player2,
	})
}
