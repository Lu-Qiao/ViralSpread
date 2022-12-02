package main

import (
	"gowut/gwu"
	"strconv"
)

func OpenWeb(allInputsChan chan Inputs) {
	// Create and build a window
	win := gwu.NewWindow("main", "Inputs")
	win.Style().SetFullWidth()
	win.SetHAlign(gwu.HACenter)
	win.SetCellPadding(10)

	// Page instruction
	win.Add(gwu.NewLabel("Please enter all parameters:"))

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
	alphaTB := gwu.NewTextBox("80")
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

	// add botton to promt simulation
	btn := gwu.NewButton("Submit & Simulate")
	win.Add(btn)
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

		// insert inputs into channel
		allInputsChan <- allInputs
	}, gwu.ETypeClick)

	// Create and start a GUI server (omitting error check)
	server := gwu.NewServer("guitest", "localhost:8081")
	server.SetText("Viral Spread - Umbrella Corporation")
	server.AddWin(win)
	server.Start("") // Also opens windows list in browser
}