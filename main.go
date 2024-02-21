package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/nfnt/resize"
)

func compressAndConvertImage(inputPath, outputPath string, maxHeight uint, jpegQuality int, wg *sync.WaitGroup) {
	defer wg.Done()

	inputFile, err := os.Open(inputPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inputFile.Close()

	img, _, err := image.Decode(inputFile)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	originalHeight := img.Bounds().Dy()
	if uint(originalHeight) > maxHeight {
		img = resize.Resize(0, maxHeight, img, resize.Lanczos3)
	}

	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}

	outputPath = strings.TrimSuffix(outputPath, filepath.Ext(outputPath)) + ".jpg"
	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	jpegOpts := &jpeg.Options{Quality: jpegQuality}
	if err := jpeg.Encode(outputFile, img, jpegOpts); err != nil {
		fmt.Println("Error encoding JPEG:", err)
		return
	}

	fmt.Println("Processed and converted to JPEG:", outputPath)
}

func processDirectory(inputDir, outputDir string, maxHeight uint, jpegQuality int) {
	var wg sync.WaitGroup

	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(strings.ToLower(info.Name()), ".png") {
			relPath, err := filepath.Rel(inputDir, path)
			if err != nil {
				return err
			}

			outputPath := filepath.Join(outputDir, relPath)
			wg.Add(1)
			go compressAndConvertImage(path, outputPath, maxHeight, jpegQuality, &wg)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error walking through directory:", err)
	}

	wg.Wait()
}

func main() {
	var (
		inputDir    string
		outputDir   string
		maxHeight   uint
		jpegQuality int
	)

	flag.StringVar(&inputDir, "i", "", "Path to the input directory")
	flag.StringVar(&outputDir, "ot", "", "Path to the output directory")
	flag.UintVar(&maxHeight, "h", 300, "Maximum height of the images")
	flag.IntVar(&jpegQuality, "q", 80, "JPEG quality (1-100)")
	flag.Parse()

	if inputDir == "" || outputDir == "" {
		fmt.Println("Input and output directories are required.")
		flag.PrintDefaults()
		return
	}

	processDirectory(inputDir, outputDir, maxHeight, jpegQuality)
}
