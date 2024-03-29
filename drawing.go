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
// currentBoard (Board) and width of the image
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
			case currentBoard[i][j].state == "Dead" || currentBoard[i][j].state == "Dead1" || currentBoard[i][j].state == "Dead2":
				img.Set(i, j, color.RGBA{0, 0, 0, 0xff})
			// uninfected = white
			case currentBoard[i][j].state == "Uninfected":
				img.Set(i, j, color.RGBA{255, 255, 255, 0xff})

			case currentBoard[i][j].state == "Infectious1":
				img.Set(i, j, color.RGBA{255, 0, 0, 0xff})
				// infected = green
			case currentBoard[i][j].state == "Infected1":
				img.Set(i, j, color.RGBA{245, 194, 66, 0xff})

			case currentBoard[i][j].state == "Infectious2":
				img.Set(i, j, color.RGBA{82, 9, 219, 0xff})
				// infected = green
			case currentBoard[i][j].state == "Infected2":
				img.Set(i, j, color.RGBA{130, 237, 189, 0xff})

			default:
				// Use zero value.
			}
		}
	}

	return img
}

// SaveCellDataToCSV
// Inputs: timeSteps (float64), cell counts over time, and filename (string)
func SaveCellDataToCSV(timeSteps float64, cellTimePoints [][]int, filename string) {
	// create csv file
	csvFile, err := os.Create(filename + ".out.csv")
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

// SaveCoCellDataToCSV
// Inputs: timeSteps (float64), cell counts over time, and filename (string)
func SaveCoCellDataToCSV(timeSteps float64, cellTimePoints [][]int, filename string) {
	// create csv file
	csvFile, err := os.Create(filename + "_coinfection.out.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	// initiate writer
	csvwriter := csv.NewWriter(csvFile)
	// write header to csv file
	header := []string{"Time (day)", "Number of normal cells", "Number of target cells1", "Number of infectious cells1", "Number of target cells2",
		"Number of infectious cells2", "Number of dead cells caused by virus 1", "Number of dead cells caused by virus 2"}
	_ = csvwriter.Write(header)
	// write data
	for i := range cellTimePoints {
		time := fmt.Sprintf("%f", float64(i)*timeSteps)
		N := fmt.Sprintf("%d", cellTimePoints[i][0])
		T1 := fmt.Sprintf("%d", cellTimePoints[i][1])
		I1 := fmt.Sprintf("%d", cellTimePoints[i][2])
		T2 := fmt.Sprintf("%d", cellTimePoints[i][3])
		I2 := fmt.Sprintf("%d", cellTimePoints[i][4])
		D1 := fmt.Sprintf("%d", cellTimePoints[i][5])
		D2 := fmt.Sprintf("%d", cellTimePoints[i][6])
		row := []string{time, N, T1, I1, T2, I2, D1, D2}
		_ = csvwriter.Write(row)
	}
	csvwriter.Flush()
	csvFile.Close()
}

// SaveEffectivenessDataToCSV()
// Inputs: cell count after 30 days of infection for different drug effectiveness
func SaveEffectivenessDataToCSV(finalCell [][]int) {
	// create csv file
	csvFile, err := os.Create("effectiveness.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	// initiate writer
	csvwriter := csv.NewWriter(csvFile)
	// write header to csv file
	header := []string{"Effectiveness", "Number of normal cells", "Number of target cells", "Number of infectious cells", "Number of dead cells"}
	_ = csvwriter.Write(header)
	// write data
	for i := range finalCell {
		time := fmt.Sprintf("%f", float64(i)/100)
		N := fmt.Sprintf("%d", finalCell[i][0])
		T := fmt.Sprintf("%d", finalCell[i][1])
		I := fmt.Sprintf("%d", finalCell[i][2])
		D := fmt.Sprintf("%d", finalCell[i][3])
		row := []string{time, N, T, I, D}
		_ = csvwriter.Write(row)
	}
	csvwriter.Flush()
	csvFile.Close()
}
