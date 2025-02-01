package main

import (
	"fmt"
	"os"
	"path/filepath"
	"resize-thumb/internal/entity"
	"time"
)

func main() {

	inputDir := "images"
	outputDir := "thumbnails"

	now := time.Now()

	images := []entity.ImageFile{}

	loadImages(inputDir, outputDir, &images)

	fmt.Println("Number of images: ", len(images))

	wp := entity.NewWorkerPool(images, 5)

	wp.Start()

	fmt.Println("Time taken to process images: ", time.Since(now).Milliseconds(), "ms")

}

func loadImages(inputDir string, outputDir string, imagesFiles *[]entity.ImageFile) {

	files, err := os.ReadDir(inputDir)

	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(outputDir, os.ModePerm)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			inputFullPath := filepath.Join(inputDir, file.Name())
			outputFullPath := filepath.Join(outputDir, file.Name())

			newImageFile := entity.NewImageFile(inputFullPath, outputFullPath)

			*imagesFiles = append(*imagesFiles, *newImageFile)
		}
	}

}
