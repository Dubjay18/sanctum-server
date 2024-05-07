package dao

type Room struct {
	ID      int                `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms map[int]*Room `json:"rooms"`
}

func NewHub() *Hub {
	return &Hub{
		Rooms: make(map[int]*Room),
	}
}
