package dao

type Room struct {
	ID      int                `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room `json:"rooms"`
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}
func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			if _, ok := h.Rooms[cl.RoomId]; ok {
				r := h.Rooms[cl.RoomId]
				if _, ok := r.Clients[cl.ID]; !ok {
					r.Clients[cl.ID] = cl
				}
			}
		case cl := <-h.Unregister:
			if _, ok := h.Rooms[cl.RoomId]; ok {
				r := h.Rooms[cl.RoomId]
				if _, ok := r.Clients[cl.ID]; ok {
					if len(r.Clients) != 0 {
						h.Broadcast <- &Message{
							Content:  "user left the chat",
							RoomId:   cl.RoomId,
							Username: cl.Username,
						}
					}
					delete(r.Clients, cl.ID)
					close(cl.Message)
				}
			}

		case m := <-h.Broadcast:
			if _, ok := h.Rooms[m.RoomId]; ok {
				r := h.Rooms[m.RoomId]
				for _, cl := range r.Clients {
					cl.Message <- m
				}
			}

		}
	}
}
