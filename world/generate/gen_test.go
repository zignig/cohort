package generate

import (
	"fmt"
	"testing"
)

func TestTiles(t *testing.T) {
	w := NewWorld(32, 32, Empty)
	r := Rander{0.37, Water}
	w.Scan(r)
	//r = Rander{0.1, Sand}
	//w.Scan(r)
	c := Caver{}
	w.ReScan(c, 40)
	fmt.Println(w)

}
