<script setup lang="ts">
import type { Piece } from '..//models/piece';

const props = defineProps<{
    piece: Piece
}>()

</script>

<template>
    <div v-if="piece" class="piece" :class="`${props.piece.color}-${props.piece.type}`"></div>
</template>

<style scoped>
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
    background-image: url('../assets/pieces/white/white-pawn.png');
}

.piece.white-rook {
    background-image: url('../assets/pieces/white/white-rook.png');
}

.piece.white-knight {
    background-image: url('../assets/pieces/white/white-knight.png');
}

.piece.white-bishop {
    background-image: url('../assets/pieces/white/white-bishop.png');
}

.piece.white-king {
    background-image: url('../assets/pieces/white/white-king.png');
}

.piece.white-queen {
    background-image: url('../assets/pieces/white/white-queen.png');
}

/*black pieces*/
.piece.black-pawn {
    background-image: url('../assets/pieces/black/black-pawn.png');
}

.piece.black-rook {
    background-image: url('../assets/pieces/black/black-rook.png');
}

.piece.black-knight {
    background-image: url('../assets/pieces/black/black-knight.png');
}

.piece.black-bishop {
    background-image: url('../assets/pieces/black/black-bishop.png');
}

.piece.black-king {
    background-image: url('../assets/pieces/black/black-king.png');
}

.piece.black-queen {
    background-image: url('../assets/pieces/black/black-queen.png');
}
</style>
