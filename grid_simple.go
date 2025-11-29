package main

import "math/rand"

type Grid struct {
	rows int
	cols int
	cells []uint8
}

func NewGrid( rows, cols int) *Grid {
	return &Grid{
		rows: rows,
		cols: cols,
		cells: make([]uint8, rows*cols),
	}
}

func (g *Grid) Size() (int, int) {
	return g.rows, g.cols
}

func (g *Grid) IndexFor(x, y int) int {
	if (x < 0 || x > g.cols-1 || y < 0 || y > g.rows-1) {
		return -1
	}
	return y*g.cols + x
}

func (g *Grid) Get(x, y int) uint8 {
	index := g.IndexFor(x, y)
	if (index == -1) {
		return 0
	}
	return g.cells[index]
}

func (g *Grid) Set(x, y int, value uint8) {
	index := g.IndexFor(x, y)
	if (index == -1) {
		return
	}
	g.cells[index] = value
}

func (g *Grid) SumOfNeighbors(x, y int) uint8 {
	return g.Get(x-1, y-1) + g.Get(x, y-1) + g.Get(x+1, y-1) +
		   g.Get(x-1, y  )                 + g.Get(x+1, y  ) +
		   g.Get(x-1, y+1) + g.Get(x, y+1) + g.Get(x+1, y+1)
}

func (g *Grid) Clear() {
	for i := range g.cells {
		g.cells[i] = 0
	}
}

func (g *Grid) Randomize(probability float32) {
	for i := range g.cells {
		if rand.Float32() < probability {
			g.cells[i] = 1
		} else {
			g.cells[i] = 0
		}
	}
}