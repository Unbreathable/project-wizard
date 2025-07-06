package service

import (
	"maps"
	"sync"

	"slices"

	"github.com/Liphium/project-wizard/backend/game"
	"github.com/google/uuid"
)

const CharacterAmount = 4

type LobbyMode string

// All the elements
const (
	LobyMode1vs1 LobbyMode = "1vs1"
)

var LobbyModes = []LobbyMode{
	LobyMode1vs1,
}

type Lobby struct {
	mutex           *sync.Mutex
	mode            LobbyMode          // lobby mode
	id              string             // lobby id for interactions
	token           string             // lobby token for join requests
	playerIds       []string           // needed to generate unique ids
	spectatorTokens []string           // to whom results should be send
	players         map[string]*Player // needed to generate unique ids
	teamIds         []string           // slice of ids of teams
	teams           map[string]*Team   // team id -> *Team
	game            *Game              // game instance (nil until game starts)
}

type LobbyInfo struct {
	Id    string    `json:"id"`
	Mode  LobbyMode `json:"mode"`
	Token string    `json:"-"`
}

var lobbies sync.Map

// Creates a new lobby and returns the id
func CreateLobby(mode LobbyMode) (lobbyId string) {

	// Create unique lobby id
	lobbyId = uuid.New().String()
	for {
		if _, ok := lobbies.Load(lobbyId); !ok {
			break
		}
		lobbyId = uuid.New().String()
	}

	// Store lobby
	lobbies.Store(lobbyId, &Lobby{
		mutex: &sync.Mutex{},
		id:    lobbyId,
		mode:  mode,
		game:  nil,
		token: uuid.New().String(),
	})
	return lobbyId
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

func (lobby *Lobby) NewPlayer(name string) *Player {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	id := uuid.New().String()
	for slices.Contains(lobby.playerIds, id) {
		id = uuid.New().String()
	}

	token := uuid.New().String()
	for slices.Contains(lobby.spectatorTokens, token) {
		token = uuid.New().String()
	}

	// Add player to team
	lobby.players[id] = &Player{
		mutex:       &sync.Mutex{},
		relatedTeam: nil,
		id:          id,
		name:        name,
		ready:       false,
		readyTurn:   false,
		token:       token,
		gamePlayer:  nil,
	}
	lobby.playerIds = append(lobby.playerIds, id)
	lobby.spectatorTokens = append(lobby.playerIds, token)

	return lobby.players[id]
}

// Add Spectator from lobby
func (lobby *Lobby) AddSpectator(name string) {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	token := uuid.New().String()
	for slices.Contains(lobby.spectatorTokens, token) {
		token = uuid.New().String()
	}

	lobby.spectatorTokens = append(lobby.spectatorTokens, token)
}

// Remove Spectator from lobby
func (lobby *Lobby) RemoveSpectator(token string) {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	lobby.spectatorTokens = slices.DeleteFunc(lobby.spectatorTokens, func(id string) bool { return id == token })
}

// Adds new team to lobby with initial players
func (lobby *Lobby) NewTeam(size int) *Team {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	id := uuid.New().String()
	for slices.Contains(lobby.teamIds, id) {
		id = uuid.New().String()
	}

	team := &Team{
		mutex:        &sync.Mutex{},
		id:           id,
		size:         size,
		relatedLobby: lobby,
	}

	lobby.teamIds = append(lobby.teamIds, id)
	lobby.teams[id] = team

	return team
}

// Removes the team from the lobby
func (lobby *Lobby) RemoveTeam(teamId string) {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	team := lobby.teams[teamId]

	// Set related team in players
	for _, v := range team.GetPlayers() {
		v.SetTeam(nil)
	}

	// Remove Team
	lobby.teamIds = slices.DeleteFunc(lobby.teamIds, func(id string) bool { return id == teamId })
	delete(lobby.teams, teamId)
}

// Returns the team by id (can be nil)
func (lobby *Lobby) GetTeam(teamId string) *Team {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	return lobby.teams[teamId]
}

// Returns the all teams
func (lobby *Lobby) GetTeams() map[string]*Team {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	return lobby.teams
}

// Returns the all spectators
func (lobby *Lobby) GetSpectator() []string {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	return lobby.spectatorTokens
}

// Returns the all teams
func (lobby *Lobby) GetInfo() LobbyInfo {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	return LobbyInfo{
		Id:    lobby.id,
		Mode:  lobby.mode,
		Token: lobby.token,
	}
}

// Returns all players in lobby
func (lobby *Lobby) GetPlayers() []*Player {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	players := []*Player{}
	for _, v := range lobby.teams {
		players = append(players, slices.Collect(maps.Values(v.GetPlayers()))...)
	}

	return players
}

// Returns player by id
func (lobby *Lobby) GetPlayer(playerId string) *Player {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	return lobby.players[playerId]
}

// Return if all teams are full
func (lobby *Lobby) IsFull() bool {
	lobby.mutex.Lock()
	defer lobby.mutex.Unlock()

	full := true
	for _, v := range lobby.teams {
		if !v.IsFull() {
			full = false
			break
		}
	}

	return full
}

// Returns if a game is going on
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
	}

	// Collect player token for game start event
	players := []*Player{}
	for _, v := range lobby.teams {
		players = append(players, slices.Collect(maps.Values(v.GetPlayers()))...)
	}
	playerTokens := []string{}
	for _, v := range players {
		playerTokens = append(playerTokens, v.GetInfo().Token)
	}

	// Collect characters todo
	returnChars := map[string]map[string][]game.Character{}
	for _, t := range lobby.GetTeams() {
		for _, p := range t.GetPlayers() {
			for _, c := range p.GetGamePlayerState().Characters {
				returnChars[t.GetId()][p.GetInfo().Id] = append(returnChars[t.GetId()][p.GetInfo().Id], *c)
			}
		}
	}

	// Send game start event
	Instance.Send(playerTokens, GameStartEvent(returnChars))
}
