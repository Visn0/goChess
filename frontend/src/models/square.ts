import { ref, type Ref } from 'vue'
import { Color, File, Rank } from './constants'
import type { Piece } from './piece'

class Square {
    public readonly id: string
    public readonly color: string
    public readonly file: File
    public readonly rank: Rank

    private _piece: Ref<Piece | null>
    public get piece(): Piece | null {
        return this._piece.value
    }
    public getPiece(): Piece | null {
        return this._piece.value
    }

    private _isValidMove: Ref<boolean>
    public isValidMove(): boolean {
        return this._isValidMove.value
    }

    constructor(file: File, rank: Rank, piece: Piece | null) {
        this.file = file
        this.rank = rank
        this._piece = ref(piece)
        this._isValidMove = ref(false)

        this.id = `${String.fromCharCode(65 + file)}${rank + 1}`
        this.color = (file + rank) % 2 === 0 ? `${Color.BLACK}-square` : `${Color.WHITE}-square`
    }

    public isEmpty(): boolean {
        return this.piece === null
    }

    public setPiece(piece: Piece | null) {
        this._piece.value = piece
    }

    public setAsValidMove() {
        this._isValidMove.value = true
    }

    public unsetAsValidMove() {
        this._isValidMove.value = false
    }

    public equals(s: Square): boolean {
        return this.file === s.file && this.rank === s.rank
    }
}

export { Square }
