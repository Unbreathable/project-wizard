package magic_scripts

import (
	"log"

	"github.com/Liphium/magic/mconfig"
	"github.com/Liphium/project-wizard/backend/util/requests"
)

func RunLobbyCreate(p *mconfig.Plan) {
	res, err := requests.PostRequestURL("http://localhost:"+p.Environment["PORT"]+"/lobby/create", requests.Map{})
	if err != nil {
		log.Fatalln("couldn't do create lobby request:", err)
	}
	if !requests.ValueOr(res, "success", false) {
		log.Fatalln("request error:", requests.ValueOr(res, "message", "?"))
	}

	log.Printf("Lobby id: %s", requests.ValueOr(res, "lobby_id", "no id"))
	log.Printf("Player id: %s", requests.ValueOr(res, "player_id", "no id"))
}
