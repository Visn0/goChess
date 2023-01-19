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
            {{ room.id }}
        </div>
        <ul class="list-group list-group-flush">
            <template v-if="room.players.length > 0">
                <li v-for="p in room.players" :key="p.id" class="list-group-item">
                    {{ p.id }}
                </li>
                <li v-show="room.players.length < 2" class="list-group-item p-0 bg-dark">
                    <button
                        type="button"
                        :disabled="room.isRoomFull()"
                        class="btn btn-green btn-sm w-100 h-100 m-0 text-left"
                        @click="joinRoom"
                    >
                        Join
                    </button>
                </li>
            </template>
            <template v-else>No players yet</template>
        </ul>
    </div>
</template>
