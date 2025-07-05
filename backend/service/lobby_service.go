package service

import (
	"maps"
	"sync"

	"github.com/Liphium/project-wizard/backend/game"
	"github.com/google/uuid"
)

const CharacterAmount = 4

type Lobby struct {
	mutex     *sync.Mutex
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
		playerID: {
			Name:  name,
			ID:    playerID,
			Ready: false,
			Token: uuid.New().String(),
		},
		player2ID: {
			Name:  "",
			ID:    player2ID,
			Ready: false,
			Token: uuid.New().String(),
		},
	}

	// Store lobby
	lobbies.Store(lobbyId, &Lobby{
		mutex:     &sync.Mutex{},
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

// Returns pointer to game
func (lobby *Lobby) GetGame() *Game {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	return lobby.game
}

// Returns all players
func (lobby *Lobby) GetPlayers() []Player {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	players := []Player{}

	for p := range maps.Values(lobby.players) {
		players = append(players, p)
	}

	return players
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

func (lobby *Lobby) StartGame() {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	lobby.game = &Game{
		mutex:        &sync.Mutex{},
		relatedLobby: lobby,
		playersReady: map[string]bool{
			lobby.playersId[0]: false,
			lobby.playersId[1]: false,
		},
		playerActions: map[string][]game.GameAction{
			lobby.playersId[0]: {},
			lobby.playersId[1]: {},
		},
		playerSwaps: map[string][]int{
			lobby.playersId[0]: {},
			lobby.playersId[1]: {},
		},
	}
	player1, _ := lobby.players[lobby.playersId[0]]
	player2, _ := lobby.players[lobby.playersId[1]]
	charsP1 := []game.Character{}
	charsP2 := []game.Character{}

	for _, v := range player1.GamePlayer.Characters {
		charsP1 = append(charsP1, *v)
	}
	for _, v := range player2.GamePlayer.Characters {
		charsP2 = append(charsP2, *v)
	}

	// Send game start event
	Instance.Send([]string{player1.Token, player2.Token}, GameStartEvent(map[string][]game.Character{
		lobby.playersId[0]: charsP1,
		lobby.playersId[1]: charsP2,
	}))
}
