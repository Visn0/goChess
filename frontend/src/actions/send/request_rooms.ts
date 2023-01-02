import { ConnectionRepository } from '../../connection_repository/connection_repository'

class RoomPlayer {
    id: string

    public toDivHTML(): string {
        const htmlStr = `<li class="list-group-item">${this.id}</li>`
        return htmlStr
    }

    public static createFromJSON(src: any): RoomPlayer {
        const dst: RoomPlayer = Object.assign(new RoomPlayer(), src)
        return dst
    }
}

class Room {
    public id: string
    public players: Array<RoomPlayer>

    public toDivHTML(): string {
        let room = `
        <div id="room-${this.id}" class="card w-100 bg-dark text-white">
            <div class="card-header p-0 bg-green d-grid">
                <button type="button" class="btn" class="">
                    Room: ${this.id}
                </button>
            </div>
            <ul class="list-group list-group-flush">
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
}

class RequestRoomsReponse {
    public rooms: Array<Room>
}
function RequestRoomsAction(repository: ConnectionRepository, modalBodyID: string) {
    repository
        .sendHTTPRequest('GET', 'rooms', null)
        .then((response) => response.json())
        .then((jsonBody: RequestRoomsReponse) => {
            let rooms: Array<Room> = jsonBody.rooms.map((i) => Room.createFromJSON(i))
            rooms = rooms.sort((r1, r2) => r1.id > r2.id ? 1 : -1)
            renderRooms(rooms, modalBodyID)
        })
        .catch((err) => console.log(err))
}

function renderRooms(rooms: Array<Room>, modalBodyID: string) {
    const modalBodyElement = document.getElementById(modalBodyID) as HTMLElement
    modalBodyElement.innerHTML = ''

    for (let i = 0; i < rooms.length; i++) {
        modalBodyElement.innerHTML += rooms[i].toDivHTML()
        if (i < rooms.length - 1) {
            modalBodyElement.innerHTML += '\n<br>'
        }
    }
}

export { RequestRoomsAction }
