package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct{
	scale uint8
	width int
	height int
	updateInterval int
	backgroundColor color.Color
	liveCellColor color.Color
	grid GridIface
	buffer GridIface
	count int
	isPaused bool
}

func NewGame(scale uint8, width, height int, randomizeProbability float32) *Game {
	g := &Game{
		scale: scale,
		width: width,
		height: height,
		updateInterval: 20,
		backgroundColor: color.RGBA{102, 102, 102, 1},
		liveCellColor: color.RGBA{102, 187, 102, 1},
		grid: NewGrid(height, width),
		buffer: NewGrid(height, width),
	}
	g.grid.Randomize(randomizeProbability)
	return g
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.isPaused = !g.isPaused
	}
	
	if g.isPaused {
		return nil
	}
    g.count++
    if g.count == g.updateInterval {
        for x := 0; x < g.width; x++ {
            for y := 0; y < g.height; y++ {
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
    screen.Fill(g.backgroundColor)

    for x := 0; x < g.width; x++ {
        for y := 0; y < g.height; y++ {
			if g.grid.Get(x,y) == 1 {
				screen.Set(x, y, g.liveCellColor)
			}
		}
    }

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return g.width, g.height
}

func main() {
	log.Println("Starting Game of Life...")

	game := NewGame(8, 160, 120, 0.5)
	

	ebiten.SetWindowSize(game.width*int(game.scale), game.height*int(game.scale))
	ebiten.SetWindowTitle("Game of Life")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

