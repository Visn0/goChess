package room

import websocket "github.com/gofiber/websocket/v2"

type wsConn = websocket.Conn

type Player struct {
	Ws *wsConn
	ID string
}

type PlayerPublicInfo struct {
	ID string `json:"id"`
}

func (p *Player) GetPublicInfo() *PlayerPublicInfo {
	return &PlayerPublicInfo{
		ID: p.ID,
	}
}
