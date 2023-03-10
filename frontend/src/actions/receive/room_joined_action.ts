import { Room, Rooms } from '@/models/room'
import { RoomPlayer } from '@/models/room_player'
import type { ReceiveAction } from './receive_action'

class RoomJoinedParams {
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

class RoomJoinedAction implements ReceiveAction {
    private rooms: Rooms

    constructor(rooms: Rooms) {
        this.rooms = rooms
    }

    public Invoke(body: string) {
        const p: RoomJoinedParams = JSON.parse(body)
        if (p.httpCode === 400) {
            alert('Room already exists.')
            return
        }

        let players = Object.assign(new Array<RoomPlayer>(), p.room.players)
        players = players.map((p) => RoomPlayer.createFromJSON(p))
        const myRoom = new Room(p.room.id, players)

        this.rooms.setMyRoom(myRoom)
    }
}

export { RoomJoinedAction }
