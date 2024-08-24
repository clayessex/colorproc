package colors

import "math"

// From: https://en.wikipedia.org/wiki/Color_difference
// sometimes called "redmean"
func Distance(x, y Rgb) float64 {
	rmean := (float64(x.R) + float64(y.R)) / 2.0
	r := float64(x.R) - float64(y.R)
	g := float64(x.G) - float64(y.G)
	b := float64(x.B) - float64(y.B)
	wr := (2.0 + rmean/256.0) * r * r
	wg := 4.0 * g * g
	wb := (2.0 + (255.0-rmean)/256.0) * b * b
	return math.Sqrt(wr + wg + wb)
}

// Simple Hue based distance calculation between two Hsl values
// The choice of 84% yields 57.6 degrees of Hue and then 8% of each Saturation
// and Luminance
func Distance2(a, b Hsl) float64 {
	x := float64(a.H-b.H) / 360.0
	y := float64(a.S - b.S)
	z := float64(a.L - b.L)
	return 0.84*x + 0.08*y + 0.08*z
}
