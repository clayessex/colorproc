package main

import (
	"fmt"
	"math"
	"strconv"
)

func (h Hex) ToRgb() (Rgb, error) {
	s := string(h)
	if s[0] == '#' {
		s = s[1:]
	}

	v, err := strconv.ParseInt(s, 16, 32)
	if err != nil {
		return Rgb{}, err
	}

	r := uint8(v >> 16 & 0xff)
	g := uint8(v >> 8 & 0xff)
	b := uint8(v & 0xff)

	return Rgb{r, g, b}, nil
}

func (h Hex) ToHsl() (Hsl, error) {
	rgb, err := h.ToRgb()
	if err != nil {
		return Hsl{}, err
	}
	return rgb.ToHsl(), nil
}

func (rgb Rgb) ToHex() Hex {
	s := fmt.Sprintf("#%x%x%x", rgb.R, rgb.G, rgb.B)
	return Hex(s)
}

func (rgb Rgb) ToHsl() Hsl {
	vmax := max(rgb.R, rgb.G, rgb.B)
	vmin := min(rgb.R, rgb.G, rgb.B)

	// normalize
	nr := float64(rgb.R) / 255.0
	ng := float64(rgb.G) / 255.0
	nb := float64(rgb.B) / 255.0
	nvmax := float64(vmax) / 255.0
	nvmin := float64(vmin) / 255.0

	h := 0.0
	s := 0.0
	l := (nvmax + nvmin) / 2.0

	if vmax != vmin {
		c := nvmax - nvmin
		s = c / (1.0 - math.Abs(2.0*l-1.0))

		switch vmax {
		case rgb.R:
			h = (ng - nb) / c
			h = math.Mod(h, 6.0)
			if ng < nb {
				h += 6.0
			}
		case rgb.G:
			h = (nb-nr)/c + 2.0
		case rgb.B:
			h = (nr-ng)/c + 4.0
		}

		h *= 60.0
	}

	return Hsl{float32(h), float32(s), float32(l)}
}

func (hsl Hsl) ToRgb() Rgb {
	nh := float64(hsl.H)
	ns := float64(hsl.S)
	nl := float64(hsl.L)

	if hsl.S == 0 {
		v := denorm(nl)
		return Rgb{v, v, v}
	}

	calc := func(n float64) float64 {
		k := math.Mod(n+(nh/30.0), 12)
		a := ns * min(nl, 1.0-nl)
		return nl - a*max(-1.0, min(k-3.0, 9.0-k, 1.0))
	}
	r, g, b := calc(0.0), calc(8.0), calc(4.0)
	return Rgb{denorm(r), denorm(g), denorm(b)}
}

func (hsl Hsl) ToHex() Hex {
	rgb := hsl.ToRgb()
	return rgb.ToHex()
}

func denorm(v float64) uint8 {
	return uint8(math.Round(v * 255.0))
}
