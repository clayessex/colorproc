package main

import (
	"fmt"
	"log"
	"slices"

	"github.com/clayessex/colorproc/colors"
)

// var ColorNames = []struct {
//         name string
//         rgb  string
// }{
//         {name: "100 Mph", rgb: "#c93f38"},
//         {name: "18th Century Green", rgb: "#a59344"},

type ColorInfo struct {
	name string
	rgb  colors.Hex
	hsl  colors.Hsl
	pos  float64
}

func main() {
	colorlist := make([]ColorInfo, 0, 30000)
	midpoint := colors.Hsl{H: 180.0, S: 0.5, L: 0.5}

	for _, v := range ColorNames {
		rgb := colors.Hex(v.rgb)
		hsl, err := rgb.ToHsl()
		if err != nil {
			log.Printf("invalid rgb value: %s\n", v.rgb)
			continue
		}
		pos := distance(hsl, midpoint)
		colorlist = append(colorlist, ColorInfo{v.name, rgb, hsl, pos})
	}

	slices.SortFunc(colorlist, func(a, b ColorInfo) int {
		if a.pos < b.pos {
			return -1
		}
		if a.pos > b.pos {
			return 1
		}
		return 0
	})

	for _, v := range colorlist {
		fmt.Printf("%s, %v, \"%s\", %f\n", v.rgb, v.hsl, v.name, v.pos+0.5)
	}
}

// The choice of 84% yields 57.6 degrees of Hue and then 8% of each Saturation
// and Luminance
func distance(a, b colors.Hsl) float64 {
	x := float64(a.H-b.H) / 360.0
	y := float64(a.S - b.S)
	z := float64(a.L - b.L)
	return 0.84*x + 0.08*y + 0.08*z
}
