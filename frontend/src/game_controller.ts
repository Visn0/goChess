import { eventTopics } from './constants'
import { Board } from './board'
import { File, Rank } from './constants'
import { Square } from './square'
import { ConnectionRepository } from './backend/connection_repository'
import { ReceiveAction } from './actions/receive/receive_action'
import { CreateRoomAction } from './actions/send/create_room_action'
import { RoomCreatedAction } from './actions/receive/room_created_action'
import { MovesReceivedAction } from './actions/receive/moves_received_action'
import { RequestMovesAction } from './actions/send/request_moves_action'
import { MovePieceAction } from './actions/send/move_piece_action'
import { PieceMovedAction } from './actions/receive/piece_moved_action'

class GameController {
    private repository: ConnectionRepository
    private _board: Board
    public get board(): Board {
        return this._board
    }

    private _srcSquare: SrcSquare | null
    public get srcSquare(): SrcSquare | null {
        return this._srcSquare
    }

    private receiveActions: Map<string, ReceiveAction>

    constructor(board: Board, repository: ConnectionRepository) {
        this._board = board
        this.repository = repository

        document.addEventListener(eventTopics.OnSquareClick, (e: Event) => {
            this.onSquareClick(e as CustomEvent)
        })

        this._srcSquare = null

        this.registerReceiveActions()
    }

    private registerReceiveActions() {
        this.receiveActions = new Map<string, ReceiveAction>([
            ['room-created', new RoomCreatedAction()],
            ['moves-received', new MovesReceivedAction(this)],
            ['piece-moved', new PieceMovedAction(this.board)]
        ])
    }

    public createRoom(name: string, password: string) {
        CreateRoomAction(this.repository, name, password)
    }

    public openWebSocketConnetion() {
        this.repository.openConnection()
        this.repository.addOnMessageEventListener(this.onWebSocketMessage.bind(this))
    }

    private onWebSocketMessage(event: MessageEvent) {
        const body = JSON.stringify(event.data)

        class Params {
            action: string
        }

        const p: Params = JSON.parse(body)
        const action = this.receiveActions.get(p.action)
        console.log(p.action, action)

        action?.Invoke(body)
    }

    private onSquareClick(event: CustomEvent) {
        const file: File = event.detail.file
        const rank: Rank = event.detail.rank

        const square = this._board.getSquare(file, rank)
        if (this.isSrcSquareSelected()) {
            MovePieceAction(this.repository, this, square)
            this.unselectSrcSquare()
            return
        }

        if (square.isEmpty()) {
            return
        }

        this.selectSquare(square)
    }

    private selectSquare(square: Square) {
        this._srcSquare = new SrcSquare(square)
        RequestMovesAction(this.repository, square)
        console.log(this.isSrcSquareSelected())
    }

    private unselectSrcSquare() {
        this._srcSquare = null
    }

    private isSrcSquareSelected(): boolean {
        return this._srcSquare !== null
    }
}

class SrcSquare {
    public square: Square
    private validMoves: Array<Square>

    constructor(square: Square) {
        this.square = square
        this.validMoves = new Array<Square>(0)
    }

    public setValidMoves(moves: Array<Square>) {
        this.validMoves = moves
    }

    public canInnerPieceMoveTo(dst: Square): boolean {
        const found = this.validMoves.find((m: Square) => m.file === dst.file && m.rank === dst.rank)
        return found !== undefined
    }
}

export { GameController }
