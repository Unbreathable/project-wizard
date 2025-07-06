package service

import (
	"fmt"
	"slices"
	"sync"
)

type Team struct {
	mutex        *sync.Mutex
	relatedLobby *Lobby
	size         int                // amount of allowed players
	id           string             // lobby id for interactions
	playerIds    []string           // slice of ids of teams
	players      map[string]*Player // player id -> *Player
}

func (team *Team) GetPlayers() map[string]*Player {
	team.mutex.Lock()
	defer team.mutex.Unlock()
	return team.players
}

func (team *Team) GetPlayer(playerId string) *Player {
	team.mutex.Lock()
	defer team.mutex.Unlock()
	return team.players[playerId]
}

func (team *Team) GetId() string {
	team.mutex.Lock()
	defer team.mutex.Unlock()
	return team.id
}

func (team *Team) GetSize() int {
	team.mutex.Lock()
	defer team.mutex.Unlock()
	return team.size
}

func (team *Team) AddPlayer(player *Player) error {
	team.mutex.Lock()
	defer team.mutex.Unlock()

	if len(team.playerIds)+1 > team.size {
		return fmt.Errorf("team is full")
	}

	playerInfo := player.GetInfo()

	// Add player to team
	team.playerIds = append(team.playerIds, playerInfo.Id)
	team.players[playerInfo.Id] = player

	// Set related team in player
	player.SetTeam(team)
	return nil
}

func (team *Team) RemovePlayer(playerId string) {
	team.mutex.Lock()
	defer team.mutex.Unlock()

	player := team.players[playerId]

	// Delete player from team
	team.playerIds = slices.DeleteFunc(team.playerIds, func(id string) bool { return id == playerId })
	delete(team.players, playerId)

	// Remove related team in player
	player.SetTeam(nil)
}

func (team *Team) IsFull() bool {
	team.mutex.Lock()
	defer team.mutex.Unlock()

	if len(team.playerIds) >= team.size {
		return true
	}
	return false
}
