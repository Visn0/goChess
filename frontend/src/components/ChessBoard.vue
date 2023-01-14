<script setup lang="ts">
import type { Board } from '../models/board'

defineProps<{
    board: Board
}>()

</script>

<template>
    <div class="chess-board">
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

/*white pieces*/
.piece.white-pawn {
    background-image: url('../pieces/white/white-pawn.png');
}

.piece.white-rook {
    background-image: url('../pieces/white/white-rook.png');
}

.piece.white-knight {
    background-image: url('../pieces/white/white-knight.png');
}

.piece.white-bishop {
    background-image: url('../pieces/white/white-bishop.png');
}

.piece.white-king {
    background-image: url('../pieces/white/white-king.png');
}

.piece.white-queen {
    background-image: url('../pieces/white/white-queen.png');
}

/*black pieces*/
.piece.black-pawn {
    background-image: url('../pieces/black/black-pawn.png');
}

.piece.black-rook {
    background-image: url('../pieces/black/black-rook.png');
}

.piece.black-knight {
    background-image: url('../pieces/black/black-knight.png');
}

.piece.black-bishop {
    background-image: url('../pieces/black/black-bishop.png');
}

.piece.black-king {
    background-image: url('../pieces/black/black-king.png');
}

.piece.black-queen {
    background-image: url('../pieces/black/black-queen.png');
}
</style>
