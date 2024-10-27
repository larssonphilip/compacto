package utils

import (
	"image"
	"image/color"
	"testing"
)

func TestColorDistance(t *testing.T) {
	color1 := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	color2 := color.RGBA{R: 0, G: 255, B: 0, A: 255}

	distance := colorDistance(color1, color2)

	if distance <= 0 {
		t.Errorf("Expected positive distance, got %f", distance)
	}
}

func TestKMeans(t *testing.T) {
	colors := []color.RGBA{
		{R: 255, G: 0, B: 0, A: 255},
		{R: 0, G: 255, B: 0, A: 255},
		{R: 0, G: 0, B: 255, A: 255},
	}

	k := 2
	clusters := kMeans(colors, k)
	if len(clusters) != k {
		t.Errorf("Expected %d clsuters, got %d", k, len(clusters))
	}
}

func TestQuantizationError(t *testing.T) {
	oldPixel := color.RGBA{R: 100, G: 150, B: 200, A: 255}
	newPixel := color.RGBA{R: 90, G: 140, B: 210, A: 255}

	rError, gError, bError := quantizationError(oldPixel, newPixel)
	if rError != 10 || gError != 10 || bError != -10 {
		t.Errorf("Expected (10, 10, -10), got (%d, %d, %d)", rError, gError, bError)
	}
}

func TestSpreadError(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))
	initialColor := color.RGBA{R: 100, G: 100, B: 100, A: 255}
	img.Set(1, 1, initialColor)

	spreadError(img, 1, 1, 10, 10, 10, 0.5)

	adjustedColor := img.RGBAAt(1, 1)
	expectedColor := color.RGBA{R: 105, G: 105, B: 105, A: 255}

	if adjustedColor != expectedColor {
		t.Errorf("Expected color %v, got %v", expectedColor, adjustedColor)
	}
}

func TestDitherImage(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	img.Set(0, 0, color.RGBA{R: 100, G: 150, B: 200, A: 255})

	palette := []color.RGBA{
		{R: 90, G: 140, B: 190, A: 255},
		{R: 110, G: 160, B: 210, A: 255},
	}

	ditheredImage := ditherImage(img, palette)
	ditheredColor := ditheredImage.At(0, 0).(color.RGBA)
	expectedColor := findNearestColor(color.RGBA{R: 100, G: 150, B: 200, A: 255}, palette)

	if ditheredColor != expectedColor {
		t.Errorf("Expected color %v, got %v", expectedColor, ditheredColor)
	}
}

func TestRemapImageToPalette(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))
	img.Set(0, 0, color.RGBA{R: 100, G: 100, B: 100, A: 255})
	img.Set(1, 0, color.RGBA{R: 150, G: 150, B: 150, A: 255})
	img.Set(0, 1, color.RGBA{R: 200, G: 200, B: 200, A: 255})
	img.Set(1, 1, color.RGBA{R: 250, G: 250, B: 250, A: 255})

	palette := []color.RGBA{
		{R: 90, G: 90, B: 90, A: 255},
		{R: 160, G: 160, B: 160, A: 255},
	}

	remappedImage := remapImageToPalette(img, palette)

	for y := 0; y < 2; y++ {
		for x := 0; x < 2; x++ {
			remappedColor := remappedImage.At(x, y).(color.RGBA)
			expectedColor := findNearestColor(img.RGBAAt(x, y), palette)
			if remappedColor != expectedColor {
				t.Errorf("At (%d, %d): Expected color %v, got %v", x, y, expectedColor, remappedColor)
			}
		}
	}
}
