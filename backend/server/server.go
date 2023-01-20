package server

import (
	"chess/server/domain"
	"chess/server/infrastructure"
	"chess/server/shared"
	"flag"
	"log"
	"sync"

	"github.com/buger/jsonparser"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	websocket "github.com/gofiber/websocket/v2"
)

type subEvent struct {
	Connection *shared.WsConn
	Body       []byte
}

type Server struct {
	addr string
	port string
	app  *fiber.App

	wsConnections      map[*shared.WsConn]struct{}
	wsConnectionsMutex sync.Mutex

	register   chan subEvent
	unregister chan subEvent

	roomManager *domain.RoomManager
}

func NewServer(addr, port string) *Server {
	return &Server{
		addr:               addr,
		port:               port,
		app:                fiber.New(),
		wsConnections:      make(map[*shared.WsConn]struct{}),
		wsConnectionsMutex: sync.Mutex{},
		register:           make(chan subEvent),
		unregister:         make(chan subEvent),
		roomManager:        domain.NewRoomManager(),
	}
}

// The server instantiates the middleware (proxy)
func (s *Server) initMiddleware() {
	s.app.Use(cors.New())
	s.app.Use(func(c *fiber.Ctx) error {
		// ONLY ALLOW LOCAL REQUESTS
		if !c.IsFromLocal() {
			log.Println("Blocked request")
			return nil
		}
		if websocket.IsWebSocketUpgrade(c) { // Returns true if the client requested upgrade to the WebSocket protocol
			log.Println("Websocket upgraded")
			return c.Next()
		}
		log.Println("HTTP request")
		log.Println(c.BaseURL(), c.OriginalURL())
		return c.Next()
	})
}

// The server instantiates the router (routes for http and ws conections).
func (s *Server) initRouter() {
	s.initWebsocket()
	s.initHttp()
}

func (s *Server) Run() {
	s.initMiddleware()
	s.initRouter()
	log.Println("Init server")

	fullAddr := flag.String("addr", s.port, "http service address")
	// flag.Parse()
	log.Println("Start run server")
	log.Println(s.app.Listen(*fullAddr))
	log.Println("Run have failed")
}

// Configures the route for ws requests and handles them
func (s *Server) initWebsocket() {
	s.app.All("/ws", websocket.New(func(c *shared.WsConn) {
		log.Println("New ws connection")

		_, message, err := c.ReadMessage()
		if err != nil {
			// Error reading because of an unexpected disconnect (probably)
			log.Println("Some error:", err)
			return
		}
		log.Println("Get message.")

		reqAction, err := jsonparser.GetString(message, "action")
		if err != nil {
			log.Println("Error getting action:", err)
			return
		}
		reqBody, _, _, err := jsonparser.Get(message, "body")
		if err != nil {
			log.Println("Error getting body:", err)
			return
		}

		repository := infrastructure.NewBackendConnectionRepository(c)
		switch reqAction {
		case "create-room":
			log.Println("Request create room")
			controller := infrastructure.NewCreateRoomWsController(s.roomManager, repository)
			err = controller.Invoke(reqBody)
			if err != nil {
				// TODO: return message in the websocket
				log.Println(err)
			}
			// roomActions.WsCreateRoom(s.roomManager, reqBody, c)
		case "join-room":
			log.Println("Request join room")
			// roomActions.WsJoinRoom(s.roomManager, reqBody, c)
		}
	}))
}

func (s *Server) initHttp() {
	s.app.Get("/rooms", func(ctx *fiber.Ctx) error {
		log.Println("############# ROOMS endpoint")
		// return roomActions.HttpGetRooms(ctx, s.roomManager)
		return nil
	})
}
