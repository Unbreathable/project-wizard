package service

type GameState uint

const (
	GameStateTeamSelect      GameState = 0
	GameStateCharacterSelect GameState = 1
	GameStateRunning         GameState = 2
	GameStateSimulating      GameState = 3
	GameStateEnd             GameState = 4
)

func (g GameState) GetPtr() *GameState {
	return &g
}
