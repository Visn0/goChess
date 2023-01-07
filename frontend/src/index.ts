import { BackendConnectionRepository } from './connection_repository/backend_connection_repository'
import { ConnectionRepository } from './connection_repository/connection_repository'
/* eslint-disable  @typescript-eslint/no-unused-vars */
import { MockConnectionRepository } from './connection_repository/mock_connection_repository'
/* eslint-enable  @typescript-eslint/no-unused-vars */
import { Board } from './board'
import { Color, constants } from './constants'
import { GameController } from './game_controller'
import { RequestRoomsAction } from './actions/send/request_rooms'
import { Rooms } from './room'
import { JoinRoomAction } from './actions/send/join_room_action'
import { RouteActions } from './actions/receive/route_actions'
import { MovesReceivedAction } from './actions/receive/moves_received_action'
import { PieceMovedAction } from './actions/receive/piece_moved_action'
import { RoomCreatedAction } from './actions/receive/room_created_action'
import { ReceiveAction } from './actions/receive/receive_action'
import { CreateRoomAction } from './actions/send/create_room_action'

const playerID = 'MiPlayerID'
const rooms = new Rooms('modal-list-rooms-body')
const board: Board = new Board(document.getElementById('chess-board') as HTMLElement)
/* eslint-disable capitalized-comments */
// const repository: ConnectionRepository = new MockConnectionRepository()
/* eslint-enable capitalized-comments */
const repository: ConnectionRepository = new BackendConnectionRepository('localhost', '8081', 'ws')
const gameController: GameController = new GameController(rooms, board, repository)

const routeActions: RouteActions = new RouteActions(
    new Map<string, ReceiveAction>([
        ['create-room', new RoomCreatedAction(rooms)],
        ['request-moves', new MovesReceivedAction(gameController)],
        ['move-piece', new PieceMovedAction(gameController)]
    ])
)
repository.addOnWebSocketMessageEventListener(routeActions.route())

window.onload = () => {
    board.initFromFenNotation(constants.StartingPosition)
    board.render(Color.WHITE)

    RequestRoomsAction(repository, rooms)
    setInterval(() => RequestRoomsAction(repository, rooms), 10000)
}

const btnsCreateRoom = document.getElementsByName('btn-create-room')
btnsCreateRoom.forEach((btn) => {
    btn.onclick = () => {
        CreateRoomAction(repository, 'userID', 'roomID', 'roomPassword')
    }
})

document.addEventListener('join-room-event', (e: Event) => {
    const ce = e as CustomEvent
    JoinRoomAction(repository, playerID, ce.detail, 'roomPassword')
})

export {}
