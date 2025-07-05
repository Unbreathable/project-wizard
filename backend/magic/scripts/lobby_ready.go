package magic_scripts

import (
	"log"

	"github.com/Liphium/magic/mconfig"
	"github.com/Liphium/project-wizard/backend/util/requests"
)

func RunLobbyReady(p *mconfig.Plan, playerId string, lobbyId string) requests.Map {
	res, err := requests.PostRequestURL("http://localhost:"+p.Environment["PORT"]+"/lobby/ready", requests.Map{"player_id": playerId, "lobby_id": lobbyId})
	if err != nil {
		log.Fatalln("couldn't do ready lobby request:", err)
	}
	if !requests.ValueOr(res, "success", false) {
		log.Fatalln("request error:", requests.ValueOr(res, "message", "?"))
	}
	return res
}
