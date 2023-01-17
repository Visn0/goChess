import type { CreateRoomMessage } from '@/actions/send/create_room_action'
import type { JoinRoomMessage } from '@/actions/send/join_room_action'
import type { MovePieceMessage } from '@/actions/send/move_piece_action'
import { File, Rank } from '../constants'
import type { ConnectionRepository, Message } from './connection_repository'

class MockConnectionRepository implements ConnectionRepository {
    private readonly onMessageEventTopic: string = 'mock-repository-onMessageEvent-topic'

    public openWebSocketConnection(onOpenListener?: (e: Event) => any | null) {
        console.log('=> Connection opened')
        if (onOpenListener) {
            onOpenListener(new CustomEvent(''))
        }
    }

    public closeWebSocketConnection() {
        console.log('=> Connection closed')
    }

    public addOnWebSocketMessageEventListener(fn: (e: MessageEvent) => void) {
        document.addEventListener(this.onMessageEventTopic, (e: Event) => {
            fn(e as MessageEvent)
        })
    }

    public sendWebSocketMessage(message: Message) {
        const me = this.messageEventFactory(message)
        document.dispatchEvent(me)
    }

    private messageEventFactory(message: Message): MessageEvent {
        switch (message.action) {
            case 'create-room':
                return this.createRoomCreatedMessage(message)
            case 'join-room':
                return this.createRoomJoinedMessage(message)
            case 'request-moves':
                return this.createMovesReceivedMessage()
            case 'move-piece':
                return this.createPieceMovedMessage(message)
            default:
                return new MessageEvent(this.onMessageEventTopic, {
                    data: { message: `INVALID ACTION: ${message.action}` }
                })
        }
    }

    private createRoomCreatedMessage(message: Message): MessageEvent {
        const m = message as CreateRoomMessage
        const data = JSON.stringify({
            action: 'create-room',
            httpCode: 200,
            room: {
                id: m.body.roomID,
                players: [
                    {
                        id: 'mockUser'
                    }
                ]
            }
        })

        return new MessageEvent(this.onMessageEventTopic, { data: data })
    }

    private createRoomJoinedMessage(message: Message): MessageEvent {
        const m = message as JoinRoomMessage
        const data = JSON.stringify({
            action: 'create-room',
            httpCode: 200,
            room: {
                id: m.body.roomID,
                players: [
                    {
                        id: 'mockUser'
                    },
                    {
                        id: m.body.playerID
                    }
                ]
            }
        })

        return new MessageEvent(this.onMessageEventTopic, { data: data })
    }

    private createMovesReceivedMessage(): MessageEvent {
        const moves = []
        for (let f = File.A; f <= File.H; f++) {
            for (let r = Rank._1; r <= Rank._8; r++) {
                moves.push({ file: f, rank: r })
            }
        }

        const data = JSON.stringify({ action: 'request-moves', validMoves: moves })
        return new MessageEvent(this.onMessageEventTopic, { data: data })
    }

    private createPieceMovedMessage(message: Message): MessageEvent {
        const m = message as MovePieceMessage
        const data = JSON.stringify({
            action: 'move-piece',
            src: {
                file: m.body.src.file,
                rank: m.body.src.rank
            },
            dst: {
                file: m.body.dst.file,
                rank: m.body.dst.rank
            }
        })

        return new MessageEvent(this.onMessageEventTopic, { data: data })
    }

    /* eslint-disable @typescript-eslint/no-unused-vars */
    public sendHTTPRequest(method: string, path: string, body: any): Promise<Response> {
        switch (path) {
            case 'rooms':
                return new Promise<Response>((resolve, reject): Response => {
                    const bodyStr = JSON.stringify({
                        rooms: [
                            {
                                id: 'room1',
                                players: []
                            },
                            {
                                id: 'room2',
                                players: [{ id: 'player1' }]
                            }
                        ]
                    })
                    const bodyRes = new Blob([bodyStr], {
                        type: 'application/json'
                    })
                    const options = { status: 200 }
                    const res = new Response(bodyRes, options)

                    resolve(res)
                    return res
                })

            default:
                return new Promise<Response>(() => null)
        }
    }
    /* eslint-enable @typescript-eslint/no-unused-vars */
}

export { MockConnectionRepository }
