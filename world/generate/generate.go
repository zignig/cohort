package generate

import "math/rand"

// procedural world builder
// structures borrowed from
// github.com/beefsack/go-astar

type Tile struct {
	Kind int
	X, Y int
	W    World
}

const (
	Empty = iota
	Grass
	Sand
	Water
	Road
	Rail
	Blocker
)

var TypeRunes = map[int]rune{
	Empty:   ' ',
	Grass:   '.',
	Sand:    'S',
	Water:   '#',
	Road:    'A',
	Rail:    'R',
	Blocker: '☒',
}

type World map[int]map[int]*Tile

// SetTile sets a tile at the given coordinates in the world.
func (w World) SetTile(t *Tile, x, y int) {
	if w[x] == nil {
		w[x] = map[int]*Tile{}
	}
	w[x][y] = t
	t.X = x
	t.Y = y
	t.W = w
}

// Tile gets the tile at the given coordinates in the world.
func (w World) Tile(x, y int) *Tile {
	if w[x] == nil {
		return nil
	}
	return w[x][y]
}

func NewWorld(x, y int, toType int) *World {
	// reed rand on each new world
	//rand.Seed(time.Now().UTC().UnixNano())
	w := &World{}
	for i := 0; i <= x; i++ {
		for j := 0; j <= y; j++ {
			w.SetTile(&Tile{Kind: toType}, i, j)
		}
	}
	return w
}

type Modifier interface {
	Mod(*Tile)
}

// run a function on each tile in the world
func (w World) Scan(m Modifier) {
	width := len(w)
	height := len(w[0])
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			t := w.Tile(i, j)
			m.Mod(t)
		}
	}

}

// run scan multiple times
func (w World) ReScan(m Modifier, repeat int) {
	for i := 0; i <= repeat; i++ {
		w.Scan(m)
	}
}

// print out world rep
func (w World) String() (s string) {
	width := len(w)
	height := len(w[0])
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			t := w.Tile(i, j)
			s = s + string(TypeRunes[t.Kind])
		}
		s = s + "\r\n"
	}

	return s
}

// random modifier
type Rander struct {
	prob   float64
	toType int
}

func (r Rander) Mod(t *Tile) {
	f := rand.Float64()
	if f < r.prob {
		t.Kind = r.toType
	}
}
