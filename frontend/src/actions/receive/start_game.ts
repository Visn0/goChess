import type { Color } from '@/models/constants'
import { ReceiveParams } from './receive_params'
import type { Game } from '@/models/game'

class StartGameParams extends ReceiveParams {
    playerColor: Color
    duration: number
}

class StartGameAction {
    private game: Game

    constructor(game: Game) {
        this.game = game
    }

    public Invoke(body: string) {
        const p: StartGameParams = JSON.parse(body)
        this.game.start(p.playerColor, p.duration)
    }
}

export { StartGameAction }
