package game_routes

import (
	"github.com/Liphium/project-wizard/backend/game"
	"github.com/Liphium/project-wizard/backend/integration"
	"github.com/Liphium/project-wizard/backend/service"
	"github.com/gofiber/fiber/v2"
)

type GameRemoveTurnRequest struct {
	LobbyId  string `json:"lobby_id" validate:"required"`
	PlayerId string `json:"player_id" validate:"required"`
	Token    string `json:"token" validate:"required"`

	TurnActions []game.GameAction `json:"turn_actions" validate:"required"`
	TurnSwap    []int             `json:"turn_swap" validate:"required"`
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
	if lobby.GetPlayer(req.PlayerId).GetInfo().Token != req.Token {
		return integration.InvalidRequest(c, "bad token")
	}

	game := lobby.GetGame()
	if game == nil {
		return integration.InvalidRequest(c, "no game")
	}

	if game.IsReady() {
		return integration.InvalidRequest(c, "turn is running")
	}

	if !game.IsPlayerReady(req.PlayerId) {
		return integration.InvalidRequest(c, "already not ready")
	}

	// Remove players actions
	game.RemovePlayerActions(req.PlayerId)

	// Unready player
	lobby.GetPlayer(req.PlayerId).SetReadyTurn(false)

	// Collect player status for event
	playerReady := []service.PlayerReady{}
	for _, v := range lobby.GetPlayers() {
		info := v.GetInfo()
		playerReady = append(playerReady, service.PlayerReady{
			Id:    info.Id,
			Ready: info.ReadyTurn,
		})
	}

	// Send game status change event to players
	service.Instance.Send(lobby.GetSpectator(), service.GameInfoEvent(playerReady))

	return integration.SuccessfulRequest(c)
}
