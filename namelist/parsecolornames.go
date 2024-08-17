package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	ListFilename   = "colornames.csv.gz"
	OutputFilename = "colornames.go.out"
)

type Color struct {
	name string
	rgb  string
}

func main() {
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
		rgb := fields[1]

		colorNames = append(colorNames, Color{name, rgb})

		if maxLines--; maxLines == 0 {
			log.Print("reached max lines")
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	writeColors(OutputFilename, colorNames)
}

// Writes to a file (named filename) the contents of colors as a Go const
// struct slice
func writeColors(filename string, colors []Color) {
	fout, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()

	fout.WriteString(`
const ColorNames = []struct{
    name string
    rgb string
}{
`)

	for _, v := range colors {
		fmt.Fprintf(fout, ` { name: "%s", rgb: "%v" },`, v.name, v.rgb)
		fmt.Fprintf(fout, "\n")
	}

	fout.WriteString("}\n")
}
