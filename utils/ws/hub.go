package ws

type Broadcast struct {
	msg    []byte
	client *Client
}

type Hub struct {
	// clients map[*Client]struct{}

	// map[channels ID][]*Client
	channels map[string]map[*Client]struct{}

	alone map[*Client]bool

	broadcast chan *Broadcast

	register chan *Client

	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		// clients:    make(map[*Client]struct{}),
		channels:   make(map[string]map[*Client]struct{}),
		alone:      make(map[*Client]bool),
		broadcast:  make(chan *Broadcast),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.Register(client)
		case client := <-h.unregister:
			h.Unregister(client)
		case b := <-h.broadcast:
			h.DoBroadcast(b)
		}
	}
}

func (h *Hub) Register(client *Client) {
	if len(client.channel) != 0 {
		h.regisChannel(client)
		return
	} else {
		h.regisAlone(client)
		return
	}
}

func (h *Hub) regisAlone(c *Client) {
	if len(c.channel) != 0 {
		h.regisChannel(c)
		return
	}
	h.alone[c] = true
}

func (h *Hub) regisChannel(c *Client) {
	if len(c.channel) == 0 {
		h.regisAlone(c)
		return
	}
	if _, ok := h.channels[c.channel]; !ok {
		h.channels[c.channel] = make(map[*Client]struct{})
	}
	h.channels[c.channel][c] = struct{}{}
}

func (h *Hub) Unregister(client *Client) {
	if len(client.channel) != 0 {
		h.unregisterChannel(client)
	} else {
		h.unregisterAlone(client)
	}
	close(client.send)
	return
}
func (h *Hub) unregisterAlone(client *Client) {
	if _, ok := h.alone[client]; ok {
		delete(h.alone, client)
	}
}

func (h *Hub) unregisterChannel(client *Client) {
	if _, ok := h.channels[client.channel]; !ok {
		return
	}
	if _, ok := h.channels[client.channel][client]; ok {
		delete(h.channels[client.channel], client)
	}
}

func (h *Hub) DoBroadcast(b *Broadcast) {
	if len(b.client.channel) != 0 {
		if !h.HasChannel(b.client.channel) {
			return
		}
		for client, _ := range h.channels[b.client.channel] {
			select {
			case client.send <- b.msg:
			default:
				close(client.send)
				delete(h.channels[client.channel], client)
			}
		}
	} else {
		select {
		case b.client.send <- b.msg:
		default:
			close(b.client.send)
			delete(h.alone, b.client)
		}
	}
}

func (h *Hub) HasChannel(channel string) bool {
	if _, ok := h.channels[channel]; ok {
		return true
	}
	return false
}
