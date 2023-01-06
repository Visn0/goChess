import { ConnectionRepository, Message } from '../../connection_repository/connection_repository'

class JoinRoomMessage implements Message {
    action: string
    body: {
        roomID: string,
        playerID: string
    }

    constructor(roomID: string, playerID: string) {
        this.action = 'join-room'
        this.body = {
            roomID: roomID,
            playerID: playerID
        }
    }
}

function JoinRoomAction(repository: ConnectionRepository, roomID: string, playerID: string) {
    const m = new JoinRoomMessage(roomID, playerID)

    repository.sendWebSocketMessage(m)
}

export { JoinRoomAction, JoinRoomMessage }
