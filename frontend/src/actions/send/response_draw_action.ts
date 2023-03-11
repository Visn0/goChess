import type { ConnectionRepository, Message } from '@/models/connection_repository/connection_repository'

class ResponseDrawMessage implements Message {
    action: string
    body: {
        drawAccepted: boolean
    }

    constructor(drawAccepted: boolean) {
        this.action = 'drawResponse'
        this.body = { drawAccepted: drawAccepted }
    }
}

function ResponseDrawAction(repository: ConnectionRepository, drawAccepted: boolean) {
    const m = new ResponseDrawMessage(drawAccepted)
    repository.sendWebSocketMessage(m)
}

export { ResponseDrawAction, ResponseDrawMessage }
