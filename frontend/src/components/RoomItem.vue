<script setup lang="ts">
import type { Room } from '@/models/room'

const props = defineProps<{
    room: Room
}>()

const emit = defineEmits(['joinRoom'])
function joinRoom() {
    emit('joinRoom', props.room.id)
}
</script>

<template>
    <div class="card bg-dark text-white border border-secondary">
        <div class="card-header bg-dark">
            <div class="position-absolute" style="margin-top: -5px">
                <button
                    type="button"
                    :disabled="room.isRoomFull()"
                    class="btn bg-green btn-sm text-left"
                    @click="joinRoom"
                >
                    Join
                </button>
            </div>
            <div>{{ room.id }}</div>
        </div>
        <ul class="list-group list-group-flush">
            <template v-if="room.players.length > 0">
                <li v-for="p in room.players" :key="p.id" class="list-group-item">
                    {{ p.id }}
                </li>
            </template>
            <template v-else>No players yet</template>
        </ul>
    </div>
</template>
