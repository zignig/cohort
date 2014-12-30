package world

//import (
//	"fmt"
//)

// sector logic
// currently 2 dimensional array , consider octree

// find the current sector of a player
func (pos *V3) Sector() (x int, y int) {
	//fmt.Println(pos.X, pos.Y, pos.Z)
	secx := (pos.X + (SectorSize / 2)) / SectorSize
	secz := (pos.Z + (SectorSize / 2)) / SectorSize
	//fmt.Println("into => [", secx, ",", secz, "]")
	return int(secx), int(secz)
}
