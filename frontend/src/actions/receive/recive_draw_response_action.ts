import type { Game } from '@/models/game'
import { EndGameReason } from '@/models/constants'

class DrawResponseParams {
    public drawResponse: boolean
}
class ResponseDrawAction {
    private game: Game

    constructor(game: Game) {
        this.game = game
    }

    public Invoke(body: string) {
        const p: DrawResponseParams = JSON.parse(body)

        if (p.drawResponse) {
            this.game.setEndGameReason(EndGameReason.DRAW)
            this.game.setEndGame(true)
            this.game.repository.closeWebSocketConnection()
        }
    }
}

export { ResponseDrawAction }
