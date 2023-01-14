interface Message {
    action: string
    body: any
}

interface ConnectionRepository {
    openWebSocketConnection(onOpenListener?: (e: Event) => any | null): void
    closeWebSocketConnection(): void
    addOnWebSocketMessageEventListener(fn: (e: MessageEvent) => void): void
    sendWebSocketMessage(message: Message): void

    sendHTTPRequest(method: string, path: string, body: any): Promise<Response>
}

export { ConnectionRepository, Message }
