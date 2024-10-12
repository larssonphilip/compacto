package images

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"github.com/InfinityTools/go-imagequant"
)

func loadImage(filePath string) (image.Image, string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", fmt.Errorf("Failed to open image file: %w", err)
	}
	defer file.Close()

	ext := filepath.Ext(filePath)
	switch ext {
	case ".jpg", ".jpeg":
		img, err := jpeg.Decode(file)
		return img, "jpeg", err
	case ".png":
		img, err := png.Decode(file)
		return img, "png", err
	case ".gif":
		img, err := gif.Decode(file)
		return img, "gif", err
	default:
		return nil, "", fmt.Errorf("Unsupported image format: %s", ext)
	}
}

func saveImage(img image.Image, format, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("Failed to create output file: %w", err)
	}
	defer file.Close()
	switch format {
	case "jpeg":
		err = jpeg.Encode(file, img, nil)
	case "png":
		err = png.Encode(file, img)
	case "gif":
		err = gif.Encode(file, img, nil)
	default:
		err = fmt.Errorf("Unsupported image format: %s", format)
	}
	return err
}

func CompressPngImage(inputPath, outputPath string) {
	att := imagequant.CreateAttributes()
	defer att.Release()

	img, format, err := loadImage(inputPath)
	if err != nil {
		fmt.Errorf("Failed to load image: %w", err)
	}

	qimg := att.CreateImage(img, 0.0)

	qresult, _ := att.QuantizeImage(qimg)

	imgOut, _ := att.WriteRemappedImage(qresult, qimg)

	saveImage(imgOut, format, outputPath)
}
