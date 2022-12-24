import { Square, File, Rank } from './square'

class Board {
    container: HTMLElement
    squares: Array<Array<Square>>

    constructor(container: HTMLElement) {
        this.container = container
        this.squares = new Array<Array<Square>>()
    }

    public reset() {
        this.squares = new Array<Array<Square>>(8)
        let board = '<table class="board"'

        for (let r = Rank._8; r >= Rank._1; r--) {
            const rank = new Array<Square>(8)
            board += '<tr>'
            for (let f = File.A; f <= File.H; f++) {
                const square = new Square(f, r)
                rank[f] = square

                board += square.toTableCellHTML()
            }
            board += '</tr>'

            this.squares[r] = rank
        }

        board += '</table>'
        this.container.innerHTML = board
    }
}

export { Board }
