import { ConnectionRepository, Message } from './connection_repository'

class BackendConnectionRepository implements ConnectionRepository {
    private host: string
    private port: string
    private wsPath: string

    private connection: WebSocket | null

    constructor(host: string, port: string, wsPath: string) {
        this.host = host
        this.port = port
        this.wsPath = wsPath
        this.connection = null
    }

    public openWebSocketConnection() {
        this.closeWebSocketConnection()

        const protocol = window.location.protocol.includes('s') ? 'wss' : 'ws'
        let url = `${protocol}://${this.host}:${this.port}`
        if (this.wsPath !== '') {
            url += `/${this.wsPath}`
        }

        this.connection = new WebSocket(url)
    }

    public closeWebSocketConnection() {
        if (this.connection !== null) {
            this.connection.close()
        }
    }

    public addOnWebSocketMessageEventListener(fn: (e: MessageEvent) => void) {
        if (this.connection === null) {
            return
        }

        this.connection.onmessage = fn
    }

    public sendWebSocketMessage(message: Message) {
        this.connection?.send(JSON.stringify(message))
    }

    public sendHTTPRequest(method: string, path: string, body: any): any {
        const protocol = window.location.protocol.includes('s') ? 'https' : 'http'
        let url = `${protocol}://${this.host}:${this.port}`
        if (path !== '') {
            url += `/${path}`
        }

        const params = {
            headers: {
                // 'content-type': 'application/json; charset=UTF-8',
            },
            body: body,
            method: method
        }


        console.log(url, params)
        fetch(url, params)
            .then((res) => {
                console.log(res)
                return res.body
            })
            .then((res) => console.log(res))
            .catch((error) => console.log(error))
    }
}

export { BackendConnectionRepository }
