package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"slices"

	"github.com/clayessex/colorproc/build/colornames"
	"github.com/clayessex/colorproc/colors"
)

const (
	OutputPath     = "../build/colornamesorted/"
	OutputFilename = "colornamesorted.go"
)

type ColorInfo struct {
	name string
	hex  colors.Hex
	rgb  colors.Rgb
	hsl  colors.Hsl
	pos  float64
}

func main() {
	colorlist := make([]ColorInfo, 0, 30000)

	for _, v := range colornames.List {
		hex := colors.Hex(v.Rgb)
		rgb, err := hex.ToRgb()
		if err != nil {
			log.Printf("invalid rgb value: %s\n", v.Rgb)
			continue
		}
		hsl := rgb.ToHsl()
		pos := 0.0
		colorlist = append(colorlist, ColorInfo{v.Name, hex, rgb, hsl, pos})
	}

	SortHue(colorlist)

	WriteColorNames(OutputPath, OutputFilename, colorlist)
}

func SortHue(list []ColorInfo) {
	midpoint := colors.Rgb{R: 127, G: 127, B: 127}

	for i := 0; i < len(list); i++ {
		list[i].pos = colors.Distance(midpoint, list[i].rgb)
	}

	slices.SortFunc(list, func(a, b ColorInfo) int {
		if a.pos < b.pos {
			return -1
		}
		if a.pos > b.pos {
			return 1
		}
		return 0
	})
}

func WriteColorNames(filepath, filename string, s []ColorInfo) {
	if _, err := os.Stat(filepath); errors.Is(err, fs.ErrNotExist) {
		os.Mkdir(filepath, 0775)
	}

	f, err := os.Create(filepath + filename)
	if err != nil {
		log.Fatal("unable to create: ", filepath+filename)
	}
	defer f.Close()

	f.WriteString(`
package colornamesorted
var List = []struct {
        Name string
        Rgb  string
}{
`)

	lines := 6

	for _, v := range s {
		fmt.Fprintf(f, `    {Name: "%s", Rgb: "%s"},`, v.name, v.rgb)
		fmt.Fprintf(f, "\n")
		lines++
	}

	f.WriteString("}\n")
	lines++

	fmt.Printf("Wrote %v lines\n", lines)
}
