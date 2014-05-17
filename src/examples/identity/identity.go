package main

import (
	"cmplximage"
	"image"
	"image/png"
	"os"
)

func Identity(z complex128) complex128 {
	return z
}

func main() {
	res := image.Rect(0, 0, 1600, 1600)

	domain := cmplximage.NewCmplxRect(complex(-1, -1), complex(1, 1))
	img := cmplximage.Draw(cmplximage.HSLWheelMap(Identity), res, domain)
	file, err := os.Create("identity.png")
	if err != nil {
		return
	}
	defer file.Close()
	png.Encode(file, img)
}
