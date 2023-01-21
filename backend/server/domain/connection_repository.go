package domain

import "chess/server/shared"

type ConnectionRepository interface {
	SendWebSocketMessage(interface{}) error
	GetWebSocketConnection() *shared.WsConn
}
