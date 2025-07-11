package routes

import (
	"encoding/json"
	"log"
	"os"

	game_routes "github.com/Liphium/project-wizard/backend/routes/game"
	info_routes "github.com/Liphium/project-wizard/backend/routes/info"
	lobby_routes "github.com/Liphium/project-wizard/backend/routes/lobby"
	"github.com/Liphium/project-wizard/backend/service"
	"github.com/Liphium/project-wizard/neogate"
	"github.com/gofiber/fiber/v2"
)

type NeogateTokenAttachment struct {
	PlayerId string `json:"player_id"`
	LobbyId  string `json:"lobby_id"`
}

func SetupRoutes(router fiber.Router) {

	// neogate
	router.Route("/gateway", setupNeo)

	// lobby routes
	router.Route("/lobby", lobby_routes.LobbyRoutes)

	// game routes
	router.Route("/game", game_routes.GameRoutes)

	router.Route("/info", info_routes.InfoRoutes)
}

func validatedNeoHandler[T struct{}](handler func(*neogate.Context, T) neogate.Event) func(*neogate.Context, T) neogate.Event {
	return func(c *neogate.Context, content T) neogate.Event {
		if err := service.Validate.Struct(content); err != nil {
			return neogate.ErrorResponse(c, "invalid request", err)
		}

		return handler(c, content)
	}
}

// Setup neogate
func setupNeo(router fiber.Router) {

	neogate.CreateHandlerFor(service.Instance, "lobby_create", func(c *neogate.Context, _ interface{}) neogate.Event {
		return neogate.ErrorResponse(c, "server error", nil)
	})

	// Create the gateway
	service.Instance = neogate.Setup(neogate.Config{
		Secret: []byte(os.Getenv("JWT_SECRET")),

		// Handle client disconnect
		ClientDisconnectHandler: func(client *neogate.Client) {
			neoAtt, err := decodeSession(client.Session)
			if err == nil {
				lobby, ok := service.GetLobby(neoAtt.LobbyId)
				if ok {
					players := lobby.GetPlayers()
					service.RemoveLobby(neoAtt.LobbyId)
					for _, p := range players {
						data, err := encodeSession(NeogateTokenAttachment{PlayerId: p.ID, LobbyId: neoAtt.LobbyId})
						if err == nil {
							service.Instance.Disconnect(p.Token, data)
						}
					}
				}
			}
		},

		// Handle enter network
		ClientConnectHandler: func(client *neogate.Client, key string) bool {
			return false
		},

		// Handle client entering network
		ClientEnterNetworkHandler: func(client *neogate.Client, attachment string) bool {
			// Send an event to notify of connection success
			service.Instance.SendEventToClient(client, neogate.Event{
				Name: "ng_success",
			})
			att, err := decodeSession(attachment)
			if err != nil {
				return true
			}
			data, err := lobby_routes.GetLobbyInfo(att.LobbyId)
			if err != nil {
				return true
			}
			service.Instance.SendOne(client.ID, lobby_routes.LobbyChangeEvent(data))

			return false
		},

		// Check the delivered token
		CheckToken: func(token, attachments string) (neogate.ClientInfo, bool) {
			var clientInfo neogate.ClientInfo

			clientInfo = neogate.ClientInfo{
				Account: token,
				Session: attachments,
			}

			neoAtt, err := decodeSession(attachments)
			if err != nil {
				return clientInfo, false
			}

			lobby, ok := service.GetLobby(neoAtt.LobbyId)
			if !ok {
				return clientInfo, false
			}
			if lobby.GetPlayerTokenById(neoAtt.PlayerId) != token {
				return clientInfo, false
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

func decodeSession(attachments string) (NeogateTokenAttachment, error) {
	var neoAtt NeogateTokenAttachment
	err := json.Unmarshal([]byte(attachments), &neoAtt)
	return neoAtt, err
}

func encodeSession(att NeogateTokenAttachment) (string, error) {
	bytes, err := json.Marshal(att)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
