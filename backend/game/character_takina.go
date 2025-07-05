package game

func NewTakina(id uint, player *GamePlayer) Character {
	return Character{
		ID:            id,
		relatedPlayer: player,
		Name:          "Takina",
		Origin:        "Lycoris Recoil",
		Elements:      []Element{ElementFire, ElementDark},
		Health:        250,
		Actions: map[uint]Action{
			1: {
				ID:          1,
				Name:        "Dodge",
				Element:     ElementFire,
				Description: "Dodge any attack cast on her",
				ManaCost:    40,
				Before: func(current, target *Character) {
					current.AddStatusEffect(StatusEffectDodge(1))
				},
			},
		},
	}
}
