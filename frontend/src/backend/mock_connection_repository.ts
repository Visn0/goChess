import { MovePieceMessage } from '../actions/send/move_piece_action'
import { File, Rank } from '../constants'
import { ConnectionRepository, Message } from './connection_repository'

class MockConnectionRepository implements ConnectionRepository {
    private readonly onMessageEventTopic: string = 'mock-repository-onMessageEvent-topic'

    public openConnection() {
        console.log('=> Connection opened')
    }

    public closeConnection() {
        console.log('=> Connection closed')
    }

    public addOnMessageEventListener(fn: (e: MessageEvent) => void) {
        document.addEventListener(this.onMessageEventTopic, (e: Event) => {
            fn(e as MessageEvent)
        })
    }

    public sendMessage(message: Message) {
        const me = this.messageEventFactory(message)
        document.dispatchEvent(me)
    }

    private messageEventFactory(message: Message): MessageEvent {
        switch (message.action) {
            case 'create-room':
                return new MessageEvent(this.onMessageEventTopic, { data: { action: 'room-created' } })
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

    private createMovesReceivedMessage(): MessageEvent {
        const moves = []
        for (let f = File.A; f <= File.H; f++) {
            for (let r = Rank._1; r <= Rank._8; r++) {
                moves.push({ file: f, rank: r })
            }
        }

        const data = { action: 'moves-received', validMoves: moves }
        return new MessageEvent(this.onMessageEventTopic, { data: data })
    }

    private createPieceMovedMessage(message: Message): MessageEvent {
        const m = message as MovePieceMessage
        const data = {
            action: 'piece-moved',
            src: {
                file: m.body.src.file,
                rank: m.body.src.rank
            },
            dst: {
                file: m.body.dst.file,
                rank: m.body.dst.rank
            }
        }
        return new MessageEvent(this.onMessageEventTopic, { data: data })
    }
}

export { MockConnectionRepository }
