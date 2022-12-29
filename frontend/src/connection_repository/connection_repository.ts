interface Message {
    action: string
    body: any
}

interface ConnectionRepository {
    openConnection(): void
    closeConnection(): void
    addOnMessageEventListener(fn: (e: MessageEvent) => void): void
    sendMessage(message: Message): void
}

export { ConnectionRepository, Message }
