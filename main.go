package main

import (
	"image"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

const SIZE_MAX_WIDTH = 420

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

	// Save the resized image.
	out, err := os.Create("resized.png")
	if err != nil {
		log.Fatalf("failed create file: %s", err)
	}
	defer out.Close()

	err = png.Encode(out, m)
	if err != nil {
		log.Fatalf("failed encode image: %s", err)
	}
}
