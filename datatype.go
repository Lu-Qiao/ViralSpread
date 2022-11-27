package main

// cell
// state string
// conc of virus float
// position orderedPair
type Cell struct {
	// states: "Infectious", "Infected",
	// "dead", and "Uninfected"
	state     string
	concVirus float64
	position  orderedPair
}

// tissue
type Tissue [][]Cell

type orderedPair struct {
	x, y int
}

// Use Board to set the position of cells
type Board [][]*Cell

// orderedPair
// int
