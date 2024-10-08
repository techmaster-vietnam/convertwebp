package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/davidbyttow/govips/v2/vips"
)

const lossyQuality = 80

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
func ConvertFolderToWebP(inputDir string, outputDir string, convertFunction convertFunctionType) error {
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
			outputFile := filepath.Join(outputDir, strings.TrimSuffix(fileName, filepath.Ext(fileName))+".webp")
			// Gọi hàm chuyển đổi
			if err := convertFunction(path, outputFile, lossyQuality); err != nil {
				fmt.Println(err)
			}
		}
		return nil
	})
}

func init() {
	config := vips.Config{
		ConcurrencyLevel: 2,
		MaxCacheSize:     100,
		CollectStats:     false,
	}
	vips.Startup(&config)
	vips.LoggingSettings(LoggingHandlerFunction, vips.LogLevelCritical)

}

func main() {
	inputDir := filepath.Join(".", "in")
	outputDir := filepath.Join(".", "out")
	//ConvertFolderToWebP(inputDir, outputDir, KolesaConvert)
	//ConvertFolderToWebP(inputDir, outputDir, ChaiConvert)
	//ConvertFolderToWebP(inputDir, outputDir, gm_convert)
	//ConvertFolderToWebP(inputDir, outputDir, cweb_convert)
	ConvertFolderToWebP(inputDir, outputDir, VipsConvert)
	//ConvertAnimatedGifToWebP("/Users/cuong/CODE/techmasterweb/convertwebp/in/car.gif", "/Users/cuong/CODE/techmasterweb/convertwebp/out/car.webp", lossyQuality)
	defer vips.Shutdown()
}
