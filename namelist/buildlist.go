package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/clayessex/colorproc/colors"
)

const (
	ListFilename   = "colornames.csv.gz"
	OutputFilename = "colornames.go.out"
)

type Color struct {
	name string
	hex  colors.Hex
	hsl  colors.Hsl
	dist float64
}

// The choice of 84% yields 57.6 degrees of Hue and then 8% of each Saturation
// and Luminance
func distance(a, b colors.Hsl) float64 {
	x := float64(a.H-b.H) / 360.0
	y := float64(a.S - b.S)
	z := float64(a.L - b.L)
	return 0.84*x + 0.08*y + 0.08*z
}

func main() {
	fout, err := os.Create(OutputFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()

	f, err := os.Open(ListFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	z, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	defer z.Close()

	colorNames := make([]Color, 0, 30000)
	midpoint := colors.Hsl{180.0, 0.5, 0.5}

	scanner := bufio.NewScanner(z)
	maxLines := 20
	first := true
	for scanner.Scan() {
		if first {
			first = false
			continue
		}

		fields := strings.Split(scanner.Text(), ",")
		if len(fields) < 2 {
			log.Printf("Line is not valid CSV: [%s]", scanner.Text())
			continue
		}

		name := fields[0]
		hex := colors.Hex(fields[1])
		hsl, err := hex.ToHsl()
		if err != nil {
			log.Printf("Line color is not parsable: [%s]", scanner.Text())
		}
		dist := distance(hsl, midpoint)

		colorNames = append(colorNames, Color{name, hex, hsl, dist})

		if maxLines--; maxLines == 0 {
			log.Print("reached max lines")
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	slices.SortFunc(colorNames, func(a, b Color) int {
		if a.dist < b.dist {
			return -1
		} else if a.dist > b.dist {
			return 1
		} else {
			return 0
		}
	})

	fout.WriteString(`
const ColorNames = []struct{
    name string
    hex Hex
    hsl Hsl
}{
`)

	for _, v := range colorNames {
		hslstring := fmt.Sprintf("Hsl{%f, %f, %f}", v.hsl.H, v.hsl.S, v.hsl.L)
		fmt.Fprintf(fout, ` name: "%s", hex: "%v", hsl: %v, dist: %f`, v.name, v.hex, hslstring, v.dist)
		fmt.Fprintf(fout, "\n")
	}

	fout.WriteString("}\n")
}
