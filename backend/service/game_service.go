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

func (g *Game) IsReady() bool {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	for _, v := range g.playersReady {
		if !v {
			return v
		}
	}
	return true
}

func (g *Game) IsPlayerReady(playerID string) bool {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	ready, ok := g.playersReady[playerID]
	if !ok {
		return ok
	}
	return ready
}

func (g *Game) SetPlayerReady(playerID string, ready bool) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.playersReady[playerID] = ready
}

// Removes the players actions
func (g *Game) RemovePlayerActions(playerId string) {
	g.playerActions[playerId] = []game.GameAction{}
	g.playerSwaps[playerId] = []int{}
}

// Verify the a player's actions to make sure they are valid in the simulation.
func (g *Game) VerifyPlayerActions(playerId string, actions []game.GameAction, swap []int) bool {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	player, _, err := g.relatedLobby.GetPlayerById(playerId)
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

	neededMana := 0
	oversights := 0
	normal := 0
	for _, action := range actions {
		character := player.GamePlayer.GetCharacterById(action.CharacterId)
		if len(character.Actions) < int(action.ActionId)+1 || character.IsDead() {
			return false
		}
		attack := character.Actions[action.ActionId]
		if attack.Oversight {
			oversights++

		} else {
			normal++
		}
		neededMana += attack.ManaCost

		// Check if the selected slots are correct
		targetPlayer, _, err := g.relatedLobby.GetPlayerById(action.Target)
		if err != nil {
			return false
		}
		targetChars := targetPlayer.GamePlayer.Characters
		if len(targetChars) < action.Slot+1 || targetChars[action.Slot].IsDead() {
			return false
		}
	}

	// Check mana and action type
	if neededMana > player.GamePlayer.Mana || oversights > OversightsPerTurn || normal > NormalActionsPerTurn {
		return false
	}

	return true
}

func (g *Game) StartTurn() {
	p1, err := g.relatedLobby.GetPlayer(1)
	p2, err := g.relatedLobby.GetPlayer(2)
	err = game.RunSimulation([]*game.GamePlayer{p1.GamePlayer, p2.GamePlayer}, g.playerActions, g.playerSwaps)
	if err != nil {
		// TODO: error handling
	}

	charsP1 := []game.Character{}
	for _, c := range p1.GamePlayer.Characters {
		charsP1 = append(charsP1, *c)
	}

	charsP2 := []game.Character{}
	for _, c := range p2.GamePlayer.Characters {
		charsP2 = append(charsP2, *c)
	}

	// Send result of turn to clients
	Instance.Send([]string{p1.Token, p2.Token}, GameUpdateEvent(SimulationResultEvent{
		Swaps:   g.playerSwaps,
		Actions: g.playerActions,
		Result: map[string]SimulationResult{
			p1.ID: {
				Mana:       p1.GamePlayer.Mana,
				ID:         p1.ID,
				Characters: charsP1,
			},
			p2.ID: {
				Mana:       p2.GamePlayer.Mana,
				ID:         p2.ID,
				Characters: charsP2,
			},
		},
	}))
}
