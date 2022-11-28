package main

import (
	"math"
	"math/rand"
	"time"
)

// SimulateViralSpread simulates the viral spread system over numGens generations starting
// with initialBoard using a time step of timeStep seconds.
// Input: a Board object initialBoard, a int of generations parameter numGens, a float64 time interval timeStep,
// parameters for cell and virus, and initial number of target cells and infectious cells
// Output: a slice of numGens + 1 total Board objects.
func SimulateViralSpread(initialBoard Board, numGens int, timeSteps float64, parameters Parameters, initialT, initialI int) []Board {
	timePoints := make([]Board, numGens+1)
	timePoints[0] = initialBoard

	// now range over the number of generations and update the Board each time
	for i := 1; i <= numGens; i++ {
		timePoints[i] = UpdateBoard(timePoints[i-1], timeSteps, parameters, initialT, initialI)
	}

	return timePoints

}

// UpdateBoard updates current board with new deltaT and new daltaI
// Input: a board object currentBoard, a float64 for timeSteps, a parameters including different necessary parameters for cells and virus,
// two int for T and I which are target cells and infected cells
// Output: a borad object which is an updated board from current board
func UpdateBoard(currentBoard Board, timeSteps float64, parameters Parameters, T, I int) Board {
	// Copy Board and store it in newBoard
	newBoard := CopyBoard(currentBoard)
	// Calculate deltaT and deltaI
	deltaT := CalculateDeltaT(T, I, timeSteps, parameters)
	deltaI := CalculateDeltaI(T, I, timeSteps, parameters)
	// Update the states of infectious cells and target cells
	UpdateState(newBoard, deltaT, deltaI)

	for i := range newBoard {
		for j := range newBoard[i] {
			if newBoard[i][j].state == "Infected" {
				UpdateCell(i, j, newBoard, timeSteps, parameters)
			}
		}
	}
	return newBoard
}

// CopyBoard is to do deep copy for current board to make sure each field would be copied
// Input: a board object for currentBoard
// Output: a copy of currentBoard
func CopyBoard(currentBoard Board) Board {
	// initialize an empty board according to the input board
	newBoard := make([][]Cell, len(currentBoard))
	for i := range currentBoard {
		newBoard[i] = make([]Cell, len(currentBoard[0]))
	}
	// deep copy all values in the board
	for i := range newBoard {
		for j := range newBoard[0] {
			newBoard[i][j].state = currentBoard[i][j].state
			newBoard[i][j].concVirus = currentBoard[i][j].concVirus
		}
	}
	return newBoard
}

// CalculateDeltaT is to calculate deltaT for untreated cell to cell model
// Input: two int for target cells and infected cells, float64 for timeSteps
// and a parameters object including several parameters that will be used in the calculation
// Output: a int object for deltaT
func CalculateDeltaT(T, I int, timeSteps float64, parameters Parameters) int {
	transmission := CalculateCellTransmission(T, I, parameters)

	deltaT := (parameters.lambda - transmission - parameters.dT) * timeSteps
	return int(deltaT)
}

// CalculateDeltaI is to calculate deltaI for untreated cell to cell model
// Input: two int for target cells and infected cells, float64 for timeSteps
// and a parameters object including several parameters that will be used in the calculation
// Output: a int object for deltaI
func CalculateDeltaI(T, I int, timeSteps float64, parameters Parameters) int {
	transmission := CalculateCellTransmission(T, I, parameters)

	deltaI := (transmission - parameters.delta*float64(I)) * timeSteps
	return int(deltaI)
}

// CalculateCellTransmission (omega*T*I)
// Input:
// Output:
func CalculateCellTransmission(T, I int, parameters Parameters) float64 {
	transmission := 0.0
	if parameters.treatment == "blockcell" || parameters.treatment == "blockvirus" {
		transmission = (1 - parameters.epsilonCell) * parameters.omega * float64(I) * float64(T)
	} else {
		transmission = parameters.omega * float64(I) * float64(T)
	}
	return transmission
}

// UpdateState updates the state of infection cells and target cells
// Input: a board object for currentBoard, two int objects for deltaT and deltaI which are generated from CalculateDeltaI and CalculateDeltaT
func UpdateState(currentBoard Board, deltaT, deltaI int) {
	// Update the state of infectious cells at currentBoard
	UpdateInfectiousCells(currentBoard, deltaI)
	// Update the state of target cells at currentBoard
	UpdateTargetCells(currentBoard, deltaT)
}

// UpdateInfectiousCells collects all the infectious cells and then randomly selects the number of absolute deltaI
// Input: a board object for current board, a int for deltaI which is calculated from CalculateDeltaI
func UpdateInfectiousCells(currentBoard Board, deltaI int) {
	// Create a list to store the index of infectious cells
	listInfectiousCells := make([]OrderedPair, 0)
	// Loop through currentBoard to find infectious cells
	for i := range currentBoard {
		for j := range currentBoard[i] {
			if currentBoard[i][j].state == "Infectious" {
				// If cell is infection
				// Set the index of infectious cells to OrderedPair
				var newOrderedPair OrderedPair
				newOrderedPair.x = i
				newOrderedPair.y = j
				// Store OrderedPair in the list
				listInfectiousCells = append(listInfectiousCells, newOrderedPair)
			}
		}
	}
	// Set seed
	rand.Seed(time.Now().UnixNano())
	// Randomly select deltaI times of infectious cells and change their state to dead
	for i := 0; i > deltaI; i-- {
		randIndex := rand.Intn(len(listInfectiousCells))
		// Change state of cell from infectious to dead
		currentBoard[listInfectiousCells[randIndex].x][listInfectiousCells[randIndex].y].state = "Dead"
	}
}

// UpdateTargetCells is to update the cells that will be susceptible to infection
// Input: a board object of currentBoard, a int for deltaT which is calculated from CalculateDeltaT
func UpdateTargetCells(currentBoard Board, deltaT int) {
	// Create a list to store the index of infectious cells
	listInfectiousCells := make([]OrderedPair, 0)
	// Loop through currentBoard to find infectious cells
	for i := range currentBoard {
		for j := range currentBoard[i] {
			if currentBoard[i][j].state == "Infectious" {
				// If cell is infection
				// Set the index of infectious cells to OrderedPair
				var newOrderedPair OrderedPair
				newOrderedPair.x = i
				newOrderedPair.y = j
				// Store OrderedPair in the list
				listInfectiousCells = append(listInfectiousCells, newOrderedPair)
			}
		}
	}

	// Randomly select deltaT times of infectious cells
	// and randomly choose a cell that will be affected by this infectious cells
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < deltaT; i++ {
		// Choose one infectious cells
		randIndex := rand.Intn(len(listInfectiousCells))
		// Randomly select a cell around this infectious cells
		// And the cell will be affected by infectious cells
		var randDirection OrderedPair
		randDirection.x = int(math.Pow(-1, float64(rand.Intn(100))))
		randDirection.y = int(math.Pow(-1, float64(rand.Intn(100))))
		xIndex := listInfectiousCells[randIndex].x + randDirection.x
		yIndex := listInfectiousCells[randIndex].y + randDirection.y
		// Avoid index out of range
		if xIndex < 0 {
			xIndex += 2
		} else if xIndex >= len(currentBoard) {
			xIndex -= 2
		}
		if yIndex < 0 {
			yIndex += 2
		} else if yIndex >= len(currentBoard[i]) {
			yIndex -= 2
		}
		// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		if currentBoard[xIndex][yIndex].state == "Uninfected" {
			currentBoard[xIndex][yIndex].state = "Infected"
		}
	}
}


// UpdateCell
// Input:
func UpdateCell(i, j int, currentBoard Board, timeSteps float64, parameters Parameters) {
	deltaR := 0.0
	if parameters.treatment == "blockvirus" || parameters.treatment == "blockboth" {
		deltaR = UpdateVirusConcentrationNoTreatment(i, j, currentBoard, timeSteps, parameters)
	} else {
		deltaR = UpdateVirusConcentrationBlockVirus(i, j, currentBoard, timeSteps, parameters)
	}

	currentBoard[i][j].concVirus += deltaR
	if currentBoard[i][j].concVirus >= parameters.threshold {
		currentBoard[i][j].state = "Infectious"
	}
}

// UpdateVirusConcentrationNoTreatment
// Input:
// Output:
func UpdateVirusConcentrationNoTreatment(i, j int, currentBoard Board, timeSteps float64, parameters Parameters) float64 {
	return currentBoard[i][j].concVirus * (parameters.alpha*(1-currentBoard[i][j].concVirus/parameters.rCap) - parameters.gamma - parameters.rho) * timeSteps
}

// UpdateVirusConcentrationBlockVirus
// Input:
// Output:
func UpdateVirusConcentrationBlockVirus(i, j int, currentBoard Board, timeSteps float64, parameters Parameters) float64 {
	return currentBoard[i][j].concVirus * ((1-parameters.epsilonVirus)*parameters.alpha*(1-currentBoard[i][j].concVirus/parameters.rCap) - parameters.gamma - parameters.rho) * timeSteps
}
