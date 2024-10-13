package main

import (
	"flag"
	"fmt"
	"os"

	imagecompressor "compacto.app/compacto/internal/images"
)

func main() {
	fmt.Println("Compacto v0.0.1")

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./compacto compress-image [--quality 1-100] [--speed 1-10] <inputPath> <outputPath>")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "compress-image":
		compressImageCmd := flag.NewFlagSet("compress-image", flag.ExitOnError)

		quality := compressImageCmd.Int("quality", 100, "Quality of the image")
		speed := compressImageCmd.Int("speed", 2, "Speed of the image compression")
		dither := compressImageCmd.Float64("dither", 1.0, "Dither of the image compression")

		err := compressImageCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Error parsing flags: ", err)
			os.Exit(1)
		}

		if compressImageCmd.NArg() < 2 {
			fmt.Println("Usage: ./compacto compress-image [--quality 1-100] [--speed 1-10] [--dither 0.0-1.0] <inputPath> <outputPath>")
			os.Exit(1)
		}
		inputPath := compressImageCmd.Args()[0]
		outputPath := compressImageCmd.Args()[1]

		qualityMin := *quality - 15
		qualityMax := *quality

		fmt.Printf("Compressing image %s to %s with quality %d and speed %d\n", inputPath, outputPath, *quality, *speed)
		imagecompressor.CompressPngImage(inputPath, outputPath, qualityMin, qualityMax, *speed, *dither)
		fmt.Println("Image compressed")
	case "compress-video":
		// compressVideoCmd := flag.NewFlagSet("compress-video", flag.ExitOnError)
		fmt.Println("Video compression not implemented yet")
	case "compress-pdf":
		// compressPdfCmd := flag.NewFlagSet("compress-pdf", flag.ExitOnError)
		fmt.Println("PDF compression not implemented yet")
	}
}
