import type { File, Rank } from '../../constants'
import type { GameController } from '../../game_controller'
import type { Square } from '../../square'

class PieceMovedParams {
    public src: CoordinateParams
    public dst: CoordinateParams
}

class CoordinateParams {
    public file: File
    public rank: Rank
}

class PieceMovedAction {
    private gameController: GameController

    constructor(gameController: GameController) {
        this.gameController = gameController
    }

    public Invoke(body: string) {
        const p: PieceMovedParams = JSON.parse(body)

        const srcSquare: Square = this.gameController.board.getSquare(p.src.file, p.src.rank)

        this.gameController.board.setSquarePiece(p.dst.file, p.dst.rank, srcSquare.piece)
        this.gameController.board.setSquarePiece(srcSquare.file, srcSquare.rank, null)
        this.gameController.unselectSrcSquare()
    }
}

export { PieceMovedAction }
