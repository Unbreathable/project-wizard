package magic_tests

import (
	"testing"

	"github.com/Liphium/magic/mconfig"
	magic_scripts "github.com/Liphium/project-wizard/backend/magic/scripts"
	magic_util "github.com/Liphium/project-wizard/backend/magic/scripts/util"
	"github.com/Liphium/project-wizard/backend/util/requests"
)

// Do not call this function anything with Test, it will cause errors
func MagicTestLobby(t *testing.T, p *mconfig.Plan) {
	lobby1Username1 := "lobby1julian"
	lobby1Username2 := "lobby1jan"
	lobby2Username1 := "lobby2"

	lobbyId1 := ""
	playerId2 := ""

	// Create 2 lobbies
	t.Run("creating lobbies", func(t *testing.T) {
		res := magic_scripts.RunLobbyCreate(p, lobby1Username1)
		magic_util.AssertEq(t, requests.ValueOr(res, "success", false), true)

		lobbyId1 = requests.ValueOr(res, "lobby_id", "no id")
		requests.ValueOr(res, "player_id", "no id")

		res = magic_scripts.RunLobbyCreate(p, lobby2Username1)
		magic_util.AssertEq(t, requests.ValueOr(res, "success", false), true)
	})

	// Join lobby 1
	t.Run("joining a lobby", func(t *testing.T) {
		res := magic_scripts.RunLobbyJoin(p, lobby1Username2, lobbyId1)
		magic_util.AssertEq(t, requests.ValueOr(res, "success", false), true)

		playerId2 = requests.ValueOr(res, "player_id", "no id")
	})

	// Ready player 2
	t.Run("ready a player", func(t *testing.T) {
		magic_scripts.RunLobbyReady(p, playerId2, lobbyId1)
	})

	// Unready player 2
	t.Run("ready a player", func(t *testing.T) {
		magic_scripts.RunLobbyUnready(p, playerId2, lobbyId1)
	})
}
