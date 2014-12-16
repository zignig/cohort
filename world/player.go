package world

import (
	"fmt"
	"time"
)

//
type Player struct {
	name  string
	pos   V3
	rot   E4
	alive bool
	stat  *gridStatus
	// channel for messages to player
	InMess  chan []byte
	OutMess chan []byte
	Closer  chan bool
	// reverse reference for world
	world *World
}

// generates a new player and hand back a
func (w *World) NewPlayer() (thePlayer *Player) {
	fmt.Println("Create New Player")
	p := &Player{}
	p.stat = NewGridStatus()
	// mesage channel for player actions
	// mesage from client
	p.InMess = make(chan []byte, 5)
	// message to client
	p.OutMess = make(chan []byte, 5)
	// closer for player loop
	p.Closer = make(chan bool)
	// world reference
	p.world = w
	// add player to world.
	w.playerLock.Lock()
	w.players[p] = true
	w.playerChan <- p
	w.playerLock.Unlock()

	return p
}

// internal player loop
func (p *Player) Run() {
	fmt.Println("starting player")
	ticker := time.NewTicker(time.Second * 5).C
	pm := &playMessage{}
	for {
		select {
		case <-ticker:
			fmt.Println("do player stuff")
		case m := <-p.InMess:
			pm.Decode(m)
			// update the player
			p.Update(pm)
		case <-p.Closer:
			fmt.Println("close player")
			return
		}
	}

}

func (p *Player) Update(pm *playMessage) {
	// update player based on location
	switch pm.Class {
	case "location":
		{
			loc := pm.Data.(*PosMessage)
			loc.Pos.Sector()
		}
	}

}

// find the current sector of a player
func (pos *V3) Sector() (x int, y int) {
	fmt.Println(pos.X, pos.Y, pos.Z)
	secx := int((pos.X + (SectorSize / 2)) / SectorSize)
	secz := int((pos.Z + (SectorSize / 2)) / SectorSize)
	fmt.Println("into => [", secx, ",", secz, "]")
	return 0, 0
}
