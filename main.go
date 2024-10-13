package main

import (
	"flag"
	"fmt"
	"os"

	imagecompressor "compacto.app/compacto/internal/images"
)

func main() {
	fmt.Println("Compacto v0.0.1")
	fmt.Println("Arguments: ", os.Args)
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./compacto <compressionType> <inputPath> <outputPath>")
		os.Exit(1)
	}

	compressImageCmd := flag.NewFlagSet("compress-image", flag.ExitOnError)
	// compressVideoCmd := flag.NewFlagSet("compress-video", flag.ExitOnError)
	// compressPdfCmd := flag.NewFlagSet("compress-pdf", flag.ExitOnError)

	quality := flag.Int("quality", 100, "Quality of the image")
	speed := flag.Int("speed", 2, "Speed of the image compression")
	flag.Parse()

	fmt.Println("Quality: ", *quality)

	switch os.Args[1] {
	case "compress-image":
		fmt.Println("Parsing flags for compress-image subcommand...")
		fmt.Printf("Arguments passed to Parse: %v\n", os.Args[2:])

		err := compressImageCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Error parsing flags: ", err)
			os.Exit(1)
		}

		fmt.Printf("Remaining args after Parse: %v\n", compressImageCmd.Args())

		if compressImageCmd.NArg() < 2 {
			fmt.Println("Usage: ./compacto compress-image <inputPath> <outputPath> [--quality 1-100] [--speed 1-10]")
			os.Exit(1)
		}
		inputPath := compressImageCmd.Args()[0]
		outputPath := compressImageCmd.Args()[1]
		fmt.Printf("Parsed flags: quality=%d, speed=%d\n", *quality, *speed)
		qualityMin := *quality - 15
		qualityMax := *quality

		fmt.Printf("Compressing image %s to %s with qualityMin %d, qualityMax %d and speed %d\n", inputPath, outputPath, qualityMin, qualityMax, *speed)
		imagecompressor.CompressPngImage(inputPath, outputPath, qualityMin, qualityMax, *speed)
		fmt.Println("Image compressed")
	case "compress-video":
		fmt.Println("Video compression not implemented yet")
	case "compress-pdf":
		fmt.Println("PDF compression not implemented yet")
	}
}
