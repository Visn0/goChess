package server

import (
	"fmt"
	"log"
)

type WsHandler func([]byte) error

type WsRouter struct {
	handlers map[string]WsHandler
}

func NewWsRouter() *WsRouter {
	return &WsRouter{
		handlers: make(map[string]WsHandler),
	}
}

func (r *WsRouter) Add(path string, wsHandler WsHandler) {
	if _, ok := r.handlers[path]; ok {
		panic(fmt.Sprintf("There already exists a path %q", path))
	}

	r.handlers[path] = wsHandler
}

func (r *WsRouter) Handle(path string, body []byte) {
	handler, ok := r.handlers[path]
	if !ok {
		log.Printf("Path not found: %q", path)
		return
	}

	log.Printf("=> Running handler for path %q.\n", path)
	err := handler(body)
	if err != nil {
		log.Printf("=> Error running handler for path %q: %v", path, err)
	}
}
