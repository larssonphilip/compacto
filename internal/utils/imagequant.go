package utils

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"time"
)

type ColorCluster struct {
	Pixels   []color.RGBA
	Centroid color.RGBA
}

// Calculate the distance between two colors
func colorDistance(color1, color2 color.RGBA) float64 {
	rMean := float64(color1.R+color2.R) / 2
	r := float64(color1.R - color2.R)
	g := float64(color1.G - color2.G)
	b := float64(color1.B - color2.B)

	// Euclidean distance formula
	return math.Sqrt((((512 + rMean) * r * r) / 256) + 4*g*g + (((767 - rMean) * b * b) / 256))
}

// Clusters colors into k clusters using the k-means algorithm
func kMeans(colors []color.RGBA, k int) []color.RGBA {
	rand.NewSource(time.Now().UnixNano())
	clusters := make([]ColorCluster, k)
	for i := 0; i < k; i++ {
		clusters[i].Centroid = colors[rand.Intn(len(colors))]
	}

	changed := true
	for changed {
		for i := range clusters {
			clusters[i].Pixels = nil
		}

		for _, color := range colors {
			minDistance := math.MaxFloat64
			minIndex := 0
			for i, cluster := range clusters {
				distance := colorDistance(color, cluster.Centroid)
				if distance < minDistance {
					minDistance = distance
					minIndex = i
				}
			}
			clusters[minIndex].Pixels = append(clusters[minIndex].Pixels, color)
		}

		changed = false
		for i, cluster := range clusters {
			var rSum, gSum, bSum, count int
			for _, pixel := range cluster.Pixels {
				rSum += int(pixel.R)
				gSum += int(pixel.G)
				bSum += int(pixel.B)
				count++
			}
			if count > 0 {
				newCentroid := color.RGBA{
					R: uint8(rSum / count),
					G: uint8(gSum / count),
					B: uint8(bSum / count),
					A: 255,
				}
				if newCentroid != cluster.Centroid {
					clusters[i].Centroid = newCentroid
					changed = true
				}
			}
		}
	}

	centroids := make([]color.RGBA, k)
	for i, cluster := range clusters {
		centroids[i] = cluster.Centroid
	}

	return centroids
}

// Dithers an image using the Floyd-Steinberg algorithm
func ditherImage(img image.Image, palette []color.RGBA) image.Image {
	bounds := img.Bounds()
	dithered := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			oldPixel := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
			newPixel := findNearestColor(oldPixel, palette)
			dithered.Set(x, y, newPixel)

			// Calculate the error
			rQuantError, gQuantError, bQuantError := quantizationError(oldPixel, newPixel)

			// Distributing the error to the right neighbour
			spreadError(dithered, x+1, y, rQuantError, gQuantError, bQuantError, 7.0/16.0)

			// Distributing the error to the bottom left neighbour
			spreadError(dithered, x-1, y+1, rQuantError, gQuantError, bQuantError, 3.0/16.0)

			// Distributing the error to the bottom neighbour
			spreadError(dithered, x, y+1, rQuantError, gQuantError, bQuantError, 5.0/16.0)

			// Distributing the error to the bottom right neighbour
			spreadError(dithered, x+1, y+1, rQuantError, gQuantError, bQuantError, 1.0/16.0)
		}
	}

	return dithered
}

// Calculate the quantization error between two pixels
func quantizationError(oldPixel, newPixel color.RGBA) (rDiff, gDiff, bDiff int) {
	rDiff = int(oldPixel.R) - int(newPixel.R)
	gDiff = int(oldPixel.G) - int(newPixel.G)
	bDiff = int(oldPixel.B) - int(newPixel.B)

	return rDiff, gDiff, bDiff
}

// Spread the quantization error to the neighbour pixel
func spreadError(img *image.RGBA, x, y int, rError, gError, bError int, factor float64) {
	if x < 0 || x >= img.Bounds().Max.X || y < 0 || y >= img.Bounds().Max.Y {
		return
	}

	originalColor := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
	r := clamp(int(float64(originalColor.R) + float64(rError)*factor))
	g := clamp(int(float64(originalColor.G) + float64(gError)*factor))
	b := clamp(int(float64(originalColor.B) + float64(bError)*factor))

	img.Set(x, y, color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: originalColor.A})
}

// Clamps a value between 0 and 255
func clamp(value int) int {
	if value < 0 {
		return 0
	}
	if value > 255 {
		return 255
	}
	return value
}

// Remap an image to a given palette
func remapImageToPalette(img image.Image, palette []color.RGBA) image.Image {
	bounds := img.Bounds()
	remapped := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
			nearestColor := findNearestColor(originalColor, palette)
			remapped.Set(x, y, nearestColor)
		}
	}

	return remapped
}

// Find the nearest color in a palette to a given color
func findNearestColor(color color.RGBA, palette []color.RGBA) color.RGBA {
	nearestColor := palette[0]
	minDistance := colorDistance(color, nearestColor)
	for _, paletteColor := range palette {
		distance := colorDistance(color, paletteColor)
		if distance < minDistance {
			minDistance = distance
			nearestColor = paletteColor
		}
	}

	return nearestColor
}
