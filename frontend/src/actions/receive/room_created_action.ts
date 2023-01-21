import { Room, Rooms } from '@/models/room'
import { RoomPlayer } from '@/models/room_player'
import { ReceiveParams } from './receive_params'

class RoomCreatedParams extends ReceiveParams {
    httpCode: number
    room: RoomParams
}

class RoomParams {
    id: string
    players: Array<RoomPlayerParams>
}

class RoomPlayerParams {
    id: string
}

class RoomCreatedAction {
    private rooms: Rooms

    constructor(rooms: Rooms) {
        this.rooms = rooms
    }

    public Invoke(body: string) {
        console.log(body)
        const p: RoomCreatedParams = JSON.parse(body)
        if (p.error) {
            alert(`Room already exists: ${p.error.key}`)
            return
        }

        let players = Object.assign(new Array<RoomPlayer>(), p.room.players)
        players = players.map((p) => RoomPlayer.createFromJSON(p))
        const myRoom = new Room(p.room.id, players)

        this.rooms.setMyRoom(myRoom)
    }
}

export { RoomCreatedAction }
