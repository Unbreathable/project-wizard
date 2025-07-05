package service

import (
	"fmt"
	"slices"

	"github.com/Liphium/project-wizard/backend/game"
)

type Player struct {
	ID         string           `json:"player_id"` // Player id
	Name       string           `json:"name"`
	Ready      bool             `json:"ready"`
	Token      string           `json:"-"` // Player verification
	GamePlayer *game.GamePlayer `json:"-"`
}

// int playerPos refers to the players position (player1 or player2)
func (lobby *Lobby) GetPlayer(playerPos int) (Player, error) {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	player, ok := lobby.players[lobby.playersId[playerPos-1]]
	if !ok {
		return Player{}, fmt.Errorf("position doesn't exist")
	}
	return player, nil
}

// returns player, player position, error
func (lobby *Lobby) GetPlayerById(playerId string) (Player, int, error) {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	player, ok := lobby.players[playerId]
	if !ok {
		return Player{}, 0, fmt.Errorf("playerId doesn't exist")
	}
	pos := slices.Index(lobby.playersId, playerId)
	if pos == -1 {
		return Player{}, 0, fmt.Errorf("playerId doesn't exist")
	}
	return player, pos + 1, nil
}

// int playerPos refers to the players position (player1 or player2)
func (lobby *Lobby) SetNamePlayer(playerPos int, name string) error {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	player, ok := lobby.players[lobby.playersId[playerPos-1]]
	if !ok {
		return fmt.Errorf("position doesn't exist")
	}
	player.Name = name
	lobby.players[lobby.playersId[playerPos-1]] = player
	return nil
}

func (lobby *Lobby) SetNamePlayerById(playerId string, name string) error {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	player, ok := lobby.players[playerId]
	if !ok {
		return fmt.Errorf("id doesn't exist")
	}
	player.Name = name
	lobby.players[playerId] = player
	return nil
}

// int playerPos refers to the players position (player1 or player2)
func (lobby *Lobby) SetReadyPlayer(playerPos int, ready bool) error {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	player, ok := lobby.players[lobby.playersId[playerPos-1]]
	if !ok {
		return fmt.Errorf("position doesn't exist")
	}
	player.Ready = ready
	lobby.players[lobby.playersId[playerPos-1]] = player
	return nil
}

func (lobby *Lobby) SetReadyPlayerById(playerId string, ready bool) error {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	player, ok := lobby.players[playerId]
	if !ok {
		return fmt.Errorf("id doesn't exist")
	}
	player.Ready = ready
	lobby.players[playerId] = player
	return nil
}

func (lobby *Lobby) GetPlayerTokenById(playerId string) string {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	player, ok := lobby.players[playerId]
	if !ok {
		return ""
	}
	return player.Token
}
