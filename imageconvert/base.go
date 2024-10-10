package imageconvert

import (
	"fmt"
	"os"
	"path/filepath"
)

// convertFunctionType định nghĩa kiểu của hàm chuyển đổi
type convertFunctionType func(inputFile string, outputFile string, lossyQuality int) error
type convertCopyFunctionType func(inputFile string, outputDirPath string, lossyQuality int) error

/*
Move all media files to outputDir
Handle when media files do not have extension, check if it is graphic file, get width and height then convert to WebP
If media file is PDF just copy it to outputDir
*/
func ConvertFolderToWebP(
	inputDir string,
	outputDir string,
	convertFunction convertCopyFunctionType,
	lossyQuality int) error {

	return filepath.Walk(inputDir, func(inputFile string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Kiểm tra nếu là file
		if !info.IsDir() {
			// Tạo đường dẫn output tương ứng
			relPath, err := filepath.Rel(inputDir, inputFile)
			if err != nil {
				return err
			}
			outputPath := filepath.Join(outputDir, relPath)
			outputDirPath := filepath.Dir(outputPath)

			// Tạo thư mục output nếu chưa tồn tại
			if err := os.MkdirAll(outputDirPath, os.ModePerm); err != nil {
				return err
			}

			if err := convertFunction(inputFile, outputDirPath, lossyQuality); err != nil {
				fmt.Println(err)
			}
		}
		return nil
	})
}
