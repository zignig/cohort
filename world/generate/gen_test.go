package generate

import (
	"fmt"
	"testing"
)

func TestTiles(t *testing.T) {
	w := NewWorld(40, 100, Empty)
	r := Rander{0.29, Water}
	w.Scan(r)
	c := Caver{}
	w.ReScan(c, 1)
	w.Scan(r)
	w.ReScan(c, 2)
	//r = Rander{0.4, Grass}
	//w.Scan(r)
	fmt.Println(w)

}
