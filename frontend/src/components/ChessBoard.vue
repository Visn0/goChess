<script setup lang="ts">
import { Color, Rank, File } from '@/models/constants'
import type { Board } from '../models/board'
import ChessBoardSquare from './ChessBoardSquare.vue'

defineProps<{
    board: Board
    colorDown: Color
}>()

const emit = defineEmits(['onSquareClick'])
function squareClick(file: File, rank: Rank) {
    emit('onSquareClick', file, rank)
}

const ascRanks = enumToArray(Rank)
const descRanks = ascRanks.slice().reverse()

const ascFiles = enumToArray(File)
const descFiles = ascFiles.slice().reverse()

function enumToArray(e: any): Array<any> {
    const isNumber = (value: any) => isNaN(Number(value)) === false
    const values: Array<any> = Object.keys(e)
        .filter(isNumber)
        .map((key) => Number(key))

    return values
}
</script>

<template>
    <table id="board-table">
        <template v-if="colorDown === Color.WHITE">
            <tr v-for="r in descRanks" :key="r">
                <ChessBoardSquare
                    v-for="f in ascFiles"
                    :square="board.getSquare(f, r)"
                    :display-file="r === Rank._1"
                    :display-rank="f === File.A"
                    :key="board.getSquare(f, r).id"
                    @on-square-click="squareClick"
                />
            </tr>
        </template>
        <template v-else>
            <tr v-for="r in ascRanks" :key="r">
                <ChessBoardSquare
                    v-for="f in descFiles"
                    :square="board.getSquare(f, r)"
                    :display-file="r === Rank._8"
                    :display-rank="f === File.H"
                    :key="board.getSquare(f, r).id"
                    @on-square-click="squareClick"
                />
            </tr>
        </template>
    </table>
</template>
