package main

import (
	"flag"

	"github.com/TechMaster/convertwebp/imageconvert"
	"github.com/davidbyttow/govips/v2/vips"
)

func main() {
	// Đọc từ tham số dòng lệnh
	inputDir := flag.String("in", "", "Thư mục đầu vào")
	outputDir := flag.String("out", "", "Thư mục đầu ra")
	flag.Parse()

	// Gán giá trị mặc định nếu thiếu tham số đầu vào
	if *inputDir == "" {
		*inputDir = "./in"
	}
	if *outputDir == "" {
		*outputDir = "./out"
	}

	//ConvertFolderToWebP(inputDir, outputDir, KolesaConvert)
	//ConvertFolderToWebP(inputDir, outputDir, ChaiConvert)
	//ConvertFolderToWebP(inputDir, outputDir, gm_convert)
	//ConvertFolderToWebP(inputDir, outputDir, cweb_convert)
	imageconvert.ConvertFolderToWebP(*inputDir, *outputDir, imageconvert.VipsConvert, 80)
	defer vips.Shutdown()
}
