import { File, Rank } from './utils'

class Square {
    readonly id: string
    readonly color: string
    readonly file: File
    readonly rank: Rank

    constructor(file: File, rank: Rank) {
        this.file = file
        this.rank = rank

        this.id = `${String.fromCharCode(65 + file)}${rank + 1}`
        this.color = (file + rank) % 2 === 0 ? 'black-square' : 'white-square'
    }

    toTableCellHTML(): string {
        const board = `<td id="${this.id}" class="board-square ${this.color}">${this.id}</td>`
        return board
    }
}

export { Square, File, Rank }
