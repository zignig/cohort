package world

import (
	"fmt"
	"time"
)

//
type player struct {
	name  string
	pos   V3
	alive bool
	stat  *gridStatus
	// channel for messages to player
	Message chan []byte
}

// generates a new player and hand back a
func (w *World) NewPlayer() (thePlayer *player) {
	fmt.Println("Create New Player")
	p := &player{}
	// mesage channel for player actions
	p.Message = make(chan []byte, 5)
	w.playerLock.Lock()
	w.players[p] = true
	w.playerLock.Unlock()

	return p
}

// internal player loop
func (p *player) Run() {
	fmt.Println("starting player")

	//	for {
	ticker := time.NewTicker(time.Second * 5).C
	pm := &playMessage{}
	for {
		select {
		case <-ticker:
			fmt.Println("do player stuff")
		case m := <-p.Message:
			pm.Decode(m)
			fmt.Println(pm)
		}
	}

}
