package main

import (
	"fmt"

	"github.com/zignig/viewer/assets"
)

// world structures

// make a world sectors * sectors big

const Sectors = 8
const SectorSize = 256

// 3 vector
// look for some math libs for V3
type V3 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

// 4 d euler
type E4 struct {
	X float64 `json:"_x"`
	Y float64 `json:"_y"`
	Z float64 `json:"_z"`
	W float64 `json:"_w"`
}

type PosMessage struct {
	Pos  V3     `json:"pos"`
	Rot  E4     `json:"rot"`
	Uuid string `json:"uuid"`
}

// boolean status for each player of the
// grid
type gridStatus struct {
	grid [][]bool
}

func NewGridStatus() *gridStatus {
	gs := &gridStatus{}
	grid := make([][]bool, Sectors)
	for i := range grid {
		grid[i] = make([]bool, Sectors)
	}
	gs.grid = grid
	fmt.Println(gs)
	return gs
}

//
type player struct {
	name  string
	pos   V3
	alive bool
	stat  *gridStatus
}

type entity struct {
	ref  string
	data []byte
	pos  V3
}

type Sector struct {
	ref   string
	owner string
	ents  []*entity
}

type World struct {
	players  map[*player]bool
	grid     [][]*Sector
	status   *gridStatus
	register chan *connection
	// cache for assets

	cache assets.Cache
}

func NewWorld() *World {
	w := &World{}
	grid := make([][]*Sector, Sectors)
	for i := range grid {
		grid[i] = make([]*Sector, Sectors)
	}
	w.grid = grid
	w.status = NewGridStatus()
	w.register = make(chan *connection)
	return w
}

func (w *World) run() {
	for {
		select {
		case c := <-w.register:
			fmt.Println("new world registration")
			//w.players[c] = true
			c.send <- []byte("fnordy fnord fnord fnord")
		}
	}

}
