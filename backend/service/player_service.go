package service

import (
	"sync"

	"github.com/Liphium/project-wizard/backend/game"
)

type Player struct {
	mutex       *sync.Mutex
	relatedTeam *Team
	id          string           // player id
	name        string           // username displayed
	ready       bool             // ready in lobby
	readyTurn   bool             // ready in game
	token       string           // player verification
	gamePlayer  *game.GamePlayer // used in game
}

type PlayerInfo struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Ready     bool   `json:"ready"`
	ReadyTurn bool   `json:"read_turn"`
	Token     string `json:"-"`
}

func (p *Player) GetInfo() PlayerInfo {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return PlayerInfo{
		Id:        p.id,
		Name:      p.name,
		Ready:     p.ready,
		Token:     p.token,
		ReadyTurn: p.readyTurn,
	}
}

func (p *Player) SetName(name string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.name = name
}

func (p *Player) SetReady(ready bool) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.ready = ready
}

func (p *Player) SetReadyTurn(ready bool) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.readyTurn = ready
}

func (p *Player) SetTeam(team *Team) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.relatedTeam = team
}

func (p *Player) GetGamePlayer() *game.GamePlayer {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.gamePlayer
}

func (p *Player) GetGamePlayerState() game.GamePlayer {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return *p.gamePlayer
}

func (p *Player) SetGamePlayer() *game.GamePlayer {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.gamePlayer = &game.GamePlayer{
		ID:         p.id,
		Mana:       0,
		Characters: []*game.Character{},
	}

	return p.gamePlayer
}
