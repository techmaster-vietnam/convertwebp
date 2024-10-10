package imageconvert

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/davidbyttow/govips/v2/vips"
)

/*
Kiểm tra đinh dạng inputFile:
- Nếu không phải là ảnh, copy sang outputFile
- Nếu là ảnh:
  - lấy kích thước width, height
  - resize ảnh nếu width > 1024 hoặc height > 768
  - export webp với quality = lossyQuality
  - lưu file webp
*/

func VipsConvertCopy(inputFile string, outputDirPath string, lossyQuality int) error {
	// Lấy tên file từ inputFile
	fileName := filepath.Base(inputFile)
	var outputFile string
	importParams := vips.NewImportParams()
	importParams.NumPages.Set(-1)
	imageRef, err := vips.LoadImageFromFile(inputFile, importParams)
	if err != nil {
		if err == vips.ErrUnsupportedImageFormat {
			outputFile = filepath.Join(outputDirPath, fileName)
			err := copyFile(inputFile, outputFile)
			if err != nil {
				return fmt.Errorf("failed to copy file %s to %s: %v", inputFile, outputFile, err)
			}
			return nil
		} else {
			return err
		}
	}
	defer imageRef.Close()
	//Nếu file đã ở định dạng WEBP hay PDF thì copy luôn, không convert
	if imageRef.Format() == vips.ImageTypeWEBP || imageRef.Format() == vips.ImageTypePDF {
		outputFile = filepath.Join(outputDirPath, fileName)
		err := copyFile(inputFile, outputFile)
		if err != nil {
			return fmt.Errorf("failed to copy file %s to %s: %v", inputFile, outputFile, err)
		}
		return nil
	}

	isAnimatedGif := false
	if imageRef.Format() == vips.ImageTypeGIF {
		// Kiểm tra số frame để xác định ảnh gif động
		isAnimatedGif = imageRef.Pages() > 1
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
	extension := filepath.Ext(fileName)
	if extension == "" {
		outputFile = filepath.Join(outputDirPath, fileName)
	} else {
		outputFile = filepath.Join(outputDirPath, strings.TrimSuffix(fileName, extension)+".webp")
	}
	// Lưu file webp
	if err := os.WriteFile(outputFile, webpData, 0644); err != nil {
		return fmt.Errorf("failed to save webp %s: %v", outputFile, err)
	}

	fmt.Printf("Processed and converted %s\n", inputFile)
	return nil
}
