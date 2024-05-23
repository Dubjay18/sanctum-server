package dto

type CreateRoomRequest struct {
	RoomName string `json:"room_name"`
	RoomId   string `json:"room_id"`
}
