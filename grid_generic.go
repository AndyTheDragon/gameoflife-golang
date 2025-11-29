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

func (g *GenericGrid) CreateSpaceship(spaceshipType string, xOffset, yOffset int) {
	if (xOffset == -1 || yOffset == -1) {
		xOffset = rand.Int() % g.cols
		yOffset = rand.Int() % g.rows
	}
	switch spaceshipType {
	case "glider":
		g.Set(xOffset+0, yOffset+1, 1)
		g.Set(xOffset+1, yOffset+2, 1)
		g.Set(xOffset+2, yOffset+0, 1)
		g.Set(xOffset+2, yOffset+1, 1)
		g.Set(xOffset+2, yOffset+2, 1)
	case "lightweight_spaceship":
		g.Set(xOffset+0, yOffset+1, 1)
		g.Set(xOffset+0, yOffset+4, 1)
		g.Set(xOffset+1, yOffset+0, 1)
		g.Set(xOffset+2, yOffset+0, 1)
		g.Set(xOffset+3, yOffset+0, 1)
		g.Set(xOffset+3, yOffset+4, 1)
		g.Set(xOffset+4, yOffset+0, 1)
		g.Set(xOffset+4, yOffset+1, 1)
		g.Set(xOffset+4, yOffset+2, 1)
		g.Set(xOffset+4, yOffset+3, 1)
	}
}

func (g *GenericGrid) CopyFrom(other GridIface) {
	rows, cols := other.Size()
	if g.rows != rows || g.cols != cols {
		return
	}
	for x := 0; x < cols; x++ {
		for y := 0; y < rows; y++ {
			g.Set(x, y, other.Get(x, y))
		}
	}
}

func (g *GenericGrid) CopyTo(indexFor indexFunc) GridIface {
	newGrid := NewGenericGrid(g.rows, g.cols, indexFor)
	for x := 0; x < g.cols; x++ {
		for y := 0; y < g.rows; y++ {
			newGrid.Set(x, y, g.Get(x, y))
		}
	}
	// copy(newGrid.cells, g.cells)
	return newGrid
}

