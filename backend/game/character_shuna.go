package game

func NewShuna(id uint, player *GamePlayer) Character {
	return Character{
		ID:            id,
		relatedPlayer: player,
		Name:          "Shuna",
		Origin:        "Tensura",
		Elements:      []Element{ElementLight},
		Health:        270,
		Actions: map[uint]Action{
			1: NewDamageAction(1, "Light Ray", "A powerful beam that deals damage.", 40, ElementLight),
			2: {
				ID:          2,
				Name:        "Holy Bell",
				Element:     ElementNone,
				Description: "All Dark attacks by enemies are powerless.",
				Before: func(current, target *Character) {
					for _, v := range target.relatedPlayer.Characters {
						v.AddStatusEffect(StatusEffectInvulnerableToElem(ElementDark))
					}
				},
			},
			3: {
				ID:          3,
				Name:        "Holy Shield",
				Element:     ElementNone,
				Description: "Adds a shield that can takes all the MP from the damage the enemy would've dealt to the character.",
				Oversight:   true,
				ManaCost:    30,
				Before: func(current, target *Character) {
					target.AddStatusEffect(StatusEffectInvulnerable(0))
					target.AddStatusEffect(StatusEffectDmgToMana())
				},
			},
		},
	}
}
