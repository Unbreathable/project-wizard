package info_routes

import (
	"github.com/Liphium/project-wizard/backend/game"
	"github.com/gofiber/fiber/v2"
)

type CharacterInfoResponse struct {
	Success    bool             `json:"success"`
	Characters []game.Character `json:"characters"`
}

// Route: /info/characters
func charactersInfo(c *fiber.Ctx) error {
	chars := []game.Character{}
	for _, v := range game.CharacterRegistry {
		chars = append(chars, v(nil))
	}
	return c.JSON(CharacterInfoResponse{
		Success:    true,
		Characters: chars,
	})
}
