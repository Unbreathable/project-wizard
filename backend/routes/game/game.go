package game_routes

import "github.com/gofiber/fiber/v2"

func GameRoutes(router fiber.Router) {
	router.Post("/turn", turnGame)
	router.Post("/remove_turn", removeTurnGame)
}
