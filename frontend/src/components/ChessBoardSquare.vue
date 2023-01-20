<script setup lang="ts">
import type { Square } from '@/models/square'
import { fileToString, rankToString } from '@/models/constants'
import ChessPiece from './ChessPiece.vue'

const props = defineProps<{
    square: Square
    displayRank: boolean
    displayFile: boolean
}>()
const file = fileToString(props.square.file)
const rank = rankToString(props.square.rank)

const emit = defineEmits(['onSquareClick'])
function squareClick() {
    const file = props.square.file
    const rank = props.square.rank

    emit('onSquareClick', file, rank)
}
</script>

<template>
    <td id="{{ square.id }}" class="board-square p-0" :class="square.color" @click="squareClick()">
        <span v-show="displayRank" class="position-absolute top-0 start-0 rank">{{ rank }}</span>
        <span v-show="displayFile" class="position-absolute bottom-0 end-0 file">{{ file }}</span>
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
    font-weight: 600;
}

.rank {
    margin-left: 3%;
    margin-top: 3%;
}

.file {
    margin-bottom: 3%;
    margin-right: 3%;
}

.rank,
.file {
    display: inline-block;
    font-size: 16px;
    line-height: 16px;
}

.valid-move {
    width: 100%;
    height: 100%;
    shape-outside: circle(100%);
    clip-path: circle(40%);
    background-color: rgba(45, 45, 45, 0.15);
    border: 2px solid #000;
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

    .rank,
    .file {
        font-size: 12px;
        line-height: 12px;
    }
}

@media (min-width: 576px) and (max-width: 768px) {
    .board-square {
        width: 11.5vmin;
        height: 11.5vmin;
    }

    .rank,
    .file {
        font-size: 14px;
        line-height: 14px;
    }
}

@media (min-width: 768px) and (max-width: 992px) {
    .board-square {
        width: 10.5vmin;
        height: 10.5vmin;
    }

    .rank,
    .file {
        font-size: 16px;
        line-height: 16px;
    }
}
</style>
