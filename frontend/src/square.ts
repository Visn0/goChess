import { Color, File, Rank, eventTopics } from './constants'
import { Piece } from './piece'

class Square {
    public readonly id: string
    public readonly color: string
    public readonly file: File
    public readonly rank: Rank

    private _piece: Piece | null
    public get piece(): Piece | null {
        return this._piece
    }

    constructor(file: File, rank: Rank, piece: Piece | null) {
        this.file = file
        this.rank = rank
        this._piece = piece

        this.id = `${String.fromCharCode(65 + file)}${rank + 1}`
        this.color = (file + rank) % 2 === 0 ? `${Color.BLACK}-square` : `${Color.WHITE}-square`
    }

    public toTableCellHTML(): string {
        const piece = this._piece?.toDivHTMLString() || ''
        const cell = `<td id="${this.id}" class="board-square ${this.color}">${piece}</td>`
        return cell
    }

    public addOnClickEventListener() {
        const element = document.getElementById(this.id) as HTMLElement
        const file = this.file
        const rank = this.rank
        element.onclick = () => {
            const payload = {
                detail: {
                    file: file,
                    rank: rank
                }
            }

            document.dispatchEvent(new CustomEvent(eventTopics.OnSquareClick, payload))
        }
    }

    public isEmpty(): boolean {
        return this._piece === null
    }

    private updateRender() {
        const piece = this._piece?.toDivHTMLString() || ''
        const element = document.getElementById(this.id) as HTMLElement
        element.innerHTML = piece
    }

    public setPiece(piece: Piece | null) {
        this._piece = piece
        this.updateRender()
    }
}

export { Square }
