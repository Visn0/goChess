import type { Game } from '@/models/game'
import { EndGameReason } from '@/models/constants'

class DrawResponseParams {
    public response: boolean
}
class ResponseDrawAction {
    private game: Game

    constructor(game: Game) {
        this.game = game
    }

    public Invoke(body: string) {
        const p: DrawResponseParams = JSON.parse(body)

        if (p.response) {
            this.game.setEndGameReason(EndGameReason.DRAW)
            this.game.setEndGame(true)
            this.game.repository.closeWebSocketConnection()
        }
    }
}

export { ResponseDrawAction }
