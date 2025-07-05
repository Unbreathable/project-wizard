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

	// verify player token
	if lobby.GetPlayerTokenById(req.PlayerId) != req.Token {
		return integration.InvalidRequest(c, "bad token")
	}

	if err := lobby.SetReadyPlayerById(req.PlayerId, false); err != nil {
		return integration.InvalidRequest(c, "invalid player id")
	}

	// Send lobby join event to players
	data, err := GetLobbyInfo(req.LobbyId)
	if err != nil {
		return integration.InvalidRequest(c, err.Error())
	}
	p1, err := lobby.GetPlayer(1)
	if err != nil {
		return integration.InvalidRequest(c, "server error")
	}
	p2, err := lobby.GetPlayer(2)
	if err != nil {
		return integration.InvalidRequest(c, "server error")
	}
	service.Instance.Send([]string{p1.Token, p2.Token}, LobbyChangeEvent(data))

	return integration.SuccessfulRequest(c)
}
