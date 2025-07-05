package game

// All characters (ID -> Creation Function)
var CharacterRegistry = map[uint]func(*GamePlayer) Character{
	1: newCharacter(1, NewChisato),
	2: newCharacter(2, NewTakina),
	3: newCharacter(3, NewRimuru),
	4: newCharacter(4, NewShuna),
}

func newCharacter(id uint, characterFun func(uint, *GamePlayer) Character) func(*GamePlayer) Character {
	return func(gp *GamePlayer) Character {
		return characterFun(id, gp)
	}
}

type Character struct {
	relatedPlayer *GamePlayer     `json:"-"`
	ID            uint            `json:"id"` // Character id (unique for every character)
	Name          string          `json:"name"`
	Origin        string          `json:"origin"` // Which anime, game or whereever they came from
	Elements      []Element       `json:"elements"`
	Health        int             `json:"health"`
	Actions       map[uint]Action `json:"actions"`
	StatusEffects []StatusEffect  `json:"status_effects"`
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

// Add status effect to character
func (c *Character) AddStatusEffect(effect StatusEffect) {
	if c.StatusEffects == nil {
		c.StatusEffects = []StatusEffect{}
	}

	c.StatusEffects = append(c.StatusEffects, effect)
}
