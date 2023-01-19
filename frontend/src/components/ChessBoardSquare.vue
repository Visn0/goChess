<script setup lang="ts">
import type { Square } from '@/models/square'
import ChessPiece from './ChessPiece.vue'

const props = defineProps<{
    square: Square
}>()

const emit = defineEmits(['onSquareClick'])
function squareClick() {
    const file = props.square.file
    const rank = props.square.rank

    emit('onSquareClick', file, rank)
}
</script>

<template>
    <td id="{{ square.id }}" class="board-square p-0" :class="square.color" @click="squareClick()">
        <div class="piece">
            <ChessPiece v-if="square.piece" :piece="square.piece" :selected="square.isSelected()" />
        </div>
        <div v-if="square.isValidMove()" class="position-absolute top-0 h-100 w-100">
            <div class="valid-move"></div>
        </div>
    </td>
</template>

<style scoped>
.board-square {
    width: 9vmin;
    height: 9vmin;
    position: relative;
    align-items: center;
    justify-content: center;
}

.valid-move {
    width: 100%;
    height: 100%;
    shape-outside: circle(100%);
    clip-path: circle(40%);
    background-color: rgba(45, 45, 45, 0.15);
    border: 2px solid #000;
}

.board-square.white-square {
    background-color: #fff;
}

.board-square.black-square {
    /* background-color: #7fa650; */
    background-color: #65a371;
}

.piece {
    margin: 0;
    padding: 0;
    align-items: center;
    justify-content: center;

    height: 100%;
    width: 100%;
    background-repeat: no-repeat;
    background-size: 100% 100%;
}

@media (max-width: 576px) {
    .board-square {
        width: 12.5vmin;
        height: 12.5vmin;
    }
}

@media (min-width: 576px) and (max-width: 768px) {
    .board-square {
        width: 11.5vmin;
        height: 11.5vmin;
    }
}

@media (min-width: 768px) and (max-width: 992px) {
    .board-square {
        width: 10.5vmin;
        height: 10.5vmin;
    }
}
</style>
