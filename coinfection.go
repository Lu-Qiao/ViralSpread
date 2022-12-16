package main

import (
	"fmt"
	"gifhelper"
	"math/rand"
	"os"
	"time"
)

// SimulateCo is the highest level function for simulate coinfection
// Input: an Inputs object that contains all input parameters
func SimulateCo(allInputs1, allInputs2 Inputs) {
	// Copy all parameters from inputs
	width := allInputs1.width
	mode := allInputs1.mode
	numInfectious := allInputs1.numInfectious
	initialPosition := allInputs1.initialPosition
	numGens := allInputs1.numGens
	timeSteps := allInputs1.timeSteps
	imageFrequency := allInputs1.imageFrequency
	parameters1 := allInputs1.parameters
	parameters2 := allInputs2.parameters

	// Initialize tissue
	Tissue := InitializeTissue(width)
	// If mode equal random
	if mode == "random" {
		// then call RandomStart to random start the infection
		RandomCoStart(Tissue, numInfectious, parameters1.threshold)
	} else if mode == "assign" {
		// Assign specific position for infection
		AssignCoStart(Tissue, initialPosition, parameters1.threshold)
	}
	// Simulations
	fmt.Println("Simulating system.")

	timePoints, cellTimePoints := SimulateCoinfection(Tissue, numGens, numInfectious, timeSteps, parameters1, parameters2)

	fmt.Println("Viral Spread has been simulated!")

	// create filename according to inputs
	filename := os.Args[0] + "_" + mode + "_" + parameters1.treatment + "_treatment"

	// Save data
	fmt.Println("Ready to plot graph!")

	SaveCoCellDataToCSV(timeSteps, cellTimePoints, filename)
	// PlotCellData(timeSteps, cellTimePoints)

	fmt.Println("Graph drawn!")

	// Make GIF
	fmt.Println("Ready to draw images.")
	// Create images
	images := AnimateSystem(timePoints, width, imageFrequency)

	fmt.Println("Images drawn!")

	fmt.Println("Making GIF...")
	// To Gif
	gifhelper.ImagesToGIF(images, filename+"_coninfection")

	fmt.Println("Animated GIF produced!")

	fmt.Println("GIF saved successfully!")

	fmt.Printf("Waiting for next simulation...\n\n")
}

// SimulateCoinfection simulates the viral spread system with two viruses over numGens generations
// starting with initialBoard using a time step of timeStep seconds.
// Input: a Board object initialBoard, a int of generations parameter numGens, a float64 time interval timeStep,parameters for cell and virus, and initial number of target cells and infectious cells
// Output: a slice of numGens + 1 total Board objects.
func SimulateCoinfection(initialBoard Board, numGens, numInfectious int, timeSteps float64, parameters1, parameters2 Parameters) ([]Board, [][]int) {
	timePoints := make([]Board, numGens+1)
	timePoints[0] = initialBoard
	// Count number of cells (each type) in each generation
	// order: normal cell, target cell1, infectious cell1, target cell2, infectious cell2, dead cell1, dead cell2
	cellTimePoints := make([][]int, numGens+1)
	// i is generation, and j matchs to the order of cell types
	for i := range cellTimePoints {
		cellTimePoints[i] = make([]int, 7)
	}
	// the number of uninfectious cells at the beginning
	cellTimePoints[0][0] = len(initialBoard)*len(initialBoard[0]) - 2*numInfectious
	// the number of infectious cells at the beginning
	cellTimePoints[0][2] = numInfectious
	cellTimePoints[0][4] = numInfectious

	// now range over the number of generations and update the Board each time
	for i := 1; i <= numGens; i++ {
		timePoints[i], cellTimePoints[i] = UpdateCoBoard(timePoints[i-1], timeSteps, parameters1, parameters2)
	}

	return timePoints, cellTimePoints
}

// UpdateCoBoard updates current board with new deltaT and new daltaI
// Input: a board object currentBoard, a float64 for timeSteps, a parameters
// including different necessary parameters for cells and virus, two int for T
// and I which are target cells and infected cells
// Output: a borad object which is an updated board from current board
func UpdateCoBoard(currentBoard Board, timeSteps float64, parameters1, parameters2 Parameters) (Board, []int) {
	// Copy Board and store it in newBoard
	newBoard := CopyBoard(currentBoard)
	// get number of different cells: N, T1, I1, T2, I2, D
	cellNumber := GetCoCellNumber(currentBoard)
	// Calculate deltaT and deltaI
	deltaT := CalculateCoDeltaT(cellNumber[1]+cellNumber[3], cellNumber[2]+cellNumber[4], timeSteps, parameters1, parameters2)
	deltaI1 := CalculateDeltaI(cellNumber[1]+cellNumber[3], cellNumber[2], timeSteps, parameters1)
	deltaI2 := CalculateDeltaI(cellNumber[1]+cellNumber[3], cellNumber[4], timeSteps, parameters2)
	// Update the states of infectious cells and target cells
	UpdateCoState(newBoard, deltaT, deltaI1, deltaI2)
	// range through each square of the board
	for i := range newBoard {
		for j := range newBoard[i] {
			// if cell state is infected
			if newBoard[i][j].state == "Infected1" {
				// updates cell
				UpdateCoCell(i, j, newBoard, timeSteps, parameters1, 1)
			} else if newBoard[i][j].state == "Infected2" {
				// updates cell
				UpdateCoCell(i, j, newBoard, timeSteps, parameters2, 2)
			}
		}
	}
	return newBoard, cellNumber
}

// GetCoCellNumber counts the cells in the currentBoard
// Input: current board (Board)
// output: number of different cells in the current boards (int, int)
func GetCoCellNumber(currentBoard Board) []int {
	// get number of different cells: N, T1, I1, T2, I2, D1, D2
	cellNumber := make([]int, 7)
	// order: normal cell, target cell, infectious cell, dead cell
	// range through each square, and count each type of cell
	for i := range currentBoard {
		for j := range currentBoard[i] {
			if currentBoard[i][j].state == "Infected1" {
				cellNumber[1]++
			}
			if currentBoard[i][j].state == "Infectious1" {
				cellNumber[2]++
			}
			if currentBoard[i][j].state == "Infected2" {
				cellNumber[3]++
			}
			if currentBoard[i][j].state == "Infectious2" {
				cellNumber[4]++
			}
			if currentBoard[i][j].state == "Dead1" {
				cellNumber[5]++
			}
			if currentBoard[i][j].state == "Dead2" {
				cellNumber[6]++
			}
		}
	}
	// uninfected cells can be counted by total cells - other cells type
	cellNumber[0] = len(currentBoard)*len(currentBoard[0]) - cellNumber[1] - cellNumber[2] - cellNumber[3] - cellNumber[4] - cellNumber[5] - cellNumber[6]

	return cellNumber
}

// CalculateCoDeltaT is to calculate deltaT for untreated cell to cell model
// Input: two int for target cells and infected cells, float64 for timeSteps and a
// parameters object including several parameters that will be used in the calculation
// Output: a int object for deltaT
func CalculateCoDeltaT(T, I int, timeSteps float64, parameters1, parameters2 Parameters) int {
	// calculate transmission
	transmission1 := CalculateCellTransmission(T, I, parameters1)
	transmission2 := CalculateCellTransmission(T, I, parameters2)
	// use math equation to calculate deltaT
	deltaT := (parameters1.lambda - transmission1 - transmission2 - parameters1.dT*float64(T)) * timeSteps
	return int(deltaT)
}

// UpdateCoState updates the state of infection cells and target cells
// Input: a board object for currentBoard, two int objects for deltaT and deltaI
// which are generated from CalculateDeltaI and CalculateDeltaT
func UpdateCoState(currentBoard Board, deltaT, deltaI1, deltaI2 int) {
	// Update the state of infectious cells at currentBoard
	UpdateCoInfectiousCells(currentBoard, deltaI1, deltaI2)
	// Update the state of target cells at currentBoard
	UpdateCoTargetCells(currentBoard, deltaT)
}

// UpdateCoInfectiousCells collects all the infectious cells and then randomly selects the number of absolute deltaI
// Input: a board object for current board, a int for deltaI which is calculated
// from CalculateDeltaI
func UpdateCoInfectiousCells(currentBoard Board, deltaI1, deltaI2 int) {
	listInfectiousCells1, listInfectiousCells2 := FindCoInfectiousCells(currentBoard)
	// Set seed
	rand.Seed(time.Now().UnixNano())
	// Randomly select deltaI times of infectious cells and change their state to dead
	if len(listInfectiousCells1) != 0 {
		for i := 0; i > deltaI1; i-- {
			randIndex := rand.Intn(len(listInfectiousCells1))
			// Change state of cell from infectious to dead1
			currentBoard[listInfectiousCells1[randIndex].x][listInfectiousCells1[randIndex].y].state = "Dead1"
		}
	}
	if len(listInfectiousCells2) != 0 {
		for i := 0; i > deltaI2; i-- {
			randIndex := rand.Intn(len(listInfectiousCells2))
			// Change state of cell from infectious to dead2
			currentBoard[listInfectiousCells2[randIndex].x][listInfectiousCells2[randIndex].y].state = "Dead2"
		}
	}
}

// FindCoInfectiousCells store the index of two types of infectious cells on the currentBoard
// Input: a board object for current board
// Output: a slice of OrderedPair
func FindCoInfectiousCells(currentBoard Board) ([]OrderedPair, []OrderedPair) {
	// Create a list to store the index of infectious cells
	listInfectiousCells1 := make([]OrderedPair, 0)
	listInfectiousCells2 := make([]OrderedPair, 0)
	// Loop through currentBoard to find infectious cells
	for i := range currentBoard {
		for j := range currentBoard[i] {
			if currentBoard[i][j].state == "Infectious1" {
				// If cell is infection
				// Set the index of infectious cells to OrderedPair
				var newOrderedPair OrderedPair
				newOrderedPair.x = i
				newOrderedPair.y = j
				// Store OrderedPair in the list
				listInfectiousCells1 = append(listInfectiousCells1, newOrderedPair)
			} else if currentBoard[i][j].state == "Infectious2" {
				// If cell is infection
				// Set the index of infectious cells to OrderedPair
				var newOrderedPair OrderedPair
				newOrderedPair.x = i
				newOrderedPair.y = j
				// Store OrderedPair in the list
				listInfectiousCells2 = append(listInfectiousCells2, newOrderedPair)
			}
		}
	}
	return listInfectiousCells1, listInfectiousCells2
}

// UpdateCoCell updates the state and virusconcentation of the current cells based on the treatment
// Input: intergers for index, a board object for current board, timeStep as float64, parameters
func UpdateCoCell(i, j int, currentBoard Board, timeSteps float64, parameters Parameters, state int) {
	deltaR := 0.0
	// under different condition, it has different equation
	// if simulate treatment model
	if parameters.treatment == "blockvirus" || parameters.treatment == "blockboth" {
		deltaR = UpdateVirusConcentrationBlockVirus(i, j, currentBoard, timeSteps, parameters)
	} else {
		// model withour treatment
		deltaR = UpdateVirusConcentrationNoTreatment(i, j, currentBoard, timeSteps, parameters)
	}
	// update virus concentration
	currentBoard[i][j].concVirus += deltaR
	// if virus concentration in cell exceeds threshold, then turn to infectious
	if currentBoard[i][j].concVirus >= parameters.threshold {
		if state == 1 {
			currentBoard[i][j].state = "Infectious1"
		} else if state == 2 {
			currentBoard[i][j].state = "Infectious2"
		}
	}
}

// RandomCoInfectCell is to randomly select a cell around infectious cell to be infected
// Input: a board object of currentBoard, a OrderedPair object of the position of infectious cell, and a list object of the positions of the cells that are around infectious cells
// Output: a board object that updates the cell which is infected by infectious cell
func RandomCoInfectCell(currentBoard Board, infectCell OrderedPair, cellAround []OrderedPair, state string) Board {
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
		currentBoard[beInfectedCell.x][beInfectedCell.y].state = state
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
			currentBoard = RandomCoInfectCell(currentBoard, infectCell, cellAround, state)
		} else {
			// Delete the cell
			cellAround = append(cellAround[:selectIndex], cellAround[selectIndex+1:]...)
			// Choose again
			currentBoard = RandomCoInfectCell(currentBoard, infectCell, cellAround, state)
		}
	}
	return currentBoard
}

// UpdateCoTargetCells update the state of target cells
// Two viruses will take random turns to infect surrrounding cells.
// To update target cell, infect cells which are near infectious cells
// Input: a borad object for current board, int for deltaT
func UpdateCoTargetCells(currentBoard Board, deltaT int) {
	// Create a list to store the index of infectious cells
	listInfectiousCells1 := make([]OrderedPair, 0)
	listInfectiousCells2 := make([]OrderedPair, 0)
	// Loop through currentBoard to find infectious cells
	for i := range currentBoard {
		for j := range currentBoard[i] {
			if currentBoard[i][j].state == "Infectious1" {
				// If cell is infection
				// Set the index of infectious cells to OrderedPair
				var newOrderedPair OrderedPair
				newOrderedPair.x = i
				newOrderedPair.y = j
				// Store OrderedPair in the list
				listInfectiousCells1 = append(listInfectiousCells1, newOrderedPair)
			} else if currentBoard[i][j].state == "Infectious2" {
				// If cell is infection
				// Set the index of infectious cells to OrderedPair
				var newOrderedPair OrderedPair
				newOrderedPair.x = i
				newOrderedPair.y = j
				// Store OrderedPair in the list
				listInfectiousCells2 = append(listInfectiousCells2, newOrderedPair)
			}
		}
	}
	// Randomly select deltaT times of infectious cells
	// and randomly choose a cell that will be affected by this infectious cells
	if len(listInfectiousCells1) != 0 && len(listInfectiousCells2) != 0 {
		for i := 0; i < deltaT; i++ {
			// randomly set the order for two virus to infect
			rand.Seed(time.Now().UnixNano())
			order := rand.Intn(10)
			if order <= 5 {
				var infectiousCell, up, down, left, right OrderedPair
				// Randomly select one infectious cell
				randIndex := rand.Intn(len(listInfectiousCells1))
				infectiousCell = listInfectiousCells1[randIndex]
				// set position of cells near infectious cell
				up.x, down.x, left.x, right.x = infectiousCell.x-1, infectiousCell.x+1, infectiousCell.x, infectiousCell.x
				up.y, down.y, left.y, right.y = infectiousCell.y, infectiousCell.y, infectiousCell.y-1, infectiousCell.y+1
				cellAround := []OrderedPair{up, down, left, right}
				// Randomly select a cell near infectious cell using RandomInfectCell
				// update board
				currentBoard = RandomCoInfectCell(currentBoard, infectiousCell, cellAround, "Infected1")
			} else {
				var infectiousCell, up, down, left, right OrderedPair
				// Randomly select one infectious cell
				randIndex := rand.Intn(len(listInfectiousCells2))
				infectiousCell = listInfectiousCells2[randIndex]
				// set position of cells near infectious cell
				up.x, down.x, left.x, right.x = infectiousCell.x-1, infectiousCell.x+1, infectiousCell.x, infectiousCell.x
				up.y, down.y, left.y, right.y = infectiousCell.y, infectiousCell.y, infectiousCell.y-1, infectiousCell.y+1
				cellAround := []OrderedPair{up, down, left, right}
				// Randomly select a cell near infectious cell using RandomInfectCell
				// update board
				currentBoard = RandomCoInfectCell(currentBoard, infectiousCell, cellAround, "Infected2")
			}
		}
	} else if len(listInfectiousCells1) != 0 {
		for i := 0; i < deltaT; i++ {
			var infectiousCell, up, down, left, right OrderedPair
			// Randomly select one infectious cell
			randIndex := rand.Intn(len(listInfectiousCells1))
			infectiousCell = listInfectiousCells1[randIndex]
			// set position of cells near infectious cell
			up.x, down.x, left.x, right.x = infectiousCell.x-1, infectiousCell.x+1, infectiousCell.x, infectiousCell.x
			up.y, down.y, left.y, right.y = infectiousCell.y, infectiousCell.y, infectiousCell.y-1, infectiousCell.y+1
			cellAround := []OrderedPair{up, down, left, right}
			// Randomly select a cell near infectious cell using RandomInfectCell
			// update board
			currentBoard = RandomCoInfectCell(currentBoard, infectiousCell, cellAround, "Infected1")
		}
	} else if len(listInfectiousCells2) != 0 {
		for i := 0; i < deltaT; i++ {
			var infectiousCell, up, down, left, right OrderedPair
			// Randomly select one infectious cell
			randIndex := rand.Intn(len(listInfectiousCells2))
			infectiousCell = listInfectiousCells2[randIndex]
			// set position of cells near infectious cell
			up.x, down.x, left.x, right.x = infectiousCell.x-1, infectiousCell.x+1, infectiousCell.x, infectiousCell.x
			up.y, down.y, left.y, right.y = infectiousCell.y, infectiousCell.y, infectiousCell.y-1, infectiousCell.y+1
			cellAround := []OrderedPair{up, down, left, right}
			// Randomly select a cell near infectious cell using RandomInfectCell
			// update board
			currentBoard = RandomCoInfectCell(currentBoard, infectiousCell, cellAround, "Infected2")
		}
	}
}
