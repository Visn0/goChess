<script setup lang="ts">
import ChessBoard from '@/components/ChessBoard.vue'
import ChessTimer from '@/components/ChessTimer.vue'
import PlayerNickname from '@/components/PlayerNickname.vue'
import ChessPiece from '@/components/ChessPiece.vue'
import type { Board } from '@/models/board'
import { Color, constants, PieceType, type File, type Rank } from '@/models/constants'
import type { Game } from '@/models/game.js'
import router from '@/router'
import { useGameStore } from '@/stores/game'
import { usePlayerIDStore } from '@/stores/playerID'
import { onBeforeMount, ref, watch } from 'vue'
import { Piece } from '@/models/piece'
import { AbandonAction } from '@/actions/send/abandon_action'

const myPlayerIDStore = usePlayerIDStore()
const gameStore = useGameStore()
const myPlayerID = myPlayerIDStore.id
const oponentPlayerID = 'OponentPlayerID'
let board: Board
let game: Game
let colorDown = ref(Color.WHITE)
onBeforeMount(() => {
    if (gameStore.isEmpty) {
        router.push({ name: 'rooms' })
        return
    }
    game = gameStore.game as Game
    board = game.board
    board.initFromFenNotation(constants.StartingPosition)

    watch(game.getMyColor.bind(game), () => {
        colorDown.value = game.getMyColor()
    })
})

function squareClick(file: File, rank: Rank) {
    const square = board.getSquare(file, rank)
    game.selectSquare(square)
}

function promotePiece(pieceType: PieceType) {
    game.promotePiece(pieceType)
}

function cancelPromotion() {
    game.cancelPromotion()

    const modal = document.getElementById('promotion-modal') as HTMLElement
    modal.hidden = true
}

function abandon() {
    AbandonAction(game.repository)
}
</script>

<template>
    <main v-if="!gameStore.isEmpty">
        <div
            id="promotion-modal"
            class="modal"
            tabindex="-1"
            style="display: block"
            :hidden="!game.isPromotionPending()"
        >
            <div class="modal-dialog modal-dialog-centered">
                <div class="modal-content bg-dark text-light">
                    <div class="modal-header">
                        <h5 class="modal-title">Promote pawn to:</h5>
                        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                    </div>
                    <div class="modal-body d-flex justify-content-between">
                        <button
                            type="button"
                            @click="promotePiece(PieceType.ROOK)"
                            class="rounded promote-piece white-piece"
                        >
                            <ChessPiece :piece="new Piece(colorDown, PieceType.ROOK)" :selected="false" />
                        </button>
                        <button
                            type="button"
                            @click="promotePiece(PieceType.KNIGHT)"
                            class="rounded promote-piece black-piece"
                        >
                            <ChessPiece :piece="new Piece(colorDown, PieceType.KNIGHT)" :selected="false" />
                        </button>
                        <button
                            type="button"
                            @click="promotePiece(PieceType.BISHOP)"
                            class="rounded promote-piece white-piece"
                        >
                            <ChessPiece :piece="new Piece(colorDown, PieceType.BISHOP)" :selected="false" />
                        </button>
                        <button
                            type="button"
                            @click="promotePiece(PieceType.QUEEN)"
                            class="rounded promote-piece black-piece"
                        >
                            <ChessPiece :piece="new Piece(colorDown, PieceType.QUEEN)" :selected="false" />
                        </button>
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-secondary" @click="cancelPromotion()">Cancel</button>
                    </div>
                </div>
            </div>
        </div>

        <div class="vh-100 position-relative">
            <div class="container h-auto position-absolute top-50 start-50 translate-middle game">
                <div class="py-1 d-flex justify-content-between">
                    <PlayerNickname :nickname="oponentPlayerID" />
                    <ChessTimer :timer="game.getOpponentTimer()" />
                </div>
                <div class="d-flex justify-content-center">
                    <ChessBoard :board="board" :color-down="colorDown" @on-square-click="squareClick" />
                </div>
                <div class="py-1 d-flex justify-content-between">
                    <PlayerNickname :nickname="myPlayerID" />
                    <ChessTimer :timer="game.getOwnTimer()" />
                </div>

                <!-- Buttons -->
                <div class="mt-1">
                    <button type="button" class="col-4 btn btn-dark btn-sm border border-light m-2" @click="abandon()">
                        Abandon
                    </button>
                    <button type="button" class="col-4 btn btn-dark btn-sm border border-light m-2">Draw</button>
                </div>
            </div>
        </div>
    </main>
</template>

<style scoped>
.modal {
    background-color: rgba(0, 0, 0, 0.7);
}

.promote-piece {
    width: 75px;
    height: 75px;
}

.game {
    width: 72vmin !important;
    max-width: 72vmin !important;
}

@media (max-width: 281px) {
    .game {
        width: 100vmin !important;
        max-width: 100vmin !important;
    }

    .btn {
        font-size: 0.45em;
    }
}

@media (min-width: 281px) and (max-width: 576px) {
    .game {
        width: 100vmin !important;
        max-width: 100vmin !important;
    }

    .btn {
        font-size: 0.6em;
    }
}

@media (min-width: 576px) and (max-width: 768px) {
    .game {
        width: 92vmin !important;
        max-width: 92vmin !important;
    }

    .btn {
        font-size: 0.8em;
    }
}

@media (min-width: 768px) and (max-width: 992px) {
    .game {
        width: 84vmin !important;
        max-width: 84vmin !important;
    }

    .btn {
        font-size: 0.9em;
    }
}
</style>
