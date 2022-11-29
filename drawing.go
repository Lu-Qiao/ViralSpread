package main

import (
	"image"
	"image/color"
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

	// // from image package documentation
	// // Create a png file
	// f, err := os.Create("initial.png")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := png.Encode(f, img); err != nil {
	// 	f.Close()
	// 	log.Fatal(err)
	// }

	// if err := f.Close(); err != nil {
	// 	log.Fatal(err)
	// }

	// we want to return an image!
	return img
}
