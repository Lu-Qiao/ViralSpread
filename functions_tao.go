package main

// SimulateViralSpread simulates the viral spread system over numGens generations starting
// with initialBoard using a time step of timeStep seconds.
// Input: a Board object initialBoard, a int of generations parameter numGens, and a float64 time interval timeStep.
// Output: a slice of numGens + 1 total Board objects.
func SimulateViralSpread(initialBoard Board, numGens int, timeSteps float64) []Board {
	timePoints := make([]Board, numGens+1)
	timePoints[0] = initialBoard

	// now range over the number of generations and update the Board each time
	for i := 1; i <= numGens; i++ {
		timePoints[i] = UpdateBoard(timePoints[i-1], timeSteps)
	}

	return timePoints

}

// UpdateBoard
// Input:
// Output:
func UpdateBoard(currentBoard Board, timeSteps float64) Board {
	newBoard := CopyBoard(currentBoard)

	for i := range newBoard {
		for j := range newBoard[i] {
			_ = j
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
