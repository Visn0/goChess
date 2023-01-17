<script setup lang="ts">
import { MovesReceivedAction } from '@/models/actions/receive/moves_received_action'
import { PieceMovedAction } from '@/models/actions/receive/piece_moved_action'
import type { ReceiveAction } from '@/models/actions/receive/receive_action'
import { RoomCreatedAction } from '@/models/actions/receive/room_created_action'
import { RoomJoinedAction } from '@/models/actions/receive/room_joined_action'
import { RouteActions } from '@/models/actions/receive/route_actions'
import { CreateRoomAction } from '@/models/actions/send/create_room_action'
import { RequestRoomsAction } from '@/models/actions/send/request_rooms'
import { Board } from '@/models/board'
import { BackendConnectionRepository } from '@/models/connection_repository/backend_connection_repository'
import { constants } from '@/models/constants'
import { Game } from '@/models/game.js'
import { Rooms } from '@/models/room'
import { usePlayerIDStore } from '@/stores/playerID'
import RoomListing from '@/components/RoomListing.vue'
import RoomItem from '@/components/RoomItem.vue'
import { watch } from 'vue'
import { JoinRoomAction } from '@/models/actions/send/join_room_action'
import router from '@/router'
import { useGameStore } from '@/stores/game'

const playerIDStore = usePlayerIDStore()

const playerID = playerIDStore.id
const rooms = new Rooms()
const board = new Board()
board.initFromFenNotation(constants.StartingPosition)

/* eslint-disable capitalized-comments */
// const repository: ConnectionRepository = new MockConnectionRepository()
/* eslint-enable capitalized-comments */
const repository = new BackendConnectionRepository('localhost', '8081', 'ws')
const game = new Game(rooms, board, repository)

const routeActions: RouteActions = new RouteActions(
    new Map<string, ReceiveAction>([
        ['create-room', new RoomCreatedAction(rooms)],
        ['join-room', new RoomJoinedAction(rooms)],
        ['request-moves', new MovesReceivedAction(game)],
        ['move-piece', new PieceMovedAction(game)]
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
    CreateRoomAction(repository, playerID, `room-${Date.now().toString()}`, 'roomPassword')
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
        <div class="container h-auto position-absolute top-30 start-50 translate-middle d-flex justify-content-center">
            <div class="row w-100 d-flex justify-content-center">
                <div class="col-sm-12 col-lg-6">
                    <div class="d-flex justify-content-center my-2 text-light">{{ playerID }}</div>
                    <template v-if="rooms.myRoom">
                        <RoomItem :room="rooms.myRoom" />
                        <div class="h-0 p-0 rounded border border-light my-2 mb-4"></div>
                    </template>
                    <RoomListing :rooms="rooms" @join-room="joinRoom" :key="componentKey" />
                    <button type="button" class="w-100 btn bg-green" @click="createRoom">Create room</button>
                </div>
            </div>
        </div>
    </main>
</template>

<style scoped>
.top-30 {
    top: 30% !important;
}
</style>
