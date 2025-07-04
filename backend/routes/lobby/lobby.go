package lobby_routes

import "github.com/gofiber/fiber/v2"

func LobbyRoutes(router fiber.Router) {
	router.Post("/create", createLobby)
}
