import { type Ref, ref, shallowRef } from 'vue'
import { RoomPlayer } from './room_player'

class Room {
    public readonly id: string

    private _players: Array<RoomPlayer>
    public get players(): Array<RoomPlayer> {
        return this._players
    }
    private _myRoom: boolean
    public get myRoom(): boolean {
        return this._myRoom
    }

    constructor(id?: string, players?: Array<RoomPlayer>) {
        this.id = id === null || id === undefined ? '' : id
        this._players = players === null || players === undefined ? new Array(0) : players
        this._myRoom = false
    }

    public static createFromJSON(src: any): Room {
        const dst: Room = new Room(src.id)
        dst._players = src.players.map((p: any) => RoomPlayer.createFromJSON(p))
        return dst
    }

    public setMyRoom() {
        this._myRoom = true
    }

    public isRoomFull(): boolean {
        return this.players.length >= 2
    }
}

class Rooms {
    private _rooms: Ref<Array<Room>>
    public getRooms(): Array<Room> {
        return this._rooms.value
    }
    private _myRoom: Ref<Room | null>
    public get myRoom(): Room | null {
        return this._myRoom.value
    }
    public getMyRoom(): Room | null {
        return this._myRoom.value
    }

    constructor() {
        this._rooms = shallowRef(Array<Room>())
        this._myRoom = ref(null)
    }

    public setRooms(rooms: Array<Room>) {
        if (!rooms || rooms.length === 0) {
            this._rooms.value = Array<Room>()
            return
        }

        this._rooms.value = rooms.map((i) => Room.createFromJSON(i))
        this._rooms.value = this._rooms.value.sort((r1, r2) => {
            // We want most recent not full rooms first
            if (r1.isRoomFull()) {
                return 1
            }

            if (r2.isRoomFull()) {
                return -1
            }

            return r1.id > r2.id ? -1 : 1
        })
    }

    public setMyRoom(myRoom: Room) {
        this._myRoom.value = myRoom
        myRoom.setMyRoom()
    }
}

export { Room, Rooms }
