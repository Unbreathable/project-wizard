package backend_starter

import (
	"os"

	"github.com/Liphium/project-wizard/backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	app.Route("/", routes.SetupRoutes)

	app.Listen(os.Getenv("LISTEN") + ":" + os.Getenv("PORT"))
}
