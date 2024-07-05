package world

import (
	"encoding/json"
	"life/cells"
	"os"
)

// World represents the game world.
type World struct {
	Area [][]cells.Cells
}

// NewWorld creates a new instance of World with the specified width, height, and maximum number of cells.
// width : int : The width of the world.
// height : int : The height of the world.

// Returns : *World : The new instance of World.
func NewWorld(width, height int) *World {
	area := make([][]cells.Cells, width)
	for i := range area {
		area[i] = make([]cells.Cells, height)
	}
	world := &World{Area: area}

	return world
}

// Reset sets all cells in the world to be dead.
// w : *World : The world to reset.
// Returns : *World : The world with all cells set to dead.
func Reset(w *World) *World {
	for i := 0; i < len(w.Area); i++ {
		for j := 0; j < len(w.Area[i]); j++ {
			w.Area[i][j].Alive = false
		}
	}
	return w
}

// Update updates the state of the world based on the rules of the game of life.
// w : *World : The world to update.
func (w *World) Update() {
	next := make([][]cells.Cells, len(w.Area))
	for i := range next {
		next[i] = make([]cells.Cells, len(w.Area[0]))
		for j := range next[i] {
			next[i][j] = w.Area[i][j]
		}
	}

	for i := 0; i < len(w.Area); i++ {
		for j := 0; j < len(w.Area[i]); j++ {
			liveNeighbours := w.liveNeighbours(i, j)
			if w.Area[i][j].Alive {
				if liveNeighbours < 2 || liveNeighbours > 3 {
					next[i][j].Alive = false
				} else {
					next[i][j].Alive = true
				}
			} else {
				if liveNeighbours == 3 {
					next[i][j].Alive = true
				}
			}
		}
	}
	w.Area = next
}

// liveNeighbours returns the number of live neighbors for the cell at the specified coordinates.
// x : int : The x-coordinate of the cell.
// y : int : The y-coordinate of the cell.
// Returns : int : The number of live neighbors for the cell.
func (w *World) liveNeighbours(x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			ni := x + i
			nj := y + j
			if ni >= 0 && ni < len(w.Area) && nj >= 0 && nj < len(w.Area[0]) && w.Area[ni][nj].Alive {
				count++
			}
		}
	}
	return count
}

// LiveCellCount returns the number of live cells in the world.
// Returns : int : The number of live cells in the world.
func (w *World) LiveCellCount() int {
	count := 0
	for i := 0; i < len(w.Area); i++ {
		for j := 0; j < len(w.Area[i]); j++ {
			if w.Area[i][j].Alive {
				count++
			}
		}
	}
	return count
}

// SaveFile saves the world to a file with the specified filename.
// filename : string : The name of the file to save the world to.
// Returns : error : An error if the save operation fails.
func (w *World) SaveFile(filename string) error {
	data, err := json.Marshal(w)
	if err != nil {
		return (err)
	}
	return os.WriteFile(filename, data, 0644)
}

// LoadFile loads the world from a file with the specified filename.
// filename : string : The name of the file to load the world from.
// Returns : error : An error if the load operation fails.
func (w *World) LoadFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &w)
}
