package infrastructure

import (
	"chess/server/shared"
	"chess/server/shared/chesserror"
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
	err := r.ws.WriteJSON(msg)
	if err != nil {
		return chesserror.NewError(chesserror.GenericError, "Error writing message into Websocket").WithCause(err)
	}
	return nil
}

func (r *BackendConnectionRepository) ReadMessage() ([]byte, error) {
	_, msg, err := r.ws.ReadMessage()
	return msg, err
}
