package main

import (
	"fmt"

	"github.com/zignig/viewer/world/generate"
)

func main() {
	fmt.Println("tile world generator")
	w := generate.NewWorld(50, 150, generate.Empty)
	r := generate.Rander{0.42, generate.Water}
	w.Scan(r)
	c := generate.Caver{}
	w.ReScan(c, 2)

	fmt.Println(w)
}
