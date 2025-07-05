package lobby_routes

import (
	"fmt"

	"github.com/Liphium/project-wizard/backend/service"
	"github.com/Liphium/project-wizard/neogate"
	"github.com/gofiber/fiber/v2"
)

type LobbyInfoEventData struct {
	Player1         service.Player `json:"player_1"`
	Player2         service.Player `json:"player_2"`
	CharacterAmount int            `json:"character_amount"`
	Running         bool           `json:"running"`
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
			"player_1":         data.Player1,
			"player_2":         data.Player2,
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

	p1, err := lobby.GetPlayer(1)
	if err != nil {
		return LobbyInfoEventData{}, fmt.Errorf("server error")
	}
	p2, err := lobby.GetPlayer(2)
	if err != nil {
		return LobbyInfoEventData{}, fmt.Errorf("server error")
	}

	return LobbyInfoEventData{
		Player1:         p1,
		Player2:         p2,
		CharacterAmount: service.CharacterAmount,
		Running:         lobby.IsRunning(),
	}, nil
}
