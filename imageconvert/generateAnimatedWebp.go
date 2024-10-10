package imageconvert

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/davidbyttow/govips/v2/vips"
)

func GenerateAnimatedWebP(inputDir string, outputDir string) {
	// Load images as frames from inputDir
	files, err := os.ReadDir(inputDir)
	if err != nil {
		log.Fatalf("failed to read directory %s: %v", inputDir, err)
	}

	var imageRefs []*vips.ImageRef

	for _, file := range files {
		if !file.IsDir() {
			ext := strings.ToLower(filepath.Ext(file.Name()))
			if ext == ".jpg" || ext == ".png" {
				framePath := filepath.Join(inputDir, file.Name())
				img, err := vips.NewImageFromFile(framePath)
				if err != nil {
					log.Fatalf("failed to load image %s: %v", framePath, err)
				}
				defer img.Close()
				imageRefs = append(imageRefs, img)
			}
		}
	}

	// Join images into a single image
	joinedImage := imageRefs[0] // Start with the first image

	if err != nil {
		log.Fatalf("failed to initialize joined image: %v", err)
	}
	defer joinedImage.Close()

	err = joinedImage.ArrayJoin(imageRefs[1:], 1) // Join images vertically
	if err != nil {
		log.Fatalf("failed to join images: %v", err)
	}

	// Set delay for each frame
	delays := make([]int, len(imageRefs))
	for i := range delays {
		delays[i] = 100
	}
	joinedImage.SetPageHeight(imageRefs[0].Height())
	joinedImage.SetPageDelay(delays)

	// Save the joined image as an animated WebP
	options := vips.WebpExportParams{
		StripMetadata: false,
		Quality:       50,
		Lossless:      false,
	}

	webpData, imageMetadata, err := joinedImage.ExportWebp(&options)
	if err != nil {
		log.Fatalf("failed to export webp: %v", err)
	}
	fmt.Println("Image Metadata pages :", imageMetadata.Pages)
	// Save the animated WebP to file
	outputFilePath := filepath.Join(outputDir, "output.webp")
	err = os.WriteFile(outputFilePath, webpData, 0644)
	if err != nil {
		log.Fatalf("failed to save webp: %v", err)
	}

	fmt.Println("Animated WebP created successfully: output.webp")
}
