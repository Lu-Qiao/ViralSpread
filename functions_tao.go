package main

import (
	"canvas"
	"image"
	"image/color"
)

// AnimateSystem takes a slice of Tissue objects along with a canvas width
// parameter and a frequency parameter. It generates a slice of images corresponding to
// drawing each Tissue whose index is divisible by the frequency parameter
// on a canvasWidth x canvasWidth canvas
func AnimateSystem(timePoints []Tissue, canvasWidth, imageFrequency int) []image.Image {
	images := make([]image.Image, 0)

	for i := range timePoints {
		if i%imageFrequency == 0 { // only draw if current index of sky
			// is divisible by some parameter frequency
			images = append(images, DrawToCanvas(timePoints[i], canvasWidth))
		}
	}

	return images
}

// DrawToCanvas generates the image corresponding to a canvas after drawing a Sky
// object's boids on a square canvas that is canvasWidth pixels x canvasWidth pixels
func DrawToCanvas(currentTissue Tissue, canvasWidth int) image.Image {
	// set a new square canvas
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	for i := 0; i < canvasWidth; i++ {
		for j := 0; j < canvasWidth; j++ {
			switch {
			case currentTissue[i][j].state == "Infectious":
				c.img.Set(i, j, color.RGBA{255, 0, 0, 0xff})
			case currentTissue[i][j].state == "Infected":
				c.img.Set(i, j, color.RGBA{215, 222, 33, 0xff})
			case currentTissue[i][j].state == "Dead":
				c.img.Set(i, j, color.RGBA{0, 0, 0, 0xff})
			case currentTissue[i][j].state == "Uninfected":
				c.img.Set(i, j, color.RGBA{255, 255, 255, 0xff})
			default:
				// Use zero value.
			}
		}
	}

	// we want to return an image!
	return c.GetImage()
}
