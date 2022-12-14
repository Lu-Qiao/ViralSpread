package main

import (
	"math/rand"
	"time"
)

// InitializeTissue
// Input
// Output
func InitializeTissue(width int) Board {
	initialTissue := make([][]Cell, width)
	for r := range initialTissue {
		initialTissue[r] = make([]Cell, width)
	}
	for row := range initialTissue {
		for col := range initialTissue[0] {
			initialTissue[row][col].state = "Uninfected"
		}
	}
	return initialTissue
}

// AssignStart
// input
func AssignStart(initialTissue Board, initialPosition OrderedPair, conc float64) {
	initialTissue[initialPosition.x][initialPosition.y].state = "Infectious"
	initialTissue[initialPosition.x][initialPosition.y].concVirus = conc
	initialTissue[initialPosition.x+1][initialPosition.y].state = "Infectious"
	initialTissue[initialPosition.x+1][initialPosition.y].concVirus = conc
}

// RandomStart
// Input
// Output
func RandomCoStart(initialTissue Board, numPositions int, conc float64) {
	for i := 0; i < numPositions; i++ {
		time.Sleep(time.Millisecond)
		rand.Seed(time.Now().UnixNano())
		row := rand.Intn(len(initialTissue))
		time.Sleep(time.Millisecond)
		rand.Seed(time.Now().UnixNano())
		col := rand.Intn(len(initialTissue))
		rand.Seed(time.Now().UnixNano())
		state := rand.Intn(10)
		if state <= 5 {
			initialTissue[row][col].state = "Infectious1"
		} else {
			initialTissue[row][col].state = "Infectious2"
		}
		initialTissue[row][col].concVirus = conc
	}
}

func AssignCoStart(initialTissue Board, initialPosition OrderedPair, conc float64) {
	initialTissue[initialPosition.x][initialPosition.y-5].state = "Infectious1"
	initialTissue[initialPosition.x][initialPosition.y-5].concVirus = conc
	initialTissue[initialPosition.x+1][initialPosition.y-5].state = "Infectious1"
	initialTissue[initialPosition.x+1][initialPosition.y-5].concVirus = conc
	initialTissue[initialPosition.x][initialPosition.y+5].state = "Infectious2"
	initialTissue[initialPosition.x][initialPosition.y+5].concVirus = conc
	initialTissue[initialPosition.x+1][initialPosition.y+5].state = "Infectious2"
	initialTissue[initialPosition.x+1][initialPosition.y+5].concVirus = conc
}

// RandomStart
// Input
// Output
func RandomStart(initialTissue Board, numPositions int, conc float64) {
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
}
