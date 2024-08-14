package colors

import (
	"fmt"
	"math"
)

type Rgb struct {
	R uint8
	G uint8
	B uint8
}

type Hsl struct {
	H float32 // [0, 360]
	S float32 // [0, 1]
	L float32 // [0, 1]
}

type Hex string

func (hex Hex) String() string {
	return string(hex)
}

func (rgb Rgb) String() string {
	return fmt.Sprintf("RGB(%v, %v, %v)", rgb.R, rgb.G, rgb.B)
}

func pct(v float32) int {
	return int(math.Round(float64(v) * 100.0))
}

func (hsl Hsl) String() string {
	return fmt.Sprintf("HSL(%v°, %v%%, %v%%)", int(math.Round(float64(hsl.H))), pct(hsl.S), pct(hsl.L))
}

func (hsl Hsl) F3String() string {
	return fmt.Sprintf("HSL(%.3f°, %.3f%%, %.3f%%)", hsl.H, hsl.S*100.0, hsl.L*100.0)
}
