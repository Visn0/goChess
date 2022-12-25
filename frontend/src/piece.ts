import { Color, PieceType } from './constants'

class Piece {
    readonly color: Color
    readonly type: PieceType

    constructor(color: Color, type: PieceType) {
        this.color = color
        this.type = type
    }

    public toDivHTMLString(): string {
        const str = `<div class="piece ${this.color}-${this.type}"></div>`
        return str
    }
}

function PieceFactory(fenNotationChar: string): Piece {
    switch (fenNotationChar[0]) {
        // White Pieces
        case 'P':
            return new Piece(Color.WHITE, PieceType.PAWN)
        case 'R':
            return new Piece(Color.WHITE, PieceType.ROOK)
        case 'N':
            return new Piece(Color.WHITE, PieceType.KNIGHT)
        case 'B':
            return new Piece(Color.WHITE, PieceType.BISHOP)
        case 'Q':
            return new Piece(Color.WHITE, PieceType.QUEEN)
        case 'K':
            return new Piece(Color.WHITE, PieceType.KING)

        // Black Pieces
        case 'p':
            return new Piece(Color.BLACK, PieceType.PAWN)
        case 'r':
            return new Piece(Color.BLACK, PieceType.ROOK)
        case 'n':
            return new Piece(Color.BLACK, PieceType.KNIGHT)
        case 'b':
            return new Piece(Color.BLACK, PieceType.BISHOP)
        case 'q':
            return new Piece(Color.BLACK, PieceType.QUEEN)
        case 'k':
            return new Piece(Color.BLACK, PieceType.KING)
        default:
            throw `Invalid string in fen notation: ${fenNotationChar}`
    }
}

export { Piece, PieceFactory }
