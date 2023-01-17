import type { File, Rank } from '../../constants'
import type { Game } from '../../game'
import type { Square } from '../../square'

class MovesReceivedParams {
    validMoves: Array<CoordinateParams>
}

class CoordinateParams {
    public file: File
    public rank: Rank
}

class MovesReceivedAction {
    private game: Game

    constructor(game: Game) {
        this.game = game
    }

    public Invoke(body: string) {
        const p: MovesReceivedParams = JSON.parse(body)
        console.log('Moves received:', p)

        const board = this.game.board
        const validMoves: Array<Square> = p.validMoves.map((m) => board.getSquare(m.file, m.rank))

        console.log('srcSquare?.setValidMoves: ', this.game.srcSquare)
        this.game.srcSquare?.setValidMoves(validMoves)
    }
}

export { MovesReceivedAction }
