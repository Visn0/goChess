package infrastructure

import (
	"chess/server/shared"
)

type BackendConnectionRepository struct {
	ws *shared.WsConn
}

func NewBackendConnectionRepository(ws *shared.WsConn) *BackendConnectionRepository {
	return &BackendConnectionRepository{
		ws: ws,
	}
}

func (r *BackendConnectionRepository) SendWebSocketMessage(msg interface{}) error {
	// b, err := json.Marshal(msg)
	// log.Println("===>", string(b))
	// if err != nil {
	// 	return err
	// }
	return r.ws.WriteJSON(msg)
}

func (r *BackendConnectionRepository) GetWebSocketConnection() *shared.WsConn {
	return r.ws
}
