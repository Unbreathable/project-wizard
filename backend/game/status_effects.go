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
				result.DamageToCharacter = util.Ptr(0)
				return &result
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

// Create a status effect that makes that negates all damage
func StatusEffectInvulnerable(turns int) StatusEffect {
	return StatusEffect{
		ID:             "invulnerable",
		Name:           "Invulnerable",
		Description:    "The character doesn't recieve any damage.",
		Visible:        true,
		TurnsRemaining: turns,
		OnHit: func(current, from *Character, action *Action, result ActionResult) *ActionResult {
			result.DamageToCharacter = util.Ptr(0)
			return &result
		},
	}
}

// Create a status effect that applies damage to the character a redused rate
func StatusEffectReducedDamage(percentage float64) StatusEffect {
	return StatusEffect{
		Visible:        false,
		TurnsRemaining: 0,
		OnHit: func(current, from *Character, action *Action, result ActionResult) *ActionResult {
			result.DamageToCharacter = util.Ptr(int(float64(action.Damage) * percentage))
			return &result
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
			return nil
		},
	}
}

// Create a status effect that removes mana from the targeted player
func StatusEffectStealMana() StatusEffect {
	return StatusEffect{
		Visible:        false,
		TurnsRemaining: 0,
		OnHit: func(current, from *Character, action *Action, result ActionResult) *ActionResult {
			current.relatedPlayer.Mana += *result.DamageToCharacter
			from.relatedPlayer.Mana -= *result.DamageToCharacter
			return nil
		},
	}
}

// Create a status effect that removed damage from the attack and heals the person
func StatusEffectTurnDmgHeal(damageRed int) StatusEffect {
	return StatusEffect{
		Visible:        false,
		TurnsRemaining: 0,
		OnHit: func(current, from *Character, action *Action, result ActionResult) *ActionResult {
			if *result.DamageToCharacter <= damageRed {
				result.HealToCharacter = util.Ptr(*result.DamageToCharacter)
				result.DamageToCharacter = util.Ptr(0)
				return &result
			} else {
				result.DamageToCharacter = util.Ptr(*result.DamageToCharacter - damageRed)
				result.HealToCharacter = util.Ptr(damageRed)
				return &result
			}
		},
	}
}

// Create a status effect that makes that negates all damage from a certain element
func StatusEffectInvulnerableToElem(element Element) StatusEffect {
	return StatusEffect{
		Visible:        false,
		TurnsRemaining: 0,
		OnHit: func(current, from *Character, action *Action, result ActionResult) *ActionResult {
			if result.DamageElement == element {
				result.DamageToCharacter = util.Ptr(0)
				return &result
			}
			return nil
		},
	}
}

// Create a status effect that turns damage to mana
func StatusEffectDmgToMana() StatusEffect {
	return StatusEffect{
		TurnsRemaining: 0,
		OnHit: func(current, from *Character, action *Action, result ActionResult) *ActionResult {
			current.relatedPlayer.Mana += *result.DamageToCharacter
			return nil
		},
	}
}
