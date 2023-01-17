import type { ConnectionRepository, Message } from '@/models/connection_repository/connection_repository'
import type { File, Rank } from '@/models/constants'
import type { Game } from '@/models/game'
import type { Square } from '@/models/square'

class MovePieceMessage implements Message {
    action: string
    body: {
        src: {
            file: File
            rank: Rank
        }
        dst: {
            file: File
            rank: Rank
        }
    }

    constructor(src: Square, dst: Square) {
        this.action = 'move-piece'
        this.body = {
            src: {
                file: src.file,
                rank: src.rank
            },
            dst: {
                file: dst.file,
                rank: dst.rank
            }
        }
    }
}

function MovePieceAction(repository: ConnectionRepository, game: Game, dst: Square) {
    const srcSquare = game.srcSquare
    if (!srcSquare?.canInnerPieceMoveTo(dst)) {
        return
    }

    const m = new MovePieceMessage(srcSquare.square, dst)
    repository.sendWebSocketMessage(m)
}

export { MovePieceAction, MovePieceMessage }
