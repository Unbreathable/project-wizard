package game

import "github.com/Liphium/project-wizard/backend/util"

type StatusEffect struct {
	ID             string `json:"id"`          // Only needed when visible
	Name           string `json:"name"`        // Only needed when visible
	Description    string `json:"description"` // Only needed when visible
	Visible        bool   `json:"visible"`
	TurnsRemaining int    `json:"turns_remaining"` // How many turns the effect will still stay

	OnHit func(current *Character, from *Character, action *Action, result ActionResult) *ActionResult `json:"-"` // Apply the status effect when the character gets hit
}

// Create a dodge status effect (ignores attacks on current user, invisible)
func StatusEffectDodge(turns int) StatusEffect {
	return StatusEffect{
		Visible:        false,
		TurnsRemaining: turns,
		OnHit: func(current, from *Character, action *Action, result ActionResult) *ActionResult {
			if action.Damage > 0 && !action.Oversight {
				return &ActionResult{
					DamageToCharacter: util.Ptr(0),
				}
			}
			return nil
		},
	}
}
