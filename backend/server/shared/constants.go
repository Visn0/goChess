package shared

import (
	"sync"

	websocket "github.com/gofiber/websocket/v2"
)

type WsConn struct {
	*websocket.Conn
	lock sync.Mutex
}

func (ws *WsConn) WriteJSON(data interface{}) error {
	ws.lock.Lock()
	defer ws.lock.Unlock()
	return ws.Conn.WriteJSON(data)
}
