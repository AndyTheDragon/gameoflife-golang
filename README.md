# Conway's Game of Life implemented in Go
Inspired by an exercise at school, i've made a version of Game of Life (https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) in Go.

## Install instructions
* Make sure Go is installed on your computer (https://go.dev/doc/install)
* Clone or copy repository to you computer
* Run `go run .` in the directory

## Mixing it up
* The NewGame function takes 4 parameters: scale, width, height and probability.
* The scale adjusts the size of each cell, 1 is very small, 8 is big.
* The width is the number of cells in each row.
* The height is the number of cells in each column.
* The probability is the probability that each cell is alive when starting the simulation.
