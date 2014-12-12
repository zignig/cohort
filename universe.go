package main

import (
	"fmt"

	"github.com/zignig/viewer/assets"
	"github.com/zignig/viewer/util"
	"github.com/zignig/viewer/world"
)

type universe struct {
	conf  *util.Config
	world *world.World
	h     *hub
	cache *assets.Cache
}

func AndLetThereBeLight(config *util.Config) *universe {
	fmt.Println("FATOOOOMPSH")
	u := &universe{}
	u.conf = config
	u.cache = assets.NewCache()
	u.world = world.NewWorld(config, u.cache)
	return u
}

func (u *universe) String() (s string) {
	return "REALLY BIG"
}
