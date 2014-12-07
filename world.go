package main

import (
	"fmt"
)

// world structures

// make a world sectors * sectors big

const sectors = 8

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
	grid := make([][]bool, sectors)
	for i := range grid {
		grid[i] = make([]bool, sectors)
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

type sector struct {
	ref   string
	owner string
	ents  []*entity
}

type world struct {
	players []player
	grid    [][]*sector
}

func NewWorld() *world {
	w := &world{}
	grid := make([][]*sector, sectors)
	for i := range grid {
		grid[i] = make([]*sector, sectors)
	}
	w.grid = grid
	return w
}
