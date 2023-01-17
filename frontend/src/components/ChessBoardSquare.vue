<script setup lang="ts">
import type { Square } from '@/models/square'
import ChessPiece from './ChessPiece.vue'
import { watch } from 'vue'

const props = defineProps<{
    square: Square
}>()

const emit = defineEmits(['onSquareClick'])
function squareClick() {
    const file = props.square.file
    const rank = props.square.rank

    emit('onSquareClick', file, rank)
}

let componentKey = 0
watch(props.square.isValidMove.bind(props.square), () => {
    componentKey += 1
})

watch(props.square.getPiece.bind(props.square), () => {
    componentKey += 1
})
</script>

<template>
    <td
        id="{{ props.square.id }}"
        class="board-square"
        :class="props.square.color"
        @click="squareClick()"
        :key="componentKey"
    >
        <div id="{{ props.square.id }}-piece" class="piece">
            <ChessPiece v-if="props.square.piece" :piece="props.square.piece" />
        </div>
        <div
            v-if="props.square.isValidMove()"
            id="{{ props.square.id }}-valid-move"
            class="board-square valid-move"
        ></div>
    </td>
</template>

<style scoped>
.board-square {
    width: 10vmin;
    height: 10vmin;
    position: relative;
    align-items: center;
    justify-content: center;
}

.board-square.valid-move {
    position: absolute;
    top: 0;
    left: auto;

    shape-outside: circle(100%);
    clip-path: circle(40%);
    background-color: rgba(45, 45, 45, 0.15);
    border: 2px solid #000;
}

.board-square.white-square {
    background-color: #fff;
}

.board-square.black-square {
    background-color: #7fa650;
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
</style>
