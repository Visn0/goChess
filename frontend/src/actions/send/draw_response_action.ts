import type { ConnectionRepository, Message } from '@/models/connection_repository/connection_repository'

class DrawResponseMessage implements Message {
    action: string
    body: {
        response: boolean
    }

    constructor() {
        this.action = 'drawResponse'
    }
}

function DrawResponseAction(repository: ConnectionRepository, response: boolean) {
    const m = new DrawResponseMessage()
    m.body.response = response
    repository.sendWebSocketMessage(m)
}

export { DrawResponseAction, DrawResponseMessage }