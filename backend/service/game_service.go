package service

import (
	"fmt"
	"slices"
	"sync"

	"github.com/Liphium/project-wizard/backend/game"
)

const OversightsPerTurn = 1
const NormalActionsPerTurn = 1

type Fight struct {
	playerActions map[string][]game.GameAction
	playerSwaps   map[string][]int
}

type Game struct {
	mutex        *sync.Mutex
	relatedLobby *Lobby

	playerActions map[string][]game.GameAction
	playerSwaps   map[string][]int
}

// Are all players ready for next turn
func (g *Game) IsReady() bool {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	for _, p := range g.relatedLobby.GetPlayers() {
		if !p.GetInfo().ReadyTurn {
			return false
		}
	}
	return true
}

// Is the player ready for the next turn
func (g *Game) IsPlayerReady(playerID string) bool {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	return g.relatedLobby.GetPlayer(playerID).GetInfo().ReadyTurn
}

// Removes the players actions
func (g *Game) RemovePlayerActions(playerId string) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.playerActions[playerId] = []game.GameAction{}
	g.playerSwaps[playerId] = []int{}
}

// Verify the a player's actions to make sure they are valid in the simulation.
func (g *Game) VerifyPlayerActions(playerId string, actions []game.GameAction, swap []int) bool {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	player := g.relatedLobby.GetPlayer(playerId)

	// Check swap amount
	chars := player.GetGamePlayerState().Characters
	characterLen := len(chars)
	if len(swap) == 2 {
		if swap[0] == swap[1] || swap[0] > characterLen-1 || swap[1] > characterLen-1 {
			return false
		}
		if chars[swap[0]].IsDead() || chars[swap[1]].IsDead() {
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
		character := player.GetGamePlayer().GetCharacterById(action.CharacterId)
		if character == nil || len(character.Actions) < int(action.ActionId)+1 || character.IsDead() {
			return false
		}

		if slices.ContainsFunc(character.StatusEffects, func(effect game.StatusEffect) bool {
			return effect.ID == "stun"
		}) {
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
		tt := g.relatedLobby.GetTeam(action.TargetTeam)
		if tt == nil {
			return false
		}
		targetPlayer := tt.GetPlayer(action.Target)
		if targetPlayer == nil {
			return false
		}

		targetChars := targetPlayer.GetGamePlayer().Characters
		if len(targetChars) < action.Slot+1 || targetChars[action.Slot].IsDead() {
			return false
		}

	}

	// Check mana and action type
	if neededMana > player.GetGamePlayerState().Mana || oversights > OversightsPerTurn || normal > NormalActionsPerTurn {
		return false
	}

	return true
}

// Current implementation only for 1vs1
func (g *Game) StartTurn() error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	switch g.relatedLobby.GetInfo().Mode {
	case LobbyMode1vs1:
		teams := g.relatedLobby.GetTeams()
		if len(teams) != 2 {
			return fmt.Errorf("bad teams")
		}

		var p1 *Player
		var p2 *Player

		t1 := teams[0]
		t2 := teams[1]

		if len(t1.GetPlayers()) != 1 || len(t2.GetPlayers()) != 1 {
			return fmt.Errorf("bad players")
		}

		p1 = t1.GetPlayers()[0]
		p2 = t2.GetPlayers()[0]

		err := game.RunSimulation([]*game.GamePlayer{p1.GetGamePlayer(), p2.GetGamePlayer()}, g.playerActions, g.playerSwaps)
		if err != nil {
			return err
		}

		// Send result of turn to clients
		Instance.Send(g.relatedLobby.GetSpectator(), GameUpdateEvent(SimulationResultEvent{
			Swaps:   g.playerSwaps,
			Actions: g.playerActions,
			Result: map[string]SimulationResult{
				p1.GetInfo().Id: {
					Mana:       p1.GetGamePlayer().Mana,
					ID:         p1.GetInfo().Id,
					Characters: p1.GetGamePlayer().GetCharacters(),
				},
				p2.GetInfo().Id: {
					Mana:       p2.GetGamePlayer().Mana,
					ID:         p2.GetInfo().Id,
					Characters: p2.GetGamePlayer().GetCharacters(),
				},
			},
		}))
	}
	return nil
}
