import type { Color } from '@/models/constants'
import { ReceiveParams } from './receive_params'
import type { Game } from '@/models/game'
import type { ReceiveAction } from './receive_action'

class StartGameParams extends ReceiveParams {
    playerColor: Color
    opponentName: string
    duration: number
}

class StartGameAction implements ReceiveAction {
    private game: Game

    constructor(game: Game) {
        this.game = game
    }

    public Invoke(body: string) {
        const p: StartGameParams = JSON.parse(body)
        this.game.start(p.opponentName, p.playerColor, p.duration)
    }
}

export { StartGameAction }
