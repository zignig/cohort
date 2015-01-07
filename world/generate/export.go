package generate

import (
	"encoding/json"
)

// quick faker export

var BaseTile = map[int]string{
	Empty:   "sand",
	Grass:   "grass",
	Sand:    "sand",
	Water:   "water",
	Road:    "road-straight-low",
	Rail:    "lot",
	Blocker: "lot",
}

// single tile
type tile struct {
	Name   string
	Rotate int
}

type tileGrid struct {
	Ref  string
	Grid [][]*tile
}

func (w World) Export() (b []byte) {

	gs := &tileGrid{}
	grid := make([][]*tile, len(w))
	for i := range grid {
		grid[i] = make([]*tile, len(w[0]))
	}
	gs.Grid = grid
	gs.Ref = "QmPDtmEz4zTpYvUXe4ZunzCEW5YGA1fG9qTasaiqkZ82HZ"

	width := len(w)
	height := len(w[0])
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			t := w.Tile(i, j)
			grid[i][j] = &tile{BaseTile[t.Kind], 0}
		}
	}
	b, _ = json.Marshal(&gs)
	return
}
