package assets

import (
	"encoding/json"
	"fmt"
)

const Sectors = 8

type V3 struct {
	x float64
	y float64
	z float64
}

type Reference struct {
	Ips    string
	IsName bool `json:",omitempty"`
	Path   string
}

type WorldStore struct {
	Title string
	Grid  [][]Reference
}

func NewWorldStore() *WorldStore {
	ws := &WorldStore{}
	grid := make([][]Reference, Sectors)
	for i := range grid {
		grid[i] = make([]Reference, Sectors)
	}
	ws.Grid = grid
	return ws
}

type SectorStore struct {
	Name   string
	Assets []AssetItem
}

type AssetItem struct {
	Ref Reference
	Pos V3
	Rot V3
}

func export() {
	ws := NewWorldStore()
	j, err := json.MarshalIndent(ws, "", "\t")
	if err != nil {
		fmt.Println("json error", err.Error())
	}
	fmt.Println(string(j))
}
