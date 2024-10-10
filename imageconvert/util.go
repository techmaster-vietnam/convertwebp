package imageconvert

import (
	"fmt"
	"io"
	"math"
	"os"

	"github.com/davidbyttow/govips/v2/vips"
)

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// Copy the file content from source to destination
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	// Flush the file to ensure all data is written
	err = destinationFile.Sync()
	if err != nil {
		return err
	}

	return nil
}

func resizeImage(imageRef *vips.ImageRef) error {
	// Lấy kích thước ảnh
	width := imageRef.Width()
	height := imageRef.Height()

	// Kiểm tra và thay đổi kích thước nếu cần
	if width > 1024 || height > 768 {
		scale := math.Min(1024.0/float64(width), 768.0/float64(height))
		if err := imageRef.Resize(scale, vips.KernelAuto); err != nil {
			return fmt.Errorf("failed to resize image: %v", err)
		}
	}
	return nil
}

func LoggingHandlerFunction(messageDomain string, messageLevel vips.LogLevel, message string) {
	fmt.Printf("Domain: %s, Level: %d, Message: %s\n", messageDomain, messageLevel, message)
}
