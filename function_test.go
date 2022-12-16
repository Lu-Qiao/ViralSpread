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
			t.Errorf("Expect %d col of board in row %d, got %d", t1.width, i, len(t1.board[i]))
			countErrorCol++
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
				countErrorCell++
			}
		}
	}

	if countErrorCell == 0 {
		fmt.Println("Test InitializeTissue: states of cells are correct!")
	}

}

func TestAssignStart(t *testing.T) {
	type Test struct {
		board          Board
		width          int
		position       OrderedPair
		answer         int
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
			countErrorPosition++
		}
	}

	if countErrorPosition == 0 {
		fmt.Println("Test AssignStart: positions of infectious cells are correct!")
	}
}

func TestRanomStart(t *testing.T) {
	type Test struct {
		board       Board
		width       int
		numPosition int
		answer      int
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
				countInfectiousCell1++
			}
		}
	}

	if countInfectiousCell1 != t1.answer {
		t.Errorf("Expect %d of cells be infectious, but got %d", t1.answer, countInfectiousCell1)

	} else {
		fmt.Println("Test RandomStart: number of infectious cells is correct!")
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

func TestCalculateCellTransmission(t *testing.T) {
	type Test struct {
		parameters Parameters
		t, i       int
		answer     float64
	}

	// Test model with blockcell treatment
	var t1 Test
	t1.parameters.treatment = "blockcell"
	t1.parameters.epsilonCell = 2
	t1.parameters.omega = 5
	t1.t = 20
	t1.i = 40
	t1.answer = 8000.0
	outPut1 := CalculateCellTransmission(t1.t, t1.i, t1.parameters)

	if t1.answer != outPut1 {
		t.Errorf("Expected cell transmission rate with blockcell treatment is %f, got %f", t1.answer, outPut1)
	} else {
		fmt.Println("Test CalculateCellTransmissions: cell transmission rate with blockcell treatment is correct!")
	}

	// Test model without treatment
	var t2 Test
	t2.parameters.omega = 5
	t2.t = 20
	t2.i = 40
	t2.answer = 4000.0
	outPut2 := CalculateCellTransmission(t2.t, t2.i, t2.parameters)

	if t2.answer != outPut2 {
		t.Errorf("Expected cell transmission rate with untreatment model is %f, got %f", t2.answer, outPut2)
	} else {
		fmt.Println("Test CalculateCellTransmissions: cell transmission rate with untreatment model is correct!")
	}

	// Test model with blockboth treatment
	var t3 Test
	t3.parameters.treatment = "blockboth"
	t3.parameters.epsilonCell = 2
	t3.parameters.omega = 5
	t3.t = 20
	t3.i = 40
	t3.answer = 8000.0
	outPut3 := CalculateCellTransmission(t3.t, t3.i, t3.parameters)

	if t3.answer != outPut3 {
		t.Errorf("Expected cell transmission rate with blockboth treatment is %f, got %f", t3.answer, outPut3)
	} else {
		fmt.Println("Test CalculateCellTransmissions: cell transmission rate with blockboth treatment is correct!")
	}

}

func TestCalculateDeltaI(t *testing.T) {
	type Test struct {
		parameters Parameters
		t, i       int
		answer     int
	}

	var t1 Test
	t1.t = 20
	t1.i = 10
	t1.parameters.delta = 4
	t1.parameters.omega = 5
	t1.answer = 960

	outPut1 := CalculateDeltaI(t1.t, t1.i, 1, t1.parameters)

	if t1.answer != outPut1 {
		t.Errorf("Expected DeltaI is %d, got %d", t1.answer, outPut1)
	} else {
		fmt.Println("Test CalculateDeltaI: DeltaI is correct!")
	}

}

func TestCalculateDeltaT(t *testing.T) {
	type Test struct {
		parameters Parameters
		t, i       int
		answer     int
	}

	var t1 Test
	t1.t = 20
	t1.i = 10
	t1.parameters.dT = 4
	t1.parameters.omega = 5
	t1.parameters.lambda = 10000
	t1.answer = 8920

	outPut1 := CalculateDeltaT(t1.t, t1.i, 1, t1.parameters)

	if t1.answer != outPut1 {
		t.Errorf("Expected DeltaT is %d, got %d", t1.answer, outPut1)
	} else {
		fmt.Println("Test CalculateDeltaT: DeltaT is correct!")
	}

}

func TestGetCellNumber(t *testing.T) {
	type Test struct {
		board  Board
		answer []int
	}

	// Test if function can correctly calculate number of each type of cells
	var t1 Test
	var countError1 int
	t1.board = InitializeTissue(5)
	t1.answer = []int{25, 0, 0, 0}
	name := []string{"Uninfected", "Infected", "Infectious", "Dead"}
	outPut1 := GetCellNumber(t1.board)

	for i := 0; i < len(outPut1); i++ {
		if t1.answer[i] != outPut1[i] {
			t.Errorf("Expect %d of %s cell, got %d", t1.answer[i], name[i], outPut1[i])
			countError1++
		}
	}
	if countError1 == 0 {
		fmt.Println("Test GetCellNumber: number of each type of cells is correct!")
	}

	var t2 Test
	var countError2 int
	t2.board = InitializeTissue(5)
	t2.answer = []int{10, 5, 5, 5}

	for i := 0; i < len(t2.board); i++ {
		for j := 0; j < len(t2.board[i]); j++ {
			if (i+j)%5 == 1 {
				t2.board[i][j].state = "Infected"
			} else if (i+j)%5 == 2 {
				t2.board[i][j].state = "Infectious"
			} else if (i+j)%5 == 3 {
				t2.board[i][j].state = "Dead"
			}

		}
	}

	outPut2 := GetCellNumber(t2.board)

	for i := 0; i < len(outPut2); i++ {
		if t2.answer[i] != outPut2[i] {
			t.Errorf("Expect %d of %s cell, got %d", t2.answer[i], name[i], outPut2[i])
			countError2++
		}
	}
	if countError2 == 0 {
		fmt.Println("Test GetCellNumber: number of each type of cells is correct!")
	}

}

func TestCopyBoard(t *testing.T) {
	type Test struct {
		board           Board
		startInfectious OrderedPair
	}

	var t1 Test
	t1.board = InitializeTissue(5)
	t1.startInfectious.x, t1.startInfectious.y = 2, 2
	AssignStart(t1.board, t1.startInfectious, 5)

	outPut1 := CopyBoard(t1.board)

	// Check number of row
	if len(t1.board) != len(outPut1) {
		t.Errorf("Expect %d row of board, got %d", len(t1.board), len(outPut1))
	} else {
		fmt.Println("Test CopyBoard: number of row of board is correct!")
	}

	// Check number of col
	var countErrorCol int
	for i := 0; i < len(t1.board); i++ {
		if len(t1.board[i]) != len(outPut1[i]) {
			t.Errorf("Expect %d col of board in row %d, got %d", len(t1.board[i]), i, len(outPut1[i]))
			countErrorCol++
		}
	}

	if countErrorCol == 0 {
		fmt.Println("Test CopyBoard: number of col of board is correct!")
	}

	// Check state and concVirus of each cell
	var countErrorCell int
	for i := 0; i < len(t1.board); i++ {
		for j := 0; j < len(t1.board[i]); j++ {
			if (i == 2 || i == 3) && j == 2 {
				if outPut1[i][j].state != "Infectious" {
					t.Errorf("Expect cell at %d, %d of board is infectious, but got %s", i, j, outPut1[i][j].state)
					countErrorCell++
				}
				if outPut1[i][j].concVirus != 5 {
					t.Errorf("Expect cell virus concentration at %d, %d of board is %f, but got %f", i, j, t1.board[i][j].concVirus, outPut1[i][j].concVirus)
					countErrorCell++
				}

			} else {
				if t1.board[i][j].state != "Uninfected" {
					t.Errorf("Expect cell at %d, %d of board is uninfected, but got %s", i, j, outPut1[i][j].state)
					countErrorCell++
				}
				if outPut1[i][j].concVirus != 0 {
					t.Errorf("Expect cell virus concentration at %d, %d of board is %f, but got %f", i, j, t1.board[i][j].concVirus, outPut1[i][j].concVirus)
					countErrorCell++
				}
			}
		}
	}

	if countErrorCell == 0 {
		fmt.Println("Test CopyBoard: fields of cells are correct!")
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

func TestGetCoCellNumber(t *testing.T) {
	type Test struct {
		board  Board
		answer []int
	}

	// Coinfection jas six states of cells, so make sure function could correctly count each type
	// Test if function can correctly calculate number of each type of cells
	var t1 Test
	var countError1 int
	t1.board = InitializeTissue(5)
	t1.answer = []int{25, 0, 0, 0, 0, 0, 0}
	name := []string{"Uninfected", "Infected1", "Infectious1", "Infected2", "Infectious2", "Dead1", "Dead2"}
	outPut1 := GetCoCellNumber(t1.board)

	for i := 0; i < len(outPut1); i++ {
		if t1.answer[i] != outPut1[i] {
			t.Errorf("Expect %d of %s cell, got %d", t1.answer[i], name[i], outPut1[i])
			countError1++
		}
	}
	
	if countError1 == 0 {
		fmt.Println("Test GetCoCellNumber: number of each type of cells is correct!")
	}

	var t2 Test
	var countError2 int
	t2.board = InitializeTissue(5)
	t2.answer = []int{3, 3, 4, 3, 5, 4, 3}

	for i := 0; i < len(t2.board); i++ {
		for j := 0; j < len(t2.board[i]); j++ {
			if (i+j)%7 == 1 {
				t2.board[i][j].state = "Infected1"
			} else if (i+j)%7 == 2 {
				t2.board[i][j].state = "Infected2"
			} else if (i+j)%7 == 3 {
				t2.board[i][j].state = "Infectious1"
			} else if (i+j)%7 == 4 {
				t2.board[i][j].state = "Infectious2"
			}  else if (i+j)%7 == 5 {
				t2.board[i][j].state = "Dead1"
			}  else if (i+j)%7 == 6 {
				t2.board[i][j].state = "Dead2"
			}

		}
	}

	outPut2 := GetCoCellNumber(t2.board)

	for i := 0; i < len(outPut2); i++ {
		if t2.answer[i] != outPut2[i] {
			t.Errorf("Expect %d of %s cell, got %d", t2.answer[i], name[i], outPut2[i])
			countError2++
		}
	}

	if countError2 == 0 {
		fmt.Println("Test GetCoCellNumber: number of each type of cells is correct!")
	}
}
