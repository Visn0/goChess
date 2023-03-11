import type { ConnectionRepository, Message } from '@/models/connection_repository/connection_repository'

class RequestDrawMessage implements Message {
    action: string
    body: null

    constructor() {
        this.action = 'draw-request'
    }
}

function RequestDrawAction(repository: ConnectionRepository) {
    const m = new RequestDrawMessage()
    repository.sendWebSocketMessage(m)
}

export { RequestDrawAction, RequestDrawMessage }
