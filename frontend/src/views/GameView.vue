<script setup lang="ts">
import ChessBoard from '@/components/ChessBoard.vue'
import PlayerNickname from '@/components/PlayerNickname.vue'
import type { Board } from '@/models/board'
import { Color, constants, type File, type Rank } from '@/models/constants'
import type { Game } from '@/models/game.js'
import router from '@/router'
import { useGameStore } from '@/stores/game'
import { usePlayerIDStore } from '@/stores/playerID'
import { onBeforeMount } from 'vue'

const myPlayerIDStore = usePlayerIDStore()
const gameStore = useGameStore()

const myPlayerID = myPlayerIDStore.id
const oponentPlayerID = 'OponentPlayerID'
let board: Board
let game: Game

onBeforeMount(() => {
    if (gameStore.isEmpty) {
        console.log('Game is empty. Redirecting to /rooms.')
        router.push({ name: 'rooms' })
        return
    }
    game = gameStore.game as Game
    board = game.board
    board.initFromFenNotation(constants.StartingPosition)
})

function squareClick(file: File, rank: Rank) {
    const square = board.getSquare(file, rank)
    game.selectSquare(square)
}
</script>

<template>
    <main v-if="!gameStore.isEmpty">
        <div class="vh-100 position-relative">
            <div class="container h-auto position-absolute top-50 start-50 translate-middle game">
                <div class="py-1 d-flex justify-content-between">
                    <PlayerNickname :nickname="oponentPlayerID" />
                </div>
                <div class="d-flex justify-content-center">
                    <ChessBoard :board="board" :color-down="Color.WHITE" @on-square-click="squareClick" />
                </div>
                <div class="py-1 d-flex justify-content-between">
                    <PlayerNickname :nickname="myPlayerID" />
                </div>

                <!-- Buttons -->
                <div class="mt-1">
                    <button type="button" class="col-sm btn btn-dark btn-sm border border-light m-2">Abandon</button>
                    <button type="button" class="col-sm btn btn-dark btn-sm border border-light m-2">Draw</button>
                </div>
            </div>
        </div>
    </main>
</template>

<style>
.game {
    width: 72vmin !important;
    max-width: 72vmin !important;
}

@media (max-width: 576px) {
    .game {
        width: 100vmin !important;
        max-width: 100vmin !important;
    }
}

@media (min-width: 576px) and (max-width: 768px) {
    .game {
        width: 92vmin !important;
        max-width: 92vmin !important;
    }
}

@media (min-width: 768px) and (max-width: 992px) {
    .game {
        width: 84vmin !important;
        max-width: 84vmin !important;
    }
}
</style>
