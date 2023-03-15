import type { Game } from '@/models/game'
import { EndGameReason } from '@/models/constants'

class DrawResponseParams {
    public drawResponse: boolean
}
class ReceiveDrawResponseAction {
    private game: Game

    constructor(game: Game) {
        this.game = game
    }

    public Invoke(body: string) {
        const p: DrawResponseParams = JSON.parse(body)
        console.log(p)

        if (p.drawResponse) {
            this.game.setEndGameReason(EndGameReason.DRAW)
            this.game.setEndGame(true)
            this.game.repository.closeWebSocketConnection()
        } else {
            this.game.setEndGameReason(EndGameReason.DRAWDECLINED)
            this.game.setEndGame(true)
        }
    }
}

export { ReceiveDrawResponseAction }
