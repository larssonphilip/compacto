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
		if err != nil {
			return nil, "", fmt.Errorf("Failed to decode JPEG image: %w", err)
		}
		return img, "jpeg", err
	case ".png":
		img, err := png.Decode(file)
		if err != nil {
			return nil, "", fmt.Errorf("Failed to decode PNG image: %w", err)
		}
		return img, "png", nil
	case ".gif":
		img, err := gif.Decode(file)
		if err != nil {
			return nil, "", fmt.Errorf("Failed to decode GIF image: %w", err)
		}
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

func CompressPngImage(inputPath, outputPath string, qualityMin, qualityMax, speed int, dither float64) {
	attributes := imagequant.CreateAttributes()
	defer attributes.Release()

	attributes.SetQuality(qualityMin, qualityMax)
	attributes.SetSpeed(speed)

	img, format, err := loadImage(inputPath)
	if err != nil {
		error := fmt.Errorf("Failed to load image: %w", err)
		fmt.Println(error)
		return
	}

	qimg := attributes.CreateImage(img, 0.0)

	qresult, err := attributes.QuantizeImage(qimg)
	if err != nil {
		error := fmt.Errorf("Failed to quantize image: %w", err)
		fmt.Println(error)
		return
	}

	dithererr := attributes.SetDitheringLevel(qresult, float32(dither))

	if dithererr != nil {
		error := fmt.Errorf("Failed to set dithering level: %w", dithererr)
		fmt.Println(error)
	}

	imgOut, _ := attributes.WriteRemappedImage(qresult, qimg)

	saveImage(imgOut, format, outputPath)
}
