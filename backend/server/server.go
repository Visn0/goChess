package server

import (
	"encoding/json"
	"flag"
	"log"
	"sync"

	"github.com/buger/jsonparser"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	websocket "github.com/gofiber/websocket/v2"
)

type wsConn = websocket.Conn

type subEvent struct {
	Connection *wsConn
	Body       []byte
}

type Server struct {
	addr string
	port string
	app  *fiber.App

	wsConnections      map[*wsConn]struct{}
	wsConnectionsMutex sync.Mutex

	register   chan subEvent
	unregister chan subEvent

	rooms map[string]*Room
}

func NewServer(addr, port string) *Server {
	return &Server{
		addr:               addr,
		port:               port,
		app:                fiber.New(),
		wsConnections:      make(map[*wsConn]struct{}),
		wsConnectionsMutex: sync.Mutex{},
		register:           make(chan subEvent),
		unregister:         make(chan subEvent),
		rooms: map[string]*Room{
			"room1": NewRoom(),
		},
	}
}

// The server instantiates the middleware (the proxy)
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
		return c.SendStatus(fiber.StatusAccepted)
	})
}

// The server instantiates the router (routes for http and ws conections).
func (s *Server) initRouter() {
	s.initWebsocket()

	s.app.Get("/", func(ctx *fiber.Ctx) error {
		log.Println("############# ROOMS endpoint")
		return ctx.JSON(s)
	})
}

func (s *Server) Run() {
	s.initMiddleware()
	s.initRouter()
	log.Println("Init server")

	go s.handleSubscriptions()
	fullAddr := flag.String("addr", s.port, "http service address")
	// flag.Parse()
	log.Println("Start run server")
	log.Println(s.app.Listen(*fullAddr))
	log.Println("Run have failed")
}

// Configures the route for ws requests and handles them
func (s *Server) initWebsocket() {
	s.app.Get("/ws", websocket.New(func(c *wsConn) {

		log.Println("New ws connection")

		messageType, message, err := c.ReadMessage()
		log.Println(messageType)
		if err != nil {
			// Error reading because of an unexpected disconnect (probably)
			log.Println("Some error:", err)
			// TODO: remove connection from our lists
			return
		}
		log.Println("Get message.")

		reqAction, _ := jsonparser.GetString(message, "action")
		reqBody, _, _, _ := jsonparser.Get(message, "body")

		switch reqAction {
		case "create-room":
			log.Println("New subscription")
			s.register <- subEvent{
				Connection: c,
				Body:       reqBody,
			}
			req := RequestCreateRoom{}
			err := json.Unmarshal(reqBody, &req)
			if err != nil {
				log.Println("Error unmarshalling request create room:", err)
				return
			}

			room, ok := s.rooms[req.Name]
			if !ok {
				room = NewRoom()
				player := &Player{
					ws:   c,
					name: "player1",
				}
				room.AddPlayer(player)
				s.rooms[req.Name] = room
				log.Println("Room created")
				go room.HandleGame()
			} else {
				_ = room
				log.Println("Room already exists")
			}
		}
	}))
}

func (s *Server) addSubscription(event subEvent) {
	s.wsConnectionsMutex.Lock()
	defer s.wsConnectionsMutex.Unlock()
	s.wsConnections[event.Connection] = struct{}{}
	log.Println("connection registered")
}

func (s *Server) removeSubscription(event subEvent) {
	s.wsConnectionsMutex.Lock()
	defer s.wsConnectionsMutex.Unlock()
	delete(s.wsConnections, event.Connection)
	log.Println("connection unregistered")
}

type RequestCreateRoom struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (s *Server) handleSubscriptions() {
	defer func() {
		close(s.register)
		close(s.unregister)
	}()
	for {
		select {
		case event := <-s.register:
			log.Println("handleSubscription received new register")
			s.addSubscription(event)
		case event := <-s.unregister:
			log.Println("handleSubscription received new unregister")
			s.removeSubscription(event)
		}
	}
}
