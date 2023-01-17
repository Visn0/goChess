import type { File, Rank } from '@/models/constants'
import type { Game } from '@/models/game'
import type { Square } from '@/models/square'

class PieceMovedParams {
    public src: CoordinateParams
    public dst: CoordinateParams
}

class CoordinateParams {
    public file: File
    public rank: Rank
}

class PieceMovedAction {
    private game: Game

    constructor(game: Game) {
        this.game = game
    }

    public Invoke(body: string) {
        const p: PieceMovedParams = JSON.parse(body)

        const srcSquare: Square = this.game.board.getSquare(p.src.file, p.src.rank)

        this.game.board.setSquarePiece(p.dst.file, p.dst.rank, srcSquare.piece)
        this.game.board.setSquarePiece(srcSquare.file, srcSquare.rank, null)
        this.game.unselectSrcSquare()
    }
}

export { PieceMovedAction }
