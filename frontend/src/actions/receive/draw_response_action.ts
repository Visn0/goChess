import type { Game } from '@/models/game'

class DrawResponseParams {
    public response: boolean
}
class DrawResponseAction {
    private game: Game

    constructor(game: Game) {
        this.game = game
    }

    public Invoke(body: string) {
        const p: DrawResponseParams = JSON.parse(body)

        if(p.response)
        {
            this.game.setEndGameReason('draw')
            this.game.setEndGame(true)
            this.game.repository.closeWebSocketConnection()
        }
    }
}

export { DrawResponseAction }