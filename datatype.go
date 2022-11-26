package main

// cell
// state string
// conc of virus float
// position orderedPair
type Cell struct {
	state     string
	concVirus float64
	position  orderedPair
}

// tissue
// width
type Tissue struct {
	width int
	board Board
}

type orderedPair struct {
	x, y int
}

// Use Board to set the position of cells
type Board [][]*Cell

// orderedPair
// int
