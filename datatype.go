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

// Parameters
type Parameters struct {
	// parameters for cells
	lambda, omega, dT, delta, threshold float64
	// parameters for virus
	rCap, alpha, gamma, rho float64
}

type OrderedPair struct {
	x, y int
}
