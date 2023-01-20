import type { Board } from './board'
import type { ConnectionRepository } from './connection_repository/connection_repository'
import { RequestMovesAction } from '@/actions/send/request_moves_action'
import { MovePieceAction } from '@/actions/send/move_piece_action'
import { Color, PieceType, Rank } from './constants'
import type { Piece } from './piece'
import { ref, type Ref } from 'vue'
import type { Square } from './square'

class Game {
    private _repository: ConnectionRepository
    public get repository(): ConnectionRepository {
        return this._repository
    }

    private _board: Board
    public get board(): Board {
        return this._board
    }

    private _srcSquare: SrcSquare | null
    public get srcSquare(): SrcSquare | null {
        return this._srcSquare
    }

    private dstPromotedPawn: Square | null
    private _pendingPromotion: Ref<boolean>
    public isPromotionPending(): boolean {
        return this._pendingPromotion.value
    }

    constructor(board: Board, repository: ConnectionRepository) {
        this._board = board
        this._repository = repository
        this.dstPromotedPawn = null
        this._pendingPromotion = ref(false)

        this._srcSquare = null
    }

    public selectSquare(square: Square) {
        if (this._srcSquare !== null) {
            if (this._srcSquare.equals(square)) {
                this.unselectSrcSquare()
                return
            }

            if (!this._srcSquare.canInnerPieceMoveTo(square)) {
                this.unselectSrcSquare()
                return
            }

            if (this.canPromote(this._srcSquare, square)) {
                this.dstPromotedPawn = square
                this._pendingPromotion.value = true
                // show piece options to player
                return
            }

            MovePieceAction(this.repository, this, square, null)
            this.unselectSrcSquare()
            return
        }

        if (square.isEmpty()) {
            return
        }

        square.setAsSelected()
        this._srcSquare = new SrcSquare(square)

        RequestMovesAction(this.repository, square)
    }

    private canPromote(src: SrcSquare, dst: Square): boolean {
        const piece = src.square.piece as Piece
        if (piece.type !== PieceType.PAWN) {
            return false
        }

        if (dst.rank === Rank._8) {
            return piece.color === Color.WHITE
        }

        if (dst.rank === Rank._1) {
            return piece.color === Color.BLACK
        }

        return false
    }

    public promotePiece(pieceType: PieceType) {
        if (this.isPromotionPending() && this.dstPromotedPawn) {
            MovePieceAction(this.repository, this, this.dstPromotedPawn, pieceType)
        }

        this._pendingPromotion.value = false
        this.dstPromotedPawn = null
    }

    public cancelPromotion() {
        this._pendingPromotion.value = false
        this.dstPromotedPawn = null
    }

    public unselectSrcSquare() {
        this._srcSquare?.square.unsetAsSelected()
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
