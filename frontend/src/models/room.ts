import { RoomPlayer } from './room_player'

class Room {
    public id: string
    public players: Array<RoomPlayer>
    public myRoom: boolean

    constructor(id?: string, players?: Array<RoomPlayer>) {
        this.id = id === null || id === undefined ? '' : id
        this.players = players === null || players === undefined ? new Array(0) : players
        this.myRoom = false
    }

    public toDivHTML(): string {
        let room = `
            <div id="${this.id}" class="card w-100 bg-dark text-white">
                <div class="card-header bg-green mx-1">                                        
        `

        if (!this.myRoom) {
            room += `
                    <div class="position-absolute" style="margin-top: -5px;">
                        <button id="join-btn-${this.id}" class="btn btn-success btn-sm text-left"
                            onclick="document.dispatchEvent(new CustomEvent('join-room-event', { detail: '${this.id}' }))">
                            Join
                        </button>
                    </div>
            `
        }

        room += `
                <div>
                    ${this.id}
                </div>                    
            </div>
            <ul class="list-group list-group-flush row mx-1">
        `

        if (this.players.length > 0) {
            this.players.forEach((p) => (room += p.toDivHTML()))
        } else {
            room += '<li class="list-group-item">No players yet.</li>'
        }
        return (room += '</ul>\n</div>')
    }

    public static createFromJSON(src: any): Room {
        const dst: Room = Object.assign(new Room(), src)
        dst.players = dst.players.map((p) => RoomPlayer.createFromJSON(p))
        return dst
    }

    public setMyRoom() {
        this.myRoom = true
    }
}

class Rooms {
    private container: HTMLElement
    private rooms: Array<Room>
    private myRoom: Room | null

    constructor(containerID: string) {
        this.container = document.getElementById(containerID) as HTMLElement
        this.rooms = Array<Room>()
        this.myRoom = null
    }

    public render() {
        let rooms = this.myRoom ? this.myRoom.toDivHTML() : ''

        if (rooms !== '' && this.rooms.length > 0) {
            rooms += '\n<li class="list-group-item border-top border-secondary my-4 p-0"></li>'
        }

        for (let i = 0; i < this.rooms.length; i++) {
            rooms += this.rooms[i].toDivHTML()
            if (i < this.rooms.length - 1) {
                rooms += '\n<br>'
            }
        }

        this.container.innerHTML = rooms
    }

    public setRooms(rooms: Array<Room>) {
        this.rooms = rooms.map((i) => Room.createFromJSON(i))
        this.rooms = this.rooms.sort((r1, r2) => (r1.id > r2.id ? 1 : -1))
    }

    public setMyRoom(myRoom: Room) {
        this.myRoom = myRoom
        myRoom.setMyRoom()
    }
}

export { Room, Rooms }
