import { Board } from '../board'
import { File, Rank } from '../constants'
import { Square } from '../square'

class MovePieceParams {
    public src: CoordinateParams
    public dst: CoordinateParams
}

class CoordinateParams {
    public file: File
    public rank: Rank
}

function MovePiece(board: Board, body: string): void {
    const p: MovePieceParams = JSON.parse(body)

    const srcSquare: Square = board.getSquare(p.src.file, p.src.rank)

    board.setSquarePiece(p.dst.file, p.dst.rank, srcSquare.piece)
    board.setSquarePiece(srcSquare.file, srcSquare.rank, null)
}

export { MovePiece }
