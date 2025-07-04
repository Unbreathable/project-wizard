package service

import (
	"sync"

	"github.com/google/uuid"
)

type Lobby struct {
	mutex   sync.Mutex
	id      string
	player1 Player
	player2 Player
}

type Player struct {
	ID   string `json:"player_id"`
	Name string `json:"name"`
}

var lobbies sync.Map

// Creates a new lobby and returns the id
func CreateLobby(name string) (lobbyId string, playerID string) {

	// Create unique lobby id
	lobbyId = uuid.New().String()
	for {
		if _, ok := lobbies.Load(lobbyId); !ok {
			break
		}
		lobbyId = uuid.New().String()
	}

	// Create player ids
	playerID = uuid.New().String()
	player2ID := uuid.New().String()
	for {
		if player2ID != playerID {
			break
		}
		player2ID = uuid.New().String()
	}

	// Store lobby
	lobbies.Store(lobbyId, &Lobby{
		mutex: sync.Mutex{},
		id:    lobbyId,
		player1: Player{
			Name: name,
			ID:   playerID,
		},
		player2: Player{
			Name: "",
			ID:   player2ID,
		},
	})
	return lobbyId, playerID
}

// Removes lobby by uuid
func RemoveLobby(lobbyId string) {
	lobbies.Delete(lobbyId)
}

// Loads lobby struct by uuid
func GetLobby(lobbyId string) (*Lobby, bool) {
	value, ok := lobbies.Load(lobbyId)
	if !ok {
		return &Lobby{}, false
	}
	return value.(*Lobby), true
}

func (lobby *Lobby) IsFull() bool {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()
	return lobby.player2.Name != ""
}

func (lobby *Lobby) GetPlayer1() Player {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()
	return lobby.player1
}

func (lobby *Lobby) GetPlayer2() Player {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()
	return lobby.player2
}

func (lobby *Lobby) SetNamePlayer1(name string) {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	lobby.player1.Name = name
}

func (lobby *Lobby) SetNamePlayer2(name string) {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	lobby.player2.Name = name
}
