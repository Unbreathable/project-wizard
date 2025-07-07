package lobby_routes

import (
	"github.com/Liphium/project-wizard/backend/integration"
	"github.com/Liphium/project-wizard/backend/service"
	"github.com/gofiber/fiber/v2"
)

type LobbyUnreadyRequest struct {
	LobbyId  string `json:"lobby_id" validate:"required"`
	PlayerId string `json:"player_id" validate:"required"`
	Token    string `json:"token" validate:"required"`
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

	if lobby.IsRunning() {
		return integration.InvalidRequest(c, "game is running")
	}

	player := lobby.GetPlayer(req.PlayerId)

	// verify player token
	if player.GetInfo().Token != req.Token {
		return integration.InvalidRequest(c, "bad token")
	}

	if !player.GetInfo().Ready {
		return integration.InvalidRequest(c, "player isnt ready")
	}

	// unready player
	player.SetReady(false)

	// Send lobby change event to players
	data, err := GetLobbyInfo(req.LobbyId)
	if err != nil {
		return integration.InvalidRequest(c, err.Error())
	}
	service.Instance.Send(lobby.GetPlayersTokens(), LobbyChangeEvent(data))

	return integration.SuccessfulRequest(c)
}
