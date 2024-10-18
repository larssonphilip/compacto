package utils

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"time"
)

type ColorCluster struct {
	Centroid color.RGBA
	Pixels   []color.RGBA
}

func colorDistance(color1, color2 color.RGBA) float64 {
	rmean := float64(color1.R+color2.R) / 2
	r := float64(color1.R - color2.R)
	g := float64(color1.G - color2.G)
	b := float64(color1.B - color2.B)
	return math.Sqrt((((512 + rmean) * r * r) / 256) + 4*g*g + (((767 - rmean) * b * b) / 256))
}

func kMeans(colors []color.RGBA, k int) []color.RGBA {
	// rand.Seed(time.Now().UnixNano())
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
