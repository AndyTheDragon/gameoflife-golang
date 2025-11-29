package main

import (
	"fmt"
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
var grid [width][height]uint8 = [width][height]uint8{}
var buffer [width][height]uint8 = [width][height]uint8{}
var count int = 0
var isPaused = false

type Game struct{}

func (g *Game) Update() error {
	if (inpututil.IsKeyJustPressed(ebiten.KeySpace)) {
		isPaused = !isPaused
	}
	
	if isPaused {
		return nil
	}
    count++
    if count == 20 {
        // same logic as your old update()
        for x := 1; x < width-1; x++ {
            for y := 1; y < height-1; y++ {
                buffer[x][y] = 0
                neighbours := grid[x-1][y-1] + grid[x][y-1] + grid[x+1][y-1] +
                    grid[x-1][y] + 0 + grid[x+1][y] +
                    grid[x-1][y+1] + grid[x][y+1] + grid[x+1][y+1]

                if grid[x][y] == 0 && neighbours == 3 {
                    buffer[x][y] = 1
                } else if neighbours < 2 || neighbours > 3 {
                    buffer[x][y] = 0
                } else {
                    buffer[x][y] = grid[x][y]
                }
            }
        }
        grid, buffer = buffer, grid
        count = 0
    }
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    screen.Fill(backgroundColor)

    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            for i := 0; i < scale; i++ {
                for j := 0; j < scale; j++ {
                    if grid[x][y] == 1 {
                        screen.Set(x*scale+i, y*scale+j, liveCellColor)
                    }
                }
            }
        }
    }
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return width*scale, height*scale
}

func main() {
	fmt.Println("Starting Game of Life...")

	// random initial state
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			if rand.Float32() < 0.5 {
				grid[x][y] = 1
			}
		}
	}

	ebiten.SetWindowSize(width*scale, height*scale)
	ebiten.SetWindowTitle("Game of Life")

	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

