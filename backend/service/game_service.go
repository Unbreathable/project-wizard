package service

import (
	"sync"

	"github.com/Liphium/project-wizard/backend/game"
)

const OversightsPerTurn = 1
const NormalActionsPerTurn = 1

type Game struct {
	mutex        sync.Mutex
	relatedLobby *Lobby

	playersReady  map[string]bool
	playerActions map[string][]game.GameAction
	playerSwaps   map[string][]int
}

func (game *Game) IsReady() bool {
	game.mutex.Lock()
	defer game.mutex.Unlock()

	for _, v := range game.playersReady {
		if !v {
			return v
		}
	}
	return true
}

// Verify the a player's actions to make sure they are valid in the simulation.
func (game *Game) VerifyPlayerActions(playerId string) bool {
	game.mutex.Lock()
	defer game.mutex.Unlock()

	player, _, err := game.relatedLobby.GetPlayerById(playerId)
	if err != nil {
		return false
	}

	// Check swap amount
	characterLen := len(player.GamePlayer.Characters)
	swaps := game.playerSwaps[playerId]
	if len(swaps) == 2 {
		if swaps[0] == swaps[1] || swaps[0] > characterLen-1 || swaps[1] > characterLen-1 {
			return false
		}
		if player.GamePlayer.Characters[swaps[0]].IsDead() || player.GamePlayer.Characters[swaps[1]].IsDead() {
			return false
		}
	}
	if len(swaps) != 2 && len(swaps) != 0 {
		return false
	}

	return true

}
