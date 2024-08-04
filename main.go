package main

import (
	"fmt"

	"github.com/samber/lo"
)

/*
 1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
 2. Any live cell with two or three live neighbours lives on to the next generation.
 3. Any live cell with more than three live neighbours dies, as if by overpopulation.
 4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

|-----|
|x| | |
|x|x| |
| | | |
|-----|

*/

func main() {
	world := NewWorld([]Point{
		{1, 3},
		{3, 3},
		{3, 2},
		{6, 2},
		{6, 6},
	})
	world.print()
}

type Point struct {
	x, y int
}

type World struct {
	liveCells []Point
}

func NewWorld(seed []Point) *World {
	return &World{
		liveCells: seed,
	}
}

func (w *World) isCellAlive(cell Point) bool {
	_, found := lo.Find(w.liveCells, func(x Point) bool {
		return x == cell
	})
	return found
}

func (w *World) print() {
	// TODO: update this to work when dimensions go into the negative
	rows, cols := w.getDimensions()
	fmt.Printf("rows: %d ; cols: %d\n", rows, cols)
	fmt.Printf("  ")
	for col := 1; col < cols+1; col++ {
		fmt.Printf("%d", col)
	}
	fmt.Println()
	for row := 1; row < rows+1; row++ {
		fmt.Printf("%02d", row)
		for col := 1; col < cols+1; col++ {
			var cell string
			if w.isCellAlive(Point{row, col}) {
				cell = "■"
			} else {
				cell = " "
			}

			fmt.Print(cell)
		}

		fmt.Println()
	}
}

func (w *World) getDimensions() (rows, cols int) {
	maxX := lo.MaxBy(w.liveCells, func(cell Point, max Point) bool {
		return cell.x > max.x
	})
	maxY := lo.MaxBy(w.liveCells, func(cell Point, max Point) bool {
		return cell.y > max.y
	})
	rows = maxY.y
	cols = maxX.x
	return
}

func (w *World) tick() {
	// get the dimensions
	// new dims are old+1 in each direction
	// for each cell in the new dim, check against the rules to find if its alive or dead:
	// check if alive - one ruleset; if dead, another
	// discard old world, put in the new one.
}
