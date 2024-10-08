package imageconvert

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// convertFunctionType định nghĩa kiểu của hàm chuyển đổi
type convertFunctionType func(inputFile string, outputFile string, lossyQuality int) error

func getOutputPath(inputFile, outputDir string) string {
	ext := filepath.Ext(inputFile)
	return filepath.Join(outputDir, strings.TrimSuffix(filepath.Base(inputFile), ext)+".webp")
}

// isImageFile kiểm tra nếu file có phần mở rộng là ảnh
func isImageFile(filePath string) bool {
	ext := strings.ToLower(filepath.Ext(filePath))
	return ext == ".gif" || ext == ".jpeg" || ext == ".jpg" || ext == ".png"
}

// ConvertFolderToWebP quét tất cả các file ảnh trong inputDir và chuyển đổi chúng sang WebP
func ConvertFolderToWebP(
	inputDir string,
	outputDir string,
	convertFunction convertFunctionType,
	lossyQuality int) error {

	return filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Kiểm tra nếu là file và có phần mở rộng là ảnh
		if !info.IsDir() && isImageFile(path) {
			// Tạo đường dẫn output tương ứng
			relPath, err := filepath.Rel(inputDir, path)
			if err != nil {
				return err
			}
			outputPath := filepath.Join(outputDir, relPath)
			outputDirPath := filepath.Dir(outputPath)

			// Tạo thư mục output nếu chưa tồn tại
			if err := os.MkdirAll(outputDirPath, os.ModePerm); err != nil {
				return err
			}
			// Lấy tên file từ inputFile
			fileName := filepath.Base(path)
			// Tạo đường dẫn outputFile mới
			outputFile := filepath.Join(outputDirPath, strings.TrimSuffix(fileName, filepath.Ext(fileName))+".webp")
			// Gọi hàm chuyển đổi
			if err := convertFunction(path, outputFile, lossyQuality); err != nil {
				fmt.Println(err)
			}
		}
		return nil
	})
}
