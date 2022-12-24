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

enum File {
    A = 0,
    B = 1,
    C = 2,
    D = 3,
    E = 4,
    F = 5,
    G = 6,
    H = 7
}

enum Rank {
    _1 = 0,
    _2 = 1,
    _3 = 2,
    _4 = 3,
    _5 = 4,
    _6 = 5,
    _7 = 6,
    _8 = 7
}

export { Square, File, Rank }
