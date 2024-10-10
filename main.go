package main

import (
	"flag"

	"github.com/TechMaster/convertwebp/imageconvert"
	"github.com/davidbyttow/govips/v2/vips"
)

func main() {
	// Đọc từ tham số dòng lệnh
	inputDir := flag.String("in", "", "input folder")
	outputDir := flag.String("out", "", "output folder")
	lossyQuality := flag.Int("lossyQuality", 80, "lossy quality")
	flag.Parse()

	// Gán giá trị mặc định nếu thiếu tham số đầu vào
	if *inputDir == "" {
		*inputDir = "./in"
	}
	if *outputDir == "" {
		*outputDir = "./out"
	}

	imageconvert.ConvertFolderToWebP(*inputDir, *outputDir, imageconvert.VipsConvertCopy, *lossyQuality)
	defer vips.Shutdown()
}
