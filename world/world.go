package world

import (
	"fmt"
	"sync"
	"time"

	"github.com/zignig/cohort/assets"
	"github.com/zignig/cohort/util"
)

// world structures

// make a world sectors * sectors big

const Sectors = 12    // square layout X*X
const SectorSize = 32 // prepare for tiles

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

// boolean status for each player of the
// grid
type gridStatus struct {
	grid [][]bool
}

// grid printer
func (g *gridStatus) String() (s string) {
	for i := range g.grid {
		s = s + fmt.Sprint("|")
		for j := range g.grid[i] {
			stat := g.grid[i][j]
			if stat {
				s = s + fmt.Sprint("O", "|")
			} else {
				s = s + fmt.Sprint("X", "|")
			}
		}
		s = s + fmt.Sprintln("")
	}
	return s
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

type World struct {
	players map[*Player]bool
	grid    [][]*assets.SectorStore
	status  *gridStatus
	// tiles
	tiles  *tileGrid
	cache  *assets.Cache
	config *util.Config
	ref    string
	// lock for player map
	playerLock sync.Mutex
	// player messages into world
	playerChan chan *Player
	// load assets into player ( browser )
	loaderChan chan *Player

	// data structs from ipfs ( side load and check )
	ws *assets.WorldStore
}

func NewWorld(config *util.Config, cache *assets.Cache) *World {
	w := &World{}
	grid := make([][]*assets.SectorStore, Sectors)
	for i := range grid {
		grid[i] = make([]*assets.SectorStore, Sectors)
	}
	w.tiles = NewTileGrid()
	w.grid = grid
	w.players = make(map[*Player]bool)
	w.playerChan = make(chan *Player)
	w.loaderChan = make(chan *Player)
	w.status = NewGridStatus()
	//w.register = make(chan *connection)
	w.config = config
	w.cache = cache
	w.ref = config.Ref
	return w
}

func (w *World) Run() {
	err := w.Load()
	if err != nil {
		fmt.Println("world load fail bailing")
		return
	}
	ticker := time.NewTicker(time.Second * 60).C
	for {
		select {
		case <-ticker:
			{
				// update world here
				fmt.Println("world ticker ", time.Now())
				fmt.Println("# of players ", len(w.players))
			}
		case p := <-w.playerChan:
			{
				// new player arrives
				fmt.Println("arrrrg boink")
				w.SendTiles(p)
				p.OutMess <- []byte("this is a test")
			}
		case lc := <-w.loaderChan:
			{
				// send object references to the client
				//fmt.Println("send data to player to load sector")
				w.LoadSector(lc)
				//lc.OutMess <- []byte("load sector")
			}
		}
	}

}
