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
let board: Board
let game: Game
let colorDown = ref(Color.WHITE)
let oponentPlayerID = ref('Player 2')
onBeforeMount(() => {
    if (gameStore.isEmpty) {
        router.push({ name: 'rooms' })
        return
    }
    game = gameStore.game as Game
    board = game.board
    board.initFromFenNotation(constants.StartingPosition)

    watch(game.started.bind(game), () => {
        colorDown.value = game.getMyColor()
        oponentPlayerID.value = game.getOpponentID()

        const modal = document.getElementById('wait-player-modal') as HTMLElement
        modal.hidden = true
    })

    watch(game.endGame.bind(game), () => {
        switch(game.getEndGameReason()) {
            case 'abandon': {
                const modal = document.getElementById('endgame-modal') as HTMLElement
                const text = document.getElementById('endgame-text') as HTMLElement
                text.innerText = "Enemy player abandoned the game"
                modal.hidden = false
                break
            }
            case 'draw': {
                const modal = document.getElementById('endgame-modal') as HTMLElement
                const text = document.getElementById('endgame-text') as HTMLElement
                text.innerText = "The game ended in a draw"
                modal.hidden = false
                break
            }
        }
        
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
    router.push({ name: 'rooms' })
}

function draw() {

}

function goRooms() {
    router.push({ name: 'rooms' })
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
                            <ChessPiece
                                :piece="new Piece(colorDown, PieceType.ROOK)"
                                :selected="false"
                                :king-check="false"
                            />
                        </button>
                        <button
                            type="button"
                            @click="promotePiece(PieceType.KNIGHT)"
                            class="rounded promote-piece black-piece"
                        >
                            <ChessPiece
                                :piece="new Piece(colorDown, PieceType.KNIGHT)"
                                :selected="false"
                                :king-check="false"
                            />
                        </button>
                        <button
                            type="button"
                            @click="promotePiece(PieceType.BISHOP)"
                            class="rounded promote-piece white-piece"
                        >
                            <ChessPiece
                                :piece="new Piece(colorDown, PieceType.BISHOP)"
                                :selected="false"
                                :king-check="false"
                            />
                        </button>
                        <button
                            type="button"
                            @click="promotePiece(PieceType.QUEEN)"
                            class="rounded promote-piece black-piece"
                        >
                            <ChessPiece
                                :piece="new Piece(colorDown, PieceType.QUEEN)"
                                :selected="false"
                                :king-check="false"
                            />
                        </button>
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-secondary" @click="cancelPromotion()">Cancel</button>
                    </div>
                </div>
            </div>
        </div>

        <div id="wait-player-modal" class="modal" tabindex="-1" style="display: block">
            <div class="modal-dialog modal-dialog-centered">
                <div class="modal-content bg-dark text-light">
                    <h5 class="modal-title">Waiting for player</h5>
                </div>
            </div>
        </div>

        <!--End game modal-->
        <div id="endgame-modal" class="modal" tabindex="-1" style="display: block" hidden="true">
            <div class="modal-dialog modal-dialog-centered">
                <div class="modal-content bg-dark text-light">
                    <div class="modal-body">
                        <h5 id="endgame-text" class="modal-title"></h5>
                        <button type="button" class="mt-2 w-100 btn btn-sm btn-green" @click="goRooms()">
                            Go rooms
                        </button>
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
                    <button type="button" class="col-4 btn btn-dark btn-sm border border-light m-2" @click="abandon()">Abandon</button>
                    <button type="button" class="col-4 btn btn-dark btn-sm border border-light m-2" @click="draw()">Draw</button>
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
