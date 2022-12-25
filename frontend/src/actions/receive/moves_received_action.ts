import { File, Rank } from '../../constants'
import { GameController } from '../../game_controller'
import { Square } from '../../square'

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
        console.log('params:', p)
        const board = this.gameController.board
        const validMoves: Array<Square> = p.validMoves.map((m) => board.getSquare(m.file, m.rank))

        this.gameController.srcSquare?.setValidMoves(validMoves)
        console.log(this.gameController.srcSquare)
    }
}

export { MovesReceivedAction }
