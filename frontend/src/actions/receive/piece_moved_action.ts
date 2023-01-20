import type { File, PieceType, Rank } from '@/models/constants'
import type { Game } from '@/models/game'
import { Piece } from '@/models/piece'
import type { Square } from '@/models/square'

class PieceMovedParams {
    public src: CoordinateParams
    public dst: CoordinateParams
    public promoteTo: PieceType
}

class CoordinateParams {
    public file: File
    public rank: Rank
}

class PieceMovedAction {
    private game: Game

    constructor(game: Game) {
        this.game = game
    }

    public Invoke(body: string) {
        const p: PieceMovedParams = JSON.parse(body)
        console.log('==> Piece moved: ', p)

        const src: Square = this.game.board.getSquare(p.src.file, p.src.rank)
        const srcPiece = src.piece as Piece

        const dstPieceType = p.promoteTo ? new Piece(srcPiece.color, p.promoteTo) : srcPiece

        this.game.board.setSquarePiece(p.dst.file, p.dst.rank, dstPieceType)
        this.game.board.setSquarePiece(src.file, src.rank, null)
        this.game.unselectSrcSquare()
    }
}

export { PieceMovedAction }
