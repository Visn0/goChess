import { Color, constants } from './constants'
import { MovePiece } from './actions/move_piece'
import { Board } from './board'

class GameController {
    private serverURL: string
    private wsConn: WebSocket | null

    private board: Board

    constructor(board: Board, host: string, path: string) {
        this.wsConn = null
        this.board = board

        const protocol = window.location.protocol.includes('s') ? 'wss' : 'ws'
        this.serverURL = `${protocol}://${host}/${path}`
    }

    public start() {
        this.openWebSocketConnetion()
        this.board.initFromFenNotation(constants.StartingPosition)
        this.board.render(Color.WHITE)
    }

    private openWebSocketConnetion() {
        if (this.wsConn !== null) {
            this.wsConn.close()
        }

        this.wsConn = new WebSocket(this.serverURL)
        this.wsConn.onmessage = this.onWebSocketMessage.bind(this)
    }

    private onWebSocketMessage(event: MessageEvent) {
        const body = event.data
        class Params {
            action: string
        }

        const params: Params = JSON.parse(body)
        this.router(params.action, body)
    }

    private router(action: string, body: string) {
        switch (action) {
            case 'move-piece':
                MovePiece(this.board, body)
                break
            default:
                console.log('################')
                console.log('ACTION: ' + action)
                console.log('BODY: ' + action)
                console.log('################')
        }
    }
}

export { GameController }
