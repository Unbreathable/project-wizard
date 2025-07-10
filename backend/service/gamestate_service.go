package service

type GameState = string

const (
	GameStateTeamSelect      GameState = "team_select"
	GameStateCharacterSelect GameState = "char_select"
	GameStateRunning         GameState = "running"
	GameStateSimulating      GameState = "simulating"
	GameStateEnd             GameState = "end"
)
