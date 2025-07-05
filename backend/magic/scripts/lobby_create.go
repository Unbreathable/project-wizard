package magic_scripts

import (
	"log"

	"github.com/Liphium/magic/mconfig"
	"github.com/Liphium/project-wizard/backend/util/requests"
)

func RunLobbyCreate(p *mconfig.Plan, username string) requests.Map {
	res, err := requests.PostRequestURL("http://localhost:"+p.Environment["PORT"]+"/lobby/create", requests.Map{"name": username})
	if err != nil {
		log.Fatalln("couldn't do create lobby request:", err)
	}
	if !requests.ValueOr(res, "success", false) {
		log.Fatalln("request error:", requests.ValueOr(res, "message", "?"))
	}

	lobbyId := requests.ValueOr(res, "lobby_id", "no id")
	playerId := requests.ValueOr(res, "player_id", "no id")

	log.Printf("Lobby id: %s", lobbyId)
	log.Printf("Player id: %s", playerId)
	return res
}
