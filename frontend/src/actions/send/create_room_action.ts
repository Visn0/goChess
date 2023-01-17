import type { ConnectionRepository, Message } from '@/models/connection_repository/connection_repository'

class CreateRoomMessage implements Message {
    action: string
    body: {
        playerID: string
        roomID: string
        password: string
    }

    constructor(playerID: string, roomID: string, password: string) {
        this.action = 'create-room'
        this.body = {
            playerID: playerID,
            roomID: roomID,
            password: password
        }
    }
}

function CreateRoomAction(repository: ConnectionRepository, playerID: string, roomID: string, roomPassword: string) {
    const m = new CreateRoomMessage(playerID, roomID, roomPassword)

    repository.openWebSocketConnection(() => {
        repository.sendWebSocketMessage(m)
    })
}

export { CreateRoomAction, CreateRoomMessage }
