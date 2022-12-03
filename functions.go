package main

import (
	"math/rand"
	"time"
)

// SimulateViralSpread simulates the viral spread system over numGens generations
// starting with initialBoard using a time step of timeStep seconds.
// Input: a Board object initialBoard, a int of generations parameter numGens, a
// float64 time interval timeStep,parameters for cell and virus, and initial
// number of target cells and infectious cells
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
// Input: a board object currentBoard, a float64 for timeSteps, a parameters
// including different necessary parameters for cells and virus, two int for T
// and I which are target cells and infected cells
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

// CopyBoard is to do deep copy for current board to make sure each field would
// be copied
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
// Input: two int for target cells and infected cells, float64 for timeSteps and a
// parameters object including several parameters that will be used in the calculation
// Output: a int object for deltaT
func CalculateDeltaT(T, I int, timeSteps float64, parameters Parameters) int {
	transmission := CalculateCellTransmission(T, I, parameters)

	deltaT := (parameters.lambda - transmission - parameters.dT) * timeSteps
	return int(deltaT)
}

// CalculateDeltaI is to calculate deltaI for untreated cell to cell model
// Input: two int for target cells and infected cells, float64 for timeSteps and a
// parameters object including several parameters that will be used in the calculation
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
	if parameters.treatment == "blockcell" || parameters.treatment == "blockboth" {
		transmission = parameters.epsilonCell * parameters.omega * float64(I) * float64(T)
	} else {
		transmission = parameters.omega * float64(I) * float64(T)
	}
	return transmission
}

// UpdateState updates the state of infection cells and target cells
// Input: a board object for currentBoard, two int objects for deltaT and deltaI
// which are generated from CalculateDeltaI and CalculateDeltaT
func UpdateState(currentBoard Board, deltaT, deltaI int) {
	// Update the state of infectious cells at currentBoard
	UpdateInfectiousCells(currentBoard, deltaI)
	// Update the state of target cells at currentBoard
	UpdateTargetCells(currentBoard, deltaT)
}

// UpdateInfectiousCells collects all the infectious cells and then randomly selects
// the number of absolute deltaI
// Input: a board object for current board, a int for deltaI which is calculated
// from CalculateDeltaI
func UpdateInfectiousCells(currentBoard Board, deltaI int) {
	listInfectiousCells := FindInfectiousCells(currentBoard)
	// Set seed
	rand.Seed(time.Now().UnixNano())
	// Randomly select deltaI times of infectious cells and change their state to dead
	if len(listInfectiousCells) != 0 {
		for i := 0; i > deltaI; i-- {
			randIndex := rand.Intn(len(listInfectiousCells))
			// Change state of cell from infectious to dead
			currentBoard[listInfectiousCells[randIndex].x][listInfectiousCells[randIndex].y].state = "Dead"
		}
	}
}
// CountInfectiousCells store the index of infectious cells on the currentBoard
// Input: a board object for current board
// Output: a slice of OrderedPair
func FindInfectiousCells(currentBoard Board) []OrderedPair{
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
return listInfectiousCells
}

// UpdateCell updates the state and virusconcentation of the current cells based on the treatment
// Input: intergers for index, a board object for current board, timeStep as float64, parameters
func UpdateCell(i, j int, currentBoard Board, timeSteps float64, parameters Parameters) {
	deltaR := 0.0
	if parameters.treatment == "blockvirus" || parameters.treatment == "blockboth" {
		deltaR = UpdateVirusConcentrationBlockVirus(i, j, currentBoard, timeSteps, parameters)
	} else {
		deltaR = UpdateVirusConcentrationNoTreatment(i, j, currentBoard, timeSteps, parameters)
	}

	currentBoard[i][j].concVirus += deltaR
	if currentBoard[i][j].concVirus >= parameters.threshold {
		currentBoard[i][j].state = "Infectious"
	}
}

// UpdateVirusConcentrationNoTreatment calculates deltaR when no treatment
// Input: intergers for index, a board object for current board, timeStep as float64, parameters
// Output: float64
func UpdateVirusConcentrationNoTreatment(i, j int, currentBoard Board, timeSteps float64, parameters Parameters) float64 {
	return currentBoard[i][j].concVirus * (parameters.alpha*(1-currentBoard[i][j].concVirus/parameters.rCap) - parameters.gamma - parameters.rho) * timeSteps
}

// UpdateVirusConcentrationBlockVirus calculates deltaR when virus replication is blocked
// Input: intergers for index, a board object for current board, timeStep as float64, parameters
// Output: float64
func UpdateVirusConcentrationBlockVirus(i, j int, currentBoard Board, timeSteps float64, parameters Parameters) float64 {
	return currentBoard[i][j].concVirus * ((1-parameters.epsilonVirus)*parameters.alpha*(1-currentBoard[i][j].concVirus/parameters.rCap) - parameters.gamma - parameters.rho) * timeSteps
}

// RandomInfectCell is to randomly select a cell around infectious cell to be infected
// Input: a board object of currentBoard, a OrderedPair object of the position of infectious cell,
// and a list object of the positions of the cells that are around infectious cells
// Output: a board object that updates the cell which is infected by infectious cell
func RandomInfectCell(currentBoard Board, infectCell OrderedPair, cellAround []OrderedPair) Board {
	// Set seed
	rand.Seed(time.Now().UnixNano())
	// Randomly select a index in cellAround list
	selectIndex := rand.Intn(len(cellAround))
	beInfectedCell := cellAround[selectIndex]
	///// If the position is out of range
	// If x is < 0 then changes position to the right side of infectious cell
	if beInfectedCell.x < 0 {
		beInfectedCell.x += 2
		// If x position is larger than width if board,
		// then change position to the left side of infectious cell
	} else if beInfectedCell.x >= len(currentBoard) {
		beInfectedCell.x -= 2
	}
	// If y is < 0 then changes position to the up side of infectious cell
	if beInfectedCell.y < 0 {
		beInfectedCell.y += 2
		// If y position is larger than width if board,
		// then change position to the down side of infectious cell
	} else if beInfectedCell.y >= len(currentBoard[0]) {
		beInfectedCell.y -= 2
	}
	// If the cell that be chosen is uninfectedm then change the state to infected and concVirus to 1.
	// That means that the cell is infected by infectious cell
	if currentBoard[beInfectedCell.x][beInfectedCell.y].state == "Uninfected" {
		currentBoard[beInfectedCell.x][beInfectedCell.y].state = "Infected"
		currentBoard[beInfectedCell.x][beInfectedCell.y].concVirus = 1
		return currentBoard
	}
	// If cellAround only has one element, then return Board
	// It means that there is no uninfected cell around it, so return board
	if len(cellAround) == 1 {
		return currentBoard
	} else {
		// If there has other cells around it,
		if selectIndex == len(cellAround)-1 {
			// Delete the chosen cell in the list
			cellAround = cellAround[0:selectIndex]
			// Use recursion to randomly select again
			currentBoard = RandomInfectCell(currentBoard, infectCell, cellAround)
		} else {
			cellAround = append(cellAround[:selectIndex], cellAround[selectIndex+1:]...)
			currentBoard = RandomInfectCell(currentBoard, infectCell, cellAround)
		}
	}
	return currentBoard
}

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
	if len(listInfectiousCells) != 0 {
		for i := 0; i < deltaT; i++ {
			var infectiousCell, up, down, left, right OrderedPair
			randIndex := rand.Intn(len(listInfectiousCells))
			infectiousCell = listInfectiousCells[randIndex]
			up.x, down.x, left.x, right.x = infectiousCell.x-1, infectiousCell.x+1, infectiousCell.x, infectiousCell.x
			up.y, down.y, left.y, right.y = infectiousCell.y, infectiousCell.y, infectiousCell.y-1, infectiousCell.y+1
			cellAround := []OrderedPair{up, down, left, right}
			currentBoard = RandomInfectCell(currentBoard, infectiousCell, cellAround)
		}
	}

}
