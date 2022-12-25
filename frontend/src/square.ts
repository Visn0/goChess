import { Color, File, Rank } from './constants'
import { Piece } from './piece'

class Square {
    readonly id: string
    readonly color: string
    readonly file: File
    readonly rank: Rank

    piece: Piece | null

    constructor(file: File, rank: Rank, piece: Piece | null) {
        this.file = file
        this.rank = rank
        this.piece = piece

        this.id = `${String.fromCharCode(65 + file)}${rank + 1}`
        this.color = (file + rank) % 2 === 0 ? `${Color.BLACK}-square` : `${Color.WHITE}-square`
    }

    toTableCellHTML(): string {
        const piece = this.piece?.toDivHTMLString() || ''
        const cell = `<td id="${this.id}" class="board-square ${this.color}">${piece} ${this.id}</td>`
        return cell
    }
}

export { Square }
