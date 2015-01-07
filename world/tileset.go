package world

// functions for handling tile sets
import (
	"container/list"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/zignig/viewer/assets"
	"github.com/zignig/viewer/world/generate"
)

// single tile
type tile struct {
	Name   string
	Rotate int
}

type tileGrid struct {
	Ref  string
	Grid [][]*tile
}

func NewTileGrid() *tileGrid {
	gs := &tileGrid{}
	grid := make([][]*tile, Sectors)
	for i := range grid {
		grid[i] = make([]*tile, Sectors)
	}
	gs.Grid = grid
	fmt.Println(gs)
	return gs
}

// sends a blank sand floor
func (w *World) SendFloor(p *Player, x int, y int) {
	offx := float64(x * SectorSize)
	offy := float64(y * SectorSize)
	// send a floor builder
	fl := &FloorMessage{}
	fl.Pos.X = offx
	fl.Pos.Z = offy
	fl.Size = SectorSize
	data, err := Encode(fl)
	if err != nil {
		fmt.Println("floor fail ", err)
		return
	}
	fmt.Println(string(data))
	p.OutMess <- data
}

// send all tiles in world
func (w *World) SendTiles(p *Player) {
	for i := range w.tiles.Grid {
		for j := range w.tiles.Grid[i] {
			w.SendTile(p, i, j)
		}
	}
}

// send a specific tile to client
func (w *World) SendTile(p *Player, x int, y int) {
	offx := float64(x * SectorSize)
	offy := float64(y * SectorSize)
	// send a floor builder
	ti := &TileMessage{}
	ti.Ref = w.config.Tile
	ti.Pos.X = offx
	ti.Pos.Y = -2
	ti.Pos.Z = offy
	ti.Name = w.tiles.Grid[x][y].Name
	data, err := Encode(ti)
	if err != nil {
		fmt.Println("floor fail ", err)
		return
	}
	fmt.Println(string(data))
	p.OutMess <- data
}

func (w *World) MakeWorld(s int) *tileGrid {
	t := generate.Basic(s)
	tg := &tileGrid{}
	err := json.Unmarshal(t, tg)
	if err != nil {
		fmt.Println(err)
	}
	return tg
}
func (w *World) GenTiles() *tileGrid {
	tiles := w.config.Tile
	st, err := w.cache.Ls(tiles + "/tiles")
	if err != nil {
		fmt.Println("FAIL resolve tile listing")
	}
	fmt.Print(string(st))
	tileList := &assets.Listing{}
	json.Unmarshal(st, tileList)
	// get a list of the links
	links := tileList.Objects[0].Links

	objs := list.New()
	for i := range links {
		TileName := links[i].Name
		if strings.HasSuffix(TileName, ".obj") {
			tn := strings.TrimSuffix(TileName, ".obj")
			fmt.Println("tile name ", tn)
			objs.PushBack(tn)
		}
	}
	fmt.Println(objs.Len(), " tiles in hash")
	tg := NewTileGrid()
	for i := range tg.Grid {
		for j := range tg.Grid[i] {
			tile := &tile{}
			tg.Grid[i][j] = tile
			f := objs.Front()
			if f == nil {
				tile.Name = "grass"
			} else {
				tile.Name = f.Value.(string)
				objs.Remove(f)
			}
		}
	}
	fmt.Println(tg)
	return tg
}
