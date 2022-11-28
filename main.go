package main

import (
	"fmt"
	"gifhelper"
	"os"
)

func main() {
	// // os.Args[0] is program name ("ViralSpread")

	// // os.Args[1] takes width
	// width, err1 := strconv.Atoi(os.Args[1])
	// if err1 != nil {
	// 	panic(err1)
	// }

	// // os.Args[2] takes start mode
	// mode := os.Args[2]

	// // os.Args[3] takes numInfectious argument
	// numInfectious, err3 := strconv.Atoi(os.Args[3])
	// if err3 != nil {
	// 	panic(err3)
	// }

	// // os.Args[4] takes initialPosition
	// // input will be a position along the diagonal
	// initialPosition, err4 := strconv.Atoi(os.Args[4])
	// if err4 != nil {
	// 	panic(err4)
	// }

	// var pos OrderedPair
	// pos.x = initialPosition
	// pos.y = initialPosition

	// // os.Args[5] takes numGens argument
	// numGens, err5 := strconv.Atoi(os.Args[5])
	// if err5 != nil {
	// 	panic(err5)
	// }
	// _ = numGens

	// // os.Args[6] takes timeSteps argument
	// timeSteps, err6 := strconv.ParseFloat(os.Args[6], 64)
	// if err6 != nil {
	// 	panic(err6)
	// }
	// _ = timeSteps

	// // os.Args[7] takes lambda argument
	// lambda, err7 := strconv.ParseFloat(os.Args[7], 64)
	// if err7 != nil {
	// 	panic(err7)
	// }

	// // os.Args[8] takes omega argument
	// omega, err8 := strconv.ParseFloat(os.Args[8], 64)
	// if err8 != nil {
	// 	panic(err8)
	// }

	// // os.Args[9] takes dT argument
	// dT, err9 := strconv.ParseFloat(os.Args[9], 64)
	// if err9 != nil {
	// 	panic(err9)
	// }

	// // os.Args[10] takes delta argument
	// delta, err10 := strconv.ParseFloat(os.Args[10], 64)
	// if err10 != nil {
	// 	panic(err10)
	// }

	// // os.Args[11] takes threshold argument
	// threshold, err11 := strconv.ParseFloat(os.Args[11], 64)
	// if err11 != nil {
	// 	panic(err11)
	// }

	// // os.Args[12] takes rCap argument
	// rCap, err12 := strconv.ParseFloat(os.Args[12], 64)
	// if err12 != nil {
	// 	panic(err12)
	// }

	// // os.Args[13] takes alpha argument
	// alpha, err13 := strconv.ParseFloat(os.Args[13], 64)
	// if err13 != nil {
	// 	panic(err13)
	// }

	// // os.Args[14] takes gamma argument
	// gamma, err14 := strconv.ParseFloat(os.Args[14], 64)
	// if err14 != nil {
	// 	panic(err14)
	// }

	// // os.Args[15] takes rho argument
	// rho, err15 := strconv.ParseFloat(os.Args[15], 64)
	// if err15 != nil {
	// 	panic(err15)
	// }

	// // os.Args[16] takes treatment argument
	// treatment := os.Args[16]

	// // os.Args[17] takes epsilonCell argument
	// epsilonCell, err17 := strconv.ParseFloat(os.Args[17], 64)
	// if err17 != nil {
	// 	panic(err17)
	// }

	// // os.Args[18] takes epsilonVirus argument
	// epsilonVirus, err18 := strconv.ParseFloat(os.Args[18], 64)
	// if err18 != nil {
	// 	panic(err18)
	// }

	// // os.Args[19] takes imageFrequency argument
	// imageFrequency, err19 := strconv.Atoi(os.Args[19])
	// if err19 != nil {
	// 	panic(err19)
	// }

	// ************************************************************************************************
	// ************************************************************************************************
	// temp: test inputs
	width := 1000
	mode := "assign"
	numInfectious := 5
	initialPosition := 500
	numGens := 10000
	timeSteps := 10.0
	lambda := 100.0
	omega := 0.001
	dT := 0.1
	delta := 0.5
	threshold := 50000.0
	rCap := 1000000.0
	alpha := 1.0
	gamma := 1.0
	rho := 1500.0
	treatment := "no"
	epsilonCell := 0.0
	epsilonVirus := 0.0
	imageFrequency := 50

	var pos OrderedPair
	pos.x = initialPosition
	pos.y = initialPosition
	// ************************************************************************************************
	// ************************************************************************************************

	// Copy all parameters from inputs
	var parameters Parameters
	// parameters for cells
	parameters.lambda = lambda
	parameters.omega = omega
	parameters.dT = dT
	parameters.delta = delta
	parameters.threshold = threshold
	// parameters for virus
	parameters.rCap = rCap
	parameters.alpha = alpha
	parameters.gamma = gamma
	parameters.rho = rho
	// parameters for treatments
	parameters.treatment = treatment
	parameters.epsilonCell = epsilonCell
	parameters.epsilonVirus = epsilonVirus

	fmt.Println("Command line arguments read successfully.")

	Tissue := InitializeTissue(width)
	if mode == "random" {
		RandomStart(Tissue, numInfectious, threshold)
	} else if mode == "assign" {
		AssignStart(Tissue, pos, threshold)
	}

	fmt.Println("Simulating system.")

	timePoints := SimulateViralSpread(Tissue, numGens, timeSteps, parameters, 0, numInfectious)

	fmt.Println("Viral Spread has been simulated!")
	fmt.Println("Ready to draw images.")

	images := AnimateSystem(timePoints, width, imageFrequency)

	fmt.Println("Images drawn!")

	fmt.Println("Making GIF...")

	var input string

	for i := range os.Args {
		input += os.Args[i] + "_"
	}

	gifhelper.ImagesToGIF(images, input)

	fmt.Println("Animated GIF produced!")

	fmt.Println("Exiting normally.")
}
