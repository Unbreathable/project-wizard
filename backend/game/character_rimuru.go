package game

func NewRimuru(id uint, player *GamePlayer) Character {
	return Character{
		ID:            id,
		relatedPlayer: player,
		Name:          "Rimuru",
		Origin:        "Tensura",
		Elements:      []Element{ElementWater, ElementDark},
		Health:        400,
		Actions: map[uint]Action{
			1: {
				ID:          1,
				Name:        "Stomach",
				Element:     ElementNone,
				Description: "In case attacked in the current turn, steal MP gotten from attack.",
				ManaCost:    30,
				Before: func(current, target *Character) {
					current.AddStatusEffect(StatusEffectStealMana())
				},
			},
			2: NewDamageAction(2, "Water Blade", "Shoots a water blade.", 40, ElementWater),
			3: {
				ID:          3,
				Name:        "Guiding Shield",
				Element:     ElementNone,
				Description: "Adds a shield in front of someone that turns 40 HP of damage into healing.",
				Oversight:   true,
				ManaCost:    20,
				Before: func(current, target *Character) {
					target.AddStatusEffect(StatusEffectTurnDmgHeal(40))
				},
			},
		},
	}
}
