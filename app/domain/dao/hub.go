package dao

type Room struct {
	ID      int                `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	// Registered clients.
	Rooms map[string]*Room `json:"rooms"`
	// Inbound messages from the clients.
	Broadcast chan *Message
	// Register requests from the clients.
	Register chan *Client
	// Unregister requests from clients.
	Unregister chan *Client
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
							Content:  cl.Username + " has left the room",
							RoomId:   cl.RoomId,
							Username: cl.Username,
						}
					}
					delete(r.Clients, cl.ID)
					close(cl.Message)
				}
			}
		case msg := <-h.Broadcast:
			if _, ok := h.Rooms[msg.RoomId]; ok {
				r := h.Rooms[msg.RoomId]
				for _, cl := range r.Clients {
					cl.Message <- msg
				}
			}

		}
	}
}
