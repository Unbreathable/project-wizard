package info_routes

import "github.com/gofiber/fiber/v2"

func InfoRoutes(router fiber.Router) {
	router.Post("/characters", charactersInfo)
}
