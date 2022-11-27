package main

// SimulateViralSpread simulates the viral spread system over numGens generations starting
// with initialBoard using a time step of timeStep seconds.
// Input: a Board object initialBoard, a int of generations parameter numGens, a float64 time interval timeStep,
// and parameters for cell and virus
// Output: a slice of numGens + 1 total Board objects.
func SimulateViralSpread(initialBoard Board, numGens int, timeSteps float64, parameters Parameters) []Board {
	timePoints := make([]Board, numGens+1)
	timePoints[0] = initialBoard

	// now range over the number of generations and update the Board each time
	for i := 1; i <= numGens; i++ {
		timePoints[i] = UpdateBoard(timePoints[i-1], timeSteps, parameters)
	}

	return timePoints

}

// UpdateBoard
// Input:
// Output:
func UpdateBoard(currentBoard Board, timeSteps float64, parameters Parameters) Board {
	newBoard := CopyBoard(currentBoard)

	UpdateState(newBoard, parameters)

	for i := range newBoard {
		for j := range newBoard[i] {
			if newBoard[i][j].state == "Infected" {
				UpdateCell(i, j, newBoard, parameters.threshold)
			}
		}
	}
	return newBoard
}

// CopyBoard
// Input:
// Output:
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

// UpdateState
// Input:
func UpdateState(currentBoard Board, parameters Parameters) {
	deltaT := UpdateTargetCells(currentBoard, parameters)
	deltaI := UpdateInfectiousCells(currentBoard, parameters)

	_ = deltaT
	_ = deltaI
}

// UpdateTargetCells
// Input:
// Output:
func UpdateTargetCells(currentBoard Board, parameters Parameters) float64 {
	deltaT := 0.0

	return deltaT
}

// UpdateInfectiousCells
// Input:
// Output:
func UpdateInfectiousCells(currentBoard Board, parameters Parameters) float64 {
	deltaI := 0.0

	return deltaI
}

// UpdateCell
// Input:
func UpdateCell(i, j int, currentBoard Board, threshold float64) {
	UpdateVirusConcentration(i, j, currentBoard)
	if currentBoard[i][j].concVirus >= threshold {
		currentBoard[i][j].state = "Infectious"
	}
}

// UpdateVirusConcentration
// Input:
func UpdateVirusConcentration(i, j int, currentBoard Board) {

}
