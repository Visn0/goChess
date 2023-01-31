import type { Game } from '@/models/game'

class AbandonAction {
    private game: Game

    constructor(game: Game) {
        this.game = game
    }

    public Invoke() {
        this.game.setAbandoned(true)
        this.game.repository.closeWebSocketConnection()
    }
}

export { AbandonAction }
