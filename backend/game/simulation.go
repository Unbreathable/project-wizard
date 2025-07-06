package game

import (
	"errors"
)

type GameAction struct {
	CharacterId uint   `json:"char_id" validate:"required"`
	ActionId    uint   `json:"action_id" validate:"required"`
	Target      string `json:"target" validate:"required"`      // Targetted player id
	TargetTeam  string `json:"target_team" validate:"required"` // Targetted team id
	Slot        int    `json:"slot" validate:"required"`        // Targetted slot id
}

func RunSimulation(players []*GamePlayer, actions map[string][]GameAction, swaps map[string][]int) error {

	// Convert all the actions made by the player to actual actions from characters
	actionsToExecute := map[string][]*Action{}
	for _, player := range players {

		// Swap characters
		player.Characters[swaps[player.ID][0]], player.Characters[swaps[player.ID][1]] = player.Characters[swaps[player.ID][1]], player.Characters[swaps[player.ID][0]]

		// Calculate actions
		for _, action := range actions[player.ID] {
			var character *Character = nil
			var actionToExecute *Action = nil
			for _, c := range player.Characters {
				if c.ID == action.CharacterId {
					character = c
					for _, a := range c.Actions {
						if a.ID == action.ActionId {
							actionToExecute = &a
						}
					}
				}
			}
			if character == nil || actionToExecute == nil {
				return errors.New("character or action couldn't be found")
			}

			if actionsToExecute[player.ID] == nil {
				actionsToExecute[player.ID] = []*Action{}
			}
			actionToExecute.originCharacter = character

			// Find target for the action
			var targettedCharacter *Character = nil
			for _, p := range players {
				if p.ID == action.Target {
					targettedCharacter = p.Characters[action.Slot]
				}
			}
			if targettedCharacter == nil {
				return errors.New("couldn't find target for action")
			}
			actionToExecute.targetCharacter = targettedCharacter

			actionsToExecute[player.ID] = append(actionsToExecute[player.ID], actionToExecute)
		}
	}

	// Prepare all the actions
	for _, actions := range actionsToExecute {
		for _, a := range actions {
			if a.Before != nil {
				a.Before(a.originCharacter, a.targetCharacter)
			}
		}
	}

	// Run all the actions
	for _, actions := range actionsToExecute {
		for _, a := range actions {
			if a.Execute != nil {
				a.latestResult = a.Execute(a.originCharacter, a.targetCharacter)
			}
		}
	}

	// TODO: Simulate status effects and results of attacks

	return nil
}
