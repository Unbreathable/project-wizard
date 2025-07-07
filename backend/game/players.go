package game

import (
	"slices"
	"sync"
)

type GamePlayer struct {
	mutex      *sync.Mutex
	ID         string
	Mana       int
	Characters []*Character
}

func (player *GamePlayer) GetCharacterById(charId uint) *Character {
	player.mutex.Lock()
	defer player.mutex.Unlock()

	ind := slices.IndexFunc(player.Characters, func(c *Character) bool {
		return c.ID == charId
	})
	if ind == -1 {
		return nil
	}
	return player.Characters[ind]
}

func (player *GamePlayer) GetCharacters() []Character {
	player.mutex.Lock()
	defer player.mutex.Unlock()

	chars := []Character{}
	for _, v := range player.Characters {
		chars = append(chars, *v)
	}
	return chars
}

func (player *GamePlayer) SetCharacters(chars []*Character) {
	player.mutex.Lock()
	defer player.mutex.Unlock()

	player.Characters = chars
}
