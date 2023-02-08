import type { ConnectionRepository, Message } from '@/models/connection_repository/connection_repository'

class DrawRequestMessage implements Message {
    action: string
    body: null

    constructor() {
        this.action = 'draw-request'
    }
}

function DrawRequestAction(repository: ConnectionRepository) {
    const m = new DrawRequestMessage()
    repository.sendWebSocketMessage(m)
}

export { DrawRequestAction, DrawRequestMessage }