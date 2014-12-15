package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zignig/viewer/world"
)

// message from player
type playMessage struct {
	Class   string          `json:"class"`
	Message json.RawMessage `json:"message"`
	Data    interface{}
}

// TODO ove to world ?
func (pm *playMessage) Decode(m []byte) {
	var dst interface{}
	//pm := &playMessage{}
	err := json.Unmarshal(m, pm)
	if err != nil {
		fmt.Println("base decode ", err)
	}
	//fmt.Println("class ", pm.Class)
	switch pm.Class {
	case "location":
		{
			dst = &world.PosMessage{}
			//fmt.Println("got location ", string(pm.Message))
		}
	}
	err = json.Unmarshal(pm.Message, dst)
	pm.Data = dst
	if err != nil {
		fmt.Println("message error ", err)
	}
	//fmt.Println(dst)
}

// player hub
type hub struct {
	// Registered connections.
	connections map[*connection]bool

	// Inbound messages from the connections.
	broadcast chan []byte

	// Register requests from the connections.
	register chan *connection

	// Unregister requests from connections.
	unregister chan *connection

	world *world.World
}

// readPump pumps messages from the websocket connection to the hub.
func (c *connection) readPump() {
	defer func() {
		u.h.unregister <- c
		c.ws.Close()
	}()
	//
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	m := &playMessage{}
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}

		//fmt.Println(string(message))
		m.Decode(message)
		fmt.Println(m.Data)
		//h.broadcast <- message
	}
}

// write writes a message with the given message type and payload.
func (c *connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// writePump pumps messages from the hub to the websocket connection.
func (c *connection) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// lifted from  The Gorilla WebSocket Authors. All rights reserved.
type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

func NewHub(w *world.World) (h *hub) {
	h = &hub{}
	h.broadcast = make(chan []byte)
	h.register = make(chan *connection)
	h.unregister = make(chan *connection)
	h.connections = make(map[*connection]bool)
	h.world = w
	return h
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			fmt.Println("new connection")
			h.connections[c] = true
			c.send <- []byte("load")
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.send)
			}
			//case m := <-h.broadcast:
			//fmt.Println(string(m))
			//for c := range h.connections {
			//	select {
			//	case c.send <- m:
			//	default:
			//		close(c.send)
			//		delete(h.connections, c)
			//	}
			//}
		}
	}
}
