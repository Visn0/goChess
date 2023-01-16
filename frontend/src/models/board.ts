import { File, Rank, constants } from './constants'
import { Piece, PieceFactory } from './piece'
import { Square } from './square'

class Board {
    private squares: Array<Array<Square>>

    constructor() {
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

    public getSquare(file: File, rank: Rank): Square {
        return this.squares[rank][file]
    }

    public setSquarePiece(file: File, rank: Rank, piece: Piece | null) {
        const square = this.squares[rank][file]
        square.setPiece(piece)
    }
}

export { Board }
