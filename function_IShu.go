package main

import (
	"math/rand"
	"time"
)

func RandomInfectCell(currentBoard Board, infectCell OrderedPair, cellAround []OrderedPair) Board {
	rand.Seed(time.Now().UnixNano())
	selectIndex := rand.Intn(len(cellAround))
	beInfectedCell := cellAround[selectIndex]
	if beInfectedCell.x < 0 {
		beInfectedCell.x += 2
	} else if beInfectedCell.x >= len(currentBoard) {
		beInfectedCell.x -= 2
	}
	if beInfectedCell.y < 0 {
		beInfectedCell.y += 2
	} else if beInfectedCell.y >= len(currentBoard[0]) {
		beInfectedCell.y -= 2
	}
	if currentBoard[beInfectedCell.x][beInfectedCell.y].state == "Uninfected" {
		currentBoard[beInfectedCell.x][beInfectedCell.y].state = "Infected"
	} else if len(cellAround) == 1 {
		return currentBoard
	} else {
		if selectIndex == len(cellAround)-1 {
			cellAround = cellAround[0:selectIndex]
			currentBoard = RandomInfectCell(currentBoard, infectCell, cellAround)
		} else {
			cellAround = cellAround[0:selectIndex]
			cellAround = append(cellAround, cellAround[selectIndex+1:]...)
			currentBoard = RandomInfectCell(currentBoard, infectCell, cellAround)
		}
	}
	return currentBoard
}

func UpdateTargetCells2(currentBoard Board, deltaT int) {
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
