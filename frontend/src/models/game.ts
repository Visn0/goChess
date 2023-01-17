import type { Board } from './board'
import type { Square } from './square'
import type { ConnectionRepository } from './connection_repository/connection_repository'
import { RequestMovesAction } from './actions/send/request_moves_action'
import { MovePieceAction } from './actions/send/move_piece_action'
import type { Rooms } from './room'

class Game {
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

    constructor(rooms: Rooms, board: Board, repository: ConnectionRepository) {
        this.rooms = rooms
        this._board = board
        this.repository = repository

        this._srcSquare = null
    }

    public selectSquare(square: Square) {
        if (this._srcSquare !== null) {
            if (this._srcSquare.equals(square)) {
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

        this._srcSquare = new SrcSquare(square)
        RequestMovesAction(this.repository, square)
    }

    public unselectSrcSquare() {
        this._srcSquare?.removeValidMoves()
        this._srcSquare = null
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
        console.log('Set valid moves:', moves)
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

export { Game }
