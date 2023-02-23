import type { ConnectionRepository, Message } from '@/models/connection_repository/connection_repository'

class ResponseDrawMessage implements Message {
    action: string
    body: {
        drawResponse: boolean
    }

    constructor(drawResponse: boolean) {
        this.action = 'response-draw'
        this.body = { drawResponse: drawResponse }
    }
}

function ResponseDrawAction(repository: ConnectionRepository, drawResponse: boolean) {
    const m = new ResponseDrawMessage(drawResponse)
    repository.sendWebSocketMessage(m)

    if (drawResponse) { 
        repository.closeWebSocketConnection() 
    }
    
}

export { ResponseDrawAction, ResponseDrawMessage }
