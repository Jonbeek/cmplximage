package cmplximage

import (
	"image/color"
	"math"
	"math/cmplx"
)

// Needed because Go doesn't have a floating point round function.
// Either way, guaranteed to fit in a uint8 and be positive.
func round(num float64) uint8 {
	return uint8(math.Floor(num + 0.5))
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
		x := (2 * real(val)) / div
		y := (2 * imag(val)) / div
		z := (add - 1.0) / div
		// Now with the calculations out of the way, convert to standard color.
		// Uniformly map from [-1,1] to [0,255]
		r := round(255 * ((x + 1) / 2))
		g := round(255 * ((y + 1) / 2))
		b := round(255 * ((z + 1) / 2))

		return color.RGBA{r, g, b, 255}
	}
}

// HSLWheelMap generates a ColorMap from a ComplexMap, using a the HSL color
// space with the Argument of the point as the hue, and a nonlinear mapping of
// the absolute value of the point as the lightness.
func HSLWheelMap(fnc ComplexMap) ColorMap {
	return func(point complex128) color.Color {
		val := fnc(point)

		add := math.Pow(cmplx.Abs(val), 2.0)
		// map to [0,1]
		L := add / (add + 1.0)

		// See wikipedia. This is a fairly bad implementation of the algorithm.

		H := (3.0 * cmplx.Phase(point) / math.Pi) + 3.0

		// Get the values being used
		C := (1.0 - math.Abs(2.0*L-1.0))
		X := C * (1.0 - math.Abs(math.Mod(H, 2.0)-1.0))

		// Giant semi-conditional, again, see wikipedia.
		var R1, G1, B1 float64
		switch math.Floor(H) {
		case 0:
			R1 = C
			G1 = X
		case 1:
			R1 = X
			G1 = C
		case 2:
			G1 = C
			B1 = X
		case 3:
			G1 = X
			B1 = C
		case 4:
			R1 = X
			B1 = C
		case 5:
			R1 = C
			B1 = X
		default:
			// Pass on this, 0 initialization is good.
		}

		// m is the "minimum" value of each color.
		m := (L - C/2.0)
		// Convert from [0,1] RGB to [0,255] RGB (and add m while we're at it).
		r := round(255.0 * (R1 + m))
		g := round(255.0 * (G1 + m))
		b := round(255.0 * (B1 + m))

		// We're done. Replace all of the above once Go adds HSL support.
		return color.RGBA{r, g, b, 255}
	}
}
