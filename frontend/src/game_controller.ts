import { Color, constants, eventTopics } from './constants'
import { MovePiece } from './actions/move_piece'
import { Board } from './board'
import { File, Rank } from './constants'
import { Square } from './square'

class SrcSquare {
    private square: Square
    private validMoves: Array<Square>

    constructor(square: Square) {
        this.square = square
        this.validMoves = new Array<Square>(0)
    }

    public setValidMoves(moves: Array<Square>) {
        this.validMoves = moves
    }

    public canInnerPieceMoveTo(dst: Square): boolean {
        const found = this.validMoves.find(m => (m.file === dst.file && m.rank === dst.rank))
        return found !== undefined
    }
}
class GameController {
    private serverURL: string
    private wsConn: WebSocket | null

    private board: Board

    private srcSquare: SrcSquare | null

    constructor(board: Board, host: string, path: string) {
        this.wsConn = null
        this.board = board

        const protocol = window.location.protocol.includes('s') ? 'wss' : 'ws'
        this.serverURL = `${protocol}://${host}/${path}`

        document.addEventListener(eventTopics.OnSquareClick, (e: Event) => {
            this.onSquareClick(e as CustomEvent)
        })

        this.srcSquare = null
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

    private onSquareClick(event: CustomEvent) {
        const file: File = event.detail.file
        const rank: Rank = event.detail.rank

        const square = this.board.getSquare(file, rank)
        if (this.isSrcSquareSelected()) {
            if (this.srcSquare?.canInnerPieceMoveTo(square)) {
                const body = {
                    src: this.srcSquare,
                    dst: square
                }

                MovePiece(this.board, JSON.stringify(body))
            }

            this.unselectSrcSquare()
            return
        }

        if (square.isEmpty()) {
            return
        }

        this.selectSquare(square)
    }

    private selectSquare(square: Square) {
        this.srcSquare = new SrcSquare(square)
        console.log('SRC selected: ', square)
    }

    private unselectSrcSquare() {
        this.srcSquare = null
    }

    private isSrcSquareSelected(): boolean {
        return this.srcSquare !== null
    }
}

export { GameController }
