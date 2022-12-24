import { Board } from './board'
import { GameController } from './game_controller'

const board: Board = new Board(document.getElementById('chess-board') as HTMLElement)
const gameController: GameController = new GameController(board, 'localhost', '')

function init() {
    gameController.start()
}

window.onload = () => {
    init()
}

export {}
