package main

// cell
// state string
// conc of virus float
type Cell struct {
	// states: "Infectious", "Infected", "dead", and "Uninfected"
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
	// parameters for treatments
	treatment                 string
	epsilonCell, epsilonVirus float64
}

// OrderedPair
type OrderedPair struct {
	x, y int
}

// All inputs from WebApp
type Inputs struct {
	width, numInfectious, numGens, imageFrequency int
	initialPosition                               OrderedPair
	mode                                          string
	timeSteps                                     float64
	parameters                                    Parameters
}

// Output for multiprocs
type Output struct {
	board      Board
	cellNumber []int
	index      int
}
