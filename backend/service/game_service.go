package service

import (
	"sync"
)

const OversightsPerTurn = 1
const NormalActionsPerTurn = 1

type GameAction struct {
	ActionId   uint
	Slot1      int
	Slot2Owner string
	Slot2      int
}

type Game struct {
	mutex        sync.Mutex
	relatedLobby *Lobby

	playersReady  map[string]bool
	playerActions map[string][]GameAction
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

func (game *Game) verifyPlayerActions(playerId string) bool {
	game.mutex.Lock()
	defer game.mutex.Unlock()

	player, _, err := game.relatedLobby.GetPlayerById(playerId)
	if err != nil {
		return false
	}

	chars := player.GamePlayer.Characters

	// Check swap amount
	swaps := game.playerSwaps[playerId]
	if (len(swaps)%2) != 0 || len(swaps) > 2 {
		return false
	}

	// Check swap position
	if swaps[0] == swaps[1] || swaps[0] > len(chars)-1 || swaps[1] > len(chars)-1 {
		return false
	}

	//

	return true

}
