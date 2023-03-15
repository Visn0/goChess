import type { ConnectionRepository, Message } from '@/models/connection_repository/connection_repository'

class AbandonMessage implements Message {
    action: string
    body: null

    constructor() {
        this.action = 'abandon'
    }
}

function AbandonAction(repository: ConnectionRepository) {
    const m = new AbandonMessage()
    repository.sendWebSocketMessage(m)
    repository.closeWebSocketConnection()
}

export { AbandonAction, AbandonMessage }
