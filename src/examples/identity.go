package main

import (
	"cmplximage"
	"image"
	"os"
	"image/png"
)

func Identity(z complex128) complex128 {
	return z
}

func Main() {
	res := image.NewRect(0, 0, 400, 400)

	domain := cmplximage.NewRect(complex(-10, -10), complex(10, 10))
	img := cmplximage.Draw(cmplximage.RiemannMap(Identity), res, domain)
	file := os.Create("identity.png")
	defer file.Close()
	png.Encode(file, img)
}
