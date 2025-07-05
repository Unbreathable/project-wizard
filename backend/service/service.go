package service

import (
	"github.com/Liphium/project-wizard/backend/game"
	"github.com/Liphium/project-wizard/neogate"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Instance *neogate.Instance
var Validate *validator.Validate = validator.New()

type SimulationResultEvent struct {
	Swaps   map[string][]int             `json:"swaps"`
	Actions map[string][]game.GameAction `json:"actions"`
	Result  map[string]SimulationResult  `json:"results"`
}

type SimulationResult struct {
	Mana       int              `json:"mana"`
	ID         string           `json:"player_id"`
	Characters []game.Character `json:"characters"`
}

// Event for game start
func GameStartEvent(data map[string][]game.Character) neogate.Event {
	return neogate.Event{
		Name: "game_start",
		Data: fiber.Map{
			"characters": data,
		},
	}
}

// Event for game infos
func GameInfoEvent() neogate.Event {
	return neogate.Event{
		Name: "game_info",
		Data: fiber.Map{}, // TODO: return info
	}
}

// Event for game start
func GameUpdateEvent(res SimulationResultEvent) neogate.Event {
	return neogate.Event{
		Name: "game_update",
		Data: fiber.Map{
			"result": res,
		},
	}
}
