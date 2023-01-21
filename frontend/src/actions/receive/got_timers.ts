import { ReceiveParams } from './receive_params'
import type { Game } from '@/models/game'

class GotTimersParams extends ReceiveParams {
    playerTime: number
    enemyTime: number
}

class GotTimersAction {
    private game: Game

    constructor(game: Game) {
        this.game = game
    }

    public Invoke(body: string) {
        const p: GotTimersParams = JSON.parse(body)

        this.game.setOwnTimer(p.playerTime)
        this.game.setOpponentTimer(p.playerTime)
    }
}

export { GotTimersAction }
