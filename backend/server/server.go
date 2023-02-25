package server

import (
	"chess/server/application"
	"chess/server/domain"
	"chess/server/infrastructure"
	"chess/server/shared"
	"flag"
	"fmt"
	"log"
	"path"
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
	app := fiber.New()
	app.Use(cors.New())

	return &Server{
		addr:               addr,
		port:               port,
		app:                app,
		wsConnections:      make(map[*shared.WsConn]struct{}),
		wsConnectionsMutex: sync.Mutex{},
		register:           make(chan subEvent),
		unregister:         make(chan subEvent),
		roomManager:        domain.NewRoomManager(),
	}
}

func (s *Server) Static(prefix, root string, singlePageApp bool, config ...fiber.Static) {
	s.app.Static(prefix, root, config...)

	if singlePageApp {
		s.app.Get("/rooms", func(ctx *fiber.Ctx) error {
			return ctx.SendFile(path.Join(root, "index.html"))
		})
		s.app.Get("/game", func(ctx *fiber.Ctx) error {
			return ctx.SendFile(path.Join(root, "index.html"))
		})
	}
}

// The server instantiates the middleware (proxy)
func (s *Server) initMiddleware() {
	s.app.Use(func(c *fiber.Ctx) error {
		// ONLY ALLOW LOCAL REQUESTS
		// if !c.IsFromLocal() {
		// 	log.Println("Blocked request")
		// 	return nil
		// }

		fmt.Printf("### HEADER Connection: %s\n", c.Context().Request.Header.Peek("Connection"))
		fmt.Printf("### HEADER Upgrade: %s\n", c.Context().Request.Header.Peek("Upgrade"))
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
	s.initHTTP()
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
	s.app.All("/ws", websocket.New(func(conn *websocket.Conn) {
		wsConn := &shared.WsConn{Conn: conn}
		log.Println("New ws connection")

		_, message, err := wsConn.ReadMessage()
		if err != nil {
			// Error reading because of an unexpected disconnect (probably)
			log.Println("Some error:", err)
			return
		}

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

		c := infrastructure.NewBackendConnectionRepository(wsConn)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		var room *domain.Room
		switch reqAction {
		case "create-room":
			log.Println("Request create room")
			createRoomController := infrastructure.NewCreateRoomWsController(s.roomManager, c)
			room, err = createRoomController.Invoke(reqBody)
			if err != nil {
				log.Println(err)
				err = c.SendWebSocketMessage(err)
				log.Println(err)
				return
			}

			s.wsRouter(room, c, true, wg)

		case "join-room":
			log.Println("Request join room")
			joinRoomController := infrastructure.NewJoinRoomWsController(s.roomManager, c)
			room, err = joinRoomController.Invoke(reqBody)
			if err != nil {
				log.Println(err)
				_ = c.SendWebSocketMessage(err)
				return
			}

			s.wsRouter(room, c, false, wg)
		}

		wg.Wait()
	}))
}

func (s *Server) wsRouter(room *domain.Room, c domain.ConnectionRepository, isHost bool, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Println("Room activated")
	var player *domain.Player
	var enemy *domain.Player

	for player == nil || enemy == nil {
		if isHost {
			player = room.Player1
			enemy = room.Player2
		} else {
			player = room.Player2
			enemy = room.Player1
		}
	}

	cEnemy := infrastructure.NewBackendConnectionRepository(enemy.Ws)
	if isHost {
		err := application.StartGameAction(player, enemy, c, cEnemy, 10*60*1000)
		if err != nil {
			log.Println("Error starting game: ", err)
			_ = c.SendWebSocketMessage(err)
			_ = cEnemy.SendWebSocketMessage(err)
			return
		}
	}
	if player.Color == domain.WHITE {
		ok := room.Game.CalculateValidMoves(domain.WHITE)
		if !ok {
			log.Println("Error calculating first valid moves")
			return
		}
	}

	for {
		// Blocking when waiting for the enemy player action
		_, message, err := player.Ws.ReadMessage()
		if err != nil {
			log.Println("Some error:", err)
			player = nil
			return
		}

		reqAction, _ := jsonparser.GetString(message, "action")
		reqBody, _, _, _ := jsonparser.Get(message, "body")

		switch reqAction {
		case "request-moves":
			getValidMovesController := infrastructure.NewGetValidMovesWsController(c, room.Game)
			err := getValidMovesController.Invoke(reqBody)
			if err != nil {
				log.Println("Error getting valid moves: ", err)
			}

		case "move-piece":
			movePieceController := infrastructure.NewMovePieceWsController(c, cEnemy, room.Game)
			err := movePieceController.Invoke(reqBody)
			if err != nil {
				log.Println("Error getting move piece: ", err)
			}

			player.StopTimer()
			enemy.StartTimer()

			fmt.Println("Moved Player: ", player.ID, " color: ", player.Color, " Time left:", player.TimeLeft())
			fmt.Println("Turn Player: ", enemy.ID, " color: ", enemy.Color, " Time left:", enemy.TimeLeft())

		case "get-timers":
			getTimersController := infrastructure.NewGetTimersWsController(c, player, enemy)
			err := getTimersController.Invoke()
			if err != nil {
				log.Println("Error getting player timers: ", err)
			}

		case "abandon":
			abandonController := infrastructure.NewAbandonWsController(cEnemy)
			err := abandonController.Invoke()
			if err != nil {
				log.Println("Error abandon game: ", err)
			}
			// Connection is closed by the client
			// TODO: Remove room from room manager
		default:
			log.Println("Unknown action")
		}
	}
}

func (s *Server) initHTTP() {
	s.app.Get("/rooms", infrastructure.NewGetRoomsHTTPController(s.roomManager).Invoke)
}
