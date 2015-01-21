package main

import (
	"fmt"

	"github.com/beefsack/go-astar"
	"github.com/zignig/cohort/world/generate"
)

func main() {
	fmt.Println("tile world generator")
	w := generate.NewWorld(40, 80, generate.Empty)
	r := generate.Rander{0.49, generate.Water}
	w.Scan(r)
	c := generate.Caver{}
	w.ReScan(c, 4)

	fmt.Println(w)
}
