package assets

import (
	"encoding/json"
	"fmt"
)

const Sectors = 8

type V3 struct {
	X float64
	Y float64
	Z float64
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

func LoadWorldStore(data []byte) (ws *WorldStore, err error) {
	err = json.Unmarshal(data, &ws)
	return ws, err
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
	Assets []*AssetItem
}

type AssetItem struct {
	Ref Reference
	Pos V3
	Rot V3
}

func export() {
	ws := &SectorStore{}
	ws.Assets = append(ws.Assets, &AssetItem{})
	j, err := json.MarshalIndent(ws, "", "\t")
	if err != nil {
		fmt.Println("json error", err.Error())
	}
	fmt.Println(string(j))
}

func dump(ws interface{}) {
	j, err := json.MarshalIndent(ws, "", "\t")
	if err != nil {
		fmt.Println("json error", err.Error())
	}
	fmt.Println(string(j))
}
