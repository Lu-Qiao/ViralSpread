package main

import (
	"math/rand"
	"time"
)

// InitializeTissue
// Input
// Output
func InitializeTissue(width int) Board {
	var initialTissue Board
	initialTissue = make([][]Cell, width)
	for r := range initialTissue {
		initialTissue[r] = make([]Cell, width)
	}
	return initialTissue
}

// AssignStart
// input
// Output
func AssignStart(initialTissue Board, initialPositions []OrderedPair, conc float64) Board {
	for _, p := range initialPositions {
		initialTissue[p.x][p.y].state = "Infectious"
		initialTissue[p.x][p.y].concVirus = conc
	}
	return initialTissue
}

// RandomStart
// Input
// Output
func RandomStart(initialTissue Board, numPositions int, conc float64) Board {
	for i := 0; i < numPositions; i++ {
		time.Sleep(time.Millisecond)
		rand.Seed(time.Now().UnixNano())
		row := rand.Intn(len(initialTissue))
		time.Sleep(time.Millisecond)
		rand.Seed(time.Now().UnixNano())
		col := rand.Intn(len(initialTissue))

		initialTissue[row][col].state = "Infectious"
		initialTissue[row][col].concVirus = conc
	}
	return initialTissue
}
