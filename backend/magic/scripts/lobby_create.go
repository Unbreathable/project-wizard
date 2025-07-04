package magic_scripts

import (
	"log"

	"github.com/Liphium/magic/mconfig"
	"github.com/Liphium/project-wizard/backend/util/requests"
)

func RunLobbyCreate(p *mconfig.Plan) {
	res, err := requests.PostRequestURL("http://localhost:"+p.Environment["PORT"]+"/lobby/create", requests.Map{})
	if err != nil {
		log.Fatalln("couldn't do refresh request:", err)
	}
	if !requests.ValueOr(res, "success", false) {
		log.Fatalln("request error:", requests.ValueOr(res, "message", "?"))
	}
}
