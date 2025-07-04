package lobby_routes

import (
	"regexp"

	"github.com/Liphium/project-wizard/backend/integration"
	"github.com/Liphium/project-wizard/backend/service"
	"github.com/gofiber/fiber/v2"
)

type LobbyJoinRequest struct {
	LobbyId string `json:"lobby_id" validate:"required"`
	Name    string `json:"name" validate:"required"`
}

type LobbyJoinResponse struct {
	Success  bool   `json:"success"`
	PlayerId string `json:"player_id"`
}

// Route: /lobby/join
func joinLobby(c *fiber.Ctx) error {
	var req LobbyJoinRequest
	if err := c.BodyParser(&req); err != nil {
		return integration.InvalidRequest(c, "request is invalid")
	}
	if err := service.Validate.Struct(req); err != nil {
		return integration.InvalidRequest(c, "request format is invalid")
	}

	if !regexp.MustCompile("[A-Za-z0-9_-]{2,20}").MatchString(req.Name) {
		return integration.InvalidRequest(c, "request format is invalid")
	}

	lobby, ok := service.GetLobby(req.LobbyId)
	if !ok {
		return integration.InvalidRequest(c, "invalid id")
	}

	if lobby.IsFull() {
		return integration.InvalidRequest(c, "lobby is full")
	}

	lobby.SetNamePlayer2(req.Name)

	return c.JSON(LobbyJoinResponse{
		Success:  true,
		PlayerId: lobby.GetPlayer2().ID,
	})
}
