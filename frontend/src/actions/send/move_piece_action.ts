import type { ConnectionRepository, Message } from '@/models/connection_repository/connection_repository'
import { pieceTypeToNumber, type File, type PieceType, type Rank } from '@/models/constants'
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
        promoteTo: number | null
    }

    constructor(src: Square, dst: Square, promoteTo: PieceType | null) {
        this.action = 'move-piece'
        this.body = {
            src: {
                file: src.file,
                rank: src.rank
            },
            dst: {
                file: dst.file,
                rank: dst.rank
            },
            promoteTo: promoteTo ? pieceTypeToNumber(promoteTo) : null
        }
    }
}

function MovePieceAction(repository: ConnectionRepository, game: Game, dst: Square, promoteTo: PieceType | null) {
    const src = game.srcSquare?.square as Square
    const m = new MovePieceMessage(src, dst, promoteTo)
    console.log('==> Move piece: ', m)
    repository.sendWebSocketMessage(m)
}

export { MovePieceAction, MovePieceMessage }
