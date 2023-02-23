<script setup lang="ts">
import { MovesReceivedAction } from '@/actions/receive/moves_received_action'
import { PieceMovedAction } from '@/actions/receive/piece_moved_action'
import type { ReceiveAction } from '@/actions/receive/receive_action'
import { RoomCreatedAction } from '@/actions/receive/room_created_action'
import { RoomJoinedAction } from '@/actions/receive/room_joined_action'
import { RouteActions } from '@/actions/receive/route_actions'
import { CreateRoomAction } from '@/actions/send/create_room_action'
import { RequestRoomsAction } from '@/actions/send/request_rooms'
import { Board } from '@/models/board'
import { BackendConnectionRepository } from '@/models/connection_repository/backend_connection_repository'
import { constants } from '@/models/constants'
import { Game } from '@/models/game.js'
import { Rooms } from '@/models/room'
import { usePlayerIDStore } from '@/stores/playerID'
import RoomListing from '@/components/RoomListing.vue'
import RoomItem from '@/components/RoomItem.vue'
import { watch } from 'vue'
import { JoinRoomAction } from '@/actions/send/join_room_action'
import router from '@/router'
import { useGameStore } from '@/stores/game'
import { GotTimersAction } from '@/actions/receive/got_timers'
import { StartGameAction } from '@/actions/receive/start_game'
import { AbandonAction } from '@/actions/receive/abandon_action'
import { ReceiveDrawRequestAction } from '@/actions/receive/receive_draw_request_action'

const playerIDStore = usePlayerIDStore()

const playerID = playerIDStore.id
const rooms = new Rooms()
const board = new Board()
board.initFromFenNotation(constants.StartingPosition)

console.log('base url: ', import.meta.env.BASE_URL)
const url = import.meta.env.VITE_APP_API_HOST
const repository = new BackendConnectionRepository(url, 'ws')
const game = new Game(board, repository)

const routeActions: RouteActions = new RouteActions(
    new Map<string, ReceiveAction>([
        ['create-room', new RoomCreatedAction(rooms)],
        ['join-room', new RoomJoinedAction(rooms)],
        ['start-game', new StartGameAction(game)],
        ['request-moves', new MovesReceivedAction(game)],
        ['move-piece', new PieceMovedAction(game)],
        ['get-timers', new GotTimersAction(game)],
        ['abandon', new AbandonAction(game)],
        ['receive-draw-request', new ReceiveDrawRequestAction(game)]
    ])
)
repository.addOnWebSocketMessageEventListener(routeActions.route())

RequestRoomsAction(repository, rooms)
const requestRoomInterval = setInterval(() => RequestRoomsAction(repository, rooms), 10000)

function goToGame() {
    clearInterval(requestRoomInterval)

    const gameStore = useGameStore()
    gameStore.set(game)
    router.push({ name: 'game' })
}

function createRoom() {
    CreateRoomAction(repository, playerID, `${Date.now()}`, 'roomPassword')
    goToGame()
}

function joinRoom(roomID: string) {
    JoinRoomAction(repository, playerID, roomID, 'roomPassword')
    goToGame()
}

let componentKey = 1
watch(rooms.getRooms.bind(rooms), () => {
    componentKey += 1
})

watch(rooms.getMyRoom.bind(rooms), () => {
    componentKey += 1
})
</script>

<template>
    <main>
        <div class="container h-auto position-absolute top-40 start-50 translate-middle room-listing">
            <div class="d-flex justify-content-center w-100">
                <div>
                    <template v-if="rooms.myRoom">
                        <RoomItem :room="rooms.myRoom" />
                        <div class="h-0 p-0 rounded border border-light my-2 mb-4"></div>
                    </template>
                    <RoomListing :rooms="rooms" @join-room="joinRoom" :key="componentKey" />
                    <button type="button" class="mt-2 w-100 btn btn-sm btn-green" @click="createRoom">
                        Create room
                    </button>
                </div>
            </div>
        </div>
    </main>
</template>

<style>
.top-40 {
    top: 45% !important;
}

.room-listing {
    width: 46.8vmin !important;
    max-width: 46.8vmin !important;
    font-size: 0.95em;
}

@media (max-width: 576px) {
    .room-listing {
        width: 96vmin !important;
        max-width: 96vmin !important;
        font-size: 0.75em;
    }

    .btn {
        font-size: 0.75em;
    }
}

@media (min-width: 576px) and (max-width: 768px) {
    .room-listing {
        width: 76vmin !important;
        max-width: 76vmin !important;
        font-size: 0.85em;
    }

    .btn {
        font-size: 0.85em;
    }
}

@media (min-width: 768px) and (max-width: 992px) {
    .room-listing {
        width: 61.6vmin !important;
        max-width: 61.6vmin !important;
        font-size: 0.87em;
    }

    .btn {
        font-size: 0.87em;
    }
}

@media (min-width: 992px) and (max-width: 1200px) {
    .room-listing {
        width: 72.6vmin !important;
        max-width: 72.6vmin !important;
    }

    .btn {
        font-size: 0.95em;
    }
}
</style>
