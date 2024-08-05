package main

import (
	"fmt"
	"time"

	"github.com/samber/lo"
)

/*
 1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
 2. Any live cell with two or three live neighbours lives on to the next generation.
 3. Any live cell with more than three live neighbours dies, as if by overpopulation.
 4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
*/

func main() {
	midpoint := 40 / 2
	world := NewWorld([]Cell{
		{midpoint, midpoint - 1},
		{midpoint, midpoint},
		{midpoint, midpoint + 1},
		{midpoint - 1, midpoint},
		{midpoint + 1, midpoint + 1},
	})

	for {

		fmt.Print("\033[H\033[2J")
		world.print()
		time.Sleep(200 * time.Millisecond)
		world.tick()

	}

}

type Cell struct {
	x, y int
}

type World struct {
	minX, maxX, minY, maxY int
	liveCells              []Cell
}

func NewWorld(seed []Cell) *World {
	return &World{
		minX:      0,
		maxX:      40,
		minY:      0,
		maxY:      40,
		liveCells: seed,
	}
}

func (w *World) isCellAlive(cell Cell) bool {
	_, found := lo.Find(w.liveCells, func(x Cell) bool {
		return x == cell
	})
	return found
}

func (w *World) print() {
	// TODO: update this to work when dimensions go into the negative
	fmt.Printf("  ")
	for col := w.minY; col < w.maxY; col++ {
		// fmt.Printf("%d", col)
	}
	fmt.Println()
	for row := w.minX; row < w.maxX; row++ {
		// fmt.Printf("%02d", row)
		for col := w.minY; col < w.maxY; col++ {
			var cell string
			if w.isCellAlive(Cell{row, col}) {
				cell = "😃"
			} else {
				cell = " "
			}

			fmt.Print(cell)
		}

		fmt.Println()
	}
}

func (w *World) countLiveNeighbors(cell Cell) int {
	x := cell.x
	y := cell.y
	// n
	ns := []Cell{
		{x, y - 1},
		{x - 1, y - 1},
		{x - 1, y},
		{x - 1, y + 1},
		{x, y + 1},
		{x + 1, y + 1},
		{x + 1, y},
		{x + 1, y - 1},
	}
	count := 0
	for _, n := range ns {
		if w.isCellAlive(n) {
			count++
		}
	}
	return count
}

func (w *World) tick() {
	// for each cell in the new dim, check against the rules to find if its alive or dead:
	newCells := []Cell{}
	for x := w.minX; x < w.maxX; x++ {
		for y := w.minY; y < w.maxY; y++ {
			liveNeighbourCount := w.countLiveNeighbors(Cell{x, y})
			if w.isCellAlive(Cell{x, y}) {
				if liveNeighbourCount == 2 || liveNeighbourCount == 3 {
					newCells = append(newCells, Cell{x, y}) // TODO do not duplicate
				}
			} else {
				if liveNeighbourCount == 3 {
					newCells = append(newCells, Cell{x, y})
				}
			}
		}

	}
	w.liveCells = newCells
}

/*
 1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
 2. Any live cell with two or three live neighbours lives on to the next generation.
 3. Any live cell with more than three live neighbours dies, as if by overpopulation.
 4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
*/
