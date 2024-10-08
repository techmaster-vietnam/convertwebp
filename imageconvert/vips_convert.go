package imageconvert

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/davidbyttow/govips/v2/vips"
)

func ConvertAnimatedGifToWebP(inputFile string, outputFile string, lossyQuality int) error {
	importParams := vips.NewImportParams()
	importParams.NumPages.Set(-1)
	imageRef, err := vips.LoadImageFromFile(inputFile, importParams)
	if err != nil {
		return err
	}
	defer imageRef.Close()

	// Cấu hình export webp
	webpParams := vips.WebpExportParams{
		StripMetadata: false,
		Lossless:      false,
		Quality:       lossyQuality,
	}
	webpData, imageMetadata, err := imageRef.ExportWebp(&webpParams)
	if err != nil {
		return fmt.Errorf("failed to export webp %s: %v", inputFile, err)
	}
	fmt.Println("Image Metadata pages :", imageMetadata.Pages)
	// Lưu file webp
	if err := os.WriteFile(outputFile, webpData, 0644); err != nil {
		return fmt.Errorf("failed to save webp %s: %v", outputFile, err)
	}

	fmt.Printf("Processed and converted %s\n", inputFile)
	return nil
}

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
func resizeImage(imageRef *vips.ImageRef) error {
	// Lấy kích thước ảnh
	width := imageRef.Width()
	height := imageRef.Height()

	// Kiểm tra và thay đổi kích thước nếu cần
	if width > 1024 || height > 768 {
		scale := 1.0
		if width > 1024 {
			scale = 1024.0 / float64(width)
		}
		if height > 768 {
			scale = 768.0 / float64(height)
		}
		if err := imageRef.Resize(scale, vips.KernelAuto); err != nil {
			return fmt.Errorf("failed to resize image: %v", err)
		}
	}
	return nil
}
func VipsConvert(inputFile string, outputFile string, lossyQuality int) error {
	// Lấy extension của file
	ext := strings.ToLower(filepath.Ext(inputFile))
	importParams := vips.NewImportParams()
	importParams.NumPages.Set(-1)
	imageRef, err := vips.LoadImageFromFile(inputFile, importParams)
	if err != nil {
		return err
	}
	defer imageRef.Close()
	isAnimatedGif := false
	if ext == ".gif" {
		// Kiểm tra số frame để xác định ảnh gif động
		if imageRef.Pages() > 1 {
			isAnimatedGif = true
		}
	}
	//Chỉ resize ảnh nếu không phải là animated gif
	if !isAnimatedGif {
		// Gọi hàm resizeImage
		if err := resizeImage(imageRef); err != nil {
			return fmt.Errorf("failed to resize image %s: %v", inputFile, err)
		}
	}

	// Cấu hình export webp
	webpParams := vips.WebpExportParams{
		StripMetadata: true,
		Lossless:      false,
		Quality:       lossyQuality,
	}
	webpData, _, err := imageRef.ExportWebp(&webpParams)
	if err != nil {
		return fmt.Errorf("failed to export webp %s: %v", inputFile, err)
	}

	// Lưu file webp
	if err := os.WriteFile(outputFile, webpData, 0644); err != nil {
		return fmt.Errorf("failed to save webp %s: %v", outputFile, err)
	}

	fmt.Printf("Processed and converted %s\n", inputFile)
	return nil
}

func LoggingHandlerFunction(messageDomain string, messageLevel vips.LogLevel, message string) {
	fmt.Printf("Domain: %s, Level: %d, Message: %s\n", messageDomain, messageLevel, message)
}
