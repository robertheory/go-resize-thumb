package entity

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

type ImageFile struct {
	inputPath  string
	outputPath string
}

func NewImageFile(inputPath, outputPath string) *ImageFile {
	return &ImageFile{
		inputPath:  inputPath,
		outputPath: outputPath,
	}
}

func (i *ImageFile) Process() error {

	file, err := os.Open(i.inputPath)

	if err != nil {
		return err
	}

	defer file.Close()

	image, _, err := image.Decode(file)

	if err != nil {
		return err
	}

	thumbnail := resize.Resize(100, 0, image, resize.Lanczos3)

	outputFile, err := os.Create(i.outputPath)

	if err != nil {
		return err
	}

	defer outputFile.Close()

	switch strings.ToLower(filepath.Ext(i.inputPath)) {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(outputFile, thumbnail, nil)

		if err != nil {
			return err
		}

	case ".png":
		err = png.Encode(outputFile, thumbnail)

		if err != nil {
			return err
		}

	}

	fmt.Println("- Processed: ", i.inputPath)

	return nil
}
