enum File {
    A = 0,
    B = 1,
    C = 2,
    D = 3,
    E = 4,
    F = 5,
    G = 6,
    H = 7
}

function fileToString(e: File): string {
    return File[e]
}

enum Rank {
    _1 = 0,
    _2 = 1,
    _3 = 2,
    _4 = 3,
    _5 = 4,
    _6 = 5,
    _7 = 6,
    _8 = 7
}

function rankToString(e: Rank): string {
    return Rank[e].charAt(1)
}

enum Color {
    BLACK = 'black',
    WHITE = 'white'
}

enum PieceType {
    PAWN = 'pawn',
    ROOK = 'rook',
    KNIGHT = 'knight',
    BISHOP = 'bishop',
    QUEEN = 'queen',
    KING = 'king',
    EMPTY = 'empty'
}

function pieceTypeToNumber(pieceType: PieceType): number {
    return Object.keys(PieceType)
        .map((key) => PieceType[key as keyof typeof PieceType])
        .indexOf(pieceType)
}

function numberToPieceType(n: number): PieceType {
    return PieceType[Object.keys(PieceType)[n] as keyof typeof PieceType]
}

const constants = {
    StartingPosition: 'rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1'
}

enum EndGameReason {
    ABANDON = 'abandon',
    DRAW_REQUEST = 'draw-request',
    DRAW = 'draw',
    CHECKMATE = 'checkmate'
}

export { File, fileToString, rankToString, Rank, Color, PieceType, pieceTypeToNumber, numberToPieceType, EndGameReason, constants }
