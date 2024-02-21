package main

import (
	"fmt"
	"image"

	// Import the PNG package for decoding
	"image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

// compressAndConvertImage resizes a PNG image and converts it to JPEG.
func compressAndConvertImage(inputPath, outputPath string, maxHeight uint, jpegQuality int) error {
	// Open the input image
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// Decode the image
	img, _, err := image.Decode(inputFile)
	if err != nil {
		return err
	}

	// Resize the image if it is taller than maxHeight
	originalHeight := img.Bounds().Dy()
	if uint(originalHeight) > maxHeight {
		img = resize.Resize(0, maxHeight, img, resize.Lanczos3)
	}

	// Ensure the output directory exists
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return err
	}

	// Change the file extension to .jpg
	outputPath = strings.TrimSuffix(outputPath, filepath.Ext(outputPath)) + ".jpg"

	// Create the output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Encode the image to JPEG format
	jpegOpts := &jpeg.Options{Quality: jpegQuality}
	if err := jpeg.Encode(outputFile, img, jpegOpts); err != nil {
		return err
	}

	fmt.Println("Processed and converted to JPEG:", outputPath)
	return nil
}

// processDirectory processes each PNG image in a directory and subdirectories.
func processDirectory(inputDir, outputDir string, maxHeight uint, jpegQuality int) error {
	return filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Check if file is a PNG image
		if strings.HasSuffix(strings.ToLower(info.Name()), ".png") {
			relPath, err := filepath.Rel(inputDir, path)
			if err != nil {
				return err
			}

			outputPath := filepath.Join(outputDir, relPath)
			return compressAndConvertImage(path, outputPath, maxHeight, jpegQuality)
		}

		return nil
	})
}

func main() {
	inputDir := "EPIC41Fantasy" // Replace with your input directory
	outputDir := "output"       // Replace with your output directory
	maxHeight := uint(330)      // Maximum height of the images
	jpegQuality := 90           // JPEG quality (1-100)

	if err := processDirectory(inputDir, outputDir, maxHeight, jpegQuality); err != nil {
		fmt.Println("Error processing directory:", err)
	}
}
