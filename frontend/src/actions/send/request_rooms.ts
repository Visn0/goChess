import { ConnectionRepository } from '../../connection_repository/connection_repository'
import { Room, Rooms } from '../../room'

class RequestRoomsReponse {
    public rooms: Array<Room>
}

function RequestRoomsAction(repository: ConnectionRepository, rooms: Rooms) {
    console.log("this")
    repository
        .sendHTTPRequest('GET', 'rooms', null)
        .then((response) => response.json())
        .then((jsonBody: RequestRoomsReponse) => {
            rooms.setRooms(jsonBody.rooms)
            rooms.render()
        })
        .catch((err) => console.log(err))
}

export { RequestRoomsAction }
