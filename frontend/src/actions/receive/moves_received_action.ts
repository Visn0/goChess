import type { File, Rank } from '@/models/constants'
import type { Game } from '@/models/game'
import type { Square } from '@/models/square'
import type { ReceiveAction } from './receive_action'

class MovesReceivedParams {
    validMoves: Array<CoordinateParams>
}

class CoordinateParams {
    public file: File
    public rank: Rank
}

class MovesReceivedAction implements ReceiveAction {
    private game: Game

    constructor(game: Game) {
        this.game = game
    }

    public Invoke(body: string) {
        const p: MovesReceivedParams = JSON.parse(body)

        console.log(p)
        const board = this.game.board
        const validMoves: Array<Square> = p.validMoves.map((m) => board.getSquare(m.file, m.rank))

        if (validMoves.length < 1) {
            this.game.unselectSrcSquare()
            return
        }

        this.game.srcSquare?.setValidMoves(validMoves)
    }
}

export { MovesReceivedAction }
