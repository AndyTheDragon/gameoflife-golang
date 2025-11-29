package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const scale = 8
const width = 160
const height = 120

var backgroundColor color.Color = color.RGBA{102, 102, 102, 1}
var liveCellColor color.Color = color.RGBA{102, 187, 102, 1}

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

type Game struct{
	grid *Grid
	buffer *Grid
	count int
	isPaused bool
}

func (g *Game) Update() error {
	if (inpututil.IsKeyJustPressed(ebiten.KeySpace)) {
		g.isPaused = !g.isPaused
	}
	
	if g.isPaused {
		return nil
	}
    g.count++
    if g.count == 20 {
        for x := 0; x < width; x++ {
            for y := 0; y < height; y++ {
                g.buffer.Set(x,y, 0)
                neighbours := g.grid.SumOfNeighbors(x, y)

                if g.grid.Get(x, y) == 0 && neighbours == 3 {
                    g.buffer.Set(x, y, 1)
                } else if neighbours < 2 || neighbours > 3 {
                    g.buffer.Set(x, y, 0)
                } else {
                    g.buffer.Set(x, y, g.grid.Get(x, y))
                }
            }
        }
        g.grid, g.buffer = g.buffer, g.grid
        g.count = 0
    }
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    screen.Fill(backgroundColor)

    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
			if g.grid.Get(x,y) == 1 {
				screen.Set(x, y, liveCellColor)
			}
		}
    }

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return width, height
}

func main() {
	log.Println("Starting Game of Life...")

	ebiten.SetWindowSize(width*scale, height*scale)
	ebiten.SetWindowTitle("Game of Life")

	game := &Game{
		grid: NewGrid(height, width),
		buffer: NewGrid(height, width),
	}
	// random initial state
	game.grid.Randomize(0.5)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

