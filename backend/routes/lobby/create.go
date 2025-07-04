package lobby_routes

import (
	"log"

	"github.com/Liphium/project-wizard/backend/integration"
	"github.com/gofiber/fiber/v2"
)

// Route: /lobby/create
func createLobby(c *fiber.Ctx) error {
	log.Println("crate lobby")
	return integration.SuccessfulRequest(c)
}
