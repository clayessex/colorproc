package colors

import (
	"fmt"
	"testing"
)

// Generate data slice by generating all types and printing in Go code format
func dontTestDumpValues(t *testing.T) {
	data := []Hex{
		"#ffffff",
		"#000000",
		"#808080",
		"#ff00ff",
		"#f7f3d9",
		"#7f1c9e",
		"#0b2226",
		"#21ff90",
		"#cebf88",
		"#8080ff",
		"#411bea",
		"#223344",
	}

	shex := func(hex Hex) string {
		return fmt.Sprintf("Hex(\"%s\")", hex)
	}
	srgb := func(rgb Rgb) string {
		return fmt.Sprintf("Rgb{R:%d,G:%d,B:%d}", rgb.R, rgb.G, rgb.B)
	}
	shsl := func(hsl Hsl) string {
		return fmt.Sprintf("Hsl{H:%.15f,S:%.15f,L:%.15f}", hsl.H, hsl.S, hsl.L)
	}

	for _, v := range data {
		rgb, err := v.ToRgb()
		if err != nil {
			panic("invalid Hex rgb value")
		}
		hsl := rgb.ToHsl()
		fmt.Printf("{ %s, %s, %s },\n", shex(v), srgb(rgb), shsl(hsl))
	}
}
