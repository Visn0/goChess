import type { Game } from '@/models/game'

class DrawRequestAction {
    private game: Game

    constructor(game: Game) {
        this.game = game
    }

    public Invoke() {
        this.game.setEndGameReason('draw-request')
        this.game.setEndGame(true)
    }
}

export { DrawRequestAction }