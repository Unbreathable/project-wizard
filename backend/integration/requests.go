package integration

import (
	"log"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
)

func SuccessfulRequest(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
	})
}

func FailedRequest(c *fiber.Ctx, message string, err error) error {

	// Print error if it isn't nil
	if err != nil {
		log.Println(c.Route().Name + " ERROR: " + message + ":" + err.Error())
		debug.PrintStack()
	}

	return c.JSON(fiber.Map{
		"success": false,
		"error":   message,
	})
}

func InvalidRequest(c *fiber.Ctx, message string) error {
	log.Println(c.Route().Name + " request is invalid. msg: " + message)
	debug.PrintStack()
	return c.SendStatus(fiber.StatusBadRequest)
}
