import type { ConnectionRepository, Message } from '@/models/connection_repository/connection_repository'

class JoinRoomMessage implements Message {
    action: string
    body: {
        playerID: string
        roomID: string
        password: string
    }

    constructor(playerID: string, roomID: string, password: string) {
        this.action = 'join-room'
        this.body = {
            playerID: playerID,
            roomID: roomID,
            password: password
        }
    }
}

function JoinRoomAction(repository: ConnectionRepository, playerID: string, roomID: string, password: string) {
    const m = new JoinRoomMessage(playerID, roomID, password)

    repository.openWebSocketConnection(() => {
        repository.sendWebSocketMessage(m)
    })
}

export { JoinRoomAction, JoinRoomMessage }
