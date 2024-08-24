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
	rgb  colors.Hex
	hsl  colors.Hsl
	pos  float64
}

func main() {
	colorlist := make([]ColorInfo, 0, 30000)
	midpoint := colors.Hsl{H: 180.0, S: 0.5, L: 0.5}

	for _, v := range colornames.List {
		rgb := colors.Hex(v.Rgb)
		hsl, err := rgb.ToHsl()
		if err != nil {
			log.Printf("invalid rgb value: %s\n", v.Rgb)
			continue
		}
		pos := colors.Distance2(hsl, midpoint)
		colorlist = append(colorlist, ColorInfo{v.Name, rgb, hsl, pos})
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

	WriteColorNames(OutputPath, OutputFilename, colorlist)
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
