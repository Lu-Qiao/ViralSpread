package main

import (
	"encoding/csv"
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
)

// AnimateSystem takes a slice of Board objects along with a width
// parameter and a frequency parameter. It generates a slice of images corresponding to
// drawing each Board whose index is divisible by the frequency parameter
// on a width x width image
func AnimateSystem(timePoints []Board, width, imageFrequency int) []image.Image {
	images := make([]image.Image, 0)

	for i := range timePoints {
		if i%imageFrequency == 0 { // only draw if current index of tissue
			// is divisible by some parameter frequency
			images = append(images, DrawToImage(timePoints[i], width))
		}
	}

	return images
}

// DrawToImage generates the image corresponding to the states of cells in the currentBoard
func DrawToImage(currentBoard Board, width int) image.Image {
	// create rectangle
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, width}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for i := 0; i < width; i++ {
		for j := 0; j < width; j++ {
			switch {
			// infectious = red
			case currentBoard[i][j].state == "Infectious":
				img.Set(i, j, color.RGBA{255, 0, 0, 0xff})
			// infected = green
			case currentBoard[i][j].state == "Infected":
				img.Set(i, j, color.RGBA{215, 222, 33, 0xff})
			// dead = black
			case currentBoard[i][j].state == "Dead":
				img.Set(i, j, color.RGBA{0, 0, 0, 0xff})
			// uninfected = white
			case currentBoard[i][j].state == "Uninfected":
				img.Set(i, j, color.RGBA{255, 255, 255, 0xff})
			default:
				// Use zero value.
			}
		}
	}

	return img
}

// SaveCellDataToCSV
// Inputs:
func SaveCellDataToCSV(timeSteps float64, cellTimePoints [][]int) {
	// create csv file
	csvFile, err := os.Create("Number of target cells and infectious cells over time.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	// initiate writer
	csvwriter := csv.NewWriter(csvFile)
	// write header to csv file
	header := []string{"Time (day)", "Number of normal cells", "Number of target cells", "Number of infectious cells", "Number of dead cells"}
	_ = csvwriter.Write(header)
	// write data
	for i := range cellTimePoints {
		time := fmt.Sprintf("%f", float64(i)*timeSteps)
		N := fmt.Sprintf("%d", cellTimePoints[i][0])
		T := fmt.Sprintf("%d", cellTimePoints[i][1])
		I := fmt.Sprintf("%d", cellTimePoints[i][2])
		D := fmt.Sprintf("%d", cellTimePoints[i][3])
		row := []string{time, N, T, I, D}
		_ = csvwriter.Write(row)
	}
	csvwriter.Flush()
	csvFile.Close()
}
