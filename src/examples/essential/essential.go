package main

import (
	"cmplximage"
	"image"
	"image/png"
	"math/cmplx"
	"os"
)

func Essential(z complex128) complex128 {
	return cmplx.Exp(1.0 / z)
}

func main() {
	res := image.Rect(0, 0, 1600, 1600)

	domain := cmplximage.NewCmplxRect(complex(-1, -1), complex(1, 1))
	img := cmplximage.Draw(cmplximage.HSLWheelMap(Essential), res, domain)
	file, err := os.Create("essential.png")
	if err != nil {
		return
	}
	defer file.Close()
	png.Encode(file, img)
}
