package imageconvert

import (
	"fmt"
	"os/exec"
	"strconv"
)

func cweb_convert(inputFile string, outputFile string, lossyQuality int) error {
	cmd := exec.Command("cwebp", "-q", strconv.Itoa(lossyQuality), inputFile, "-o", outputFile)
	// Chạy lệnh
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Error decoding image %s : %w", inputFile, err)
	}

	fmt.Printf("Processed and converted %s to %s\n", inputFile, outputFile)
	return nil
}
