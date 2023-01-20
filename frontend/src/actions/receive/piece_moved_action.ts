import { PieceType, type File, type Rank } from '@/models/constants'
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
        const dst = this.game.board.getSquare(p.dst.file, p.dst.rank)

        if (this.isPawnPassant(src, dst)) {
            this.movePassant(src, dst)
            return
        }

        const dstPieceType = p.promoteTo ? new Piece(srcPiece.color, p.promoteTo) : srcPiece

        this.game.board.setSquarePiece(p.dst.file, p.dst.rank, dstPieceType)
        this.game.board.setSquarePiece(src.file, src.rank, null)
        this.game.unselectSrcSquare()
    }

    private isPawnPassant(src: Square, dst: Square): boolean {
        if (src.piece?.type !== PieceType.PAWN) {
            return false
        }

        // Passant must be diagonal (different file)
        if (src.file === dst.file) {
            return false
        }

        return dst.isEmpty()
    }

    private movePassant(src: Square, dst: Square) {
        // Remove captured pawn
        this.game.board.setSquarePiece(dst.file, src.rank, null)

        // Move killer pawn
        this.game.board.setSquarePiece(dst.file, dst.rank, src.piece)
        this.game.board.setSquarePiece(src.file, src.rank, null)

        this.game.unselectSrcSquare()
    }
}

export { PieceMovedAction }
