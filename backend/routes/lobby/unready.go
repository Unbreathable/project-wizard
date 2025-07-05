package lobby_routes

import (
	"github.com/Liphium/project-wizard/backend/integration"
	"github.com/Liphium/project-wizard/backend/service"
	"github.com/gofiber/fiber/v2"
)

type LobbyUnreadyRequest struct {
	LobbyId  string `json:"lobby_id" validate:"required"`
	PlayerId string `json:"player_id" validate:"required"`
}

// Route: /lobby/unready
func unreadyLobby(c *fiber.Ctx) error {
	var req LobbyUnreadyRequest
	if err := c.BodyParser(&req); err != nil {
		return integration.InvalidRequest(c, "request is invalid")
	}
	if err := service.Validate.Struct(req); err != nil {
		return integration.InvalidRequest(c, "request format is invalid")
	}

	lobby, ok := service.GetLobby(req.LobbyId)
	if !ok {
		return integration.InvalidRequest(c, "invalid lobby id")
	}

	if err := lobby.SetReadyPlayerById(req.PlayerId, false); err != nil {
		return integration.InvalidRequest(c, "invalid player id")
	}
	return integration.SuccessfulRequest(c)
}
