package main

import (
	"os"
	"strconv"
)

func main() {
	// os.Args[0] is program name ("ViralSpread")

	// os.Args[1] takes numInfectious argument
	numInfectious, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil {
		panic(err1)
	}
	_ = numInfectious

	// os.Args[2] takes numGens argument
	numGens, err2 := strconv.Atoi(os.Args[2])
	if err2 != nil {
		panic(err2)
	}
	_ = numGens

	// os.Args[3] takes timeSteps argument
	timeSteps, err3 := strconv.ParseFloat(os.Args[3], 64)
	if err3 != nil {
		panic(err3)
	}
	_ = timeSteps

	// os.Args[4] takes lambda argument
	lambda, err4 := strconv.ParseFloat(os.Args[4], 64)
	if err4 != nil {
		panic(err4)
	}
	_ = lambda

	// os.Args[5] takes omega argument
	omega, err5 := strconv.ParseFloat(os.Args[5], 64)
	if err5 != nil {
		panic(err5)
	}
	_ = omega

	// os.Args[6] takes dT argument
	dT, err6 := strconv.ParseFloat(os.Args[6], 64)
	if err6 != nil {
		panic(err6)
	}
	_ = dT

	// os.Args[7] takes delta argument
	delta, err7 := strconv.ParseFloat(os.Args[7], 64)
	if err7 != nil {
		panic(err7)
	}
	_ = delta

	// os.Args[8] takes threshold argument
	threshold, err8 := strconv.ParseFloat(os.Args[8], 64)
	if err8 != nil {
		panic(err8)
	}
	_ = threshold
}
