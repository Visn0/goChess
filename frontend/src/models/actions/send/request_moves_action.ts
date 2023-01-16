import type { ConnectionRepository, Message } from '../../connection_repository/connection_repository'
import type { PieceType, File, Rank } from '../../constants'
import type { Square } from '../../square'

class RequestMovesMessage implements Message {
    action: string
    body: {
        file: File
        rank: Rank
        piece: PieceType
    }

    constructor(file: File, rank: Rank, piece: PieceType) {
        this.action = 'request-moves'
        this.body = {
            file: file,
            rank: rank,
            piece: piece
        }
    }
}

function RequestMovesAction(repository: ConnectionRepository, square: Square) {
    const m = new RequestMovesMessage(square.file, square.rank, square.piece?.type as PieceType)
    console.log('request moves: ', m)
    repository.sendWebSocketMessage(m)
}

export { RequestMovesAction, RequestMovesMessage }
