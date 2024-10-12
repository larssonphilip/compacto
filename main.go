package main

import (
	"fmt"

	"compacto.app/compacto/internal/images"
)

func main() {
	images.CompressPngImage("./test/images/bigimage.png", "./test/images/smallimage.png")
	fmt.Println("Image compressed")
}
