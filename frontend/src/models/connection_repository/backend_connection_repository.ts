import type { ConnectionRepository, Message } from './connection_repository'

class BackendConnectionRepository implements ConnectionRepository {
    private url: string
    private wsPath: string

    private connection: WebSocket | null
    private onWebSocketMessageListener: (e: MessageEvent) => void

    constructor(url: string, wsPath: string) {
        this.url = url
        this.wsPath = wsPath
        this.connection = null
    }

    public openWebSocketConnection(onOpenListener?: (e: Event) => any | null) {
        this.closeWebSocketConnection()

        const protocol = window.location.protocol.includes('s') ? 'wss' : 'ws'
        let url = `${protocol}://${this.url}`
        if (this.wsPath !== '') {
            url += `/${this.wsPath}`
        }

        console.log('WS url: ', url)
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

    public async sendWebSocketMessage(message: Message) {
        if (!this.connection) {
            return
        }

        const sleep = (ms: number) => new Promise((resolve) => setTimeout(resolve, ms))
        for (let i = 0; i < 10; i++) {
            if (this.connection.readyState === WebSocket.OPEN) {
                this.connection.send(JSON.stringify(message))
                return
            }

            await sleep(200)
        }
    }

    public sendHTTPRequest(method: string, path: string, body: any): Promise<Response> {
        const protocol = window.location.protocol.includes('s') ? 'https' : 'http'
        let url = `${protocol}://${this.url}`
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
