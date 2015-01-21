package world

import (
	"fmt"
	//	"time"
	"github.com/zignig/cohort/assets"
)

//
type Player struct {
	name  string
	pos   V3
	rot   E4
	alive bool
	stat  *gridStatus
	// current sector
	grx, gry int
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
	//ticker := time.NewTicker(time.Second * 5).C
	pm := &playMessage{}
	for {
		select {
		//case <-ticker:
		//fmt.Println(p.stat)
		case m := <-p.InMess:
			// decode the player message
			pm.Decode(m)
			// update the player
			p.Update(pm)
		case <-p.Closer:
			fmt.Println("close player")
			p.world.playerLock.Lock()
			delete(p.world.players, p)
			p.world.playerLock.Unlock()
			return
		}
	}

}

func (p *Player) Update(pm *playMessage) {
	// update player based on location
	switch pm.Class {
	case "location":
		{
			loc := pm.data.(*PosMessage)
			x, y := loc.Pos.Sector()
			//fmt.Println("player in ", x, ",", y)
			if (x >= 0) && (y >= 0) {
				p.grx = x
				p.gry = y
				status := p.stat.grid[x][y]
				//fmt.Printf("%v", p.stat)
				// sector has not been visisted
				if !status {
					fmt.Println("activate ", x, y)
					p.stat.grid[x][y] = true
					p.world.loaderChan <- p
				}
			}
		}
	}

}

func (p *Player) SendSector(ss *assets.SectorStore, x int, y int) {
	offx := float64(x * SectorSize)
	offy := float64(y * SectorSize)

	if ss == nil {
		return
	}

	// send each asset to client
	for i := range ss.Assets {
		theAsset := ss.Assets[i]
		fmt.Println("send asset to client")
		fmt.Println(theAsset)
		lm := &LoaderMessage{}
		lm.Path = ss.Ref + theAsset.Path
		lm.Pos = theAsset.Pos
		// calculate the sector offset

		lm.Pos.X = lm.Pos.X + offx
		lm.Pos.Z = lm.Pos.Z + offy

		lm.Rot = theAsset.Rot
		data, err := Encode(lm)
		if err != nil {
			fmt.Println("encode fail ", err)
		}
		fmt.Println(string(data))
		p.OutMess <- data
	}
}
