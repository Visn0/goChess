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
        <ul class="list-group room-listing overflow-auto bg-dark rounded border-dark">
            <li class="list-group-item bg-transparent" v-for="r in props.rooms.getRooms()" :key="r.id">
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
    margin-top: 20%;
    max-height: 80vmin !important;
}
</style>
