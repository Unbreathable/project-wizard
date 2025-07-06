package game

import "slices"

type GamePlayer struct {
	ID         string
	Mana       int
	Characters []*Character
}

func (player *GamePlayer) GetCharacterById(charId uint) *Character {
	ind := slices.IndexFunc(player.Characters, func(c *Character) bool {
		return c.ID == charId
	})
	if ind == -1 {
		return nil
	}
	return player.Characters[ind]
}

func (player *GamePlayer) GetCharacters() []Character {
	chars := []Character{}
	for _, v := range player.Characters {
		chars = append(chars, *v)
	}
	return chars
}
