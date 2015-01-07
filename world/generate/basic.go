package generate

import (
	"fmt"
)

// make a simple swamp land and retrun it to the server
func Basic(s int) []byte {
	w := NewWorld(s, s, Empty)
	r := Rander{0.42, Water}
	w.Scan(r)
	c := Caver{}
	w.ReScan(c, 2)
	//r = Rander{0.4, Grass}
	//w.Scan(r)
	fmt.Println(w)
	return w.Export()

}
