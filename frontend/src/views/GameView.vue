<script setup lang="ts">
import ChessBoard from '@/components/ChessBoard.vue'
import type { Board } from '@/models/board'
import { Color, constants, type File, type Rank } from '@/models/constants'
import type { Game } from '@/models/game.js'
import router from '@/router'
import { useGameStore } from '@/stores/game'
import { usePlayerIDStore } from '@/stores/playerID'
import { onBeforeMount } from 'vue'

const playerIDStore = usePlayerIDStore()
const gameStore = useGameStore()

const playerID = playerIDStore.id
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
            <div class="container h-auto position-absolute top-50 start-50 translate-middle">
                <div class="my-2 text-light">{{ playerID }}</div>
                <ChessBoard :board="board" :color-down="Color.WHITE" @on-square-click="squareClick" />

                <!-- Buttons -->
                <div class="row mt-3 text-center actions-btns">
                    <button type="button" class="col-sm btn btn-dark border border-light m-2">Abandon</button>
                    <button type="button" class="col-sm btn btn-dark border border-light m-2">Draw</button>
                </div>
            </div>
        </div>
    </main>
</template>
