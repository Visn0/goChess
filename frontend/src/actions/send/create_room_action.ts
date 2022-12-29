import { ConnectionRepository, Message } from '../../connection_repository/connection_repository'

class CreateRoomMessage implements Message {
    action: string
    body: {
        name: string
        password: string
    }

    constructor(name: string, password: string) {
        this.action = 'create-room'
        this.body = {
            name: name,
            password: password
        }
    }
}

function CreateRoomAction(repository: ConnectionRepository, roomName: string, roomPassword: string) {
    const m = new CreateRoomMessage(roomName, roomPassword)
    repository.sendWebSocketMessage(m)
}

export { CreateRoomAction, CreateRoomMessage }
