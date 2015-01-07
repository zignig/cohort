package generate

// empty modifier
type Caver struct{}

func (r Caver) Mod(t *Tile) {
	c := t.NeighbourCount()
	if c < 4 {
		t.Kind = Empty
	}
	if c > 5 {
		t.Kind = Water
	}
}

func (t *Tile) NeighbourCount() int {
	var n [3]int
	n[0] = -1
	n[1] = 0
	n[2] = 1
	count := 0
	for i := range n {
		for j := range n {
			//fmt.Println(n[i], n[j], t.X, t.Y)
			neig := t.W.Tile(t.X+n[i], t.Y+n[j])
			if neig != nil {
				//fmt.Print(neig.Kind)
				if neig.Kind != Empty {
					count = count + 1
				}
			}
		}
	}
	return count
}
