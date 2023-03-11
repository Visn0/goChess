import type { Game } from '@/models/game'
import { EndGameReason } from '@/models/constants'
import type { ReceiveAction } from './receive_action'

class AbandonAction implements ReceiveAction {
    private game: Game

    constructor(game: Game) {
        this.game = game
    }

    public Invoke() {
        this.game.setEndGameReason(EndGameReason.ABANDON)
        this.game.setEndGame(true)
        this.game.repository.closeWebSocketConnection()
    }
}

export { AbandonAction }
