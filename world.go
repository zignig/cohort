package main

type point struct {
	x int32
	y int32
}

type player struct {
	name  string
	pos   point
	alive bool
}

type entity struct {
	ref  string
	data []byte
	pos  point
}

type sector struct {
	ref   string
	owner string
	ents  []*entity
}
type world struct {
	players []player
	grid    [][]sector
}
