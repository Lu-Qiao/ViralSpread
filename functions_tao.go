package main

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

// UpdateBoard
// Input:
// Output:
func UpdateBoard(currentBoard Board, timeSteps float64, parameters Parameters, T, I int) Board {
	newBoard := CopyBoard(currentBoard)

	deltaT := CalculateDeltaT(T, I, timeSteps, parameters)
	deltaI := CalculateDeltaI(T, I, timeSteps, parameters)

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

// CalculateDeltaT
// Input:
// Output:
func CalculateDeltaT(T, I int, timeSteps float64, parameters Parameters) int {
	deltaT := (parameters.lambda - parameters.omega*float64(I)*float64(T) - parameters.dT) * timeSteps
	return int(deltaT)
}

// CalculateDeltaI
// Input:
// Output:
func CalculateDeltaI(T, I int, timeSteps float64, parameters Parameters) int {
	deltaI := (parameters.omega*float64(I)*float64(T) - parameters.delta*float64(I)) * timeSteps
	return int(deltaI)
}

// UpdateState
// Input:
func UpdateState(currentBoard Board, deltaT, deltaI int) {
	UpdateInfectiousCells(currentBoard, deltaI)
	UpdateTargetCells(currentBoard, deltaT)
}

// UpdateTargetCells
// Input:
func UpdateTargetCells(currentBoard Board, deltaT int) {

}

// UpdateInfectiousCells
// Input:
// Output:
func UpdateInfectiousCells(currentBoard Board, deltaI int) {

}

// UpdateCell
// Input:
func UpdateCell(i, j int, currentBoard Board, timeSteps float64, parameters Parameters) {
	deltaR := UpdateVirusConcentration(i, j, currentBoard, timeSteps, parameters)
	currentBoard[i][j].concVirus += deltaR
	if currentBoard[i][j].concVirus >= parameters.threshold {
		currentBoard[i][j].state = "Infectious"
	}
}

// UpdateVirusConcentration
// Input:
// Output:
func UpdateVirusConcentration(i, j int, currentBoard Board, timeSteps float64, parameters Parameters) float64 {
	deltaR := currentBoard[i][j].concVirus * (parameters.alpha*(1-currentBoard[i][j].concVirus/parameters.rCap) - parameters.gamma - parameters.rho) * timeSteps
	return deltaR
}
