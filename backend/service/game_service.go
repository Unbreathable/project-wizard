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

func (game *Game) IsPlayerReady(playerID string) bool {
	game.mutex.Lock()
	defer game.mutex.Unlock()

	ready, ok := game.playersReady[playerID]
	if !ok {
		return ok
	}
	return ready
}

func (game *Game) SetPlayerReady(playerID string, ready bool) {
	game.mutex.Lock()
	defer game.mutex.Unlock()

	game.playersReady[playerID] = ready
}

// Removes the players actions
func (game *Game) RemovePlayerActions(playerId string) {
	game.playerActions[playerId] = []game.GameAction{}
	game.playerSwaps[playerId] = []int{}
}

// Verify the a player's actions to make sure they are valid in the simulation.
func (game *Game) VerifyPlayerActions(playerId string, actions []game.GameAction, swap []int) bool {
	game.mutex.Lock()
	defer game.mutex.Unlock()

	player, _, err := game.relatedLobby.GetPlayerById(playerId)
	if err != nil {
		return false
	}

	// Check swap amount
	characterLen := len(player.GamePlayer.Characters)
	if len(swap) == 2 {
		if swap[0] == swap[1] || swap[0] > characterLen-1 || swap[1] > characterLen-1 {
			return false
		}
		if player.GamePlayer.Characters[swap[0]].IsDead() || player.GamePlayer.Characters[swap[1]].IsDead() {
			return false
		}
	}
	if len(swap) != 2 && len(swap) != 0 {
		return false
	}

	// TODO: verify actions

	// Set player status ready

	return true

}
