package service

import (
	"sync"

	"github.com/google/uuid"
)

const CharacterAmount = 4

type Lobby struct {
	mutex     sync.Mutex
	id        string            // uuid strings
	playersId []string          // uuid strings
	players   map[string]Player // uuid strings
	game      *Game
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

	playerMap := map[string]Player{
		playerID: Player{
			Name:  name,
			ID:    playerID,
			Ready: false,
		},
		player2ID: Player{
			Name:  "",
			ID:    player2ID,
			Ready: false,
		},
	}

	// Store lobby
	lobbies.Store(lobbyId, &Lobby{
		mutex:     sync.Mutex{},
		id:        lobbyId,
		playersId: []string{playerID, player2ID},
		players:   playerMap,
		game:      nil,
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
	return lobby.players[lobby.playersId[1]].Name != ""
}

func (lobby *Lobby) IsRunning() bool {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()
	return lobby.game != nil
}
