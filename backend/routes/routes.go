package routes

import (
	"log"
	"os"

	"github.com/Liphium/project-wizard/backend/service"
	"github.com/Liphium/project-wizard/neogate"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {

	router.Route("/gateway", setupNeo)
}

// Setup neogate
func setupNeo(router fiber.Router) {

	// Create the gateway
	service.Instance = neogate.Setup(neogate.Config{
		Secret: []byte(os.Getenv("JWT_SECRET")),

		// Handle client disconnect
		ClientDisconnectHandler: func(client *neogate.Client) {
			// TODO: Maybe handle disconnections a little more?
		},

		// Handle enter network
		ClientConnectHandler: func(client *neogate.Client, key string) bool {
			return false
		},

		// Handle client entering network
		ClientEnterNetworkHandler: func(client *neogate.Client, key string) bool {
			// Send an event to notify of connection success
			service.Instance.SendEventToClient(client, neogate.Event{
				Name: "ng_success",
			})

			return false
		},

		// Check the delivered token
		CheckToken: func(token, attachments string) (neogate.ClientInfo, bool) {
			var clientInfo neogate.ClientInfo

			clientInfo = neogate.ClientInfo{
				Account: "acc-1",
				Session: "acc-1",
			}
			return clientInfo, true
		},

		// Set the adapter name of the client to include the address
		ClientAdapterHandler: func(client *neogate.Client) string {
			return client.ID
		},

		ErrorHandler: func(err error) {
			log.Println("neogate error:", err)
		},

		ClientEncodingMiddleware: neogate.DefaultClientEncodingMiddleware,
		DecodingMiddleware:       neogate.DefaultDecodingMiddleware,
	})

	// Add all the routes for the gateway
	router.Route("/connect", service.Instance.MountGateway)

}
