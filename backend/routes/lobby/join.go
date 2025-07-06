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

	if lobby.IsFull() || lobby.IsRunning() {
		return integration.InvalidRequest(c, "lobby is full")
	}

	res := LobbyCreateResponse{
		Success:  true,
		LobbyId:  lobby.GetInfo().Id,
		Mode:     lobby.GetInfo().Mode,
		TeamId:   "",
		PlayerId: "",
		Token:    "",
	}

	switch lobby.GetInfo().Mode {
	case service.LobyMode1vs1:
		for _, v := range lobby.GetTeams() {
			if !v.IsFull() {
				p := lobby.NewPlayer(req.Name)
				v.AddPlayer(p) // error can be ignored because team is not full

				info := p.GetInfo()

				res.TeamId = v.GetId()
				res.PlayerId = info.Id
				res.Token = info.Token
				break
			}
		}
	}

	// TODO: add lobby info event

	// Send lobby join event to host
	data, err := GetLobbyInfo(req.LobbyId)
	if err != nil {
		return integration.InvalidRequest(c, err.Error())
	}
	service.Instance.SendOne(p1.Token, LobbyChangeEvent(data))

	return c.JSON()
}
