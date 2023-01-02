package server

import (
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

// The server instantiates the middleware (Sthe proxy)
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

	go s.handleSubscriptions()
	fullAddr := flag.String("addr", s.port, "http service address")
	// flag.Parse()
	log.Println("Start run server")
	log.Println(s.app.Listen(*fullAddr))
	log.Println("Run have failed")
}

// Configures the route for ws requests and handles them
func (s *Server) initWebsocket() {
	s.app.All("/ws", websocket.New(func(c *wsConn) {
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
			log.Println("Request create room")
			err = s.handleCreateRoom(reqBody, c)
			if err == nil {
				s.register <- subEvent{
					Connection: c,
					Body:       reqBody,
				}
			}
		case "join-room":
			log.Println("Request join room")
			s.register <- subEvent{
				Connection: c,
				Body:       reqBody,
			}
			err = s.handleJoinRoom(reqBody, c)
			if err == nil {
				s.register <- subEvent{
					Connection: c,
					Body:       reqBody,
				}
			}
		}
	}))
}

type PlayerPublicInfo struct {
	ID string `json:"id"`
}
type RoomPublicInfo struct {
	ID      string              `json:"id"`
	Players []*PlayerPublicInfo `json:"players"`
}
type ResponseHttpRooms struct {
	Rooms []*RoomPublicInfo `json:"rooms"`
}

func (s *Server) initHttp() {
	s.app.Get("/rooms", func(ctx *fiber.Ctx) error {
		log.Println("############# ROOMS endpoint")
		/////////////////////////////////////////
		roomTest1 := &Room{
			player1: &Player{
				ws: nil,
				id: "player1",
			},
			player2: &Player{
				ws: nil,
				id: "player2",
			},
			game: nil,
		}
		roomTest2 := &Room{
			player1: &Player{
				ws: nil,
				id: "player3",
			},
			player2: &Player{
				ws: nil,
				id: "player4",
			},
			game: nil,
		}
		s.rooms["roomTest1"] = roomTest1
		s.rooms["roomTest2"] = roomTest2
		/////////////////////////////////////////
		rooms := make([]*RoomPublicInfo, 0, len(s.rooms))
		for id, room := range s.rooms {
			players := make([]*PlayerPublicInfo, 0, 2)
			if room.player1 != nil {
				players = append(players, &PlayerPublicInfo{ID: room.player1.id})
			}
			if room.player2 != nil {
				players = append(players, &PlayerPublicInfo{ID: room.player2.id})
			}
			rooms = append(rooms, &RoomPublicInfo{
				ID:      id,
				Players: players,
			})
		}
		resp := ResponseHttpRooms{
			Rooms: rooms,
		}
		return ctx.JSON(resp)
	})
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
