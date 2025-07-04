package magic_scripts

import (
	"log"

	"github.com/Liphium/magic/mconfig"
	"github.com/Liphium/project-wizard/backend/service"
	"github.com/Liphium/project-wizard/backend/util/requests"
)

func RunLobbyInfo(p *mconfig.Plan, lobbyID string) {
	res, err := requests.PostRequestURL("http://localhost:"+p.Environment["PORT"]+"/lobby/info", requests.Map{"lobby_id": lobbyID})
	if err != nil {
		log.Fatalln("couldn't do lobby info request:", err)
	}
	if !requests.ValueOr(res, "success", false) {
		log.Fatalln("request error:", requests.ValueOr(res, "message", "?"))
	}
	player1 := requests.ValueOr(res, "player_1", service.Player{})
	player2 := requests.ValueOr(res, "player_2", service.Player{})
	log.Printf("Player 1 name: %s", player1.Name)
	log.Printf("Player 1 id: %s", player1.ID)
	log.Printf("Player 2 name: %s", player2.Name)
	log.Printf("Player 2 id: %s", player2.ID)
}
