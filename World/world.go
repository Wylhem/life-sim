package world

import (
	"encoding/json"
	"life/cells"
	"os"
)

type World struct {
	Area [][]cells.Cells
}

func NewWorld(width, height int, maxCells int) *World {
	area := make([][]cells.Cells, width)
	for i := range area {
		area[i] = make([]cells.Cells, height)
	}
	world := &World{Area: area}

	return world
}

func Reset(w *World) *World {
	for i := 0; i < len(w.Area); i++ {
		for j := 0; j < len(w.Area[i]); j++ {
			w.Area[i][j].Alive = false
		}
	}
	return w
}

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

func (w *World) SaveFile(filename string) error {
	data, err := json.Marshal(w)
	if err != nil {
		return (err)
	}
	return os.WriteFile(filename, data, 0644)
}

func (w *World) LoadFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &w)
}
