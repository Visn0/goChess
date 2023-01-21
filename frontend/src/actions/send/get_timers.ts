import type { ConnectionRepository, Message } from '@/models/connection_repository/connection_repository'

class GetTimersMessage implements Message {
    action: string
    body: null

    constructor() {
        this.action = 'get-timers'
    }
}

function GetTimersAction(repository: ConnectionRepository) {
    const m = new GetTimersMessage()
    repository.sendWebSocketMessage(m)
}

export { GetTimersAction, GetTimersMessage }
