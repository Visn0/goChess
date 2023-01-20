package server

import (
	"chess/server/domain"
	"chess/server/infrastructure"
	"chess/server/shared"
	"flag"
	"fmt"
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
		wg := &sync.WaitGroup{}
		wg.Add(1)
		switch reqAction {
		case "create-room":
			log.Println("Request create room")
			createRoomController := infrastructure.NewCreateRoomWsController(s.roomManager, repository)
			room, err := createRoomController.Invoke(reqBody)
			if err != nil {
				// TODO: return message in the websocket
				log.Println(err)
			}

			s.wsRouter(room, repository, true, wg)
			// roomActions.WsCreateRoom(s.roomManager, reqBody, c)
		case "join-room":
			log.Println("Request join room")
			joinRoomController := infrastructure.NewJoinRoomWsController(s.roomManager, repository)
			room, err := joinRoomController.Invoke(reqBody)
			if err != nil {
				// TODO: return message in the websocket
				log.Println(err)
			}

			s.wsRouter(room, repository, false, wg)
			// roomActions.WsJoinRoom(s.roomManager, reqBody, c)
		}

		wg.Wait()
	}))
}

func (s *Server) wsRouter(room *domain.Room, repository domain.ConnectionRepository, isHost bool, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Println("Room activated")
	var player *domain.Player
	var enemy *domain.Player
	var enemyRepository domain.ConnectionRepository

	for {
		if isHost {
			player = room.Player1
			enemy = room.Player2
		} else {
			player = room.Player2
			enemy = room.Player1
		}
		if player == nil {
			return
		}
		if enemy == nil {
			continue
		}

		if enemyRepository == nil {
			enemyRepository = infrastructure.NewBackendConnectionRepository(enemy.Ws)
		}

		if room.Game.ColotToMove != player.Color {
			continue
		}
		_, message, err := player.Ws.ReadMessage()
		if err != nil {
			log.Println("Some error:", err)
			player = nil
			return
		}
		// log.Println("Get message.")

		reqAction, _ := jsonparser.GetString(message, "action")
		reqBody, _, _, _ := jsonparser.Get(message, "body")

		switch reqAction {
		case "request-moves":
			// log.Println("Request moves")
			// application.WsGetValidMoves(r.game, reqBody, player.Ws)
			getValidMovesController := infrastructure.NewGetValidMovesWsController(repository, room.Game)
			err := getValidMovesController.Invoke(reqBody)
			if err != nil {
				log.Println("Error getting valid moves: ", err)
			}
		case "move-piece":
			if enemyRepository == nil {
				log.Println("Enemy repo is null. Enemy ID: ", enemy.ID)
			}

			movePieceController := infrastructure.NewMovePieceWsController(repository, enemyRepository, room.Game)
			err := movePieceController.Invoke(reqBody)
			if err != nil {
				log.Println("Error getting valid moves: ", err)
			}

			// log.Println("Move piece")
			// application.WsMovePiece(r.game, reqBody, player.Ws, enemy.Ws)
			player.StopTimer()
			fmt.Println("Moved Player: ", player.ID, " color: ", player.Color, " Time left:", player.TimeLeft())

			fmt.Println("Turn Player: ", enemy.ID, " color: ", enemy.Color, " Time left:", enemy.TimeLeft())
			enemy.StartTimer()
		default:
			log.Println("Unknown action")
		}
	}
}

func (s *Server) initHttp() {
	s.app.Get("/rooms", infrastructure.NewGetRoomsHttpController(s.roomManager).Invoke)
}
