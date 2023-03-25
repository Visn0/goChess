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
import { RequestDrawAction } from '@/actions/send/request_draw_action'
import { ResponseDrawAction } from '@/actions/send/response_draw_action'
import { EndGameReason } from '@/models/constants'
import ChessModal from '@/components/ChessModal.vue'

const myPlayerIDStore = usePlayerIDStore()
const gameStore = useGameStore()
const myPlayerID = myPlayerIDStore.id
let board: Board
let game: Game
let colorDown = ref(Color.WHITE)
let oponentPlayerID = ref('Player 2')
let gameModalsData = ref({
    showModal: true,
    title: 'Waiting for player two',
    buttons: new Array<{
        text: string
        action: Function
    }>()
})

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
        gameModalsData.value.showModal = false
    })

    watch(game.endGame.bind(game), () => {
        let buttons = new Array()
        switch (game.getEndGameReason()) {
            case EndGameReason.ABANDON: {
                buttons = [
                    { text: ' Close', action: () => setGameModalsData('', false, []) },
                    { text: 'Go rooms', action: goRooms }
                ]
                setGameModalsData('Enemy player abandoned the game', true, buttons)
                break
            }

            case EndGameReason.DRAW_REQUEST: {
                buttons = [
                    { text: 'Accept', action: acceptDraw },
                    { text: 'Decline', action: declineDraw }
                ]
                setGameModalsData('Do you wanna draw?', true, buttons)
                game.setEndGame(false)
                break
            }

            case EndGameReason.DRAW: {
                buttons = [
                    { text: 'Close', action: () => setGameModalsData('', false, []) },
                    { text: 'Go rooms', action: goRooms }
                ]
                setGameModalsData('The game ended in a draw', true, buttons)
                break
            }

            case EndGameReason.DRAWDECLINED: {
                buttons = [{ text: 'Close', action: () => setGameModalsData('', false, []) }]
                setGameModalsData('Oponent player decline draw', true, buttons)
                game.setEndGame(false)
                break
            }

            case EndGameReason.CHECKMATE: {
                const winner = game.isMyTurn() ? 'win' : 'lose'
                let title = 'You ' + winner + ' the game'
                buttons = [
                    { text: 'Close', action: () => setGameModalsData('', false, []) },
                    { text: 'Go rooms', action: goRooms }
                ]
                setGameModalsData(title, true, buttons)
                break
            }

            default: {
                game.setEndGame(false)
                console.log('Error in endgame switch')
            }
        }
    })
})

function setGameModalsData(title: string, showModal: boolean, buttons: Array<{ text: string; action: Function }>) {
    gameModalsData.value.title = title
    gameModalsData.value.buttons = buttons
    gameModalsData.value.showModal = showModal
}

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

function requestDraw() {
    RequestDrawAction(game.repository)
}

function acceptDraw() {
    setGameModalsData('', false, [])
    game.setEndGameReason('draw')
    ResponseDrawAction(game.repository, true)
    game.setEndGame(true)
}

function declineDraw() {
    setGameModalsData('', false, [])
    ResponseDrawAction(game.repository, false)
    game.setEndGame(false)
}

function goRooms() {
    router.push({ name: 'rooms' })
}

function createPiece(color: Color, pieceType: PieceType): Piece {
    return new Piece(color, pieceType)
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
                                :piece="createPiece(colorDown, PieceType.ROOK)"
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
                                :piece="createPiece(colorDown, PieceType.KNIGHT)"
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
                                :piece="createPiece(colorDown, PieceType.BISHOP)"
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
                                :piece="createPiece(colorDown, PieceType.QUEEN)"
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

        <!--In game Modals-->

        <ChessModal
            id="ChessActionsModal"
            :showModal="gameModalsData.showModal"
            :title="gameModalsData.title"
            :buttons="gameModalsData.buttons"
            @close="setGameModalsData('', false, [])"
        ></ChessModal>

        <!--Board view-->

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
                    <button
                        type="button"
                        class="col-4 btn btn-dark btn-sm border border-light m-2"
                        @click="requestDraw()"
                    >
                        Draw
                    </button>
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
