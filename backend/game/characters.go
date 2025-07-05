package game

type Character struct {
	relatedPlayer *GamePlayer
	ID            uint // Character id (unique for every character)
	Name          string
	Origin        string // Which anime, game or whereever they came from
	Elements      []Element
	Health        int
	Actions       []Action
	StatusEffects []StatusEffect
}

func (c *Character) IsDead() bool {
	return c.Health < 0
}

func (c *Character) RelatedPlayer() *GamePlayer {
	return c.relatedPlayer
}

// Deal damage to the character. Specify an element in case wanted.
func (c *Character) DealDamage(damage int, element *Element) {
	finalDamage := damage

	// Average the damage with the element multipliers (as specified in game design)
	if element != nil {
		var damages []int
		for _, el := range c.Elements {
			damages = append(damages, int(float64(damage)*el.GetDamageMultiplierFor(*element)))
		}
		finalDamage = 0
		for _, dmg := range damages {
			finalDamage += dmg
		}
		finalDamage /= len(damages)
	}

	c.Health -= finalDamage
}
