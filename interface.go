package main

import (
	"fmt"
	"gowut/gwu"
	"os"
	"strconv"
)

func OpenWeb() {
	win1 := CreateWindow1()
	win2 := CreateWindow2()
	win3 := CreateWindow3()
	// Create and start a GUI server (omitting error check)
	server := gwu.NewServer("guitest", "localhost:8081")
	server.SetText("Viral Spread - Umbrella Corporation")
	server.AddWin(win1)
	server.AddWin(win2)
	server.AddWin(win3)
	server.Start("") // Also opens windows list in browser
}

// CreateWindow1
// Create window to simulate number of cells vs. Time
func CreateWindow1() gwu.Window {
	// Create and build a window
	win := gwu.NewWindow("win1", "1. Simulate number of cells vs. Time")
	win.Style().SetFullWidth()
	win.SetHAlign(gwu.HACenter)
	win.SetCellPadding(10)

	// Page instruction
	instructPanel := gwu.NewHorizontalPanel()
	instructPanel.Style().SetWidthPx(800)
	instructPanel.Style().SetFontSize("200%").SetFontStyle(gwu.FontWeightBolder)
	instructPanel.Add(gwu.NewLabel("Please enter all parameters:"))
	win.Add(instructPanel)

	// create a vertical panel for general parameters
	generalPanel := gwu.NewVerticalPanel()
	generalPanel.Style().SetWidthPx(800)
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
	modeGroup := gwu.NewRadioGroup("mode")
	rbs := []gwu.RadioButton{gwu.NewRadioButton("assign", modeGroup), gwu.NewRadioButton("random", modeGroup)}
	rbs[0].SetState(true)
	for _, rb := range rbs {
		modePanel.Add(rb)
	}

	modePanel.SetCellPadding(5)
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

	initialPositionPanel.Add(gwu.NewLabel("X position: ")) // get x position
	xPositionTB := gwu.NewTextBox("50")
	xPositionTB.Style().SetWidth("30")
	xPositionTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	initialPositionPanel.Add(xPositionTB)

	initialPositionPanel.Add(gwu.NewLabel("Y position: ")) // get y position
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
	paraCellVer.Style().SetWidthPx(800)
	paraCellVer.Add(gwu.NewLabel("Parameters for cells: "))
	paraCellPanel := gwu.NewHorizontalPanel()

	paraCellPanel.Add(gwu.NewLabel("λ: ")) // lambda
	lambdaTB := gwu.NewTextBox("10000")
	lambdaTB.Style().SetWidth("50")
	lambdaTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraCellPanel.Add(lambdaTB)
	paraCellPanel.Add(gwu.NewLabel("/day, "))

	paraCellPanel.Add(gwu.NewLabel("ω: ")) // omega
	omegaTB := gwu.NewTextBox("0.001")
	omegaTB.Style().SetWidth("50")
	omegaTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraCellPanel.Add(omegaTB)
	paraCellPanel.Add(gwu.NewLabel("/day, "))

	paraCellPanel.Add(gwu.NewLabel("dT: ")) // dT
	dTTB := gwu.NewTextBox("0.02")
	dTTB.Style().SetWidth("50")
	dTTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraCellPanel.Add(dTTB)
	paraCellPanel.Add(gwu.NewLabel("/day, "))

	paraCellPanel.Add(gwu.NewLabel("δ: ")) // delta
	deltaTB := gwu.NewTextBox("0.5")
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
	paraVirusVer.Style().SetWidthPx(800)
	paraVirusVer.Add(gwu.NewLabel("Parameters for virus: "))
	paraVirusPanel := gwu.NewHorizontalPanel()

	paraVirusPanel.Add(gwu.NewLabel("Threshold: ")) // threshold
	thresholdTB := gwu.NewTextBox("300")
	thresholdTB.Style().SetWidth("50")
	thresholdTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraVirusPanel.Add(thresholdTB)
	paraVirusPanel.Add(gwu.NewLabel("unit, "))

	paraVirusPanel.Add(gwu.NewLabel("Rcap: ")) // rCap
	rCapTB := gwu.NewTextBox("500")
	rCapTB.Style().SetWidth("50")
	rCapTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraVirusPanel.Add(rCapTB)
	paraVirusPanel.Add(gwu.NewLabel("unit/day, "))

	paraVirusPanel.Add(gwu.NewLabel("α: ")) // alpha
	alphaTB := gwu.NewTextBox("80")
	alphaTB.Style().SetWidth("50")
	alphaTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraVirusPanel.Add(alphaTB)
	paraVirusPanel.Add(gwu.NewLabel("unit/day, "))

	paraVirusPanel.Add(gwu.NewLabel("γ: ")) // gamma
	gammaTB := gwu.NewTextBox("5")
	gammaTB.Style().SetWidth("50")
	deltaTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraVirusPanel.Add(gammaTB)
	paraVirusPanel.Add(gwu.NewLabel("unit/day, "))

	paraVirusPanel.Add(gwu.NewLabel("ρ: ")) // rho
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
	treatmentVer.Style().SetWidthPx(800)
	treatmentVer.Add(gwu.NewLabel("Parameters for treatment: "))
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

	// create panel for buttons
	btns := gwu.NewHorizontalPanel()

	// add botton to promt simulation
	btn := gwu.NewButton("Submit & Simulate")
	btn.Style().SetWidthPx(190)
	btn.Style().SetHeightPx(40)
	btn.Style().SetFontSize("125%")
	btns.Add(btn)

	// add botton to promt simulation
	terminate := gwu.NewButton("Terminate Program")
	terminate.Style().SetWidthPx(190)
	terminate.Style().SetHeightPx(40)
	terminate.Style().SetFontSize("125%")
	btns.Add(terminate)

	btns.SetCellPadding(10)
	win.Add(btns)
	// get inputs and start simulation
	var allInputs Inputs
	btn.AddEHandlerFunc(func(e gwu.Event) {
		// once click button, pass all inputs and start simulation!

		// takes width
		allInputs.width, _ = strconv.Atoi(widthTB.Text())

		// takes mode
		allInputs.mode = modeGroup.Selected().Text()

		// takes numInfectious
		allInputs.numInfectious, _ = strconv.Atoi(numInfectiousTB.Text())

		// take initialPosition
		allInputs.initialPosition.x, _ = strconv.Atoi(xPositionTB.Text())
		allInputs.initialPosition.y, _ = strconv.Atoi(yPositionTB.Text())

		// takes numGens
		allInputs.numGens, _ = strconv.Atoi(numGensTB.Text())

		// takes timeSteps
		allInputs.timeSteps, _ = strconv.ParseFloat(timeStepsTB.Text(), 64)

		// takes lambda
		allInputs.parameters.lambda, _ = strconv.ParseFloat(lambdaTB.Text(), 64)

		// takes omega
		allInputs.parameters.omega, _ = strconv.ParseFloat(omegaTB.Text(), 64)

		// takes dT
		allInputs.parameters.dT, _ = strconv.ParseFloat(dTTB.Text(), 64)

		// takes delta
		allInputs.parameters.delta, _ = strconv.ParseFloat(deltaTB.Text(), 64)

		// takes threshold
		allInputs.parameters.threshold, _ = strconv.ParseFloat(thresholdTB.Text(), 64)

		// takes rCap
		allInputs.parameters.rCap, _ = strconv.ParseFloat(rCapTB.Text(), 64)

		// takes alpha
		allInputs.parameters.alpha, _ = strconv.ParseFloat(alphaTB.Text(), 64)

		// takes gamma
		allInputs.parameters.gamma, _ = strconv.ParseFloat(gammaTB.Text(), 64)

		// takes rho
		allInputs.parameters.rho, _ = strconv.ParseFloat(rhoTB.Text(), 64)

		// takes treatment
		if blockCellCB.State() && blockVirusCB.State() {
			allInputs.parameters.treatment = "blockboth"
		} else if blockCellCB.State() && !blockVirusCB.State() {
			allInputs.parameters.treatment = "blockcell"
		} else if !blockCellCB.State() && blockVirusCB.State() {
			allInputs.parameters.treatment = "blockvirus"
		} else {
			allInputs.parameters.treatment = "no"
		}

		// takes epsilonCell
		allInputs.parameters.epsilonCell, _ = strconv.ParseFloat(epsilonCellTB.Text(), 64)

		// takes epsilonVirus
		allInputs.parameters.epsilonVirus, _ = strconv.ParseFloat(epsilonVirusTB.Text(), 64)

		// takes imageFrequency
		allInputs.imageFrequency, _ = strconv.Atoi(imageFrequencyTB.Text())

		// simulate GIF
		Simulate(allInputs)
	}, gwu.ETypeClick)

	terminate.AddEHandlerFunc(func(e gwu.Event) {
		// once click button, terminate program
		fmt.Println("Program terminated!")
		os.Exit(5)
	}, gwu.ETypeClick)

	return win
}

// CreateWindow2
// Create window to explore effectiveness of drugs
func CreateWindow2() gwu.Window {
	// Create and build a window
	win := gwu.NewWindow("win2", "2. Explore effectiveness of drugs")
	win.Style().SetFullWidth()
	win.SetHAlign(gwu.HACenter)
	win.SetCellPadding(10)

	// Page instruction
	instructPanel := gwu.NewHorizontalPanel()
	instructPanel.Style().SetWidthPx(800)
	instructPanel.Style().SetFontSize("200%").SetFontStyle(gwu.FontWeightBolder)
	instructPanel.Add(gwu.NewLabel("Please enter all parameters:"))
	win.Add(instructPanel)

	// create a vertical panel for general parameters
	generalPanel := gwu.NewVerticalPanel()
	generalPanel.Style().SetWidthPx(800)
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
	modeGroup := gwu.NewRadioGroup("mode")
	rbs := []gwu.RadioButton{gwu.NewRadioButton("assign", modeGroup), gwu.NewRadioButton("random", modeGroup)}
	rbs[0].SetState(true)
	for _, rb := range rbs {
		modePanel.Add(rb)
	}

	modePanel.SetCellPadding(5)
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

	initialPositionPanel.Add(gwu.NewLabel("X position: ")) // get x position
	xPositionTB := gwu.NewTextBox("50")
	xPositionTB.Style().SetWidth("30")
	xPositionTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	initialPositionPanel.Add(xPositionTB)

	initialPositionPanel.Add(gwu.NewLabel("Y position: ")) // get y position
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
	paraCellVer.Style().SetWidthPx(800)
	paraCellVer.Add(gwu.NewLabel("Parameters for cells: "))
	paraCellPanel := gwu.NewHorizontalPanel()

	paraCellPanel.Add(gwu.NewLabel("λ: ")) // lambda
	lambdaTB := gwu.NewTextBox("10000")
	lambdaTB.Style().SetWidth("50")
	lambdaTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraCellPanel.Add(lambdaTB)
	paraCellPanel.Add(gwu.NewLabel("/day, "))

	paraCellPanel.Add(gwu.NewLabel("ω: ")) // omega
	omegaTB := gwu.NewTextBox("0.001")
	omegaTB.Style().SetWidth("50")
	omegaTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraCellPanel.Add(omegaTB)
	paraCellPanel.Add(gwu.NewLabel("/day, "))

	paraCellPanel.Add(gwu.NewLabel("dT: ")) // dT
	dTTB := gwu.NewTextBox("0.02")
	dTTB.Style().SetWidth("50")
	dTTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraCellPanel.Add(dTTB)
	paraCellPanel.Add(gwu.NewLabel("/day, "))

	paraCellPanel.Add(gwu.NewLabel("δ: ")) // delta
	deltaTB := gwu.NewTextBox("0.5")
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
	paraVirusVer.Style().SetWidthPx(800)
	paraVirusVer.Add(gwu.NewLabel("Parameters for virus: "))
	paraVirusPanel := gwu.NewHorizontalPanel()

	paraVirusPanel.Add(gwu.NewLabel("Threshold: ")) // threshold
	thresholdTB := gwu.NewTextBox("300")
	thresholdTB.Style().SetWidth("50")
	thresholdTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraVirusPanel.Add(thresholdTB)
	paraVirusPanel.Add(gwu.NewLabel("unit, "))

	paraVirusPanel.Add(gwu.NewLabel("Rcap: ")) // rCap
	rCapTB := gwu.NewTextBox("500")
	rCapTB.Style().SetWidth("50")
	rCapTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraVirusPanel.Add(rCapTB)
	paraVirusPanel.Add(gwu.NewLabel("unit/day, "))

	paraVirusPanel.Add(gwu.NewLabel("α: ")) // alpha
	alphaTB := gwu.NewTextBox("80")
	alphaTB.Style().SetWidth("50")
	alphaTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraVirusPanel.Add(alphaTB)
	paraVirusPanel.Add(gwu.NewLabel("unit/day, "))

	paraVirusPanel.Add(gwu.NewLabel("γ: ")) // gamma
	gammaTB := gwu.NewTextBox("5")
	gammaTB.Style().SetWidth("50")
	deltaTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraVirusPanel.Add(gammaTB)
	paraVirusPanel.Add(gwu.NewLabel("unit/day, "))

	paraVirusPanel.Add(gwu.NewLabel("ρ: ")) // rho
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
	treatmentVer.Style().SetWidthPx(800)
	treatmentVer.Add(gwu.NewLabel("Parameters for treatment: "))
	treatmentPanel := gwu.NewHorizontalPanel()

	blockCellCB := gwu.NewCheckBox("Block cell-to-cell transmission") // treatment
	blockVirusCB := gwu.NewCheckBox("Block virus replication")
	treatmentVer.Add(blockCellCB)
	treatmentVer.Add(blockVirusCB)

	treatmentPanel.SetCellPadding(1)

	treatmentVer.Style().SetBorder2(1, gwu.BrdStyleSolid, gwu.ClrBlack)
	treatmentVer.SetCellPadding(10)
	treatmentVer.Add(treatmentPanel)
	win.Add(treatmentVer)

	// create panel for buttons
	btns := gwu.NewHorizontalPanel()

	// add botton to promt simulation
	btn := gwu.NewButton("Submit & Iterate")
	btn.Style().SetWidthPx(190)
	btn.Style().SetHeightPx(40)
	btn.Style().SetFontSize("125%")
	btns.Add(btn)

	// add botton to promt simulation
	terminate := gwu.NewButton("Terminate Program")
	terminate.Style().SetWidthPx(190)
	terminate.Style().SetHeightPx(40)
	terminate.Style().SetFontSize("125%")
	btns.Add(terminate)

	btns.SetCellPadding(10)
	win.Add(btns)
	// get inputs and start simulation
	var allInputs Inputs
	btn.AddEHandlerFunc(func(e gwu.Event) {
		// once click button, pass all inputs and start simulation!

		// takes width
		allInputs.width, _ = strconv.Atoi(widthTB.Text())

		// takes mode
		allInputs.mode = modeGroup.Selected().Text()

		// takes numInfectious
		allInputs.numInfectious, _ = strconv.Atoi(numInfectiousTB.Text())

		// take initialPosition
		allInputs.initialPosition.x, _ = strconv.Atoi(xPositionTB.Text())
		allInputs.initialPosition.y, _ = strconv.Atoi(yPositionTB.Text())

		// takes numGens
		allInputs.numGens, _ = strconv.Atoi(numGensTB.Text())

		// takes timeSteps
		allInputs.timeSteps, _ = strconv.ParseFloat(timeStepsTB.Text(), 64)

		// takes lambda
		allInputs.parameters.lambda, _ = strconv.ParseFloat(lambdaTB.Text(), 64)

		// takes omega
		allInputs.parameters.omega, _ = strconv.ParseFloat(omegaTB.Text(), 64)

		// takes dT
		allInputs.parameters.dT, _ = strconv.ParseFloat(dTTB.Text(), 64)

		// takes delta
		allInputs.parameters.delta, _ = strconv.ParseFloat(deltaTB.Text(), 64)

		// takes threshold
		allInputs.parameters.threshold, _ = strconv.ParseFloat(thresholdTB.Text(), 64)

		// takes rCap
		allInputs.parameters.rCap, _ = strconv.ParseFloat(rCapTB.Text(), 64)

		// takes alpha
		allInputs.parameters.alpha, _ = strconv.ParseFloat(alphaTB.Text(), 64)

		// takes gamma
		allInputs.parameters.gamma, _ = strconv.ParseFloat(gammaTB.Text(), 64)

		// takes rho
		allInputs.parameters.rho, _ = strconv.ParseFloat(rhoTB.Text(), 64)

		// takes treatment
		if blockCellCB.State() && blockVirusCB.State() {
			allInputs.parameters.treatment = "blockboth"
		} else if blockCellCB.State() && !blockVirusCB.State() {
			allInputs.parameters.treatment = "blockcell"
		} else if !blockCellCB.State() && blockVirusCB.State() {
			allInputs.parameters.treatment = "blockvirus"
		} else {
			allInputs.parameters.treatment = "no"
		}

		// takes imageFrequency
		allInputs.imageFrequency, _ = strconv.Atoi(imageFrequencyTB.Text())

		// iterate drug effectiveness
		ExploreEffectiveness(allInputs)
	}, gwu.ETypeClick)

	terminate.AddEHandlerFunc(func(e gwu.Event) {
		// once click button, terminate program
		fmt.Println("Program terminated!")
		os.Exit(5)
	}, gwu.ETypeClick)

	return win
}

// CreateWindow3
// Create window to simulate coinfection
func CreateWindow3() gwu.Window {
	// Create and build a window
	win := gwu.NewWindow("win3", "3. Simulate coinfection")
	win.Style().SetFullWidth()
	win.SetHAlign(gwu.HACenter)
	win.SetCellPadding(10)

	// Page instruction
	instructPanel := gwu.NewHorizontalPanel()
	instructPanel.Style().SetWidthPx(800)
	instructPanel.Style().SetFontSize("200%").SetFontStyle(gwu.FontWeightBolder)
	instructPanel.Add(gwu.NewLabel("Please enter all parameters (coinfection):"))
	win.Add(instructPanel)

	// create a vertical panel for general parameters
	generalPanel := gwu.NewVerticalPanel()
	generalPanel.Style().SetWidthPx(800)
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
	modeGroup := gwu.NewRadioGroup("mode")
	rbs := []gwu.RadioButton{gwu.NewRadioButton("assign", modeGroup), gwu.NewRadioButton("random", modeGroup)}
	rbs[0].SetState(true)
	for _, rb := range rbs {
		modePanel.Add(rb)
	}

	modePanel.SetCellPadding(5)
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

	initialPositionPanel.Add(gwu.NewLabel("X position: ")) // get x position
	xPositionTB := gwu.NewTextBox("50")
	xPositionTB.Style().SetWidth("30")
	xPositionTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	initialPositionPanel.Add(xPositionTB)

	initialPositionPanel.Add(gwu.NewLabel("Y position: ")) // get y position
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
	paraCellVer.Style().SetWidthPx(800)
	paraCellVer.Add(gwu.NewLabel("Parameters for cells: "))
	paraCellPanel := gwu.NewHorizontalPanel()

	paraCellPanel.Add(gwu.NewLabel("λ: ")) // lambda
	lambdaTB := gwu.NewTextBox("10000")
	lambdaTB.Style().SetWidth("50")
	lambdaTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraCellPanel.Add(lambdaTB)
	paraCellPanel.Add(gwu.NewLabel("/day, "))

	paraCellPanel.Add(gwu.NewLabel("ω: ")) // omega
	omegaTB := gwu.NewTextBox("0.001")
	omegaTB.Style().SetWidth("50")
	omegaTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraCellPanel.Add(omegaTB)
	paraCellPanel.Add(gwu.NewLabel("/day, "))

	paraCellPanel.Add(gwu.NewLabel("dT: ")) // dT
	dTTB := gwu.NewTextBox("0.02")
	dTTB.Style().SetWidth("50")
	dTTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraCellPanel.Add(dTTB)
	paraCellPanel.Add(gwu.NewLabel("/day, "))

	paraCellPanel.Add(gwu.NewLabel("δ: ")) // delta
	deltaTB := gwu.NewTextBox("0.5")
	deltaTB.Style().SetWidth("50")
	deltaTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraCellPanel.Add(deltaTB)
	paraCellPanel.Add(gwu.NewLabel("/day"))

	paraCellPanel.SetCellPadding(1)

	paraCellVer.Style().SetBorder2(1, gwu.BrdStyleSolid, gwu.ClrBlack)
	paraCellVer.SetCellPadding(10)
	paraCellVer.Add(paraCellPanel)
	win.Add(paraCellVer)
	// parameters for virus 1
	paraVirusVer1 := gwu.NewVerticalPanel()
	paraVirusVer1.Style().SetWidthPx(800)
	paraVirusVer1.Add(gwu.NewLabel("Parameters for virus 1: "))
	paraVirusPanel1 := gwu.NewHorizontalPanel()

	paraVirusPanel1.Add(gwu.NewLabel("Threshold 1: ")) // threshold1
	thresholdTB1 := gwu.NewTextBox("300")
	thresholdTB1.Style().SetWidth("50")
	thresholdTB1.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraVirusPanel1.Add(thresholdTB1)
	paraVirusPanel1.Add(gwu.NewLabel("unit, "))

	paraVirusPanel1.Add(gwu.NewLabel("Rcap1: ")) // rCap1
	rCapTB1 := gwu.NewTextBox("500")
	rCapTB1.Style().SetWidth("50")
	rCapTB1.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraVirusPanel1.Add(rCapTB1)
	paraVirusPanel1.Add(gwu.NewLabel("unit/day, "))

	paraVirusPanel1.Add(gwu.NewLabel("α1: ")) // alpha1
	alphaTB1 := gwu.NewTextBox("80")
	alphaTB1.Style().SetWidth("50")
	alphaTB1.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraVirusPanel1.Add(alphaTB1)
	paraVirusPanel1.Add(gwu.NewLabel("unit/day, "))

	paraVirusPanel1.Add(gwu.NewLabel("γ1: ")) // gamma1
	gammaTB1 := gwu.NewTextBox("5")
	gammaTB1.Style().SetWidth("50")
	// deltaTB1.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraVirusPanel1.Add(gammaTB1)
	paraVirusPanel1.Add(gwu.NewLabel("unit/day, "))

	paraVirusPanel1.Add(gwu.NewLabel("ρ1: ")) // rho1
	rhoTB1 := gwu.NewTextBox("5")
	rhoTB1.Style().SetWidth("50")
	deltaTB.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraVirusPanel1.Add(rhoTB1)
	paraVirusPanel1.Add(gwu.NewLabel("unit/day"))

	paraVirusPanel1.SetCellPadding(1)

	paraVirusVer1.Style().SetBorder2(1, gwu.BrdStyleSolid, gwu.ClrBlack)
	paraVirusVer1.SetCellPadding(10)
	paraVirusVer1.Add(paraVirusPanel1)
	win.Add(paraVirusVer1)
	// parameters for virus 2
	paraVirusVer2 := gwu.NewVerticalPanel()
	paraVirusVer2.Style().SetWidthPx(800)
	paraVirusVer2.Add(gwu.NewLabel("Parameters for virus 2: "))
	paraVirusPanel2 := gwu.NewHorizontalPanel()

	paraVirusPanel2.Add(gwu.NewLabel("Threshold 2: ")) // threshold2
	thresholdTB2 := gwu.NewTextBox("300")
	thresholdTB2.Style().SetWidth("50")
	thresholdTB2.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraVirusPanel2.Add(thresholdTB2)
	paraVirusPanel2.Add(gwu.NewLabel("unit, "))

	paraVirusPanel2.Add(gwu.NewLabel("Rcap2: ")) // rCap2
	rCapTB2 := gwu.NewTextBox("500")
	rCapTB2.Style().SetWidth("50")
	rCapTB2.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraVirusPanel2.Add(rCapTB2)
	paraVirusPanel2.Add(gwu.NewLabel("unit/day, "))

	paraVirusPanel2.Add(gwu.NewLabel("α2: ")) // alpha2
	alphaTB2 := gwu.NewTextBox("80")
	alphaTB2.Style().SetWidth("50")
	alphaTB2.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraVirusPanel2.Add(alphaTB2)
	paraVirusPanel2.Add(gwu.NewLabel("unit/day, "))

	paraVirusPanel2.Add(gwu.NewLabel("γ2: ")) // gamma2
	gammaTB2 := gwu.NewTextBox("5")
	gammaTB2.Style().SetWidth("50")
	// deltaTB2.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraVirusPanel2.Add(gammaTB2)
	paraVirusPanel2.Add(gwu.NewLabel("unit/day, "))

	paraVirusPanel2.Add(gwu.NewLabel("ρ2: ")) // rho2
	rhoTB2 := gwu.NewTextBox("5")
	rhoTB2.Style().SetWidth("50")
	// deltaTB2.AddSyncOnETypes(gwu.ETypeKeyUp)
	paraVirusPanel2.Add(rhoTB2)
	paraVirusPanel2.Add(gwu.NewLabel("unit/day"))

	paraVirusPanel2.SetCellPadding(1)

	paraVirusVer2.Style().SetBorder2(1, gwu.BrdStyleSolid, gwu.ClrBlack)
	paraVirusVer2.SetCellPadding(10)
	paraVirusVer2.Add(paraVirusPanel2)
	win.Add(paraVirusVer2)
	// parameters for treatment
	treatmentVer := gwu.NewVerticalPanel()
	treatmentVer.Style().SetWidthPx(800)
	treatmentVer.Add(gwu.NewLabel("Parameters for treatment: "))
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

	// create panel for buttons
	btns := gwu.NewHorizontalPanel()

	// add botton to promt simulation
	btn := gwu.NewButton("Submit & Simulate")
	btn.Style().SetWidthPx(190)
	btn.Style().SetHeightPx(40)
	btn.Style().SetFontSize("125%")
	btns.Add(btn)

	// add botton to promt simulation
	terminate := gwu.NewButton("Terminate Program")
	terminate.Style().SetWidthPx(190)
	terminate.Style().SetHeightPx(40)
	terminate.Style().SetFontSize("125%")
	btns.Add(terminate)

	btns.SetCellPadding(10)
	win.Add(btns)
	// get inputs and start simulation
	var allInputs1 Inputs
	var allInputs2 Inputs
	btn.AddEHandlerFunc(func(e gwu.Event) {
		// once click button, pass all inputs and start simulation!

		// takes width
		allInputs1.width, _ = strconv.Atoi(widthTB.Text())
		allInputs2.width, _ = strconv.Atoi(widthTB.Text())

		// takes mode
		allInputs1.mode = modeGroup.Selected().Text()
		allInputs2.mode = modeGroup.Selected().Text()

		// takes numInfectious
		allInputs1.numInfectious, _ = strconv.Atoi(numInfectiousTB.Text())
		allInputs2.numInfectious, _ = strconv.Atoi(numInfectiousTB.Text())

		// take initialPosition
		allInputs1.initialPosition.x, _ = strconv.Atoi(xPositionTB.Text())
		allInputs1.initialPosition.y, _ = strconv.Atoi(yPositionTB.Text())
		allInputs2.initialPosition.x, _ = strconv.Atoi(xPositionTB.Text())
		allInputs2.initialPosition.y, _ = strconv.Atoi(yPositionTB.Text())

		// takes numGens
		allInputs1.numGens, _ = strconv.Atoi(numGensTB.Text())
		allInputs2.numGens, _ = strconv.Atoi(numGensTB.Text())

		// takes timeSteps
		allInputs1.timeSteps, _ = strconv.ParseFloat(timeStepsTB.Text(), 64)
		allInputs2.timeSteps, _ = strconv.ParseFloat(timeStepsTB.Text(), 64)

		// takes lambda
		allInputs1.parameters.lambda, _ = strconv.ParseFloat(lambdaTB.Text(), 64)
		allInputs2.parameters.lambda, _ = strconv.ParseFloat(lambdaTB.Text(), 64)

		// takes omega
		allInputs1.parameters.omega, _ = strconv.ParseFloat(omegaTB.Text(), 64)
		allInputs2.parameters.omega, _ = strconv.ParseFloat(omegaTB.Text(), 64)

		// takes dT
		allInputs1.parameters.dT, _ = strconv.ParseFloat(dTTB.Text(), 64)
		allInputs2.parameters.dT, _ = strconv.ParseFloat(dTTB.Text(), 64)

		// takes delta
		allInputs1.parameters.delta, _ = strconv.ParseFloat(deltaTB.Text(), 64)
		allInputs2.parameters.delta, _ = strconv.ParseFloat(deltaTB.Text(), 64)

		// takes threshold
		allInputs1.parameters.threshold, _ = strconv.ParseFloat(thresholdTB1.Text(), 64)
		allInputs2.parameters.threshold, _ = strconv.ParseFloat(thresholdTB2.Text(), 64)

		// takes rCap
		allInputs1.parameters.rCap, _ = strconv.ParseFloat(rCapTB1.Text(), 64)
		allInputs2.parameters.rCap, _ = strconv.ParseFloat(rCapTB2.Text(), 64)

		// takes alpha
		allInputs1.parameters.alpha, _ = strconv.ParseFloat(alphaTB1.Text(), 64)
		allInputs2.parameters.alpha, _ = strconv.ParseFloat(alphaTB2.Text(), 64)

		// takes gamma
		allInputs1.parameters.gamma, _ = strconv.ParseFloat(gammaTB1.Text(), 64)
		allInputs2.parameters.gamma, _ = strconv.ParseFloat(gammaTB2.Text(), 64)

		// takes rho
		allInputs1.parameters.rho, _ = strconv.ParseFloat(rhoTB1.Text(), 64)
		allInputs2.parameters.rho, _ = strconv.ParseFloat(rhoTB2.Text(), 64)

		// takes treatment
		if blockCellCB.State() && blockVirusCB.State() {
			allInputs1.parameters.treatment = "blockboth"
			allInputs2.parameters.treatment = "blockboth"
		} else if blockCellCB.State() && !blockVirusCB.State() {
			allInputs1.parameters.treatment = "blockcell"
			allInputs2.parameters.treatment = "blockcell"
		} else if !blockCellCB.State() && blockVirusCB.State() {
			allInputs1.parameters.treatment = "blockvirus"
			allInputs2.parameters.treatment = "blockvirus"
		} else {
			allInputs1.parameters.treatment = "no"
			allInputs2.parameters.treatment = "no"
		}

		// takes epsilonCell
		allInputs1.parameters.epsilonCell, _ = strconv.ParseFloat(epsilonCellTB.Text(), 64)
		allInputs2.parameters.epsilonCell, _ = strconv.ParseFloat(epsilonCellTB.Text(), 64)

		// takes epsilonVirus
		allInputs1.parameters.epsilonVirus, _ = strconv.ParseFloat(epsilonVirusTB.Text(), 64)
		allInputs2.parameters.epsilonVirus, _ = strconv.ParseFloat(epsilonVirusTB.Text(), 64)

		// takes imageFrequency
		allInputs1.imageFrequency, _ = strconv.Atoi(imageFrequencyTB.Text())
		allInputs2.imageFrequency, _ = strconv.Atoi(imageFrequencyTB.Text())

		// simulate GIF
		Simulate2(allInputs1, allInputs2)
	}, gwu.ETypeClick)

	terminate.AddEHandlerFunc(func(e gwu.Event) {
		// once click button, terminate program
		fmt.Println("Program terminated!")
		os.Exit(5)
	}, gwu.ETypeClick)

	return win
}
