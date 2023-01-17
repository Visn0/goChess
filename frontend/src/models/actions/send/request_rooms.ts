import type { ConnectionRepository } from '../../connection_repository/connection_repository'
import type { Room, Rooms } from '../../room'

class RequestRoomsReponse {
    public rooms: Array<Room>
}

function RequestRoomsAction(repository: ConnectionRepository, rooms: Rooms) {
    repository
        .sendHTTPRequest('GET', 'rooms', null)
        .then((response) => response.json())
        .then((jsonBody: RequestRoomsReponse) => {
            rooms.setRooms(jsonBody.rooms)
        })
        .catch((err) => console.log(err))
}

export { RequestRoomsAction }
