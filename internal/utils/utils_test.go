package utils

import (
  "image"
  "image/color"
  "testing"
)

func TestColorDistance(t *testing.T) {
  color1 := color.RGBA(R: 255, G: 0, B: 0, A: 255)
  color2 := color.RGBA(R: 0, G: 255, B: 0, A: 255)

  distance := colorDistance(color1, color2)
  
  if distance <= 0 {
    t.Errorf("Expected positive distance, got %f", distance)
  }
}

func TestKMeans(t *testing.T) {
  colors := []color.RGBA{
    {R: 255, G: 0, B: 0, A: 255},
    {R: 0, G: 255, B: 0, A: 255},
    {R: 0, G: 0, B: 255, A: 255}
  }

  k := 2
  clusters := kMeans(colors, k)
  if len(clusters) != k {
    t.Errorf("Expected %d clsuters, got %d", len(clsuters))
  }
}
