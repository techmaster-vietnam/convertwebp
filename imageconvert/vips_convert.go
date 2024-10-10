package imageconvert

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/davidbyttow/govips/v2/vips"
)

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
