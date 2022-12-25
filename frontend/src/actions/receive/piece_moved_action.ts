import { Board } from '../../board'
import { File, Rank } from '../../constants'
import { Square } from '../../square'

class PieceMovedParams {
    public src: CoordinateParams
    public dst: CoordinateParams
}

class CoordinateParams {
    public file: File
    public rank: Rank
}

class PieceMovedAction {
    private board: Board

    constructor(board: Board) {
        this.board = board
    }

    public Invoke(body: string) {
        const p: PieceMovedParams = JSON.parse(body)

        const srcSquare: Square = this.board.getSquare(p.src.file, p.src.rank)

        this.board.setSquarePiece(p.dst.file, p.dst.rank, srcSquare.piece)
        this.board.setSquarePiece(srcSquare.file, srcSquare.rank, null)
    }
}

export { PieceMovedAction }
