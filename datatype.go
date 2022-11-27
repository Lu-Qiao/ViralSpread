package main

// cell
// state string
// conc of virus float
type Cell struct {
	// states: "Infectious", "Infected",
	// "dead", and "Uninfected"
	state     string
	concVirus float64
}

// Board
type Board [][]Cell
