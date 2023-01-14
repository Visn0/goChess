<script setup lang="ts">
import { MovesReceivedAction } from '@/models/actions/receive/moves_received_action';
import { PieceMovedAction } from '@/models/actions/receive/piece_moved_action';
import type { ReceiveAction } from '@/models/actions/receive/receive_action';
import { RoomCreatedAction } from '@/models/actions/receive/room_created_action';
import { RoomJoinedAction } from '@/models/actions/receive/room_joined_action';
import { RouteActions } from '@/models/actions/receive/route_actions';
import { CreateRoomAction } from '@/models/actions/send/create_room_action';
import { RequestRoomsAction } from '@/models/actions/send/request_rooms';
import { Board } from '@/models/board';
import { BackendConnectionRepository } from '@/models/connection_repository/backend_connection_repository';
import type { ConnectionRepository } from '@/models/connection_repository/connection_repository';
import { Color, constants } from '@/models/constants';
import { GameController } from '@/models/game_controller';
import { Rooms } from '@/models/room';
import { onMounted } from 'vue';

const playerID = 'MiPlayerID'
const rooms = new Rooms('modal-list-rooms-body')
let board: Board
/* eslint-disable capitalized-comments */
// let repository: ConnectionRepository
/* eslint-enable capitalized-comments */
let repository: ConnectionRepository
let gameController: GameController

onMounted(() => {
    board = new Board(document.getElementById('chess-board') as HTMLElement)
    /* eslint-disable capitalized-comments */
    // const repository: ConnectionRepository = new MockConnectionRepository()
    /* eslint-enable capitalized-comments */
    repository = new BackendConnectionRepository('localhost', '8081', 'ws')
    gameController = new GameController(rooms, board, repository)

    const routeActions: RouteActions = new RouteActions(
        new Map<string, ReceiveAction>([
            ['create-room', new RoomCreatedAction(rooms)],
            ['join-room', new RoomJoinedAction(rooms)],
            ['request-moves', new MovesReceivedAction(gameController)],
            ['move-piece', new PieceMovedAction(gameController)]
        ])
    )
    repository.addOnWebSocketMessageEventListener(routeActions.route())

    board.initFromFenNotation(constants.StartingPosition)
    board.render(Color.WHITE)

    RequestRoomsAction(repository, rooms)
    setInterval(() => RequestRoomsAction(repository, rooms), 10000)
})

function createRoom() {
    CreateRoomAction(repository, playerID, `room-${Date.now().toString()}`, 'roomPassword')
}

</script>

<template>
    <main>
        <div class="vh-100 position-relative">
            <div class="container h-auto position-absolute top-50 start-50 translate-middle">
                <div class="my-2 text-light">{{ playerID }}</div>
                <div id="chess-board" class="row"></div>

                <!-- Buttons -->
                <div class="row mt-3 text-center actions-btns">
                    <button type="button" class="col-sm btn btn-dark border border-light m-2" data-bs-toggle="modal"
                        data-bs-target="#modal-list-rooms" @click="createRoom()">
                        Create room
                    </button>
                    <button type="button" class="col-sm btn btn-dark border border-light m-2">
                        Abandon
                    </button>
                    <button type="button" class="col-sm btn btn-dark border border-light m-2">
                        Draw
                    </button>
                </div>
            </div>
        </div>
    </main>
</template>
