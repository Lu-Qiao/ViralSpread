package main

import (
	"fmt"
	"gifhelper"
	"os"
)

func main() {
	// create a channel to get inputs from WebApp
	allInputsChan := make(chan Inputs, 1)

	go OpenWeb(allInputsChan)
	allInputs := <-allInputsChan

	// Copy all parameters from inputs
	width := allInputs.width
	mode := allInputs.mode
	numInfectious := allInputs.numInfectious
	initialPosition := allInputs.initialPosition
	numGens := allInputs.numGens
	timeSteps := allInputs.timeSteps
	imageFrequency := allInputs.imageFrequency
	parameters := allInputs.parameters

	// Initialize tissue
	Tissue := InitializeTissue(width)
	if mode == "random" {
		RandomStart(Tissue, numInfectious, parameters.threshold)
	} else if mode == "assign" {
		AssignStart(Tissue, initialPosition, parameters.threshold)
	}

	fmt.Println("Simulating system.")

	timePoints := SimulateViralSpread(Tissue, numGens, timeSteps, parameters, 0, numInfectious)

	fmt.Println("Viral Spread has been simulated!")
	fmt.Println("Ready to draw images.")

	images := AnimateSystem(timePoints, width, imageFrequency)

	fmt.Println("Images drawn!")

	fmt.Println("Making GIF...")

	// create filename according to inputs
	filename := os.Args[0] + "_" + mode + "_" + parameters.treatment + "_treatment"

	gifhelper.ImagesToGIF(images, filename)

	fmt.Println("Animated GIF produced!")

	fmt.Println("Exiting normally.")
}
