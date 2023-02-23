import type { Board } from './board'
import type { ConnectionRepository } from './connection_repository/connection_repository'
import { RequestMovesAction } from '@/actions/send/request_moves_action'
import { MovePieceAction } from '@/actions/send/move_piece_action'
import { Color, PieceType, Rank } from './constants'
import type { Piece } from './piece'
import { ref, shallowRef, type Ref } from 'vue'
import type { Square } from './square'
import { Timer } from './timer'
import { GetTimersAction } from '@/actions/send/get_timers'

class Game {
    private _repository: ConnectionRepository
    public get repository(): ConnectionRepository {
        return this._repository
    }

    private _started: Ref<boolean>
    public started(): boolean {
        return this._started.value
    }

    private _endGame: Ref<boolean>
    public setEndGame(n: boolean) {
        this._endGame.value = n
    }
    public endGame(): boolean {
        return this._endGame.value
    }

    private endGameReason: string
    public setEndGameReason(n: string) {
        this.endGameReason = n
    }
    public getEndGameReason(): string {
        return this.endGameReason
    }

    private opponentID: string
    public getOpponentID(): string {
        return this.opponentID
    }

    private myColor: Color
    public getMyColor(): Color {
        return this.myColor
    }

    private myTurn: Ref<boolean>
    public isMyTurn(): boolean {
        return this.myTurn.value
    }
    public changeTurn() {
        if (this.myTurn.value) {
            this.ownTimer.value.pause()
            this.opponentTimer.value.start()
        } else {
            this.ownTimer.value.start()
            this.opponentTimer.value.pause()
        }
        this.myTurn.value = !this.myTurn.value
    }

    private timerInterval: number

    private ownTimer: Ref<Timer>
    public getOwnTimer(): Timer {
        return this.ownTimer.value
    }
    public setOwnTimer(milliseconds: number) {
        this.ownTimer.value.setRemainingTime(milliseconds)
    }

    private opponentTimer: Ref<Timer>
    public getOpponentTimer(): Timer {
        return this.opponentTimer.value
    }
    public setOpponentTimer(milliseconds: number) {
        this.opponentTimer.value.setRemainingTime(milliseconds)
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

    private kingCheckSquare: Square | null
    public setKingCheck(square: Square | null) {
        if (this.kingCheckSquare) {
            this.kingCheckSquare.unsetAsKingCheck()
        }

        if (square) {
            square.setAsKingCheck()
        }

        this.kingCheckSquare = square
    }

    constructor(board: Board, repository: ConnectionRepository) {
        this._board = board
        this._repository = repository
        this._started = ref(false)
        this._endGame = ref(false)
        this.endGameReason = ''
        this.opponentID = ''
        this.myColor = Color.WHITE
        this.myTurn = ref(false)
        this.ownTimer = shallowRef(new Timer(0))
        this.opponentTimer = shallowRef(new Timer(0))
        this.dstPromotedPawn = null
        this._pendingPromotion = ref(false)

        this._srcSquare = null
    }

    public start(opponentID: string, myColor: Color, durationMs: number) {
        this.opponentID = opponentID
        this.myColor = myColor
        this.ownTimer.value.durationMs = durationMs + 400
        this.opponentTimer.value.durationMs = durationMs + 400

        if (myColor === Color.WHITE) {
            this.myTurn.value = true
            this.ownTimer.value.start()
        } else {
            this.opponentTimer.value.start()
        }

        this._started.value = true
        this.timerInterval = setInterval(() => GetTimersAction(this.repository), 10000)
    }

    public selectSquare(square: Square) {
        if (!this.isMyTurn()) {
            return
        }

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
        this.unselectSrcSquare()
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
