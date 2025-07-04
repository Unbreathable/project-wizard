package magic_scripts

import (
	"log"

	"github.com/Liphium/magic/mconfig"
	"github.com/Liphium/project-wizard/backend/util/requests"
)

func RunLobbyJoin(p *mconfig.Plan, username string, lobbyId string) {
	res, err := requests.PostRequestURL("http://localhost:"+p.Environment["PORT"]+"/lobby/join", requests.Map{"name": username, "lobby_id": lobbyId})
	if err != nil {
		log.Fatalln("couldn't do join lobby request:", err)
	}
	if !requests.ValueOr(res, "success", false) {
		log.Fatalln("request error:", requests.ValueOr(res, "message", "?"))
	}

	log.Printf("Player id: %s", requests.ValueOr(res, "player_id", "no id"))
}
