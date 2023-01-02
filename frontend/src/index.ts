import { BackendConnectionRepository } from './connection_repository/backend_connection_repository'
import { ConnectionRepository } from './connection_repository/connection_repository'
import { MockConnectionRepository } from './connection_repository/mock_connection_repository'
import { Board } from './board'
import { Color, constants } from './constants'
import { GameController } from './game_controller'
import { RequestRoomsAction } from './actions/send/request_rooms'

const board: Board = new Board(document.getElementById('chess-board') as HTMLElement)
// const repository: ConnectionRepository = new MockConnectionRepository()
const repository: ConnectionRepository = new BackendConnectionRepository('localhost', '8081', 'ws')
const gameController: GameController = new GameController(board, repository)

window.onload = () => {
    board.initFromFenNotation(constants.StartingPosition)
    board.render(Color.WHITE)
    gameController.openWebSocketConnetion()

    RequestRoomsAction(repository, 'modal-list-rooms-body')
    setInterval(() => RequestRoomsAction(repository, 'modal-list-rooms-body'), 10000)
}

const btnCreateRoom = document.getElementById('btn-create-room') as HTMLElement
btnCreateRoom.onclick = () => {
    // gameController.createRoom('room', 'pass')
    repository.sendHTTPRequest('GET', 'rooms', null)
}

export {}
