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
	Grass = iota
	Sand
	Water
	Road
	Rail
	Blocker
)

var TypeRunes = map[int]rune{
	Grass:   '.',
	Sand:    'S',
	Water:   'W',
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

func NewWorld(x, y int) *World {
	w := &World{}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			w.SetTile(&Tile{Kind: Blocker}, i, j)
		}
	}
	return w
}

type Modifier interface {
	Mod(*Tile)
}

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

// empty modifier
type Rander struct{}

func (r Rander) Mod(t *Tile) {
	t.Kind = rand.Intn(len(TypeRunes))
}
