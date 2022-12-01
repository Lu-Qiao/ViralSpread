// Copyright (C) 2013 Andras Belicza. All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// A GWU example application with a single public window (no sessions).

package main

import (
	"fmt"
	"gowut/gwu"
)

type myButtonHandler struct {
	counter int
	text    string
}

func (h *myButtonHandler) HandleEvent(e gwu.Event) {
	if b, isButton := e.Src().(gwu.Button); isButton {
		b.SetText(b.Text() + h.text)
		h.counter++
		b.SetToolTip(fmt.Sprintf("You've clicked %d times!", h.counter))
		e.MarkDirty(b)
	}
}

func OpenWeb() {
	// Create and build a window
	win := gwu.NewWindow("main", "Test GUI Window")
	win.Style().SetFullWidth()
	win.SetHAlign(gwu.HACenter)
	win.SetCellPadding(10)

	// Page instruction
	win.Add(gwu.NewLabel("Please enter all parameters:"))

	// // ListBox examples
	// p := gwu.NewHorizontalPanel()
	// p.Style().SetBorder2(1, gwu.BrdStyleSolid, gwu.ClrBlack)
	// p.SetCellPadding(2)
	// p.Add(gwu.NewLabel("A drop-down list being"))
	// widelb := gwu.NewListBox([]string{"50", "100", "150", "200", "250"})
	// widelb.Style().SetWidth("50")
	// widelb.AddEHandlerFunc(func(e gwu.Event) {
	// 	widelb.Style().SetWidth(widelb.SelectedValue() + "px")
	// 	e.MarkDirty(widelb)
	// }, gwu.ETypeChange)
	// p.Add(widelb)
	// p.Add(gwu.NewLabel("pixel wide. And a multi-select list:"))
	// listBox := gwu.NewListBox([]string{"First", "Second", "Third", "Forth", "Fifth", "Sixth"})
	// listBox.SetMulti(true)
	// listBox.SetRows(4)
	// p.Add(listBox)
	// countLabel := gwu.NewLabel("Selected count: 0")
	// listBox.AddEHandlerFunc(func(e gwu.Event) {
	// 	countLabel.SetText(fmt.Sprintf("Selected count: %d", len(listBox.SelectedIndices())))
	// 	e.MarkDirty(countLabel)
	// }, gwu.ETypeChange)
	// p.Add(countLabel)
	// win.Add(p)

	// // Self-color changer check box
	// greencb := gwu.NewCheckBox("I'm a check box. When checked, I'm green!")
	// greencb.AddEHandlerFunc(func(e gwu.Event) {
	// 	if greencb.State() {
	// 		greencb.Style().SetBackground(gwu.ClrGreen)
	// 	} else {
	// 		greencb.Style().SetBackground("")
	// 	}
	// 	e.MarkDirty(greencb)
	// }, gwu.ETypeClick)
	// win.Add(greencb)

	// create a vertical panel for general parameters
	generalPanel := gwu.NewVerticalPanel()
	generalPanel.Add(gwu.NewLabel("General parameters: "))
	// width
	widthPanel := gwu.NewHorizontalPanel()
	widthPanel.Add(gwu.NewLabel("Width: "))
	widthTB := gwu.NewTextBox("100")
	widthTB.Style().SetWidth("50")
	widthTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	widthPanel.Add(widthTB)
	widthPanel.Add(gwu.NewLabel("pixel"))

	widthPanel.SetCellPadding(5)
	generalPanel.Add(widthPanel)
	// mode
	modePanel := gwu.NewHorizontalPanel()
	modePanel.Add(gwu.NewLabel("Mode: "))
	// modeGroup := gwu.NewRadioGroup()
	// modePanel.Add(gwu.NewRadioButton())
	generalPanel.Add(modePanel)
	// numInfectious
	numInfectiousPanel := gwu.NewHorizontalPanel()
	numInfectiousPanel.Add(gwu.NewLabel("Number of infectious points: "))
	numInfectiousTB := gwu.NewTextBox("2")
	numInfectiousTB.Style().SetWidth("50")
	numInfectiousTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	numInfectiousPanel.Add(numInfectiousTB)
	numInfectiousPanel.Add(gwu.NewLabel("locations"))

	numInfectiousPanel.SetCellPadding(5)
	generalPanel.Add(numInfectiousPanel)
	// initialPosition
	initialPositionPanel := gwu.NewHorizontalPanel()
	initialPositionPanel.Add(gwu.NewLabel("Initial position: "))

	initialPositionPanel.Add(gwu.NewLabel("X position")) // get x position
	xPositionTB := gwu.NewTextBox("50")
	xPositionTB.Style().SetWidth("30")
	xPositionTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	initialPositionPanel.Add(xPositionTB)

	initialPositionPanel.Add(gwu.NewLabel("Y position")) // get y position
	yPositionTB := gwu.NewTextBox("50")
	yPositionTB.Style().SetWidth("30")
	yPositionTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	initialPositionPanel.Add(yPositionTB)

	initialPositionPanel.SetCellPadding(5)
	generalPanel.Add(initialPositionPanel)
	// numGens
	numGensPanel := gwu.NewHorizontalPanel()
	numGensPanel.Add(gwu.NewLabel("Number of generations: "))
	numGensTB := gwu.NewTextBox("300")
	numGensTB.Style().SetWidth("50")
	numGensTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	numGensPanel.Add(numGensTB)
	numGensPanel.Add(gwu.NewLabel("generations"))

	numGensPanel.SetCellPadding(5)
	generalPanel.Add(numGensPanel)
	// timeSteps
	timeStepsPanel := gwu.NewHorizontalPanel()
	timeStepsPanel.Add(gwu.NewLabel("Time steps: "))
	timeStepsTB := gwu.NewTextBox("0.1")
	timeStepsTB.Style().SetWidth("50")
	timeStepsTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	timeStepsPanel.Add(timeStepsTB)
	timeStepsPanel.Add(gwu.NewLabel("day"))

	timeStepsPanel.SetCellPadding(5)
	generalPanel.Add(timeStepsPanel)
	// imageFrequency
	imageFrequencyPanel := gwu.NewHorizontalPanel()
	imageFrequencyPanel.Add(gwu.NewLabel("Image frequency: "))
	imageFrequencyTB := gwu.NewTextBox("1")
	imageFrequencyTB.Style().SetWidth("50")
	imageFrequencyTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	imageFrequencyPanel.Add(imageFrequencyTB)

	imageFrequencyPanel.SetCellPadding(5)
	generalPanel.Add(imageFrequencyPanel)

	generalPanel.Style().SetBorder2(1, gwu.BrdStyleSolid, gwu.ClrBlack)
	generalPanel.SetCellPadding(10)
	win.Add(generalPanel)
	// parameters for cells
	paraCellVer := gwu.NewVerticalPanel()
	paraCellVer.Add(gwu.NewLabel("Parameters for cells:"))
	paraCellPanel := gwu.NewHorizontalPanel()

	paraCellPanel.Add(gwu.NewLabel("lambda")) // lambda
	lambdaTB := gwu.NewTextBox("10000")
	lambdaTB.Style().SetWidth("50")
	lambdaTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraCellPanel.Add(lambdaTB)
	paraCellPanel.Add(gwu.NewLabel("/day, "))

	paraCellPanel.Add(gwu.NewLabel("omega")) // omega
	omegaTB := gwu.NewTextBox("0.001")
	omegaTB.Style().SetWidth("50")
	omegaTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraCellPanel.Add(omegaTB)
	paraCellPanel.Add(gwu.NewLabel("/day, "))

	paraCellPanel.Add(gwu.NewLabel("dT")) // dT
	dTTB := gwu.NewTextBox("0.02")
	dTTB.Style().SetWidth("50")
	dTTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraCellPanel.Add(dTTB)
	paraCellPanel.Add(gwu.NewLabel("/day, "))

	paraCellPanel.Add(gwu.NewLabel("delta")) // delta
	deltaTB := gwu.NewTextBox("5")
	deltaTB.Style().SetWidth("50")
	deltaTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraCellPanel.Add(deltaTB)
	paraCellPanel.Add(gwu.NewLabel("/day"))

	paraCellPanel.SetCellPadding(1)

	paraCellVer.Style().SetBorder2(1, gwu.BrdStyleSolid, gwu.ClrBlack)
	paraCellVer.SetCellPadding(10)
	paraCellVer.Add(paraCellPanel)
	win.Add(paraCellVer)
	// parameters for virus
	paraVirusVer := gwu.NewVerticalPanel()
	paraVirusVer.Add(gwu.NewLabel("Parameters for virus:"))
	paraVirusPanel := gwu.NewHorizontalPanel()

	paraVirusPanel.Add(gwu.NewLabel("threshold")) // threshold
	thresholdTB := gwu.NewTextBox("300")
	thresholdTB.Style().SetWidth("50")
	thresholdTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraVirusPanel.Add(thresholdTB)
	paraVirusPanel.Add(gwu.NewLabel("unit, "))

	paraVirusPanel.Add(gwu.NewLabel("rCap")) // rCap
	rCapTB := gwu.NewTextBox("500")
	rCapTB.Style().SetWidth("50")
	rCapTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraVirusPanel.Add(rCapTB)
	paraVirusPanel.Add(gwu.NewLabel("unit/day, "))

	paraVirusPanel.Add(gwu.NewLabel("alpha")) // alpha
	alphaTB := gwu.NewTextBox("0.02")
	alphaTB.Style().SetWidth("50")
	alphaTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraVirusPanel.Add(alphaTB)
	paraVirusPanel.Add(gwu.NewLabel("unit/day, "))

	paraVirusPanel.Add(gwu.NewLabel("gamma")) // gamma
	gammaTB := gwu.NewTextBox("5")
	gammaTB.Style().SetWidth("50")
	deltaTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraVirusPanel.Add(gammaTB)
	paraVirusPanel.Add(gwu.NewLabel("unit/day, "))

	paraVirusPanel.Add(gwu.NewLabel("rho")) // rho
	rhoTB := gwu.NewTextBox("5")
	rhoTB.Style().SetWidth("50")
	deltaTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraVirusPanel.Add(rhoTB)
	paraVirusPanel.Add(gwu.NewLabel("unit/day"))

	paraVirusPanel.SetCellPadding(1)

	paraVirusVer.Style().SetBorder2(1, gwu.BrdStyleSolid, gwu.ClrBlack)
	paraVirusVer.SetCellPadding(10)
	paraVirusVer.Add(paraVirusPanel)
	win.Add(paraVirusVer)
	// parameters for treatment
	treatmentVer := gwu.NewVerticalPanel()
	treatmentVer.Add(gwu.NewLabel("Parameters for treatment:"))
	treatmentPanel := gwu.NewHorizontalPanel()

	blockCellCB := gwu.NewCheckBox("Block cell-to-cell transmission") // treatment
	blockVirusCB := gwu.NewCheckBox("Block virus replication")
	treatmentVer.Add(blockCellCB)
	treatmentVer.Add(blockVirusCB)

	blockcell := gwu.NewHorizontalPanel() // epsilonCell
	blockcell.Add(gwu.NewLabel("Effectiveness of blocking cell-to-cell transmistion:"))
	epsilonCellTB := gwu.NewTextBox("0.8")
	epsilonCellTB.Style().SetWidth("50")
	epsilonCellTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	blockcell.Add(epsilonCellTB)
	treatmentVer.Add(blockcell)

	blockvirus := gwu.NewHorizontalPanel()
	blockvirus.Add(gwu.NewLabel("Effectiveness of blocking virus replication:")) // epsilonVirus
	epsilonVirusTB := gwu.NewTextBox("0.45")
	epsilonVirusTB.Style().SetWidth("50")
	epsilonVirusTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	blockvirus.Add(epsilonVirusTB)
	treatmentPanel.Add(blockvirus)

	treatmentPanel.SetCellPadding(1)

	treatmentVer.Style().SetBorder2(1, gwu.BrdStyleSolid, gwu.ClrBlack)
	treatmentVer.SetCellPadding(10)
	treatmentVer.Add(treatmentPanel)
	win.Add(treatmentVer)

	// p.Add(gwu.NewLabel("You entered:"))
	// nameLabel := gwu.NewLabel("")
	// nameLabel.Style().SetColor(gwu.ClrRed)
	// tb.AddEHandlerFunc(func(e gwu.Event) {
	// 	nameLabel.SetText(tb.Text())
	// 	e.MarkDirty(nameLabel)
	// }, gwu.ETypeChange, gwu.ETypeKeyUp)
	// p.Add(nameLabel)

	btn := gwu.NewButton("Submit & Simulate")
	btn.AddEHandler(&myButtonHandler{text: ":-)"}, gwu.ETypeClick)
	win.Add(btn)
	btnsPanel := gwu.NewNaturalPanel()
	btn.AddEHandlerFunc(func(e gwu.Event) {
		// Create and add a new button...
		newbtn := gwu.NewButton(fmt.Sprintf("Extra #%d", btnsPanel.CompsCount()))
		newbtn.AddEHandlerFunc(func(e gwu.Event) {
			btnsPanel.Remove(newbtn) // ...which removes itself when clicked
			e.MarkDirty(btnsPanel)
		}, gwu.ETypeClick)
		btnsPanel.Insert(newbtn, 0)
		e.MarkDirty(btnsPanel)
	}, gwu.ETypeClick)
	win.Add(btnsPanel)

	// Create and start a GUI server (omitting error check)
	server := gwu.NewServer("guitest", "localhost:8081")
	server.SetText("Viral Spread - Umbrella Corporation")
	server.AddWin(win)
	server.Start("") // Also opens windows list in browser
}
