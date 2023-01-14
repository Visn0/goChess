import type { File, Rank } from '../../constants'
import type { GameController } from '../../game_controller'
import type { Square } from '../../square'

class MovesReceivedParams {
    validMoves: Array<CoordinateParams>
}

class CoordinateParams {
    public file: File
    public rank: Rank
}

class MovesReceivedAction {
    private gameController: GameController

    constructor(gameController: GameController) {
        this.gameController = gameController
    }

    public Invoke(body: string) {
        const p: MovesReceivedParams = JSON.parse(body)

        const board = this.gameController.board
        const validMoves: Array<Square> = p.validMoves.map((m) => board.getSquare(m.file, m.rank))

        this.gameController.srcSquare?.setValidMoves(validMoves)
    }
}

export { MovesReceivedAction }
