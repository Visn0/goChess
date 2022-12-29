import { ConnectionRepository, Message } from '../../connection_repository/connection_repository'
import { File, Rank } from '../../constants'
import { GameController } from '../../game_controller'
import { Square } from '../../square'

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

function MovePieceAction(repository: ConnectionRepository, gameController: GameController, dst: Square) {
    const srcSquare = gameController.srcSquare
    if (!srcSquare?.canInnerPieceMoveTo(dst)) {
        return
    }

    const m = new MovePieceMessage(srcSquare.square, dst)
    repository.sendMessage(m)
}

export { MovePieceAction, MovePieceMessage }
