interface Message {
    action: string
    body: any
}

interface ConnectionRepository {
    openWebSocketConnection(): void
    closeWebSocketConnection(): void
    addOnWebSocketMessageEventListener(fn: (e: MessageEvent) => void): void
    sendWebSocketMessage(message: Message): void

    sendHTTPRequest(method: string, path: string, body: any): any
}

export { ConnectionRepository, Message }
