package generate

import (
	"fmt"
	"testing"
)

func TestTiles(t *testing.T) {
	w := NewWorld(20, 60)
	fmt.Println(w)
	r := Rander{}
	w.Scan(r)
	fmt.Println(w)
}
