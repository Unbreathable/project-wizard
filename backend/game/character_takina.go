package game

import "github.com/Liphium/project-wizard/backend/util"

func NewTakina(id uint, player *GamePlayer) Character {
	return Character{
		ID:            id,
		relatedPlayer: player,
		Name:          "Takina",
		Origin:        "Lycoris Recoil",
		Elements:      []Element{ElementFire, ElementDark},
		Health:        250,
		Actions: map[uint]Action{
			1: NewDamageAction(1, "Shoot", "Shoots her real gun at a character", 60, ElementFire),
			2: {
				ID:          2,
				Name:        "Backshot",
				Element:     ElementFire,
				Description: "Shoots her real gun with a silencer, stun the enemy hit.",
				Damage:      40,
				ManaCost:    20,
				Before: func(current, target *Character) {
					target.AddStatusEffect(StatusEffectStun(1))
				},
				Execute: func(current, target *Character) ActionResult {
					return ActionResult{
						DamageToCharacter: util.Ptr(40),
						DamageElement:     ElementFire,
					}
				},
			},
			3: {
				ID:          3,
				Name:        "Steal",
				Element:     ElementDark,
				Description: "Get rid of the same amount of MP points as health points.",
				Damage:      30,
				Before: func(current, target *Character) {
					target.AddStatusEffect(StatusEffectRemoveMana(30))
				},
				Execute: func(current, target *Character) ActionResult {
					return ActionResult{
						DamageToCharacter: util.Ptr(30),
						DamageElement:     ElementDark,
					}
				},
			},
		},
	}
}
