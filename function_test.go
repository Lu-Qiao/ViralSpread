package main

import (
	"fmt"
	"testing"
)

func TestRandomInfectCell(t *testing.T) {
	
	type Test struct {
		board Board
		startInfection OrderedPair
		answer int
	}

	//// Test if function choose one cell to be infected
	var t1 Test
	var above, below, right, left OrderedPair
	var count1 int

	t1.startInfection.x, t1.startInfection.y = 3,3
	above.x, below.x, right.x, left.x = 3, 3, 4, 2
	above.y, below.y, right.y, left.y = 2, 4, 3, 3
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
	var above1, below1, right1, left1 OrderedPair
	var count2 int
	
	t2.answer = 0
	t2.startInfection.x, t1.startInfection.y = 3,3
	above1.x, below1.x, right1.x, left1.x = 3, 3, 4, 2
	above1.y, below1.y, right1.y, left1.y = 2, 4, 3, 3
	cellAround := []OrderedPair{above, below, right, left}
	
	t2.board = InitializeTissue(5)
	t2.board[3][2].state, t2.board[3][4].state,t2.board[4][3].state = "Infected", "Infected", "Infected"
	t2.board = RandomInfectCell(t1.board, t2.startInfection, cellAround)
	
	for i := 0; i < len(cellAround); i++ {
		if t1.board[cellAround[i].x][cellAround[i].y].state == "Uninfected" {
			count2++
		}
	}
	if t1.answer != count2 {
		t.Errorf("Expect %d uninfected cell around infectious cell, got %d", t1.answer, count2)
	} else {
		fmt.Println("Test RandomInfectCell: the expected cell is infected!")
	}

}


