import { eventTopics } from './constants'
import { Board } from './board'
import { File, Rank } from './constants'
import { Square } from './square'
import { ConnectionRepository } from './connection_repository/connection_repository'
import { ReceiveAction } from './actions/receive/receive_action'
import { CreateRoomAction } from './actions/send/create_room_action'
import { RoomCreatedAction } from './actions/receive/room_created_action'
import { MovesReceivedAction } from './actions/receive/moves_received_action'
import { RequestMovesAction } from './actions/send/request_moves_action'
import { MovePieceAction } from './actions/send/move_piece_action'
import { PieceMovedAction } from './actions/receive/piece_moved_action'
import { Rooms } from './room'

class GameController {
    private repository: ConnectionRepository
    private _board: Board
    private rooms: Rooms

    public get board(): Board {
        return this._board
    }

    private _srcSquare: SrcSquare | null
    public get srcSquare(): SrcSquare | null {
        return this._srcSquare
    }

    private receiveActions: Map<string, ReceiveAction>

    constructor(rooms: Rooms, board: Board, repository: ConnectionRepository) {
        this.rooms = rooms
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
            ['create-room', new RoomCreatedAction(this.rooms)],
            ['request-moves', new MovesReceivedAction(this)],
            ['move-piece', new PieceMovedAction(this)]
        ])
    }

    public createRoom(name: string, password: string) {
        CreateRoomAction(this.repository, name, password)
    }

    public openWebSocketConnection() {
        this.repository.openWebSocketConnection()
        this.repository.addOnWebSocketMessageEventListener(this.onWebSocketMessage.bind(this))
    }

    private onWebSocketMessage(event: MessageEvent) {
        const body = event.data
        class Params {
            action: string
        }

        const p: Params = JSON.parse(body)
        const action = this.receiveActions.get(p.action)

        action?.Invoke(body)
    }

    private onSquareClick(event: CustomEvent) {
        const file: File = event.detail.file
        const rank: Rank = event.detail.rank

        const square = this._board.getSquare(file, rank)
        if (this.isSrcSquareSelected()) {
            if (this._srcSquare?.equals(square)) {
                this.unselectSrcSquare()
                return
            }

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
    }

    public unselectSrcSquare() {
        this._srcSquare?.removeValidMoves()
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
        moves.forEach((m) => m.setAsValidMove())
    }

    public removeValidMoves() {
        this.validMoves.forEach((m) => m.unsetAsValidMove())
    }

    public canInnerPieceMoveTo(dst: Square): boolean {
        const found = this.validMoves.find((m: Square) => m.file === dst.file && m.rank === dst.rank)
        return found !== undefined
    }

    public equals(s: Square): boolean {
        return this.square.equals(s)
    }
}

export { GameController }
