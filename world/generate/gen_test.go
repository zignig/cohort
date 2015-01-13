package generate

import (
	"fmt"
	"testing"
)

func TestTiles(t *testing.T) {
	w := NewWorld(64, 64, Empty)
	r := Rander{0.37, Water}
	w.Scan(r)
	c := Caver{}
	w.ReScan(c, 6)
	//r = Rander{0.4, Grass}
	//w.Scan(r)
	fmt.Println(w)
	//fmt.Println(string(w.Export()))
	//fmt.Println(Basic(20))

}
