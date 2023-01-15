<script setup lang="ts">
import { Color, Rank, File } from '@/models/constants';
import type { Board } from '../models/board'
import ChessBoardSquare from './ChessBoardSquare.vue';

defineProps<{
    board: Board
    colorDown: Color
}>()

const emit = defineEmits(['onSquareClick'])
function onSquareClickEvent(file: File, rank: Rank) {
    emit('onSquareClick', file, rank)
}

const ascRanks = enumToArray(Rank)
const descRanks = ascRanks.reverse()

const ascFiles = enumToArray(File)
const descFiles = ascFiles.reverse()

function enumToArray(e: any): Array<any> {
    const isNumber = (value: any) => isNaN(Number(value)) === false
    const values: Array<any> = Object.keys(e)
        .filter(isNumber)
        .map(key => Number(key))

    return values
}

</script>

<template>
    <div class="chess-board">
        <table id="board-table">
            <template v-if="colorDown === Color.WHITE">
                <tr v-for="r in descRanks">
                    <ChessBoardSquare v-for="f in ascFiles" :square="board.getSquare(f, r)"
                        @onSquareClick="onSquareClickEvent" />
                </tr>
            </template>
            <template v-else>
                <tr v-for="r in ascRanks">
                    <ChessBoardSquare v-for="f in descFiles" :square="board.getSquare(f, r)"
                        @onSquareClick="onSquareClickEvent" />
                </tr>
            </template>
        </table>
    </div>
</template>

<style scoped>
.board {
    margin: auto;
    border: 2px solid #000;
    border-radius: 5px;
    width: 80vmin;
    height: 80vmin;
    display: flex;
    flex-wrap: wrap;
}
</style>
