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

	// verify player token
	if lobby.GetPlayerTokenById(req.PlayerId) != req.Token {
		return integration.InvalidRequest(c, "bad token")
	}

	// Check the validity of characters
	if len(req.Characters) != service.CharacterAmount {
		return integration.InvalidRequest(c, "invalid character amount")
	}
	if hasDuplicates(req.Characters) {
		return integration.InvalidRequest(c, "bad character selection")
	}

	//create game player
	gp, err := lobby.SetGamePlayerById(req.PlayerId)
	if err != nil {
		return integration.InvalidRequest(c, "server error")
	}

	// Check the validity of characters and safe the pointers
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
	lobby.SetPlayerCharsById(req.PlayerId, ptrChars)

	if err := lobby.SetReadyPlayerById(req.PlayerId, true); err != nil {
		return integration.InvalidRequest(c, "invalid player id")
	}

	if lobby.IsRunning() {
		return integration.InvalidRequest(c, "game is running")
	}

	p1, err := lobby.GetPlayer(1)
	if err != nil {
		return integration.InvalidRequest(c, "server error")
	}
	p2, err := lobby.GetPlayer(2)
	if err != nil {
		return integration.InvalidRequest(c, "server error")
	}

	// Send lobby join event to players
	data, err := GetLobbyInfo(req.LobbyId)
	if err != nil {
		return integration.InvalidRequest(c, err.Error())
	}
	service.Instance.Send([]string{p1.Token, p2.Token}, LobbyChangeEvent(data))

	if p1.Ready && p2.Ready {
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
