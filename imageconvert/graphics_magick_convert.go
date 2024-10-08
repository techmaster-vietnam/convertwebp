package imageconvert

import (
	"fmt"
	"os/exec"
	"strconv"
)

/*
convert image to webp
*/
func gm_convert(inputFile string, outputFile string, lossyQuality int) error {
	cmd := exec.Command("gm", "convert", inputFile, "-resize", "1024x>", "-quality", strconv.Itoa(lossyQuality), outputFile)
	// Chạy lệnh
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Error decoding image %s :  %s\n", inputFile, err)
	}

	fmt.Printf("Processed and converted %s to %s\n", inputFile, outputFile)
	return nil
}
