package main

import "math/rand"

type indexFunc func(x, y, rows, cols int) int

type GenericGrid struct {
	rows int
	cols int
	cells []uint8
	indexFor indexFunc
}

func NewGenericGrid( rows, cols int, indexFor indexFunc) *GenericGrid {
	return &GenericGrid{
		rows: rows,
		cols: cols,
		cells: make([]uint8, rows*cols),
		indexFor: indexFor,
	}
}

func (g *GenericGrid) Size() (int, int) {
	return g.rows, g.cols
}

func (g *GenericGrid) Get(x, y int) uint8 {
	index := g.indexFor(x, y, g.rows, g.cols)
	if (index == -1) {
		return 0
	}
	return g.cells[index]
}

func (g *GenericGrid) Set(x, y int, value uint8) {
	index := g.indexFor(x, y, g.rows, g.cols)
	if (index == -1) {
		return
	}
	g.cells[index] = value
}

func (g *GenericGrid) SumOfNeighbors(x, y int) uint8 {
	return g.Get(x-1, y-1) + g.Get(x, y-1) + g.Get(x+1, y-1) +
		   g.Get(x-1, y  )                 + g.Get(x+1, y  ) +
		   g.Get(x-1, y+1) + g.Get(x, y+1) + g.Get(x+1, y+1)
}

func (g *GenericGrid) Clear() {
	for i := range g.cells {
		g.cells[i] = 0
	}
}

func (g *GenericGrid) Randomize(probability float32) {
	for i := range g.cells {
		if rand.Float32() < probability {
			g.cells[i] = 1
		} else {
			g.cells[i] = 0
		}
	}
}

