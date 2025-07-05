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
	Success         bool           `json:"success"`
	Player1         service.Player `json:"player_1"`
	Player2         service.Player `json:"player_2"`
	CharacterAmount int            `json:"character_amount"`
	Running         bool           `json:"running"`
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
		return integration.InvalidRequest(c, "invalid lobby id")
	}

	p1, err := lobby.GetPlayer(1)
	if err != nil {
		return integration.InvalidRequest(c, "server error")
	}
	p2, err := lobby.GetPlayer(2)
	if err != nil {
		return integration.InvalidRequest(c, "server error")
	}

	return c.JSON(LobbyInfoResponse{
		Success:         true,
		Player1:         p1,
		Player2:         p2,
		CharacterAmount: service.CharacterAmount,
		Running:         lobby.IsRunning(),
	})
}
