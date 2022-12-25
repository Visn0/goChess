import { Board } from '../board'
import { File, Rank } from '../constants'

class MovePieceParams {
    src: CoordinateParams
    dst: CoordinateParams
}

class CoordinateParams {
    file: File
    rank: Rank
}

function MovePiece(board: Board, body: string): void {
    const params: MovePieceParams = JSON.parse(body)

    console.log(params)
}

export { MovePiece }
