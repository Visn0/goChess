package domain

type ConnectionRepository interface {
	SendWebSocketMessage(interface{}) error
	GetWebSocketConnection() *wsConn
}
