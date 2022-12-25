import { Color, File, Rank, constants } from './constants'
import { PieceFactory } from './piece'
import { Square } from './square'

class Board {
    container: HTMLElement
    squares: Array<Array<Square>>

    constructor(container: HTMLElement) {
        this.container = container
        this.squares = new Array<Array<Square>>()
    }

    initFromFenNotation(notation: string) {
        this.squares = new Array<Array<Square>>(8)

        let notationIdx = 0
        for (let r = Rank._8; r >= Rank._1; r--) {
            const rank = new Array<Square>(8)

            for (let f = File.A; f <= File.H; f++, notationIdx++) {
                const c = notation[notationIdx]

                // Add empty squares
                if (this.isNumber(c)) {
                    for (let amount = parseInt(c); amount > 0; amount--, f++) {
                        rank[f] = new Square(f, r, null)
                        console.log(c, f)
                    }
                    f--
                    continue
                }

                rank[f] = new Square(f, r, PieceFactory(c))
            }

            this.squares[r] = rank
            notationIdx++ // Skip '/' char
        }
    }

    public reset() {
        this.initFromFenNotation(constants.StartingPosition)
    }

    private isNumber(s: string): boolean {
        return !Number.isNaN(Number(s))
    }

    public render(colorDown: Color) {
        let board = '<table>'
        if (colorDown === Color.WHITE) {
            for (let r = Rank._8; r >= Rank._1; --r) {
                board += '<tr>'
                for (let f = File.A; f <= File.H; ++f) {
                    board += this.squares[r][f].toTableCellHTML()
                }
                board += '</tr>\n'
            }
        } else {
            for (let r = Rank._1; r <= Rank._8; ++r) {
                board += '<tr>'
                for (let f = File.H; f >= File.A; --f) {
                    board += this.squares[r][f].toTableCellHTML()
                }
                board += '</tr>\n'
            }
        }

        board += '</table>'
        this.container.innerHTML = board
        console.log(this.squares)
    }
}

export { Board }
