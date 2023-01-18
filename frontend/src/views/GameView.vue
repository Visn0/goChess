<script setup lang="ts">
import ChessBoard from '@/components/ChessBoard.vue'
import ChessTimer from '@/components/ChessTimer.vue'
import PlayerNickname from '@/components/PlayerNickname.vue'
import type { Board } from '@/models/board'
import { Color, constants, type File, type Rank } from '@/models/constants'
import type { Game } from '@/models/game.js'
import { Timer } from '@/models/timer'
import router from '@/router'
import { useGameStore } from '@/stores/game'
import { usePlayerIDStore } from '@/stores/playerID'
import { onBeforeMount, watch } from 'vue'

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

const oponentTimer = new Timer(600)
const myTimer = new Timer(600)
myTimer.start()

watch(myTimer.isStoped.bind(myTimer), () => {
    console.log(`Timer finished: ${myTimer.toString()}`)
})
</script>

<template>
    <main v-if="!gameStore.isEmpty">
        <div class="vh-100 position-relative">
            <div class="container h-auto position-absolute top-50 start-50 translate-middle game">
                <div class="py-1 d-flex justify-content-between">
                    <PlayerNickname :nickname="oponentPlayerID" />
                    <ChessTimer :timer="oponentTimer" />
                </div>
                <div class="d-flex justify-content-center">
                    <ChessBoard :board="board" :color-down="Color.WHITE" @on-square-click="squareClick" />
                </div>
                <div class="py-1 d-flex justify-content-between">
                    <PlayerNickname :nickname="myPlayerID" />
                    <ChessTimer :timer="myTimer" />
                </div>

                <!-- Buttons -->
                <div class="mt-1">
                    <button type="button" class="col-4 action btn btn-dark btn-sm border border-light m-2">
                        Abandon
                    </button>
                    <button type="button" class="col-4 action btn btn-dark btn-sm border border-light m-2">Draw</button>
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

.action {
    font-size: 1em;
}

@media (max-width: 281px) {
    .game {
        width: 100vmin !important;
        max-width: 100vmin !important;
    }

    .action {
        font-size: 0.45em;
    }
}

@media (min-width: 281px) and (max-width: 576px) {
    .game {
        width: 100vmin !important;
        max-width: 100vmin !important;
    }

    .action {
        font-size: 0.6em;
    }
}

@media (min-width: 576px) and (max-width: 768px) {
    .game {
        width: 92vmin !important;
        max-width: 92vmin !important;
    }

    .action {
        font-size: 0.8em;
    }
}

@media (min-width: 768px) and (max-width: 992px) {
    .game {
        width: 84vmin !important;
        max-width: 84vmin !important;
    }

    .action {
        font-size: 0.9em;
    }
}
</style>
