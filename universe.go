package main

import (
	"fmt"
	"html/template"

	"github.com/zignig/viewer/assets"
	"github.com/zignig/viewer/util"
	"github.com/zignig/viewer/world"
)

type universe struct {
	conf  *util.Config
	world *world.World
	h     *hub
	cache *assets.Cache
	templ *template.Template
}

func AndLetThereBeLight(config *util.Config) *universe {
	fmt.Println("FATOOOOMPSH")
	u := &universe{}
	u.conf = config
	u.cache = assets.NewCache()
	u.world = world.NewWorld(config, u.cache)
	u.h = NewHub(u.world)
	return u
}

func (u *universe) run() {
	go u.world.Run()
	go u.h.run()
}

func (u *universe) String() (s string) {
	return "REALLY BIG"
}
