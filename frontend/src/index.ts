import { BackendConnectionRepository } from './connection_repository/backend_connection_repository'
import { ConnectionRepository } from './connection_repository/connection_repository'
import { MockConnectionRepository } from './connection_repository/mock_connection_repository'
import { Board } from './board'
import { Color, constants } from './constants'
import { GameController } from './game_controller'
import { RequestRoomsAction } from './actions/send/request_rooms'
import { Rooms } from './room'

const rooms = new Rooms('modal-list-rooms-body')
const board: Board = new Board(document.getElementById('chess-board') as HTMLElement)
// const repository: ConnectionRepository = new MockConnectionRepository()
const repository: ConnectionRepository = new BackendConnectionRepository('localhost', '8081', 'ws')
const gameController: GameController = new GameController(rooms, board, repository)

window.onload = () => {
    board.initFromFenNotation(constants.StartingPosition)
    board.render(Color.WHITE)

    gameController.openWebSocketConnection()
    RequestRoomsAction(repository, rooms)
    setInterval(() => RequestRoomsAction(repository, rooms), 10000)
}

const btnCreateRoom = document.getElementById('btn-create-room') as HTMLElement
btnCreateRoom.onclick = () => {
    gameController.createRoom('paquito', 'lepass')
}

export {}
