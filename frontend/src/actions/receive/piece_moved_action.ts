import { PieceType, File, Rank, numberToPieceType } from '@/models/constants'
import type { Game } from '@/models/game'
import { Piece } from '@/models/piece'
import type { Square } from '@/models/square'

class PieceMovedParams {
    public src: CoordinateParams
    public dst: CoordinateParams
    public promoteTo: number | null
    public kingCheck: CoordinateParams | null
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

        this.movePiece(p)

        if (p.kingCheck) {
            const kingCheckSquare = this.game.board.getSquare(p.kingCheck.file, p.kingCheck.rank)
            this.game.setKingCheck(kingCheckSquare)
        } else {
            this.game.setKingCheck(null)
        }

        this.game.changeTurn()
    }

    private movePiece(p: PieceMovedParams) {
        const src: Square = this.game.board.getSquare(p.src.file, p.src.rank)
        const dst = this.game.board.getSquare(p.dst.file, p.dst.rank)

        if (this.isCastle(src, dst)) {
            this.moveCastle(src, dst)
            return
        }

        if (this.isPawnPassant(src, dst)) {
            this.movePassant(src, dst)
            return
        }

        const srcPiece = src.piece as Piece
        const dstPieceType = p.promoteTo ? new Piece(srcPiece.color, numberToPieceType(p.promoteTo)) : srcPiece

        this.game.board.setSquarePiece(p.dst.file, p.dst.rank, dstPieceType)
        this.game.board.setSquarePiece(src.file, src.rank, null)
        this.game.unselectSrcSquare()
    }

    private isCastle(src: Square, dst: Square): boolean {
        if (src.piece?.type !== PieceType.KING) {
            return false
        }

        // Castle must be in the same rank
        if (src.rank !== dst.rank) {
            return false
        }

        // Castle is only allowed in Rank 1 (white) and 8 (black)
        if (src.rank > Rank._1 && src.rank < Rank._8) {
            return false
        }

        if (dst.rank > Rank._1 && dst.rank < Rank._8) {
            return false
        }

        const distance = src.file - dst.file
        return Math.abs(distance) >= 2
    }

    private moveCastle(src: Square, dst: Square) {
        // Move King to destination
        this.game.board.setSquarePiece(dst.file, dst.rank, src.piece)
        this.game.board.setSquarePiece(src.file, src.rank, null)

        // Move Rook next to the King
        if (dst.file === File.G) {
            const rookFileH = this.game.board.getSquare(File.H, dst.rank)
            this.game.board.setSquarePiece(File.F, dst.rank, rookFileH.piece)
            this.game.board.setSquarePiece(File.H, dst.rank, null)
        } else {
            // dst.File === File.C
            const rookFileA = this.game.board.getSquare(File.A, dst.rank)
            this.game.board.setSquarePiece(File.D, dst.rank, rookFileA.piece)
            this.game.board.setSquarePiece(File.A, dst.rank, null)
        }

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
