package game

func NewChisato(id uint, player *GamePlayer) Character {
	return Character{
		ID:            id,
		relatedPlayer: player,
		Name:          "Chisato",
		Origin:        "Lycoris Recoil",
		Elements:      []Element{ElementAir, ElementLight},
		Health:        300,
		Actions: map[uint]Action{
			1: {
				ID:          1,
				Element:     ElementNone,
				Name:        "Dodge",
				Description: "Dodge any attack cast on her.",
				ManaCost:    40,
				Before: func(current, target *Character) {
					current.AddStatusEffect(StatusEffectDodge(0))
				},
			},
			2: {
				ID:          2,
				Element:     ElementNone,
				Name:        "Fake-Shoot",
				Description: "Stuns one character for one turn with her non-lethal bullets.",
				ManaCost:    30,
				Before: func(current, target *Character) {
					current.AddStatusEffect(StatusEffectStun(1))
				},
			},
			3: {
				ID:          3,
				Element:     ElementNone,
				Name:        "Charm",
				Description: "Takes 60% of the attack of the selected character (rest is voided).",
				ManaCost:    20,
				Oversight:   true,
				Before: func(current, target *Character) {
					target.AddStatusEffect(StatusEffectInvulnerable(0))
					current.AddStatusEffect(StatusEffectReducedDamage(0.6))
				},
			},
		},
	}
}
