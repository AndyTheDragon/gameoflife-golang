package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"github.com/hajimehoshi/ebiten"

)

const scale = 2
const width = 400
const height = 300

var blue color.Color = color.RGBA{69, 145, 196, 255}
var yellow color.Color = color.RGBA{255, 230, 120, 255}
var grid [width][height]uint8 = [width][height]uint8{}
var buffer [width][height]uint8 = [width][height]uint8{}
var count int = 0

func main() {
	fmt.Println("Starting Game of Life...")
	// Game initialization and loop would go here
	for x:=1; x<width-1; x++ {
		for y:=1; y<height-1; y++ {
			if rand.Float32() < 0.5 {
				grid[x][y] = 1
			}
		}
	}
	if err := ebiten.Run(frame, width, height, 2, "Game of Life"); err != nil {
		log.Fatal(err)
	}
}

func update() error {
	for x:=1; x<width-1; x++ {
		for y:=1; y<height-1; y++ {
			buffer[x][y] = 0
			neighbours := grid[x-1][y-1] + grid[x][y-1] + grid[x+1][y-1] +
			              grid[x-1][y]   +      0       + grid[x+1][y] +
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
	return nil
}

func display(window *ebiten.Image) {
	window.Fill(blue)

	for x:=0; x<width; x++ {
		for y:=0; y<height; y++ {
			for i:=0; i<scale; i++ {
				for j:=0; j<scale; j++ {
					if grid[x][y] == 1 {
						window.Set(x*scale+i, y*scale+j, yellow)
					}
				}
			}
		}
	}
}

func frame(window *ebiten.Image) error {
	count++
	var err error = nil
	if count == 20 {
		err = update()
		count = 0
	}
	if !ebiten.IsDrawingSkipped() {
		display(window)
	}
	return err
}