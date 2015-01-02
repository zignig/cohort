package world

import (
	"errors"
	"fmt"
)

func (w *World) Load() (err error) {
	// load world here
	fmt.Println("Load World")
	baseRef := w.config.Ref
	st, err := w.cache.Resolve(baseRef)
	if err != nil {
		fmt.Println("World FAIL resolve")
		err = errors.New("World resolve Fail")
		return
	}
	data, err := w.cache.Cat(st + "/" + w.config.Path)
	if err != nil {

		fmt.Println("World doc resolve ", err)
		fmt.Println(data)
		err = errors.New("world doc fail")
		return
	}
	// save the base ref for the world
	w.ref = st
	w.ws, err = w.cache.LoadWorldStore(data)
	if err != nil {
		fmt.Println("World doc resolve")
		err = errors.New("world doc fail")
		return
	}
	return
}

func (w *World) LoadSector(p *Player) (err error) {
	// get the current sector of the player
	x := p.grx
	y := p.gry
	fmt.Println("Load Sector")
	fmt.Println("x :", x, " y :", y)
	// TODO , remove
	w.SendFloor(p, x, y)
	// has sector been loaded
	if w.status.grid[x][y] == false {
		fmt.Println("Bounce Sector")
		s := w.ws.Grid[x][y]
		fmt.Println("Sector info ", s)
		st, err2 := w.cache.Resolve(s.Ips)
		if err2 != nil {
			fmt.Println("Sector FAIL resolve")
			err = errors.New("Sector resolve Fail")
			return
		}
		data, err2 := w.cache.Cat(st + "/" + s.Path)
		if err2 != nil {
			fmt.Println("Sector load fail")
			err = errors.New("Sector load fail")
			return
		}
		fmt.Println(string(data))
		sectorData, err2 := w.cache.LoadSectorStore(data)
		sectorData.Ref = st
		fmt.Println(sectorData)
		w.grid[x][y] = sectorData
		if err2 != nil {
			fmt.Println("Sector convert fail")
			err = errors.New("Sector convert fail")
			return
		}
		w.status.grid[x][y] = true
	}
	fmt.Println(w.grid[x][y])
	// Pump data to the  player client
	p.SendSector(w.grid[x][y], x, y)
	return
}

func (w *World) SendFloor(p *Player, x int, y int) {
	offx := float64(x * SectorSize)
	offy := float64(y * SectorSize)
	// send a floor builder
	fl := &FloorMessage{}
	fl.Pos.X = offx
	fl.Pos.Z = offy
	fl.Size = SectorSize
	data, err := Encode(fl)
	if err != nil {
		fmt.Println("floor fail ", err)
		return
	}
	fmt.Println(string(data))
	p.OutMess <- data
}
