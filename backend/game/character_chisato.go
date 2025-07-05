package game

func NewChisato(id uint, player *GamePlayer) Character {
	return Character{
		ID:            id,
		relatedPlayer: player,
		Name:          "Chisato",
		Origin:        "Lycoris Recoil",
		Elements:      []Element{ElementAir, ElementLight},
		Health:        300,
		Actions: []Action{
			{
				ID:          1,
				Name:        "Dodge",
				Description: "Dodge any attack cast on her",
				ManaCost:    40,
				Before: func(current, target *Character) {
					current.AddStatusEffect(StatusEffectDodge(1))
				},
			},
		},
	}
}
