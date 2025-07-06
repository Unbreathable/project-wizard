package lobby_routes

import (
	"regexp"
	"slices"

	"github.com/Liphium/project-wizard/backend/integration"
	"github.com/Liphium/project-wizard/backend/service"
	"github.com/gofiber/fiber/v2"
)

type LobbyCreateRequest struct {
	Name string            `json:"name" validate:"required"`
	Mode service.LobbyMode `json:"mode" validate:"required"`
}

type LobbyCreateResponse struct {
	Success  bool              `json:"success"`
	LobbyId  string            `json:"lobby_id"`
	Mode     service.LobbyMode `json:"mode"`
	TeamId   string            `json:"team_id"`
	PlayerId string            `json:"player_id"`
	Token    string            `json:"token"`
}

// Route: /lobby/create
func createLobby(c *fiber.Ctx) error {
	var req LobbyCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return integration.InvalidRequest(c, "request is invalid")
	}
	if err := service.Validate.Struct(req); err != nil {
		return integration.InvalidRequest(c, "request format is invalid")
	}

	if !regexp.MustCompile("[A-Za-z0-9_-]{2,20}").MatchString(req.Name) {
		return integration.InvalidRequest(c, "request format is invalid")
	}

	// verify lobby mode
	if !slices.Contains(service.LobbyModes, req.Mode) {
		return integration.InvalidRequest(c, "request format is invalid")
	}

	// Create Lobby
	lobbyId := service.CreateLobby(req.Mode)
	lobby, _ := service.GetLobby(lobbyId)

	res := LobbyCreateResponse{
		Success:  true,
		LobbyId:  lobbyId,
		Mode:     req.Mode,
		TeamId:   "",
		PlayerId: "",
		Token:    "",
	}

	// Create teams and players according to mode
	switch req.Mode {
	case service.LobbyMode1vs1:

		// Create two teams with one player each
		team1 := lobby.NewTeam(1)
		lobby.NewTeam(1)

		player1 := lobby.NewPlayer(req.Name)

		team1.AddPlayer(player1) // error can be ignored because size is set here

		p1Info := player1.GetInfo()

		res.TeamId = team1.GetId()
		res.PlayerId = p1Info.Id
		res.Token = p1Info.Token
	}

	return c.JSON(res)
}
