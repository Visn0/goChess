import { ConnectionRepository, Message } from '../../backend/connection_repository'
import { PieceType, File, Rank } from '../../constants'
import { Square } from '../../square'

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
    repository.sendMessage(m)
}

export { RequestMovesAction, RequestMovesMessage }
