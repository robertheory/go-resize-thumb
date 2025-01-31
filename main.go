package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/nfnt/resize"
)

func main() {

	inputDir := "images"
	outputDir := "thumbnails"

	now := time.Now()

	processImages(inputDir, outputDir)

	fmt.Println("Time taken to process images: ", time.Since(now).Milliseconds(), "ms")

}

func processImages(inputDir, outputDir string) {

	files, err := os.ReadDir(inputDir)

	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(outputDir, os.ModePerm)

	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	for _, file := range files {

		if !file.IsDir() {

			wg.Add(1)

			inputFullPath := filepath.Join(inputDir, file.Name())

			outputFullPath := filepath.Join(outputDir, file.Name())

			go processThumbnail(inputFullPath, outputFullPath, &wg)
		}
	}

	wg.Wait()

	fmt.Println("All images processed successfully")

}

func processThumbnail(inputPath, outputPath string, wg *sync.WaitGroup) {

	defer wg.Done()

	file, err := os.Open(inputPath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	image, _, err := image.Decode(file)

	if err != nil {
		panic(err)
	}

	thumbnail := resize.Resize(100, 0, image, resize.Lanczos3)

	outputFile, err := os.Create(outputPath)

	if err != nil {
		panic(err)
	}

	defer outputFile.Close()

	switch strings.ToLower(filepath.Ext(inputPath)) {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(outputFile, thumbnail, nil)

		if err != nil {
			panic(err)
		}

	case ".png":
		err = png.Encode(outputFile, thumbnail)

		if err != nil {
			panic(err)
		}

	}

}
