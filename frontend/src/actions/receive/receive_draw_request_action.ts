import type { Game } from '@/models/game'
import { EndGameReason } from '@/models/constants'

class ReceiveDrawRequestAction {
    private game: Game

    constructor(game: Game) {
        this.game = game
    }

    public Invoke() {
        this.game.setEndGameReason(EndGameReason.DRAW_REQUEST)
        this.game.setEndGame(true)
    }
}

export { ReceiveDrawRequestAction }
