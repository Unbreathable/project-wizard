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

// Create a stun status effect (player isnt able to atack)
func StatusEffectStun(turns int) StatusEffect {
	return StatusEffect{
		ID:             "stun",
		Name:           "Stun",
		Description:    "The character can't perform actions.",
		Visible:        true,
		TurnsRemaining: turns,
	}
}

// Create a stun status effect (player isnt able to atack)
func StatusEffectInvulnerable(turns int) StatusEffect {
	return StatusEffect{
		ID:             "invulnerable",
		Name:           "Invulnerable",
		Description:    "The character doesn't recieve any damage.",
		Visible:        true,
		TurnsRemaining: turns,
		OnHit: func(current, from *Character, action *Action, result ActionResult) *ActionResult {
			return &ActionResult{
				DamageToCharacter: util.Ptr(0),
			}
		},
	}
}

// Create a status effect that applies damage to the character a redused rate
func StatusEffectReducedDamage(percentage float64) StatusEffect {
	return StatusEffect{
		Visible:        false,
		TurnsRemaining: 0,
		OnHit: func(current, from *Character, action *Action, result ActionResult) *ActionResult {
			return &ActionResult{
				DamageToCharacter: util.Ptr(int(float64(action.Damage) * percentage)),
			}
		},
	}
}

// Create a status effect that removes mana from the targeted player
func StatusEffectRemoveMana(amount int) StatusEffect {
	return StatusEffect{
		Visible:        false,
		TurnsRemaining: 0,
		OnHit: func(current, from *Character, action *Action, result ActionResult) *ActionResult {
			if current.relatedPlayer.Mana-amount < 0 {
				current.relatedPlayer.Mana = 0
			} else {
				current.relatedPlayer.Mana -= amount
			}
			return &result
		},
	}
}
