import { ConnectionRepository, Message } from './connection_repository'

class BackendConnectionRepository implements ConnectionRepository {
    private url: string
    private connection: WebSocket | null

    constructor(host: string, port: string, path: string) {
        const protocol = window.location.protocol.includes('s') ? 'wss' : 'ws'

        this.url = `${protocol}://${host}:${port}`
        if (path !== '') {
            this.url += `/${path}`
        }
    }

    public openConnection() {
        this.closeConnection()
        this.connection = new WebSocket(this.url)
    }

    public closeConnection() {
        if (this.connection !== null) {
            this.connection.close()
        }
    }

    public addOnMessageEventListener(fn: (e: MessageEvent) => void) {
        if (this.connection === null) {
            return
        }

        this.connection.onmessage = fn
    }

    public sendMessage(message: Message) {
        this.connection?.send(JSON.stringify(message))
    }
}

export { BackendConnectionRepository }
