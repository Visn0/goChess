package wsrouter

import (
	"fmt"
	"log"
)

type WsHandler func(*Context) error

type WsRouter struct {
	handlers map[string]WsHandler
}

func NewWsRouter(paths map[string]WsHandler) *WsRouter {
	return &WsRouter{
		handlers: paths,
	}
}

func (r *WsRouter) Add(path string, wsHandler WsHandler) {
	if _, ok := r.handlers[path]; ok {
		panic(fmt.Sprintf("There already exists a path %q", path))
	}

	r.handlers[path] = wsHandler
}

func (r *WsRouter) Handle(path string, ctx *Context) {
	handler, ok := r.handlers[path]
	if !ok {
		log.Printf("Path not found: %q", path)
		return
	}

	log.Printf("==> Running handler for path %q.\n", path)
	log.Println("Input params:")
	log.Printf("%v", string(ctx.Body))
	err := handler(ctx)
	if err != nil {
		log.Printf("=> Error running handler for path %q: %v", path, err)
	}
}
