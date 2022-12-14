package main

import (
	"fmt"
	"testing"
)

func TestInitializeTissue(t *testing.T) {
	type Test struct {
		board Board
		width int
	}

	var t1 Test
	t1.width = 5
	t1.board = InitializeTissue(t1.width)

	// Check number of row
	if len(t1.board) != t1.width {
		t.Errorf("Expect %d row of board, got %d", t1.width, len(t1.board))
	} else {
		fmt.Println("Test InitializeTissue: number of row of board is correct!")
	}
	
	// Check number of col
	var countErrorCol int
	for i := 0; i < len(t1.board); i++ {
		if len(t1.board[i]) != t1.width {
			t.Errorf("Expect %d col of board in row %d, got %d", t1.width, i, len(t1.board))
			countErrorCol ++
		} 
	}

	if countErrorCol == 0 {
		fmt.Println("Test InitializeTissue: number of col of board is correct!")
	}

	// Check state of each cell
	var countErrorCell int
	for i := 0; i < len(t1.board); i++ {
		for j := 0; j < len(t1.board[i]); j++ {
			if t1.board[i][j].state != "Uninfected" {
				t.Errorf("Expect cell at %d, %d of board is uninfected, but got %s", i, j, t1.board[i][j].state)
				countErrorCell ++
			}
		}
	}

	if countErrorCell == 0 {
		fmt.Println("Test InitializeTissue: states of cells are correct!")
	}

}

func TestAssignStart(t *testing.T) {
	type Test struct {
		board Board
		width int
		position OrderedPair
		answer int
		answerPosition []OrderedPair
	}

	var t1 Test
	t1.width = 5
	t1.answer = 2
	t1.position.x, t1.position.y = 2, 2
	t1.answerPosition = make([]OrderedPair, 0)

	t1.board = InitializeTissue(t1.width)
	AssignStart(t1.board, t1.position, 1)

	// Check if number of initial infectious cell
	for i := 0; i < len(t1.board); i++ {
		for j := 0; j < len(t1.board[i]); j++ {
			if t1.board[i][j].state == "Infectious" {
				var p OrderedPair
				p.x, p.y = i, j
				t1.answerPosition = append(t1.answerPosition, p)
			}
		}
	}

	// Check number
	if len(t1.answerPosition) != 2 {
		t.Errorf("Expect %d cells be infectious, got %d", t1.answer, t1.answerPosition)
	} else {
		fmt.Println("Test AssignStart: number of infectious cells is correct!")
	}

	// Check position
	var countErrorPosition int
	for i := 0; i < len(t1.answerPosition); i++ {
		if t1.answerPosition[i].x != t1.position.x+i || t1.answerPosition[i].y != t1.position.y {
			t.Errorf("Expect cell at %d, %d be infectious, but got %d, %d", t1.position.x+i, t1.position.y, t1.answerPosition[i].x, t1.answerPosition[i].y)
			countErrorPosition ++
		}
	}

	if countErrorPosition == 0 {
		fmt.Println("Test AssignStart: positions of infectious cells are correct!")
	}
}

func TestRanomStart(t *testing.T) {
	type Test struct {
		board Board
		width int
		numPosition int
		answer int
	}

	var t1 Test
	t1.width = 5
	t1.answer = 1
	t1.numPosition = 1

	t1.board = InitializeTissue(t1.width)
	RandomStart(t1.board, t1.numPosition, 1)

	// Count infectious cells
	var countInfectiousCell1 int
	for i := 0; i < len(t1.board); i++ {
		for j := 0; j < len(t1.board[i]); j++ {
			if t1.board[i][j].state == "Infectious" {
				countInfectiousCell1 ++
			}
		}
	}

	if countInfectiousCell1 != t1.answer {
		t.Errorf("Expect %d of cells be infectious, but got %d", t1.answer, countInfectiousCell1)

	} else {
		fmt.Println("Test RandomStart: number of infectious cells is correct!")
	}

}

func TestRandomInfectCell(t *testing.T) {
	type Test struct {
		board          Board
		startInfection OrderedPair
		answer         int
		answerPosition OrderedPair
	}

	//// Test if function choose one cell to be infected
	var t1 Test
	var above, below, right, left OrderedPair
	var count1 int

	t1.startInfection.x, t1.startInfection.y = 3, 3
	above.x, below.x, right.x, left.x = 2, 4, 3, 3
	above.y, below.y, right.y, left.y = 3, 3, 4, 2
	t1.answer = 1
	cellAround1 := []OrderedPair{above, below, right, left}

	t1.board = InitializeTissue(5)
	AssignStart(t1.board, t1.startInfection, 1)
	t1.board = RandomInfectCell(t1.board, t1.startInfection, cellAround1)

	for i := 0; i < len(cellAround1); i++ {
		if t1.board[cellAround1[i].x][cellAround1[i].y].state == "Infected" {
			count1++
		}
	}
	if t1.answer != count1 {
		t.Errorf("Expect %d infected cell around infectious cell, got %d", t1.answer, count1)
	} else {
		fmt.Println("Test RandomInfectCell: number of infected cells is correct!")
	}

	///// Test if function selects the cell we expected
	var t2 Test
	var above2, below2, right2, left2 OrderedPair
	var count2 int

	t2.answer = 0
	t2.answerPosition.x, t2.answerPosition.y = 3, 2
	t2.startInfection.x, t1.startInfection.y = 3, 3
	above2.x, below2.x, right2.x, left2.x = 2, 4, 3, 3
	above2.y, below2.y, right2.y, left2.y = 3, 3, 4, 2
	cellAround := []OrderedPair{above2, below2, right2, left2}

	t2.board = InitializeTissue(5)
	t2.board[2][3].state, t2.board[4][3].state, t2.board[3][4].state = "Infected", "Infected", "Infected"
	t2.board = RandomInfectCell(t2.board, t2.startInfection, cellAround)

	for i := 0; i < len(cellAround); i++ {
		if t2.board[cellAround[i].x][cellAround[i].y].state == "Uninfected" {
			count2++
		}
	}

	if t2.answer != count2 {
		t.Errorf("Expect %d uninfected cell around infectious cell, got %d", t2.answer, count2)
	} else {
		fmt.Println("Test RandomInfectCell: the expected cell is infected!")
	}
	
}

func TestUpdateVirusConcentrationBlockVirus(t *testing.T) {
	type Test struct {
		board          Board
		startInfection OrderedPair
		parameters     Parameters
		answer         float64
	}
	// Test if the output of calculation is correct
	var t1 Test

	t1.startInfection.x, t1.startInfection.y = 2, 2
	t1.parameters.epsilonVirus = 2.0
	t1.parameters.alpha = 3.0
	t1.parameters.rCap = 4.0
	t1.parameters.gamma = 1.0
	t1.parameters.rho = 5.0
	t1.answer = -8.25

	t1.board = InitializeTissue(5)
	AssignStart(t1.board, t1.startInfection, 1)
	outPut := UpdateVirusConcentrationBlockVirus(2, 2, t1.board, 1, t1.parameters)

	if t1.answer != outPut {
		t.Errorf("Expected virus concentration is %f, got %f", t1.answer, outPut)
	} else {
		fmt.Println("Test UpdateVirusConcentrationBlockVirus: virus concentration is updated correctly!")
	}

}

func TestUpdateVirusConcentrationNoTreatment(t *testing.T) {
	type Test struct {
		board          Board
		startInfection OrderedPair
		parameters     Parameters
		answer         float64
	}
	// Test if the output of calculation is correct
	var t1 Test

	t1.startInfection.x, t1.startInfection.y = 2, 2
	t1.parameters.alpha = 3.0
	t1.parameters.rCap = 4.0
	t1.parameters.gamma = 1.0
	t1.parameters.rho = 5.0
	t1.answer = -3.750

	t1.board = InitializeTissue(5)
	AssignStart(t1.board, t1.startInfection, 1)
	outPut := UpdateVirusConcentrationBlockVirus(2, 2, t1.board, 1, t1.parameters)

	if t1.answer != outPut {
		t.Errorf("Expected virus concentration is %f, got %f", t1.answer, outPut)
	} else {
		fmt.Println("Test UpdateVirusConcentrationNoTreatment: virus concentration is updated correctly!")
	}

}

func TestFindInfectiousCells(t *testing.T) {
	type Test struct {
		board          Board
		startInfection OrderedPair
		parameters     Parameters
		answer         int
	}
	var t1 Test

	t1.answer = 0
	t1.board = InitializeTissue(5)
	outPut1 := FindInfectiousCells(t1.board)

	if t1.answer != len(outPut1) {
		t.Errorf("Expected number of infectious cells is %d, got %d", t1.answer, len(outPut1))
	} else {
		fmt.Println("Test FindInfectiousCells: none of cells is infectious!")
	}

	var t2 Test
	t2.answer = 9
	t2.board = InitializeTissue(3)

	for i := 0; i < len(t2.board); i++ {
		for j := 0; j < len(t2.board[0]); j++ {
			t2.board[i][j].state = "Infectious"
		}
	}

	outPut2 := FindInfectiousCells(t2.board)

	if t2.answer != len(outPut2) {
		t.Errorf("Expected number of infectious cells is %d, got %d", t2.answer, len(outPut2))
	} else {
		fmt.Println("Test FindInfectiousCells: all cells are infectious!")
	}

}