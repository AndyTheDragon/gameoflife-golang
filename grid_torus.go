package main

import "math/rand"

type TorusGrid struct {
	rows int
	cols int
	cells []uint8
}

func NewTorusGrid(rows, cols int) *TorusGrid {
	return &TorusGrid{
		rows: rows,
		cols: cols,
		cells: make([]uint8, rows*cols),
	}
}

func (g *TorusGrid) Size() (int, int) {
	return g.rows, g.cols
}

func (g *TorusGrid) IndexFor(x, y int) int {
	x = (x + g.cols) % g.cols
	y = (y + g.rows) % g.rows
	return y*g.cols + x
}

func (g *TorusGrid) Get(x, y int) uint8 {
	index := g.IndexFor(x, y)
	if (index == -1) {
		return 0
	}
	return g.cells[index]
}

func (g *TorusGrid) Set(x, y int, value uint8) {
	index := g.IndexFor(x, y)
	if (index == -1) {
		return
	}
	g.cells[index] = value
}

func (g *TorusGrid) SumOfNeighbors(x, y int) uint8 {
	return g.Get(x-1, y-1) + g.Get(x, y-1) + g.Get(x+1, y-1) +
		   g.Get(x-1, y  )                 + g.Get(x+1, y  ) +
		   g.Get(x-1, y+1) + g.Get(x, y+1) + g.Get(x+1, y+1)
}

func (g *TorusGrid) Clear() {
	for i := range g.cells {
		g.cells[i] = 0
	}
}

func (g *TorusGrid) Randomize(probability float32) {
	for i := range g.cells {
		if rand.Float32() < probability {
			g.cells[i] = 1
		} else {
			g.cells[i] = 0
		}
	}
}