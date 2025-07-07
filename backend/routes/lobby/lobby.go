package lobby_routes

import (
	"fmt"

	"github.com/Liphium/project-wizard/backend/service"
	"github.com/Liphium/project-wizard/neogate"
	"github.com/gofiber/fiber/v2"
)

type LobbyInfoEventData struct {
	Teams           []TeamInfo `json:"teams"`
	CharacterAmount int        `json:"character_amount"`
	Running         bool       `json:"running"`
}

type TeamInfo struct {
	Id      string               `json:"id"`
	Size    int                  `json:"size"`
	Players []service.PlayerInfo `json:"players"`
}

func LobbyRoutes(router fiber.Router) {
	router.Post("/create", createLobby)
	router.Post("/join", joinLobby)
	router.Post("/ready", readyLobby)
	router.Post("/unready", unreadyLobby)
}

// Event for game start
func LobbyChangeEvent(data LobbyInfoEventData) neogate.Event {
	return neogate.Event{
		Name: "lobby_change",
		Data: fiber.Map{
			"players":          data.Teams,
			"character_amount": data.CharacterAmount,
			"running":          data.Running,
		},
	}
}

// Get lobby info
func GetLobbyInfo(lobbyId string) (LobbyInfoEventData, error) {
	lobby, ok := service.GetLobby(lobbyId)
	if !ok {
		return LobbyInfoEventData{}, fmt.Errorf("invalid lobby id")
	}

	teams := []TeamInfo{}

	for _, t := range lobby.GetTeams() {
		players := []service.PlayerInfo{}
		for _, p := range t.GetPlayers() {
			players = append(players, p.GetInfo())
		}
		team := TeamInfo{
			Id:      t.GetId(),
			Size:    t.GetSize(),
			Players: players,
		}
		teams = append(teams, team)
	}

	return LobbyInfoEventData{
		Teams:           teams,
		CharacterAmount: service.CharacterAmount,
		Running:         lobby.IsRunning(),
	}, nil
}
