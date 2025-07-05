package game_routes

import (
	"github.com/Liphium/project-wizard/backend/integration"
	"github.com/Liphium/project-wizard/backend/service"
	"github.com/gofiber/fiber/v2"
)

type GameRemoveTurnRequest struct {
	LobbyId  string `json:"lobby_id" validate:"required"`
	PlayerId string `json:"player_id" validate:"required"`
	Token    string `json:"token" validate:"required"`

	TurnActions []service.GameAction `json:"turn_actions" validate:"required"`
	TurnSwap    []int                `json:"turn_swap" validate:"required"`
}

// Route: /game/remove_turn
func removeTurnGame(c *fiber.Ctx) error {
	var req GameTurnRequest
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

	// verify player token
	if lobby.GetPlayerTokenById(req.PlayerId) != req.Token {
		return integration.InvalidRequest(c, "bad token")
	}

	game := lobby.GetGame()

	if game.IsReady() {
		return integration.InvalidRequest(c, "turn is running")
	}

	if !game.IsPlayerReady(req.PlayerId) {
		return integration.InvalidRequest(c, "already not ready")
	}

	// Remove players actions
	game.RemovePlayerActions(req.PlayerId)

	// Unready player
	game.SetPlayerReady(req.PlayerId, false)

	return integration.SuccessfulRequest(c)
}
