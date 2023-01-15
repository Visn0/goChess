import { ref, type Ref } from 'vue'
import { Color, File, Rank, eventTopics } from './constants'
import type { Piece } from './piece'

class Square {
    public readonly id: string
    public readonly color: string
    public readonly file: File
    public readonly rank: Rank

    private _piece: Piece | null
    public get piece(): Piece | null {
        return this._piece
    }

    private _isValidMove: Ref<boolean>
    public get isValidMove(): boolean {
        return this._isValidMove.value
    }

    private _updateRender: Ref<number>
    public get updateRender(): number {
        return this._updateRender.value
    }

    constructor(file: File, rank: Rank, piece: Piece | null) {
        this.file = file
        this.rank = rank
        this._piece = piece
        this._isValidMove = ref(false)
        this._updateRender = ref(1)

        this.id = `${String.fromCharCode(65 + file)}${rank + 1}`
        this.color = (file + rank) % 2 === 0 ? `${Color.BLACK}-square` : `${Color.WHITE}-square`
    }

    public toTableCellHTML(): string {
        const piece = this._piece?.toDivHTMLString() || ''
        const cell = `
            <td id="${this.id}" class="board-square ${this.color}">
                <div id="${this.id}-piece" class="piece">
                    ${piece}
                </div>
                <div id="${this.id}-valid-move" class="board-square valid-move" hidden>                    
                </div>                
            </td>
        `
        return cell
    }

    public isEmpty(): boolean {
        return this._piece === null
    }

    private updatePieceRender() {
        this._updateRender.value++
    }

    public setPiece(piece: Piece | null) {
        this._piece = piece
        this.updatePieceRender()
    }

    public setAsValidMove() {
        this._isValidMove.value = true
        this.updatePieceRender()
    }

    public unsetAsValidMove() {
        this._isValidMove.value = false
        this.updatePieceRender()
    }

    public equals(s: Square): boolean {
        return this.file === s.file && this.rank === s.rank
    }
}

export { Square }
