<script setup lang="ts">
import type { Rooms } from '@/models/room'
import RoomItem from './RoomItem.vue'

const props = defineProps<{
    rooms: Rooms
}>()

const emit = defineEmits(['joinRoom'])
function joinRoom(roomID: string) {
    emit('joinRoom', roomID)
}
</script>

<template>
    <template v-if="props.rooms.getRooms().length > 0">
        <ul class="list-group overflow-auto room-listing bg-dark rounded border-dark">
            <li class="list-group-item bg-transparent mx-0 px-0 pe-1" v-for="r in props.rooms.getRooms()" :key="r.id">
                <RoomItem :room="r" @join-room="joinRoom" />
            </li>
        </ul>
    </template>
    <template v-else>
        <div class="card bg-dark text-white border border-secondary">
            <ul class="list-group bg-transparent">
                <li class="list-group-item">There are no rooms yet.</li>
            </ul>
        </div>
    </template>
</template>

<style scoped>
.room-listing {
    margin-top: 15%;
    max-height: 65vmin !important;
}

::-webkit-scrollbar {
    width: 16px;
}

::-webkit-scrollbar-track {
    -webkit-box-shadow: inset 0 0 6px #1da338;
    border-radius: 10px;
}

::-webkit-scrollbar-thumb {
    border-radius: 10px;
    background-color: rgba(67, 240, 102, 0.2);
    -webkit-box-shadow: inset 0 0 6px rgba(8, 87, 24, 0.6);
}

::-webkit-scrollbar-thumb:hover {
    background-color: rgba(67, 240, 102, 0.3);
}
</style>
