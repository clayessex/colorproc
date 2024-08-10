package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
)

func main() {
	convert := flag.Bool("convert", false, "Convert a #RGB value to HSL")

	flag.Parse()

	if *convert {
		rgb := flag.Arg(0)
		if len(rgb) == 0 {
			panic("missing #RGB argument")
		}

		if rgb[0] == '#' {
			rgb = rgb[1:]
		}

		if len(rgb) != 6 {
			panic("invalid #RGB argument: " + rgb)
		}

		r, err := strconv.ParseUint(rgb[0:2], 16, 8)
		if err != nil {
			fmt.Println(err)
			return
		}
		g, err := strconv.ParseUint(rgb[2:4], 16, 8)
		if err != nil {
			fmt.Println(err)
			return
		}
		b, err := strconv.ParseUint(rgb[4:], 16, 8)
		if err != nil {
			fmt.Println(err)
			return
		}

		h, s, l := RgbToHsl(int(r), int(g), int(b))

		fmt.Printf("#%s\n", rgb)
		fmt.Printf("RGB: (%v, %v, %v)\n", r, g, b)
		fmt.Printf("HSL: (%v, %v%%, %v%%)\n", h, s, l)
	}
}

func RgbToHsl(r, g, b int) (int, int, int) {
	rn := float64(r) / 255.0
	gn := float64(g) / 255.0
	bn := float64(b) / 255.0

	vmax := max(rn, gn, bn)
	vmin := min(rn, gn, bn)
	c := vmax - vmin
	l := (vmax + vmin) / 2.0

	const ep = 1.0 / 255.0 / 2.0
	if c <= ep {
		return 0, 0, int(math.Round(l * 100.0))
	}

	var s float64
	if l < 0.5 {
		s = c / (vmax + vmin)
	} else {
		s = c / (2.0 - vmax - vmin)
	}

	var h float64
	switch max(r, g, b) {
	case r:
		h = (gn - bn) / c
		h = math.Mod(h, 6.0)
	case g:
		h = (bn-rn)/c + 2.0
	case b:
		h = (rn-gn)/c + 4.0
	}
	h *= 60.0
	if h < 0.0 {
		h += 360.0
	}

	return int(math.Round(h)),
		int(math.Round(s * 100.0)),
		int(math.Round(l * 100.0))
}
