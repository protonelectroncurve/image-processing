package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

const SIZE_MAX_WIDTH = 420

func convertToGrayscale(img image.Image) *image.Gray {
	gray := image.NewGray(img.Bounds())
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			oldColor := img.At(x, y)
			r, g, b, _ := oldColor.RGBA()
			// Convert to grayscale using the standard conversion formula.
			grayVal := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			gray.Set(x, y, color.Gray{Y: uint8(grayVal / 256)})
		}
	}
	return gray
}

func main() {
	catFile, err := os.Open("original.png")
	if err != nil {
		log.Fatalf("failed open file: %s", err)
	}
	defer catFile.Close()

	img, _, err := image.Decode(catFile)
	if err != nil {
		log.Fatalf("failed decode image: %s", err)
	}

	// resize image
	// Calculate new dimensions.
	newWidth := int(SIZE_MAX_WIDTH)
	newHeight := (int(SIZE_MAX_WIDTH) * img.Bounds().Dy()) / img.Bounds().Dx()

	// Resize the image.
	m := resize.Resize(uint(newWidth), uint(newHeight), img, resize.Lanczos3)

	// Convert to grayscale.
	grayImg := convertToGrayscale(m)

	// Save the resized image.
	out, err := os.Create("resized.png")
	if err != nil {
		log.Fatalf("failed create file: %s", err)
	}
	defer out.Close()

	err = png.Encode(out, grayImg)
	if err != nil {
		log.Fatalf("failed encode image: %s", err)
	}
}
