package lobby_routes

import (
	"github.com/Liphium/project-wizard/backend/game"
	"github.com/Liphium/project-wizard/backend/integration"
	"github.com/Liphium/project-wizard/backend/service"
	"github.com/gofiber/fiber/v2"
)

type LobbyReadyRequest struct {
	LobbyId    string `json:"lobby_id" validate:"required"`
	PlayerId   string `json:"player_id" validate:"required"`
	Token      string `json:"token" validate:"required"`
	Characters []uint `json:"character_ids" validate:"required"`
}

// Route: /lobby/ready
func readyLobby(c *fiber.Ctx) error {
	var req LobbyReadyRequest
	if err := c.BodyParser(&req); err != nil {
		return integration.InvalidRequest(c, "request is invalid")
	}
	if err := service.Validate.Struct(req); err != nil {
		return integration.InvalidRequest(c, "request format is invalid")
	}

	lobby, ok := service.GetLobby(req.LobbyId)
	if !ok {
		return integration.InvalidRequest(c, "invalid request id")
	}

	player := lobby.GetPlayer(req.PlayerId)

	// verify player token
	if player.GetInfo().Token != req.Token {
		return integration.InvalidRequest(c, "bad token")
	}

	// Check the validity of characters
	if len(req.Characters) != service.CharacterAmount {
		return integration.InvalidRequest(c, "invalid character amount")
	}
	if hasDuplicates(req.Characters) {
		return integration.InvalidRequest(c, "bad character selection")
	}

	if lobby.IsRunning() {
		return integration.InvalidRequest(c, "game is running")
	}

	if player.GetInfo().Ready {
		return integration.InvalidRequest(c, "player is ready")
	}

	//create game player
	gp := player.SetGamePlayer()

	// Check the validity of character ids and safe the pointers
	ptrChars := []*game.Character{}
	for _, char := range req.Characters {
		createFunc, ok := game.CharacterRegistry[char]
		if !ok {
			return integration.InvalidRequest(c, "bad character selection")
		}
		char := createFunc(gp)
		ptrChars = append(ptrChars, &char)
	}

	// Save chars
	gp.SetCharacters(ptrChars)

	// Ready player
	player.SetReady(true)

	// Send lobby change event to players
	data, err := GetLobbyInfo(req.LobbyId)
	if err != nil {
		return integration.InvalidRequest(c, err.Error())
	}
	service.Instance.Send(lobby.GetPlayersTokens(), LobbyChangeEvent(data))

	// start game
	if lobby.IsReady() {
		lobby.StartGame()
	}

	return integration.SuccessfulRequest(c)
}

func hasDuplicates(slice []uint) bool {
	seen := make(map[uint]bool)
	for _, value := range slice {
		if seen[value] {
			return true // Duplicate found
		}
		seen[value] = true
	}
	return false // No duplicates
}
