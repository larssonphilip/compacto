package main

import (
	"fmt"
	"os"

	"compacto.app/compacto/internal/images"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory")
	} else {
		fmt.Println("Current working directory: ", cwd)
	}

	inputPath := "test/images/bigimage.png"
	outputPath := "test/images/smallimage.png"

	images.CompressPngImage(inputPath, outputPath, 85, 100, 2)
	fmt.Println("Image compressed")
}
