package service

import (
	"sync"

	"github.com/google/uuid"
)

type Lobby struct {
	mutex   sync.Mutex
	ID      uuid.UUID
	Player1 Player
	Player2 Player
}

type Player struct {
	ID   uuid.UUID `json:"player_id"`
	Name string    `json:"name"`
}

var lobbies sync.Map

// Creates a new lobby and returns the id
func CreateLobby(name string) (lobbyId uuid.UUID, playerID uuid.UUID) {

	// Create unique lobby id
	lobbyId = uuid.New()
	for {
		if _, ok := lobbies.Load(lobbyId); !ok {
			break
		}
		lobbyId = uuid.New()
	}

	// Create player ids
	playerID = uuid.New()
	player2ID := uuid.New()
	for {
		if player2ID != playerID {
			break
		}
		player2ID = uuid.New()
	}

	// Store lobby
	lobbies.Store(lobbyId, &Lobby{
		mutex: sync.Mutex{},
		ID:    lobbyId,
		Player1: Player{
			Name: name,
			ID:   playerID,
		},
		Player2: Player{
			Name: "",
			ID:   player2ID,
		},
	})
	return lobbyId, playerID
}

// Removes lobby by uuid
func RemoveLobby(lobbyId uuid.UUID) {
	lobbies.Delete(lobbyId)
}

// Loads lobby struct by uuid
func GetLobby(lobbyId uuid.UUID) (*Lobby, bool) {
	value, ok := lobbies.Load(lobbyId)
	if !ok {
		return &Lobby{}, false
	}
	return value.(*Lobby), true
}
