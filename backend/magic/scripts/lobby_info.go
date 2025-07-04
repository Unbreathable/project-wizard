package magic_scripts

import (
	"log"

	"github.com/Liphium/magic/mconfig"
	lobby_routes "github.com/Liphium/project-wizard/backend/routes/lobby"
	"github.com/Liphium/project-wizard/backend/util/requests"
)

func RunLobbyInfo(p *mconfig.Plan, lobbyID string) {
	res, err := requests.PostRequestURLGeneric[lobby_routes.LobbyInfoResponse]("http://localhost:"+p.Environment["PORT"]+"/lobby/info", requests.Map{"lobby_id": lobbyID})
	if err != nil {
		log.Fatalln("couldn't do lobby info request:", err)
	}
	if !res.Success {
		log.Fatalln("request error")
	}
	log.Printf("Player 1 name: %s", res.Player1)
	log.Printf("Player 2 name: %s", res.Player2)
}
