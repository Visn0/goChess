package domain

type ConnectionRepository interface {
	SendWebSocketMessage(interface{}) error
	ReadMessage() ([]byte, error)
}
