import { Board } from './board'

const board: Board = new Board(document.getElementById('chess-board') as HTMLElement)

function init() {
    board.reset()
}

window.onload = () => {
    init()
}

export {}
