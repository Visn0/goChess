package application

// import (
// 	"chess/server/room"
// 	"chess/server/shared"
// 	"encoding/json"
// 	"log"
// )

// type ResponseSyncTimers struct {
// 	Own      int64 `json:"own"`
// 	Opponent int64 `json:"opponent"`
// }

// func WsSyncTimers(r room.Room) error {
// 	p1Response := &ResponseSyncTimers{
// 		Own:      r.Player1.TimeLeft(),
// 		Opponent: r.Player2.TimeLeft(),
// 	}
// 	err := sendTimersTo(r.Player1.Ws, p1Response)
// 	if err != nil {
// 		return err
// 	}

// 	p2Response := &ResponseSyncTimers{
// 		Own:      r.Player2.TimeLeft(),
// 		Opponent: r.Player1.TimeLeft(),
// 	}

// 	return sendTimersTo(r.Player1.Ws, p2Response)
// }

// func sendTimersTo(playerConn *shared.WsConn, res *ResponseSyncTimers) error {
// 	b, err := json.Marshal(res)
// 	if err != nil {
// 		log.Println("ERROR: WsSyncTimers: Error marshalling response: ", err)
// 		return err
// 	}
// 	err = playerConn.WriteJSON(b)
// 	if err != nil {
// 		log.Println("ERROR: WsSyncTimers: Error writing JSON: ", err)
// 		return err
// 	}

// 	return nil
// }
