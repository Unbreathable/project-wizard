package game_routes

import (
	"github.com/Liphium/project-wizard/backend/game"
	"github.com/Liphium/project-wizard/backend/integration"
	"github.com/Liphium/project-wizard/backend/service"

	"github.com/gofiber/fiber/v2"
)

type GameTurnRequest struct {
	LobbyId  string `json:"lobby_id" validate:"required"`
	PlayerId string `json:"player_id" validate:"required"`
	Token    string `json:"token" validate:"required"`

	TurnActions []game.GameAction `json:"turn_actions" validate:"required"`
	TurnSwap    []int             `json:"turn_swap" validate:"required"`
}

// Route: /game/turn
func turnGame(c *fiber.Ctx) error {
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
	if game == nil {
		return integration.InvalidRequest(c, "no game")
	}

	if game.IsReady() {
		return integration.InvalidRequest(c, "turn is running")
	}

	if game.IsPlayerReady(req.PlayerId) {
		return integration.InvalidRequest(c, "already ready")
	}

	// Verify Actions and swaps are possible
	if !game.VerifyPlayerActions(req.PlayerId, req.TurnActions, req.TurnSwap) {
		return integration.InvalidRequest(c, "bad actions or swaps")
	}

	// Ready player
	game.SetPlayerReady(req.PlayerId, true)

	return integration.SuccessfulRequest(c)
}
