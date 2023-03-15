package domain

import "sync"

type RoomManager struct {
	rooms map[string]*Room
	mu    sync.RWMutex
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		rooms: make(map[string]*Room),
	}
}

func (rm *RoomManager) GetRoom(id string) (*Room, bool) {
	rm.mu.RLock()
	defer rm.mu.RUnlock()
	r, exists := rm.rooms[id]
	return r, exists
}

func (rm *RoomManager) GetRooms() []*Room {
	rm.mu.RLock()
	defer rm.mu.RUnlock()
	rooms := make([]*Room, 0, len(rm.rooms))
	for _, r := range rm.rooms {
		rooms = append(rooms, r)
	}
	return rooms
}

func (rm *RoomManager) AddRoom(room *Room) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	rm.rooms[room.ID] = room
}

func (rm *RoomManager) RemoveRoom(id string) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	delete(rm.rooms, id)
}
