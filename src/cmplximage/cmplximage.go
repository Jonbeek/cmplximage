// cmplximg provides simple ways of displaying functions of one complex number
// as an image.
package cmplximage

import (
	"image"
	"image/color"
	"math"
	"math/cmplx"
)

// ComplexRect is a rectangle in the complex plane.
type ComplexRect struct {
	a complex128
	b complex128
}

func NewCmplxRect(Min, Max complex128) *ComplexRect {
	rect := new(ComplexRect)
	rect.a = Min
	rect.b = Max
	return rect
}

func (cr ComplexRect) dx() float64 {
	return math.Abs(real(cr.a) - real(cr.b))
}

func (cr ComplexRect) dy() float64 {
	return math.Abs(imag(cr.a) - real(cr.b))
}

// Other ComplexRect methods are not needed and are superfluous.
func (cr ComplexRect) bottom() float64 {
	if imag(cr.a) < imag(cr.b) {
		return imag(cr.a)
	} else {
		return imag(cr.b)
	}
}

func (cr ComplexRect) left() float64 {
	if real(cr.a) < real(cr.b) {
		return real(cr.a)
	} else {
		return real(cr.b)
	}
}

// ComplexMap is a function in the complex plane.
type ComplexMap func(point complex128) (complex128)

// ColorMap maps a point on the complex plane to a color.
type ColorMap func(point complex128) (color.Color)

// Draw creates an image of the function in the domain.
func Draw(fnc ColorMap, size image.Rectangle, domain ComplexRectangle) image.Image {
	size = size.Canon()
	// Clever vector hack to move the Min corner to 0,0
	size = size.Sub(size.Min)
	// For now, use RGBA as image type
	img := image.NewRGBA(size)
	// max x and y guaranteed to be size of rectangle
	x := size.Dx()
	y := size.Dy()
	dx := domain.dx() / x
	dy := domain.dy() / y
	// Get the initial x vals
	base_x := domain.left()
	base_y := domain.bottom()
	for i := 0; i <= x; i++ {
		for j := 0; j <= y; j++ {
			point := complex128(base_x + i*dx, base_y + j*dy)
			img.Set(i, j, fnc(point))
		}
	}
	return img
}

// RiemannMap generates a ColorMap from a ComplexMap, using the Riemann sphere.
// Red, green, and blue are respectively set to the x, y, and z coordinates.
func RiemannMap(fnc ComplexMap) ColorMap {
	return func(point complex128) color.Color {
		val := fnc(point)
		// Convert val to points on Riemann sphere, then set colors.
		// All will be in range [-1,1]
		add := math.Pow(cmplx.Abs(val), 2)
		div := 1.0 + add
		x := (2*real(val)) / div
		y := (2*imag(val)) / div
		z := (add - 1.0) / div
		// Now with the calculations out of the way, convert to standard color.
		// Uniformly map from [-1,1] to [0,255]
		r := math.Round(255 * ((x + 1) / 2))
		g := math.Round(255 * ((y + 1) / 2))
		b := math.Round(255 * ((z + 1) / 2))

		return color.RGBA(r, g, b, 255)
	}
}

