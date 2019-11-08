package life

import (
	"log"
	"math/rand"
	"strconv"
	"time"
)

// World is our representation of the board or grid of Cells
type World struct {
	emptyCell Cell
	param     Param
	step      int64
	grid      map[string]*Cell
}

var neighboursAt = [8][2]int{
	{-1, 1}, {0, 1}, {1, 1}, // above
	{-1, 0} /* self */, {1, 0}, // beside
	{-1, -1}, {0, -1}, {1, -1}, // below
}

// Init initialises the World board or grid
func (w *World) Init(board string, rules string, generateRandom bool) {
	log.Printf("Initialising World on '%s' board with '%s' rules", board, rules)
	w.emptyCell = Cell{false, false}
	w.param.Init(board, rules)
	w.step = 0
	w.grid = make(map[string]*Cell)

	if generateRandom {
		w.randomiseWorld()
	} else {
		w.predefinedWorld("../test_world.gol")
	}
}

func (w *World) randomiseWorld() {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	var x, y int
	for x = 1; x <= w.param.Width; x++ {
		for y = 1; y <= w.param.Height; y++ {
			isAlive := rand.Intn(100) < w.param.SpawnPercent
			w.addCell(x, y, isAlive)
		}
	}
}

func (w *World) predefinedWorld(fileName string) {
	var x, y int
	for x = 1; x <= w.param.Width; x++ {
		for y = 1; y <= w.param.Height; y++ {
			isAlive := false
			if x == 5 && (y == 4 || y == 5 || y == 6) {
				isAlive = true
			}
			w.addCell(x, y, isAlive)
		}
	}
}

func (w *World) addCell(x, y int, isAlive bool) {
	var cell Cell
	cell.Init(isAlive)
	w.grid[w.gridReference(x, y)] = &cell
}

func (w *World) gridReference(x, y int) string {
	x = w.wrapCoord(x, 1, w.param.Width, w.param.WrapX)
	y = w.wrapCoord(y, 1, w.param.Height, w.param.WrapY)
	return strconv.Itoa(x) + "|" + strconv.Itoa(y)
}

func (w *World) wrapCoord(val, min, max int, wrapEnabled bool) int {
	if val < min {
		if wrapEnabled {
			val = max - val
		} else {
			val = min - 1
		}
	} else if val > max {
		if wrapEnabled {
			val = val - max
		} else {
			val = max + 1
		}
	}
	return val
}

// Render generates a text display of the World state
func (w *World) Render(htmlise bool) string {
	var x, y int
	output := ""
	for y = 1; y <= w.param.Height; y++ {
		for x = 1; x <= w.param.Width; x++ {
			cell := w.cellAt(x, y)
			output += cell.String() // String() auto-stringification not recognised in '+' concatenation
		}
		if htmlise {
			output += "<br />"
		}
		output += "\n"
	}
	return output
}

func (w *World) cellAt(x, y int) *Cell {
	gridRef := w.gridReference(x, y)
	if cell, ok := w.grid[gridRef]; ok {
		return cell
	}

	return &w.emptyCell
}

// Step returns the count of iterations the World has gone through
func (w *World) Step() int64 {
	return w.step
}

func (w *World) countNeighbours(x, y int) int {
	totalNeighbours := 0
	for _, dir := range neighboursAt {
		cell := w.cellAt(x+dir[0], y+dir[1])
		if cell.IsAlive() {
			totalNeighbours++
		}
	}
	return totalNeighbours
}

// Calculate performs the next iteration of the simulation across all Cells in the World
func (w *World) Calculate() {
	for y := 1; y <= w.param.Height; y++ {
		for x := 1; x <= w.param.Width; x++ {
			cell := w.cellAt(x, y)
			nCount := w.countNeighbours(x, y)
			if cell.IsAlive() {
				willSurvive := w.param.RuleValues["s"][strconv.Itoa(nCount)]
				if !willSurvive {
					cell.UpdateState(false)
				}
			} else {
				willSpawn := w.param.RuleValues["b"][strconv.Itoa(nCount)]
				if willSpawn {
					cell.UpdateState(true)
				}
			}
		}
	}

	for y := 1; y <= w.param.Height; y++ {
		for x := 1; x <= w.param.Width; x++ {
			cell := w.cellAt(x, y)
			cell.Refresh()
		}
	}
	w.step++
}
