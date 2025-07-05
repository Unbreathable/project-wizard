package service

import (
	"github.com/Liphium/project-wizard/neogate"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Instance *neogate.Instance
var Validate *validator.Validate = validator.New()

// Event for game start
func GameStartEvent() neogate.Event {
	return neogate.Event{
		Name: "game_start",
		Data: fiber.Map{},
	}
}
