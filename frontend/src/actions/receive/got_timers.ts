import { ReceiveParams } from './receive_params'
import type { Game } from '@/models/game'
import type { ReceiveAction } from './receive_action'

class GotTimersParams extends ReceiveParams {
    playerTime: number
    enemyTime: number
}

class GotTimersAction implements ReceiveAction {
    private game: Game

    constructor(game: Game) {
        this.game = game
    }

    public Invoke(body: string) {
        const p: GotTimersParams = JSON.parse(body)

        this.game.setOwnTimer(p.playerTime)
        this.game.setOpponentTimer(p.enemyTime)

        // console.log(p)
        // console.log(this.game.getOwnTimer().toString())
        // console.log(this.game.getOpponentTimer().toString())
    }
}

export { GotTimersAction }
