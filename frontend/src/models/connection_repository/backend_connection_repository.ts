import type { ConnectionRepository, Message } from './connection_repository'

class BackendConnectionRepository implements ConnectionRepository {
    private host: string
    private port: string
    private wsPath: string

    private connection: WebSocket | null
    private onWebSocketMessageListener: (e: MessageEvent) => void

    constructor(host: string, port: string, wsPath: string) {
        this.host = host
        this.port = port
        this.wsPath = wsPath
        this.connection = null
    }

    public openWebSocketConnection(onOpenListener?: (e: Event) => any | null) {
        this.closeWebSocketConnection()

        const protocol = window.location.protocol.includes('s') ? 'wss' : 'ws'
        let url = `${protocol}://${this.host}:${this.port}`
        if (this.wsPath !== '') {
            url += `/${this.wsPath}`
        }

        this.connection = new WebSocket(url)
        if (onOpenListener) {
            this.connection.onopen = onOpenListener
        }

        this.addOnWebSocketMessageEventListener(this.onWebSocketMessageListener)
    }

    public closeWebSocketConnection() {
        if (this.connection !== null) {
            this.connection.close()
        }
    }

    public addOnWebSocketMessageEventListener(fn: (e: MessageEvent) => void) {
        if (this.connection === null) {
            this.onWebSocketMessageListener = fn
            return
        }

        this.connection.onmessage = fn
    }

    public sendWebSocketMessage(message: Message) {
        this.connection?.send(JSON.stringify(message))
    }

    public sendHTTPRequest(method: string, path: string, body: any): Promise<Response> {
        const protocol = window.location.protocol.includes('s') ? 'https' : 'http'
        let url = `${protocol}://${this.host}:${this.port}`
        if (path !== '') {
            url += `/${path}`
        }

        const headers = {
            'Content-Type': 'application/json; charset=UTF-8'
        }

        let params = null
        if (method === 'GET' || body === null || body === '' || body === undefined) {
            params = {
                headers: headers,
                method: method
            }
        } else {
            params = {
                headers: headers,
                method: method,
                body: JSON.stringify(body)
            }
        }

        return fetch(url, params)
    }
}

export { BackendConnectionRepository }
