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
		backgroundColor: color.RGBA{102, 102, 102, 255},
		liveCellColor: color.RGBA{102, 187, 102, 255},
		grid: NewTorusGrid(height, width),
		buffer: NewTorusGrid(height, width),
	}
	g.grid.Randomize(randomizeProbability)
	g.grid.CreateSpaceship("glider", -1, -1)
	g.grid.CreateSpaceship("lightweight_spaceship", -1, -1)
	return g
}

func (g *Game) setTopology(indexFunc indexFunc) {
	g.grid = g.grid.CopyTo(indexFunc)
	g.buffer = g.buffer.CopyTo(indexFunc)
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.isPaused = !g.isPaused
	}

	// topology switching
    if inpututil.IsKeyJustPressed(ebiten.KeyDigit1) {
		log.Println("Switching to Plane topology...")
        g.setTopology(IndexPlane)     // plane
    }
    if inpututil.IsKeyJustPressed(ebiten.KeyDigit2) {
		log.Println("Switching to Torus topology...")
        g.setTopology(IndexTorus)     // torus
    }
    if inpututil.IsKeyJustPressed(ebiten.KeyDigit3) {
		log.Println("Switching to Sphere topology...")
        g.setTopology(IndexSphere)    // “sphere”
    }
    if inpututil.IsKeyJustPressed(ebiten.KeyDigit4) {
		log.Println("Switching to Cylinder topology...")
        g.setTopology(IndexCylinder)  // cylinder
    }
    if inpututil.IsKeyJustPressed(ebiten.KeyDigit5) {
		log.Println("Switching to Klein bottle topology...")
        g.setTopology(IndexKlein)     // Klein bottle
    }
	if inpututil.IsKeyJustPressed(ebiten.KeyDigit6) {
		log.Println("Switching to Möbius strip topology...")
		g.setTopology(IndexMoebiusX)  // Möbius strip (horizontal twist)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDigit7) {
		log.Println("Switching to Reflective topology...")
		g.setTopology(IndexReflect)    
	}

	// launch spaceships
	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		log.Println("Launching glider spaceship...")
		g.grid.CreateSpaceship("glider", -1, -1)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		log.Println("Launching lightweight spaceship...")
		g.grid.CreateSpaceship("lightweight_spaceship", -1, -1)
	}

	// clear grid
	 if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		log.Println("Clearing grid...")
		g.grid.Clear()
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

