package imageconvert

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/chai2010/webp"
	"github.com/kolesa-team/go-webp/encoder"
)

// decodeImage decodes an image file based on its format.
func decodeImage(inputFile string) (image.Image, string, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Detect the image format from the file extension
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(inputFile), "."))
	var img image.Image
	var format string

	switch ext {
	case "jpg", "jpeg":
		img, err = jpeg.Decode(file)
		format = "jpeg"
	case "png":
		img, err = png.Decode(file)
		format = "png"
	case "gif":
		img, err = gif.Decode(file)
		format = "gif"
	default:
		return nil, "", fmt.Errorf("unsupported file format: %s", ext)
	}

	if err != nil {
		return nil, "", fmt.Errorf("failed to decode image: %w", err)
	}

	return img, format, nil
}

func KolesaConvert(inputFile string, outputFile string, lossyQuality int) error {
	img, _, err := decodeImage(inputFile)
	if err != nil {
		return fmt.Errorf("error decoding image: %w", err)
	}

	output, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer output.Close()

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, float32(lossyQuality))
	if err != nil {
		return fmt.Errorf("error creating encoder options: %w", err)
	}

	webpOptions := &webp.Options{
		Quality: options.Quality,
	}

	if err := webp.Encode(output, img, webpOptions); err != nil {
		return fmt.Errorf("error encoding to webp: %w", err)
	}

	fmt.Println("Success encoding to webp:", outputFile)
	return nil
}

func ChaiConvert(inputFile string, outputFile string, lossyQuality int) error {
	img, _, err := decodeImage(inputFile)
	if err != nil {
		return fmt.Errorf("error decoding image: %w", err)
	}

	output, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer output.Close()

	err = webp.Encode(output, img, &webp.Options{Lossless: false, Quality: float32(lossyQuality)})
	if err != nil {
		return fmt.Errorf("error encoding to webp: %w", err)
	}

	fmt.Println("Success encoding to webp:", outputFile)
	return nil
}
